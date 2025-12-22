package yufa

import "fmt"

func init() {
	pointDemo()

}

func pointDemo() {
	// 指针：一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址
	var a = 10
	var p = &a
	fmt.Printf("a = %v type: %T address: %p\n", a, a, &a)
	fmt.Printf("p = %v type: %T", p, p)

}
