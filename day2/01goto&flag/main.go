package main

import (
	"fmt"
)

func main() {

	// //跳出多层for循环
	// for i := 0; i < 10; i++ {
	// 	var flag = false
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break//跳出该层的for循环
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)
	// 	}
	// 	if flag {
	// 		break//跳出for循环，外层的循环
	// 	}
	// }
	//跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto jump //跳到我指定的那个jump标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
jump: //label标签
	fmt.Printf("for已经over了")
}
