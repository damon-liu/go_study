package jiegouti

import (
	"encoding/json"
	"fmt"
)

func init() {
	// // 结构体初始化
	// structInitDemo()

	// // 结构体方法
	// structThmeodDemo()

	// // 结构体转Json
	// struct2JsonDemo()
}

func struct2JsonDemo() {
	// 结构体转json（参数首字母必须大写）
	p1 := person{
		Name: "damon1",
		Age:  18,
	}
	jsonbyte, _ := json.Marshal(p1)
	jsonstr := string(jsonbyte)
	fmt.Println("jsonstr:  ", jsonstr)

	// json转结构体
	str := `{"Name":"damon1","Age":18,"Sex":0}`
	var p2 person
	err := json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("p2: %#v ", p2)
	}

	// 复杂结构体转换
	c := class{
		Name:    "三年一班",
		Student: make([]student, 0),
	}
	for i := 1; i <= 5; i++ {
		s := student{
			Id:     int8(i),
			Gender: "男",
			Name:   fmt.Sprintf("str_%v", i),
		}
		c.Student = append(c.Student, s)
	}
	bytejson, _ := json.Marshal(c)
	fmt.Printf("studenJson: %v \n", string(bytejson))

	str1 := `{"Name":"三年一班","Student":[{"Id":1,"Gender":"男","Name":"str_1"},{"Id":2,"Gender":"男","Name":"str_2"},{"Id":3,"Gender":"男","Name":"str_3"},{"Id":4,"Gender":"男","Name":"str_4"},{"Id":5,"Gender":"男","Name":"str_5"}]}`
	var c1 class
	json.Unmarshal([]byte(str1), &c1)
	fmt.Printf("str1: %#v", c1)
}

type class struct {
	Name    string
	Student []student
}

type student struct {
	Id     int8
	Gender string
	Name   string
}

type person struct {
	Name string `json:"id"` //结构体标签
	Age  int8   `json:"age"`
	Sex  int8
}

func (p person) PrintfPersonInfo() {
	fmt.Printf("printf person info: name is %v  age is %v sex is %v\n", p.Name, p.Age, p.Sex)
}

func (p *person) SetPersonInfo(name string, age int8, sex int8) bool {
	p.Name = name
	p.Age = age
	p.Sex = sex
	return true
}

func structThmeodDemo() {
	/*
		结构体方法:
			func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) ｛
				函数体
			｝
	*/

	p1 := person{
		Name: "damon1",
		Age:  18,
	}
	p1.PrintfPersonInfo() //printf person info: name is damon1  age is 18 sex is 0
	bl := p1.SetPersonInfo("damon2", 19, 1)
	if bl {
		p1.PrintfPersonInfo() //printf person info: name is damon2  age is 19 sex is 1
	}

}

type user struct {
	name   string
	age    int8
	sex    int8
	Hobby  []string
	Score  map[string]int8
	person person //结构体嵌套
}

func structInitDemo() {
	/*
		type 类型名 struct {
			字段名1 字段类型
			字段名2 字段类型
		}

		结构体实例是相互独立的，不会相互影响

		结构体的字段类型可以是：基本数据类型、切片、map以及结构体
		如果结构体的字段类型是：指针，slice,和map的零值都是nil，即还没有分配空间，需要先make才能使用
	*/
	// 结构体实例化
	var u1 user
	u1.name = "damon"
	u1.sex = 1
	u1.age = 18
	u1.Hobby = make([]string, 2, 5)
	u1.Hobby[0] = "sleep"
	u1.Hobby[1] = "code"
	u1.Score = make(map[string]int8)
	u1.Score["english"] = 110
	u1.Score["science"] = 100
	u1.person.Age = 18
	u1.person.Name = "sss"
	u1.person.Sex = 100
	fmt.Printf("u1: %#v  类型：%T \n", u1, u1) //u1: yufa.user{name:"damon", age:18, sex:1} 类型：yufa.user

	var u2 = new(user)
	u1.name = "damon2"
	u1.sex = 1
	u1.age = 13
	fmt.Printf("u2: %#v  类型：%T \n", u2, u2) //u2: &yufa.user{name:"", age:0, sex:0} 类型：*yufa.user

	var u3 = &user{
		name: "damon3",
		age:  20,
		sex:  0,
	}
	fmt.Printf("u3: %#v  类型：%T \n", u3, u3) //u3: &yufa.user{name:"damon3", age:20, sex:0}  类型：*yufa.user

	var u4 = user{
		name: "damon4",
		age:  21,
		sex:  0,
	}
	fmt.Printf("u4: %#v  类型：%T \n", u4, u4) //u4: yufa.user{name:"damon4", age:21, sex:0}  类型：yufa.user
}
