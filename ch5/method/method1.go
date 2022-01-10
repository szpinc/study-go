package main

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
	Age  int8
}

func (person Person) UpdateName(name string) {
	person.Name = name
}

// 获取String值
func (person Person) String() string {
	return fmt.Sprintf("[id:%d,name=%s,age=%d]", person.Age, person.Name, person.Age)
}

type Object interface {
	String() string
}

func main() {
	person := Person{1, "张三", 20}

	person.UpdateName("李四")

	fmt.Println("Person:", person)
}
