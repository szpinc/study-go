package users

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	password   string
	Nickname   string `json:"nickname"`
	Deleted    bool   `json:"deleted"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (user *User) UpdatePassword(password string) {
	user.password = password
}

func saveToFile(user *User) bool {
	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}

}

func saveToDatabase(user *User) bool {

}
