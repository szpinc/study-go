package main

import "fmt"

func main() {

	// 创建结构体
	student := Student{}

	fmt.Println("未初始化:", student)

	fmt.Printf("student:%p,name:%p,age:%p\n", &student, &student.Name, &student.Age)
	// 赋值操作
	student.Name, student.Age = "张三", 20

	fmt.Println("Student:", student)
}

type Student struct {
	Name string
	Age  int
}
