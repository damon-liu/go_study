package yufa

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
	// "github.com/shopspring/decimal"
)

func init() {
	/*

		// fmt打印
		printfDemo()

		// 申明变量
		defDemo()

		// 常量计数器
		iotaDemo()

		// 数据类型
		dataType()

	*/
}

func dataType() {
	fmt.Println("\n------------------------------------------------------  dataTypeDemo ------------------------------------------------------")
	/*
		整型：int
		有符号整型按长度分为：int8、 int16、int32、int64
		无符号整型：uint8、uint16、uint32、uint64

		浮点型：float
		float32：占用4个字节
		float64：占用32个字节

		布尔类型：bool
		默认值为：false，不允许将整型转换成布尔型（如：0转成false）

		字符串：string
		len(str) 求长度
		+ 或 fmt.Sprintf 拼接字符串
		strings.Split 分割
		strings.contains 包含
		strings.HasPrefix, strings.HasSuffix 前缀/后缀判断
		strings.index(),strings.Lastindex() 子串位置
		strings.Jion(a[]string, sep string) join操作

		字符：byte和rune
		组成每个字符串的元素，遍历字符串得到字符，用单引号定义''，分为：
		byte类型：代表一个ASCII码的一个字符
		rune类型：代表一个UTF-8字符（处理中文、日文或者复合字符串）
	*/
	fmt.Println("\n----------------------------------------------------  int demo ---------------------------------------------------")
	// int8的范围-128~127（2的-7次方~2的7次方-1），占用1个字节
	var int8num int8 = 10
	int8num = 20
	fmt.Printf("int8num = %v %T  %v个字节 \n", int8num, int8num, unsafe.Sizeof(int8num))
	// int16的范围-32768到32767（2的-15次方~2的15次方-1），占用2个字节
	var int16num int16 = 500
	fmt.Printf("int16num = %v %T %v个字节 \n", int16num, int16num, unsafe.Sizeof(int16num))
	// int16的范围-32768到32767（2的-31次方~2的31次方-1），占用4个字节
	var int32num int32 = 32769
	fmt.Printf("int32num = %v %T %v个字节 \n", int32num, int32num, unsafe.Sizeof(int32num))
	// int16的范围（2的-63次方~2的63次方-1），占用8个字节
	var int64num int64 = 955555555555555555
	fmt.Printf("int64num = %v %T %v个字节 \n", int64num, int64num, unsafe.Sizeof(int64num))

	// uint8的范围0~255（0~2的8次方-1），占用1个字节
	var uint8num uint8 = 254
	fmt.Printf("uint8num = %v %T  %v个字节 \n", uint8num, uint8num, unsafe.Sizeof(uint8num))
	// uint16的范围0~65535（0~2的16次方-1），占用2个字节
	var uint16num uint16 = 65535
	fmt.Printf("uint16num = %v %T  %v个字节 \n", uint16num, uint16num, unsafe.Sizeof(uint16num))
	// uint32的范围0~4294967295（0~2的32次方-1），占用4个字节
	var uint32num uint32 = 4294967295
	fmt.Printf("uint32num = %v %T  %v个字节 \n", uint32num, uint32num, unsafe.Sizeof(uint32num))
	// uint64的范围0~-（0~2的64次方-1），占用8个字节
	var uint64num uint64 = 1855555555555555554
	fmt.Printf("uint64num = %v %T  %v个字节 \n", uint64num, uint64num, unsafe.Sizeof(uint64num))

	// 类型转换
	var a1 int8 = 10
	var a2 int16 = 200
	// 默认的类型和操作系统有关系
	a3 := 20
	// a1+a2会报错，需要进行类型转换
	fmt.Printf("a1+a2 = %v \n, a3类型%T \n", int16(a1)+a2, a3)

	fmt.Println("\n----------------------------------------------------  float demo ---------------------------------------------------")
	var f1 float32 = 3.12
	fmt.Printf("f1 = %v %f, %T, %v个字节\n", f1, f1, f1, unsafe.Sizeof(f1))
	var f2 float64 = 154.101625
	fmt.Printf("f2 = %v %.3f, %T, %v个字节\n", f2, f2, f2, unsafe.Sizeof(f2))
	// 默认folat64
	f3 := 11.11
	fmt.Printf("f3 = %v, %T\n", f3, f3)

	// Golang中float精度丢失问题
	f4 := 1129.6
	fmt.Println("f4 = ", f4*100)
	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2)

	//类型转换
	a := 10
	c0 := float64(a)
	var c1 float32 = 23.4
	var c2 = float64(c1)
	var c3 = int8(c1)
	fmt.Printf("c0 = %v %T, c1 =%v %T, c2 = %v %T, c3 =%v %T \n", c0, c0, c1, c1, c2, c2, c3, c3)

	fmt.Println("\n----------------------------------------------------  boolen demo ---------------------------------------------------")
	var flag = true
	var flag0 bool
	fmt.Printf("flag is %v %T, flag0 is %v %T \n", flag, flag, flag0, flag0)
	if flag0 {
		fmt.Println("flag0 is true")
	} else {
		fmt.Println("flag0 is false")
	}

	fmt.Println("\n----------------------------------------------------  string demo ---------------------------------------------------")
	// 定义字符串
	var s0 string = "11-22-33"
	var s1 = "s111"
	s2 := "s222"
	fmt.Printf("s0 = %v, s1 = %v, s2 = %v", s0, s1, s2)

	// 转义字符:打印路径：D:\project\damon
	path := "D:\\project\\damon"
	fmt.Println("path: ", path)

	// 多行字符串
	str1 := `this is str
		this is str11
		this is str
		this is str`
	fmt.Println(str1)

	/*
		len(str) 求长度
		+ 或 fmt.Sprintf 拼接字符串
		strings.Split 分割
		strings.contains 包含
		strings.HasPrefix, strings.HasSuffix 前缀/后缀判断
		strings.index(),strings.Lastindex() 子串位置
		strings.Jion(a[]string, sep string) join操作
	*/
	l := len(s0)
	bl := strings.Contains(s0, "22")
	isSuffix := strings.HasSuffix(s0, "33")
	isPrefix := strings.HasPrefix(s0, "11")
	arr := strings.Split(s0, "-")
	s0j := strings.Join(arr, "*")
	fmt.Printf("s0 is %v, len is %v, is contains 22 %v, is suffix 33 %v, is prefix 11 %v , arr: %v, s0j: %v \n", s0, l, bl, isSuffix, isPrefix, arr, s0j)

	arr1 := []string{"php", "java", "golang"}
	languagestr := strings.Join(arr1, "-")
	startIdex := strings.Index(languagestr, "-")
	lastIndex := strings.LastIndex(languagestr, "-")
	fmt.Printf("languagestr is %v, start idex %v, last index %v \n", languagestr, startIdex, lastIndex)

	fmt.Println("\n----------------------------------------------------  byte demo ---------------------------------------------------")
	var bt = 'a'
	fmt.Printf("byte bt = %c %v %T \n", bt, bt, bt)
	str2 := "你好 golang"
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%v(%c) \n", str2[i], str2[i])
	// }
	// 循环字符串
	for _, avr := range str2 {
		fmt.Printf("%v(%c)", avr, avr)
	}
	fmt.Println()
	// 修改字符串
	bytestr := []byte(s0)
	bytestr[0] = '0'
	bytestr[1] = '0'
	fmt.Println(string(bytestr))
	runeStr := []rune(str2)
	runeStr[0] = '我'
	runeStr[1] = '爱'
	fmt.Println(string(runeStr))

	fmt.Println("\n----------------------------------------------------  transform data type ---------------------------------------------------")
	/*
		其他类型转换成String类型, 两种：
		fmt.Sprintf:转换格式 int为%d，float为%f，bool为%t byte为%c
		strconv包:
	*/
	var ti int = 100
	var tf float64 = 23.45
	var tb bool = true
	strTi := fmt.Sprintf("%d", ti)
	strTf := fmt.Sprintf("%f", tf)
	strTb := fmt.Sprintf("%t", tb)
	fmt.Printf("strTi=%v %T, strTf=%v %T, strTb=%v %T \n", strTi, strTi, strTf, strTf, strTb, strTb)
	strTi1 := strconv.FormatInt(int64(ti), 10)
	strTf1 := strconv.FormatFloat(tf, 'f', 2, 64)
	strTb1 := strconv.FormatBool(tb)
	fmt.Printf("strTi1=%v %T, strTf1=%v %T, strTb1=%v %T \n", strTi1, strTi1, strTf1, strTf1, strTb1, strTb1)

	/*
		String类型转换成其他类型，strconv包
	*/
	sn1 := "12"
	sn2 := "12.34"
	sn3 := "true"
	intn, _ := strconv.Atoi(sn1)
	floatn, _ := strconv.ParseFloat(sn2, 64)
	booln, _ := strconv.ParseBool(sn3)
	fmt.Printf("intn=%v %T, floatn=%v %T, booln=%v %T \n", intn, intn, floatn, floatn, booln, booln)
}

