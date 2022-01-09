package users

import (
	"time"
	"ups/config/db"
	"ups/util"

	"github.com/google/uuid"
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
	// 执行sql
	util.Execute(db.Datasource(), "insert into sys_user(id,username,password,nickname,create_time,update_time) values(?,?,?,?,?,?)",
		user.Id,
		user.Username,
		user.Password,
		user.Nickname,
		user.CreateTime,
		user.UpdateTime)
}

func buildUser(username string, password string, nickname string) *User {
	user := User{
		Username: username,
		Nickname: nickname,
	}

	uid := uuid.New().ID()
	// 设置用户id
	user.Id = string(uid)
	// 设置密码,md5
	user.Password = util.Md5(password)

	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	return &user
}
