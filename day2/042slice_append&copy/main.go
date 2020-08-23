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

	s1 = append(s1, "南京1")
	fmt.Printf("%v,%d\n", s1, cap(s1))

	s1 = append(s1, "南京2", "南京3")
	fmt.Printf("%v,%d\n", s1, cap(s1))

	//把切片ss中的元素拆开放入s1中
	ss := []string{"武汉", "南京", "上海"}
	s1 = append(s1, ss...) //...表示拆开
	fmt.Printf("%v", s1)
}
