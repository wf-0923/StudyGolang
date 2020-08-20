package main

import (
	"fmt"
)

func main() {
	//for的基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//变种2
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//无线循环
	// for {
	// 	fmt.Println(123)
	// }

	fmt.Println()
	//for range 循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)

	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}

}
