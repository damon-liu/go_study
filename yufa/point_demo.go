package yufa

import "fmt"

func init() {
	// pointDemo()
}

func pointDemo() {
	// 指针：一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址
	var a = 10
	var p = &a
	fmt.Printf("a = %v type: %T 内存地址: %p\n", a, a, &a)
	fmt.Printf("p = %v type: %T 内存地址对应的值：%v \n", p, p, *p)
	*p = 30 // 通过*p改变对应内存地址
	fmt.Printf("a = %v \n", a)

	x := 5
	fn2(x)
	fmt.Println("x = ", x)
	fn3(&x)
	fmt.Println("x = ", x)
}

func fn2(x int) {
	x = 10
}

func fn3(x *int) {
	*x = 20
}
