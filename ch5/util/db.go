package util

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	config      SqlConf
	maxPoolSize int
)

type Conn sql.DB

// 对sql.DB的封装
type Connection struct {
	Db *sql.DB
}

const (
	defaultPoolSize = 10
)

var pool []*Connection

// db配置结构体
type conf struct {
	db SqlConf `yaml:"db"`
}

// 数据源配置结构体
type SqlConf struct {
	Host     string   `yaml:"host"`     // 地址
	Port     int      `yaml:"port"`     // 端口
	User     string   `yaml:"user"`     // 用户名
	Password string   `yaml:"password"` // 密码
	Driver   string   `yaml:"driver"`   // 驱动
	Database string   `yaml:"database"` // 数据库
	poolConf poolConf `yaml:"pool"`     // 连接池配置
}

// 连接池配置结体
type poolConf struct {
	maxSize int `yaml:"maxSize"` // 连接池最大连接数
}

func init() {
	// 读取配置文件
	confData, err := ioutil.ReadFile("../conf/conf.yaml")

	if err != nil {
		panic(err)
	}

	conf := conf{}
	// 解析yaml配置文件
	e := yaml.Unmarshal(confData, &conf)

	if e != nil {
		panic(e)
	}

	// todo 从配置文件读取
	maxPoolSize = defaultPoolSize

	pool = []*Connection{}

	for i := 0; i < maxPoolSize; i++ {

		c, err := newConn()
		if err != nil {

		}
		pool = append(pool, c)
	}
}

func GetConnection() (*Connection, error) {
	// 连接池有连接，则从连接池中取出
	if len(pool) > 0 {
		return pool[0], nil
	}
	// 连接池中没有
	// 创建新连接
	c, err := newConn()

	if err != nil {
		return nil, err
	}

	pool = append(pool, c)

	return c, nil
}

func newConn() (*Connection, error) {

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Password, config.Host, config.Port, config.Database)

	d, err := sql.Open(config.Driver, datasource)

	if err != nil {
		return nil, err
	}
	conn := Connection{
		Db: d,
	}
	return &conn, nil
}

func (connection *Connection) CloseInPool() error {
	// pool还有空间，则回收连接
	if len(pool) < maxPoolSize {
		pool = append(pool, connection)
		return nil
	}
	// 连接池没有空间了，则直接关闭连接
	return connection.Db.Close()
}

func Exec(sql string, args ...interface{}) (_ sql.Result, err error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	return conn.Db.Exec(sql, args...)
}
