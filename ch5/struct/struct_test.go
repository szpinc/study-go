package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructPoint(t *testing.T) {

	// 创建结构体
	student := Student{}

	fmt.Println("未初始化:", student)

	fmt.Printf("student:%p,name:%p,age:%p\n", &student, &student.Name, &student.Age)
	// 赋值操作
	student.Name, student.Age = "张三", 20

	fmt.Println("Student:", student)
}

func TestStructJson() {

	student := Student{"1", "张三", 20}
	resultBytes, err := json.Marshal(student)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("resultBytes: %v\n", string(resultBytes))
}

type Student struct {
	Id   string
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}
