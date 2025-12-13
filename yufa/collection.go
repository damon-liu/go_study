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
		arrayDemo()

	*/
}

func arrayDemo() {
	fmt.Println("\n---------------------------  arrayDemo ---------------------------")
	array2 := [10]int{1, 2, 3, 4}
	fmt.Println("array2:  ", array2)
}

func mapdemo() {
	fmt.Println("\n---------------------------  mapdemo ----------------------------")
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
