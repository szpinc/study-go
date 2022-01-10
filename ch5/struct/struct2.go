package main

import "fmt"

// 创建结构体的方式
func main() {

	// 1. 直接声明式创建，得到的是结构体的实例
	var p1 Person
	fmt.Println("p1:", p1)
	// 2. 使用new函数创建,得到的是一个指针
	p2 := new(Person)
	fmt.Printf("使用new()创建的是一个指针:%p\n", p2)
	// 3. 创建一个Person，接收其指针
	var p3 *Person = &Person{}

	*&p3.Name = "sd"

	// (*p3).Name = "张三"
	(*p3).Age = 20
	fmt.Printf("p3的指针地址:%p,p3的实例地址:%p", &p3, p3)
}

type Person struct {
	Name string
	Age  int
}