const aa = iota

func printfDemo() {
	fmt.Println("\n------------------------------------------------------  printfDemo ------------------------------------------------------")
	var a int
	a = 1
	//  类型推导方式定义变量
	b := 2
	var c = "damon"
	fmt.Println("a: ", a, "b: ", b)
	fmt.Printf("a=%v a的类型是%T, c=%v c的类型是%T \n", a, a, c, c)
}

func defDemo() {
	fmt.Println("\n------------------------------------------------------  defDemo ------------------------------------------------------")
	// 一次申明多个变量
	var (
		username = "damon"
		sex      = 1
		age      = 18
	)
	fmt.Println("username: ", username, "sex: ", sex, "age: ", age)

	// 端变量一次性申明多个变量并赋值
	d, e, f := 1, 2.1, "ff"
	fmt.Printf("d类型：%T, e类型：%T, f类型：%T \n", d, e, f)

	// 申明常量
	const pi = 3.14159
	const (
		p0 = 1
		p1 = 1.1
		p2 = "sss"
		p3
		p4
	)
	username = "damon123"
	fmt.Printf("pi = %v,p0 = %v,p1 = %v,p2 = %v,p3 = %v,p4 = %v, username = %v \n", pi, p0, p1, p2, p3, p4, username)
}

func iotaDemo() {
	fmt.Println("\n------------------------------------------------------  iotaDemo ------------------------------------------------------")
	// 常量计数器：iota在关键字出现时将被重置未0，const中每新增一行常量申明将使iota计数一次
	const (
		b = iota
		b0
		b1
		b2
		b3 = 100
		b4
	)
	fmt.Printf("aa=%v,b=%v, b0=%v,b1=%v, b2=%v, b3=%v,b4=%v \n", aa, b, b0, b1, b2, b3, b4)
}
