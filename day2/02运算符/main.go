package main

import (
	"fmt"
)

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	//at.1/单独的语句，不能放在-的右边赋值> a-a +1
	//b--/单独的语句，不能放在-的右边赋值-> b - b - 1
	fmt.Println(a)

}
