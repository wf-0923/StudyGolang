package main

import (
	"fmt"
	"strings"
)

func main() {

	path := "\"D:\\1study\\Typora笔记\\Go\""
	fmt.Println(path)
	//多行的字符串
	s2 := `
			第一句古诗
			第二句古诗
			第三句古诗
			`
	fmt.Println(s2)

	s3 := `D:\\1study\\Typora笔记\\Go`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))

}
