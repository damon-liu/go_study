package yufa

import (
	"fmt"
)

func init() {
	/*
		定义函数类型：
			type calc func(int, int) int

		匿名函数：
			func(参数)(返回值){
				....
			}

		闭包：
			1、闭包是指有权访问另一个函数作用域中的变量的函数
			2、创建闭包的常见方式就是在一个函数内部创建另外一个函数，通过另外一个函数访问这个函数的局部
			作用：
				可以让一个变量常驻内存
				可以让一个变量不污染全局
			注意：
				由于闭包作用域返回的局部变量资源不会被立即销毁回收，所以可能回占用更多的内存。过度使用闭包会导致性能下降，建议在非常有必要的时候才使用闭包

		defer:
			在go语言函数中return语句底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前
			defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
		panic:
			可以在任何地方引发，但recover只有在defer调用的函数中有效

	*/
	// funcType()
}

func funcType() {
	// 自定义函数当作参数
	var a = 10
	var b = 4
	s1 := calc(a, b, sum)
	s2 := calc(a, b, sub)

	// 匿名函数当作参数
	s3 := calc(a, b, func(x, y int) int {
		return x * y
	})

	// 函数当作返回值
	str := "+"
	var d calcType = do(str)
	s4 := d(a, b)
	fmt.Printf("a: %v, b: %v, a+b=%v, a-b=%v, a*b=%v, a%vb=%v", a, b, s1, s2, s3, str, s4)

	// 匿名自执行函数
	func() {
		fmt.Println("匿名自执行函数...")
	}()
	s5 := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println("匿名函数中结果：", s5)

	// 闭包
	fn := adder2()
	s6 := fn(10)
	s7 := fn(10)
	s8 := fn(10)
	fmt.Printf("闭包函数中结果：%v, %v, %v \n", s6, s7, s8)
}

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 自定义方法类型
type calcType func(int, int) int

// 自定义方法当作参数
func calc(x, y int, ct calcType) int {
	return ct(x, y)
}

func do(o string) calcType {
	switch o {
	case "+":
		return sum
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}
