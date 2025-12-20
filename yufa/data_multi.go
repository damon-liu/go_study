package yufa

import (
	"encoding/json"
	"fmt"
	"sort"
)

func init() {
	/*

		// map
		mapdemo()

		// 数组
		arraydemo()

		// 值类型与引用类型
		valueAndReferenceType()
	*/

	// mapdemo()

}

func valueAndReferenceType() {
	fmt.Println("\n------------------  值类型与引用类型 demo ------------------")
	/*
		值类型：基本数据类型（int、float、bool、string）、数组、结构体
		引用类型：指针、切片、map、chan、interface、函数

		值类型赋值时，是将值复制一份给新的变量，两个变量互不影响
		引用类型赋值时，是将引用地址复制一份给新的变量，两个变量指向同一块内存，修改其中一个变量会影响另一个变量
	*/
	// int值类型
	a := 10
	b := a
	b = 20
	fmt.Printf("a=%d , b=%d \n", a, b)

	// 数组值类型
	var array1 = [3]int{1, 2, 3}
	array2 := array1
	array2[0] = 10
	fmt.Printf("array1=%v , array2=%v \n", array1, array2)

	// 切片引用类型
	var array3 = []int{1, 2, 3}
	slice1 := array3
	slice1[0] = 10
	fmt.Printf("array2=%v , slice2=%v \n", array3, slice1)
}

func arraydemo() {
	fmt.Println("\n---------------------------  array demo ---------------------------")
	array2 := [10]int{1, 2, 3, 4}
	fmt.Println("array2:  ", array2)

	array3 := [...]string{"java", "c++", "python", "php"}
	array3[0] = "golang"
	fmt.Println("array3:  ", array3)

	array4 := [...]string{1: "java", 3: "c++", 0: "python", 2: "php"}
	fmt.Println("array4:  ", array4)

	for i := 0; i < len(array4); i++ {
		fmt.Printf("array4[%d]=%s \n", i, array4[i])
	}
	for k, v := range array4 {
		fmt.Printf("array4[%d]=%s \n", k, v)
	}

	var array5 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max := array5[0]
	index5 := 0
	for i := 0; i < len(array5); i++ {
		if array5[i] > max {
			max = array5[i]
			index5 = i
		}
	}
	fmt.Printf("max=%d , index=%d \n", max, index5)

	// 多维数组
	var array6 = [...][2]string{
		{"java", "c++"},
		{"python", "php"},
		{"golang", "ruby"},
	}
	fmt.Println(array6[0][1])
	for _, v1 := range array6 {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
	}
}

func mapdemo() {
	fmt.Println("\n---------------------------  map demo ----------------------------")
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 is null !")
	}
	map1 = make(map[string]string, 10)
	map1["key1"] = "java"
	map1["key2"] = "c++"
	map1["key3"] = "python"

	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "python"

	// 删除map
	map3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	delete(map3, "two")
	jsonData, _ := json.MarshalIndent(map3, "", "  ")
	v, ok := map3["one"]
	if ok {
		fmt.Println("map3 key one value :", v)
	}
	fmt.Println("map3 :", string(jsonData))
	for k, v := range map3 {
		fmt.Printf("map3[%s]=%s \n", k, v)
	}

	// map排序
	map4 := map[int]int{
		1: 10,
		5: 50,
		2: 20,
		6: 60,
	}
	var keySlice []int
	for k := range map4 {
		keySlice = append(keySlice, k)
	}
	sort.Ints(keySlice)
	fmt.Println("keySlice： ", keySlice)
	for _, v := range keySlice {
		fmt.Printf("key = %v , value = %v \n", v, map4[v])
	}

	// 统计字符串中的每个字符的个数
	str := "hello,hello"
	mapch := make(map[string]int)
	for _, ch := range str {
		key := string(ch)
		mapch[key]++
	}
	for k, v := range mapch {
		fmt.Printf("k = %v,  count = %v \n", k, v)
	}
}
