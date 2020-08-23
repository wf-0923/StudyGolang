package main

import (
	"fmt"
)

var (
// name string
// age  int
// isok bool
)

func main() {
	fmt.Println("Hello World!123")
	fmt.Println("git测试")
	name := "wf"
	age := 21
	var isok = true

	fmt.Printf("name:%s", name) //%s:占位符，使用name这边个变量的值去替换占位符
	fmt.Println()
	fmt.Println(age)
	fmt.Print(isok)

}
