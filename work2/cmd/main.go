package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Beijing!"))
}

func startHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", sayHello)
	fmt.Println("http server start!")
	return srv.ListenAndServe()
}

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// server
	server := http.Server{Addr: ":8000"}

	// 1、server启动
	g.Go(func() error {
		return startHttpServer(&server)
	})

	g.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(ctx)
	})

	g.Go(func() error {
		// linux signal 信号的注册和处理（chan处理）
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-c:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("errgroup end exiting:%v\n", err)
	}
}
