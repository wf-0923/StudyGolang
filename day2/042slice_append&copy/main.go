package main

import (
	"fmt"
)

func main() {

	//初始化一个切片
	s1 := []string{"北京", "上海", "广州"}
	//打印这个切片
	fmt.Println(s1)
	fmt.Printf("%d\n", cap(s1))
	//给切片追加元素
	s1 = append(s1, "深圳")
	fmt.Println(s1)
	fmt.Printf("%d\n", cap(s1))

}
