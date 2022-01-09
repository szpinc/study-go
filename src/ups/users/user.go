package users

import (
	"time"
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

}
