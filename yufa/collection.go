package yufa

import (
	"encoding/json"
	"fmt"
)

func init() {
	/*

		// map
		mapdemo()

		// 数组
		arraydemo()

	*/
	// 数组
	arraydemo()
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

	map3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	jsonData, _ := json.MarshalIndent(map3, "", "  ")
	fmt.Println("map3 :", string(jsonData))
}
