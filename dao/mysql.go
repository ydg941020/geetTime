package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// user 表
type user struct {
	id   int
	age  int
	name string
}

func QueryMultiRowDemo(db *sql.DB) (*user, error) {
	var u user
	sqlStr := "select id,name,age from user limit 1"
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Not Found.")
		} else {
			fmt.Printf("query failed, err:%v\n", err)
			return nil, err
		}
	}

	return &u, nil
}
