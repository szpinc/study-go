package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// mysql配置常量
const (
	host               = "42.193.175.65"
	port               = "3306"
	user               = "go_study"
	password           = "drkJazG4fyL28hN7"
	database           = "go_study"
	maxIdleConnections = 100
	maxOpenConnections = 20
)

type config struct {
	Host               string
	Port               string
	User               string
	Password           string
	Database           string
	MaxIdleConnections int
	MaxOpenConnections int
}

func MysqlConfig() *config {
	return &config{
		Host:               host,
		Port:               port,
		User:               user,
		Password:           password,
		Database:           database,
		MaxIdleConnections: maxIdleConnections,
		MaxOpenConnections: maxOpenConnections,
	}
}

// 创建数据源
func Datasource() (*sql.DB, error) {
	config := MysqlConfig()
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	return db, nil
}
