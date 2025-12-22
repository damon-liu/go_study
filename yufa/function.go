package yufa

import (
	"fmt"
	"sort"
)

func init() {
	// function()
}

func function() {
	/*
		函数定义格式:
		1、固定参数
		func 函数名(参数)(返回值){
			....
		}

		2、可变参数
		func 函数名(...)(返回值){
			....
		}
	*/

	// 求两数和
	sum1 := sumFx(12, 3)
	fmt.Println("sum1: ", sum1)

	sum2 := sumFx2(1, 2, 3, 4)
	fmt.Println("sum2: ", sum2)

	// 返回多个值
	sum, sub := fub2(12, 3)
	fmt.Printf("sum = %v , sub = %v \n", sum, sub)

	m := map[int]string{
		1: "hello",
		4: "go!",
		3: "damon",
		2: "world",
	}
	mapKeySort(m)

}

// 固定参数
func sumFx(x int, y int) int {
	sum := x + y
	return sum
}

func sumFx1(x, y int) int {
	sum := x + y
	return sum
}

// 可变参数
func sumFx2(x ...int) int {
	fmt.Printf("%x --- %T", x, x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func sumFx3(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// return多个返回值
func fub(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func fub2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func mapKeySort(m map[int]string) map[int]string {
	var s1 []int
	for k := range m {
		s1 = append(s1, k)
	}
	sort.Ints(s1)
	m1 := make(map[int]string)
	for _, v := range s1 {
		temp := m[v]
		m1[v] = temp
		fmt.Printf("key = %v, v = %v \n", v, temp)
	}
	return m1
}
