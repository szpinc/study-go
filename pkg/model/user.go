package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/szpinc/study-go/logger"
	"github.com/szpinc/study-go/pkg/db"
	"github.com/szpinc/study-go/util"
)

var (
	log *logrus.Logger = logger.Logger() // 日志
)

// 用户信息结构体
type User struct {
	Id         string    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Nickname   string    `json:"nickname"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func Register(username string, password string, nickname string) (*User, error) {
	// 构建用户信息结构体
	user := buildUser(username, password, nickname)
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	// 执行sql
	_, err = conn.Db.Exec("insert into sys_user(id,username,password,nickname,create_time,update_time) values(?,?,?,?,?,?)",
		user.Id,
		user.Username,
		user.Password,
		user.Nickname,
		user.CreateTime,
		user.UpdateTime)

	if err != nil {
		log.Error("sql执行错误", err)
		return nil, err
	}
	return user, nil
}

// 构建用户结构体
func buildUser(username string, password string, nickname string) *User {
	user := User{
		Username: username,
		Nickname: nickname,
	}

	uid := uuid.New().ID()
	// 设置用户id
	user.Id = string(rune(uid))
	// 设置密码,md5
	user.Password = util.Md5String32([]byte(password))
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	return &user
}
