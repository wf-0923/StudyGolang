package main

// make()函数创造切片

import (
	"fmt"
)

func main() {
	//疑惑：这里的长度比容量小
	// 解惑：这里创建的是一个切片，切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	// 存在比切片长的可能，
	s1 := make([]int, 5, 10)
	fmt.Printf("s1:%v\tlen(s1):%d\tcap(s1):%d\n", s1, len(s1), cap(s1))
	s1[1] = 1
	fmt.Printf("s1:%v\tlen(s1):%d\tcap(s1):%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2:%v\tlen(s2):%d\tcap(s2):%d\n", s2, len(s2), cap(s2))

	//切片的赋值
	fmt.Printf("切片的赋值:\n")
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)
	//数组的赋值
	fmt.Printf("数组的赋值:\n")
	a1 := [...]int{1, 3, 5}
	a2 := a1
	fmt.Println(a1, a2)
	a1[0] = 1000
	fmt.Println(a1, a2)
	//比较结果：切片为引用类型：可以看到切片赋值后，源发生改变，切片也会随之变化--类比理解：快捷方式
	//         数组为数值类型：a1中值发生变化后，a2中不会随之变化，	--类比理解，文件的副本

	// 切片的遍历
	s4 = []int{1, 2, 3, 4, 5, 6}
	// 1.索引遍历
	fmt.Printf("1.索引遍历\n")
	for i := 0; i < len(s4); i++ {
		fmt.Printf("s4[%d]:%d\t", i, s4[i])
	}
	// 2.range 循环
	fmt.Printf("\n2.range 循环\n")
	for i, h := range s4 {
		fmt.Printf("s4[%d]:%d\t", i, h)
	}
	fmt.Printf("",)
}
