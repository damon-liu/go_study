package jiekou

import "fmt"

func init() {
	interfaceDemo()

	nilInterfaceDemo()
}

func nilInterfaceDemo() {

}

func interfaceDemo() {
	/*
		接口是一种数据类型，但是是抽象的类型，接口是一组methond的集合，Golang中的接口不能包含任何变量
		接口中所有的方法都没有方法提，接口定义了一个对象的行为规范，只定义规范不实现
		type 接口名 interface{
			方法名1(参数列表1) 返回值列表1
			方法名2(参数列表2) 返回值列表2
			...
		}

		空接口：接口可以不定义任何方法，没定义任何方法的接口就是空接口，没有任何约束

	*/
	p := phone{
		Name: "iphone",
	}
	c := cacmer{
		Name: "keda",
	}
	computer := computer{
		Name: "huawei",
	}
	computer.transData(p)
	computer.transData(c)
}

type usber interface {
	Start()
	Stop()
}

type computer struct {
	Name string
}

type phone struct {
	Name string
}

func (p phone) Start() { fmt.Printf("%v  is start \n", p.Name) }

func (p phone) Stop() { fmt.Printf("%v  is stop \n", p.Name) }

type cacmer struct {
	Name string
}

func (p cacmer) Start() {
	fmt.Printf("%v  is start \n", p.Name)
}

func (p cacmer) Stop() {
	fmt.Printf("%v is stop \n", p.Name)
}

func (c computer) transData(u usber) {
	u.Start()
	fmt.Printf("%v computer is working \n", c.Name)
	u.Stop()
}
