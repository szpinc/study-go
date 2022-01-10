package config

import (
	"database/sql"
	"fmt"
)

const (
	username = ""
	password = ""
	ip       = ""
	port     = "3306"
	database = ""
)

var DB *sql.DB

func InitDB() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, database)

	datasource, err := sql.Open("mysql", url)

	if err != nil {
		fmt.Printf("datasource init err: %v\n", err)
		return
	}
	// 设置最大连接数
	datasource.SetConnMaxLifetime(100)
	// 设置最大空闲连接数
	datasource.SetMaxIdleConns(20)

	if err := datasource.Ping(); err != nil {
		fmt.Printf("mysql connect err: %v\n", err)
		return
	}
	fmt.Println("mysql connect success")
}
