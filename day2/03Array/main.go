package main

import (
	"fmt"
)

var a1 [3]bool
var a2 [4]bool

func main() {
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	fmt.Println(a1, a2)

	//数组的初始化
	// 如果不初始化：默认元素都是零值（布尔型：false，整型和浮点型：0，字符串：“”）

	// 1.初始化方式 1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2
	// a11 := [11]int{0,1,2,3,4,5,6,7,8,9,10}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("a10:%T\n", a10)
	fmt.Println(a10)

	// 3.初始化方式3
	// 根据索引来初始化
	a3 := [5]int{1: 1, 2: 2}
	fmt.Printf("a3:%d\n", a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for _, h := range citys {
		fmt.Println(h)
	}
	//多维数组
	var a4 [3][3]int
	a4 = [3][3]int{
		[3]int{0, 0, 0},
		[3]int{1, 1, 1},
		[3]int{2, 2, 2},
	}
	fmt.Println(a4)

	// 多维数组的遍历
	for _, h := range a4 {

		fmt.Println(h)
		for _, v := range h {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
	// b2中修改不影响b1

	//数组练习题1：求数组[1,3,5,7,8]所有值得和
	//方法一：标准的for循环方法
	fmt.Printf("\n数组练习题1：求数组[1,3,5,7,8]所有值得和")
	L1 := [5]int{1, 3, 5, 7, 8}
	var sum int = 0
	fmt.Printf("\n方法一：标准的for循环方法\n")
	for L := 0; L < len(L1); L++ {
		sum += L1[L]
		fmt.Printf("sum加了数组L1的第[%d]项，和值为%d\n", L, sum)
	}
	fmt.Printf("数组L1的所有值的和为%d\n\n", sum)
	// 方法二：range循环的方法
	fmt.Printf("方法二：range循环的方法\n")
	var sum1 int = 0
	for L, L11 := range L1 {
		sum1 += L11
		fmt.Printf("sum加了数组L1的第[%d]项，和值为%d\n", L, sum1)
	}
	fmt.Printf("数组L1的所有值的和为%d\n\n", sum1)

	//数组练习题2:找出数组中和为指定值的两个元素的下标，
	//比如从数组[1，3，5，7，8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	fmt.Printf("\n数组练习题2:找出数组中和为指定值的两个元素的下标")
	// 方法一：标准的for循环方法
	fmt.Printf("\n方法一：标准的for循环方法\n")
	for i := 0; i < len(L1); i++ {
		for j := i + 1; j < len(L1)-1; j++ {
			if L1[i]+L1[j] == 8 {
				fmt.Printf("和为8的两个元素的下标为(%d,%d)\n", i, j)
			}
		}
	}

	//方法二：用range
	fmt.Printf("方法二：range循环的方法\n")
	L2 := L1
	for s0, s1 := range L1 {
		// fmt.Printf("%d\t", s1)
		//为了避免输出时出现（0，3）（3，0）,用L2把已经参与计算的变为很大的数避免干扰
		L2[s0] = 99999
		for s01, s2 := range L2 {
			// fmt.Printf("%d\t\n", s2)
			if s1+s2 == 8 {
				fmt.Printf("和为8的两个元素的下标为(%d,%d)\n", s0, s01)
			}
		}
	}

}
