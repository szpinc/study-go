package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   string `json:"id"` //通过打标签的方式让json字段首字母小写 (真是扯淡)
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func main() {

	student := Student{"1", "张三", 20}
	resultBytes, err := json.Marshal(student)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("resultBytes: %v\n", string(resultBytes))
}
