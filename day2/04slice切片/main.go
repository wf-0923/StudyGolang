package main

//切片slice
import (
	"fmt"
)

func main() {

	//1.切片的定义
	fmt.Printf("1.切片的定义\n ")
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放bool类型元素的切片
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)
	//长度和容量
	fmt.Printf("S1长度:%d,容量：%d\n", len(s1), cap(s1))
	fmt.Printf("S2长度:%d,容量：%d\n", len(s2), cap(s2))

	//2.由数组得到切片
	fmt.Print("2.由数组得到切片\n")
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[0:4] //基于一个数组切割，左包含右不包含，左开右闭
	fmt.Printf("S3:%d\n", s3)
	s4 := a1[1:5]
	fmt.Printf("S4:%d\tcap:%d\n", s4, cap(s4))
	s5 := a1[:4]
	s6 := a1[5:]
	s7 := a1[:]
	fmt.Printf("s5:%d\tS6:%d\tS7:%d\n", s5, s6, s7)
	//切片指向了一个底层的数组
	//切片的长度是他元素个个数
	fmt.Printf("s5:%d\tlen(s5):%d\n", s5, len(s5))
	//切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("s5的底层数组是a1：%d\tlen(a1):%d\n", a1, len(a1))
	fmt.Printf("s5:%d\tcap(s5):%d\n", s5, cap(s5))

	//3.切片再切割
	s8 := s6[2:]
	fmt.Printf("s8:%d\tlen(s8):%d\tcap(s8):%d\n\n", s8, len(s8), cap(s8))
	//切片是引用类型，都执行了底层的一个数组，
	fmt.Printf("s6:%d\n", s6)
	a1[7] = 8000 //修改了底层数组的值，切片中的值也会改变
	fmt.Printf("a1[7]:%d\n", a1[7])
	fmt.Printf("s6:%d\n", s6)
	fmt.Printf("s8:%d\n", s8)

}
