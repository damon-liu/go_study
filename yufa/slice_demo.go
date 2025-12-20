package yufa

import (
	"fmt"
)

func init() {
	// sliceDemo()

	// sortAlgorithm()
}

func sortAlgorithm() {
	fmt.Println("\n---------------------------  sort algorithm demo ---------------------------")
	/*
		常见排序算法有冒泡排序、选择排序、插入排序、归并排序、快速排序等
		Go语言标准库中提供了sort包，包含了多种排序算法的实现，可以方便地对切片和自定义数据类型进行排序
	*/
	numSlice := []int{5, 2, 6, 3, 1, 4}
	for i := 0; i < len(numSlice)-1; i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
		fmt.Println("sorted numSlice: ", numSlice)
	}
}

func sliceDemo() {
	fmt.Println("\n---------------------------  slice demo ---------------------------")
	/*
		切片是对数组的一个引用，是引用类型
		是一个拥有相同类型元素的可变长度的序列，他是基于数组类型做的一层封装
		切片的底层是一个数组，切片本身不存储任何数据，他只是对底层数组的一个引用
		语法：
			var slice_name []type = []type{value1, value2, ...}
			slice_name := []type{value1, value2, ...}
			slice_name := make([]type, length, capacity)
		长度和容量：
			长度：切片中元素的个数，可以使用len()函数获取
			容量：切片的底层数组的容量，可以使用cap()函数获取
		切片的切片：
			slice_name[low : high]  包含从low到high-1的元素
	*/
	// 切片的定义
	var slice []int
	fmt.Println("slice: ", slice, slice == nil)
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)

	slice2 := []string{"java", "c++", "python", "php"}
	fmt.Println("slice2: ", slice2)
	for _, v := range slice2 {
		fmt.Println("value: ", v)
	}

	slice3 := make([]int, 5, 10)
	fmt.Println("slice3: ", slice3, slice3 == nil)

	// 数组转切片
	array := [5]int{10, 20, 30, 40, 50}
	slice4 := array[1:4]
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice4 length: ", len(slice4))
	fmt.Println("slice4 capacity: ", cap(slice4))

	// 切片动态扩容
	var slice5 []int
	slice5 = append(slice5, 12)
	slice5 = append(slice5, 22, 32, 42)
	fmt.Printf("slice5: %v , length: %d , capacity: %d \n", slice5, len(slice5), cap(slice5))

	//  合并切片
	slice6 := []int{1, 2, 3}
	slice7 := []int{4, 5, 6}
	slice6 = append(slice6, slice7...)
	fmt.Printf("slice6: %v , length: %d , capacity: %d \n", slice6, len(slice6), cap(slice6))

	// 复制切片
	slice8 := make([]int, 1, 3)
	slice8 = append(slice8, 100, 200)
	copy(slice8, slice6)
	fmt.Printf("slice8: %v , length: %d , capacity: %d \n", slice8, len(slice8), cap(slice8))

	// 删除元素
	slice9 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 4
	slice9 = append(slice9[:index], slice9[index+1:]...)
	fmt.Printf("slice9: %v , length: %d , capacity: %d \n", slice9, len(slice9), cap(slice9))

	// 切片嵌套map
	slice10 := make([]map[string]string, 3)
	slice10[0] = map[string]string{"name": "tom", "age": "20"}
	slice10[1] = map[string]string{"name": "jack", "age": "22"}
	slice10[2] = map[string]string{"name": "mary", "age": "19"}
	fmt.Println("slice10: ", slice10)

	// 切片嵌套切片，当切片的map对象中存放一系列属性，map的value是一个切片
	slice11 := make([]map[string][]string, 2)
	slice11[0] = map[string][]string{
		"hobby": {"篮球", "足球"},
	}
	slice11[1] = map[string][]string{
		"hobby": {"看书", "旅游"},
	}
	fmt.Println("slice11: ", slice11)
}
