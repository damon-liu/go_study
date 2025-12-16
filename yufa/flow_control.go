package yufa

import (
	"fmt"
)

func init() {
	flowControl()
}

func flowControl() {
	fmt.Println("\n----------------------------------------------------  yufa/flow control ----------------------------------------------------")

	/*
		if条件判断格式：
		if 条件表达式 {
			执行语句
		} else if 条件表达式 {
			执行语句
		} else {
			执行语句
		}
	*/
	fmt.Println("\n----------------------------  if条件判断 demo ----------------------------")
	flage := true
	if flage {
		fmt.Println("flage is true")
	} else {
		fmt.Println("flage is false")
	}

	if age := 20; age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("teenager")
	}

	fmt.Println("\n----------------------------  for条件判断 demo ----------------------------")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i = ", i)
	// }

	// for i:=0; i<=50; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("even number: ", i)
	// 	}
	// }

	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1 + 2 + ... + 100 = ", sum)

	// 打印1~100之间所有是9的倍数的整数的个数及总和
	cout := 0
	sum1 := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			cout++
			sum1 += i
		}
	}
	fmt.Println("count of multiples of 9: ", cout, " sum of multiples of 9: ", sum1)

	str := "你好golang"
	for k, v := range str {
		fmt.Printf("k = %d v = %c \n", k, v)
	}

	arr := []string{"hello", "world", "golang"}
	for _, v := range arr {
		fmt.Printf(" v = %s \n", v)
	}

	fmt.Println("\n----------------------------  switch条件判断 demo ----------------------------")
	/*
		switch expression {
		case value1:
			// 执行语句
		case value2:
			// 执行语句
		default:
			// 执行语句
		}
	*/
	switch sw := "jpg"; sw {
	case "jpg":
		fmt.Println("this is a jpg file")
	case "png":
		fmt.Println("this is a png file")
	default:
		fmt.Println("unknown file type")
	}

	switch score := "A"; score {
	case "A", "B", "C":
		fmt.Println("pass")
	case "D", "E":
		fmt.Println("fail")
	}

	fmt.Println("\n----------------------------  break/continue条件判断 demo ----------------------------")
	/*
		break用于终止当前循环
		contiue跳出当前循环继续下次循环
	*/
	// flag:
	i := 0
	for i < 10 {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break flag
				break
			}
			fmt.Printf("i = %d , j = %d \n", i, j)
		}
		i++
		if i == 2 {
			continue
		}
	}

	fmt.Println("\n----------------------------  goto demo ----------------------------")
	age := 30
	if age > 24 {
		fmt.Println("成年人！")
		goto lable
	}
	fmt.Println("AAA")
	fmt.Println("BBB")
lable:
	fmt.Println("CCC")
}
