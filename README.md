# Golang语言学习

## 一、项目结构

```
go_study/
├── jiegouti/          # 结构体相关示例
│   ├── struct_demo.go    # 结构体定义、初始化、方法、JSON转换
│   ├── extends.go        # 结构体继承示例（匿名嵌入）
│   └── extends_1.go      # 结构体组合示例（具名嵌入）
├── jiekou/           # 接口相关示例
│   ├── interface_demo.go  # 接口定义和使用
│   └── duotai.go         # 多态示例
├── yufa/             # 语法基础示例
│   ├── data_base.go      # 基础数据类型
│   ├── data_multi.go     # 数组、切片、Map
│   ├── error_handler.go  # 错误处理
│   ├── flow_control.go   # 流程控制
│   ├── function.go       # 函数基础
│   ├── function_type.go  # 函数类型、匿名函数、闭包
│   ├── package_demo.go   # 包管理
│   ├── point_demo.go     # 指针
│   ├── slice_demo.go     # 切片详细示例
│   └── time_demo.go      # 时间处理
├── main.go           # 主程序入口
├── go.mod            # Go模块文件
└── README.md         # 本文档
```

## 二、语法基础

### 1. 变量声明和常量

#### 变量声明

Go语言支持多种变量声明方式：

**语法格式：**
```go
// 方式1：标准声明
var 变量名 类型 = 值

// 方式2：类型推导
var 变量名 = 值

// 方式3：短变量声明（只能在函数内使用）
变量名 := 值

// 方式4：批量声明
var (
	变量1 = 值1
	变量2 = 值2
)

// 方式5：多变量同时声明
变量1, 变量2 := 值1, 值2
```

**代码示例：**
```go
// 基本变量声明
var a int
a = 1
b := 2              // 类型推导
var c = "damon"

// 批量声明变量
var (
	username = "damon"
	sex      = 1
	age      = 18
)

// 多变量同时声明
d, e, f := 1, 2.1, "ff"
fmt.Printf("d类型：%T, e类型：%T, f类型：%T \n", d, e, f)
// 输出：d类型：int, e类型：float64, f类型：string
```

#### 常量声明

常量在声明时必须赋值，且后续不能修改。

**语法格式：**
```go
// 单个常量
const 常量名 = 值

// 批量声明
const (
	常量1 = 值1
	常量2 = 值2
)
```

**代码示例：**
```go
// 单个常量
const pi = 3.14159

// 批量声明常量
const (
	p0 = 1
	p1 = 1.1
	p2 = "sss"
	p3        // 如果省略值，则与上一行相同
	p4        // p3和p4的值都是"sss"
)
```

#### iota常量计数器

`iota`是Go语言的常量计数器，只能在常量表达式中使用。

**特点：**
- `iota`在`const`关键字出现时将被重置为0
- `const`中每新增一行常量声明将使`iota`计数一次
- 可以使用`iota`实现枚举类型

**代码示例：**
```go
const (
	b = iota  // 0
	b0        // 1
	b1        // 2
	b2        // 3
	b3 = 100  // 100（中断iota计数）
	b4        // 100（与上一行相同）
)
fmt.Printf("b=%v, b0=%v, b1=%v, b2=%v, b3=%v, b4=%v \n", b, b0, b1, b2, b3, b4)
// 输出：b=0, b0=1, b1=2, b2=3, b3=100, b4=100
```

### 2. 数据类型

#### 整型 (int)

Go语言提供了多种整数类型，按长度和有无符号进行分类。

**有符号整型：**
- `int8`：范围 -128~127，占用1个字节
- `int16`：范围 -32768~32767，占用2个字节
- `int32`：范围 -2³¹~2³¹-1，占用4个字节
- `int64`：范围 -2⁶³~2⁶³-1，占用8个字节
- `int`：根据系统平台，32位或64位

**无符号整型：**
- `uint8`：范围 0~255，占用1个字节
- `uint16`：范围 0~65535，占用2个字节
- `uint32`：范围 0~4294967295，占用4个字节
- `uint64`：范围 0~2⁶⁴-1，占用8个字节
- `uint`：根据系统平台，32位或64位

**代码示例：**
```go
var int8num int8 = 10
fmt.Printf("int8num = %v %T  %v个字节 \n", int8num, int8num, unsafe.Sizeof(int8num))
// 输出：int8num = 10 int8  1个字节

var int64num int64 = 955555555555555555
fmt.Printf("int64num = %v %T %v个字节 \n", int64num, int64num, unsafe.Sizeof(int64num))
// 输出：int64num = 955555555555555555 int64 8个字节

var uint8num uint8 = 254
fmt.Printf("uint8num = %v %T  %v个字节 \n", uint8num, uint8num, unsafe.Sizeof(uint8num))
// 输出：uint8num = 254 uint8  1个字节
```

#### 浮点型 (float)

Go语言提供两种浮点数类型：

**类型说明：**
- `float32`：32位浮点数，占用4个字节，精度约6位小数
- `float64`：64位浮点数，占用8个字节，精度约15位小数
- 默认类型为`float64`

**代码示例：**
```go
var f1 float32 = 3.12
fmt.Printf("f1 = %v %f, %T, %v个字节\n", f1, f1, f1, unsafe.Sizeof(f1))
// 输出：f1 = 3.12 3.120000, float32, 4个字节

var f2 float64 = 154.101625
fmt.Printf("f2 = %v %.3f, %T, %v个字节\n", f2, f2, f2, unsafe.Sizeof(f2))
// 输出：f2 = 154.101625 154.102, float64, 8个字节

f3 := 11.11  // 默认float64
fmt.Printf("f3 = %v, %T\n", f3, f3)
// 输出：f3 = 11.11, float64
```

#### 布尔类型 (bool)

Go语言中的布尔类型只有两个值：`true`和`false`。

**特点：**
- 默认值为`false`
- 布尔类型不能与其他类型相互转换
- 不允许将整型转换成布尔型（如：0不能转成false）

**代码示例：**
```go
var flag = true
var flag0 bool
fmt.Printf("flag is %v %T, flag0 is %v %T \n", flag, flag, flag0, flag0)
// 输出：flag is true bool, flag0 is false bool

if flag0 {
	fmt.Println("flag0 is true")
} else {
	fmt.Println("flag0 is false")
}
// 输出：flag0 is false
```

#### 字符串 (string)

Go语言中的字符串是不可变的字节序列，使用UTF-8编码。

**常用操作：**
- `len(str)`：获取字符串长度
- `+` 或 `fmt.Sprintf`：拼接字符串
- `strings.Split`：分割字符串
- `strings.Contains`：判断是否包含子串
- `strings.HasPrefix`、`strings.HasSuffix`：判断前缀/后缀
- `strings.Index()`、`strings.LastIndex()`：查找子串位置
- `strings.Join(a[]string, sep string)`：连接字符串切片
- **注意**：修改字符串需要先转换为`[]byte`或`[]rune`，修改后再转回`string`

**代码示例：**
```go
var s0 string = "11-22-33"
s1 := "s222"

// 字符串操作
l := len(s0)                           // 长度：8
bl := strings.Contains(s0, "22")       // 包含：true
isSuffix := strings.HasSuffix(s0, "33") // 后缀：true
isPrefix := strings.HasPrefix(s0, "11") // 前缀：true
arr := strings.Split(s0, "-")          // 分割：["11", "22", "33"]
s0j := strings.Join(arr, "*")          // 拼接："11*22*33"

// 多行字符串
str1 := `this is str
	this is str11
	this is str`

// 修改字符串
bytestr := []byte(s0)
bytestr[0] = '0'
bytestr[1] = '0'
fmt.Println(string(bytestr))  // 输出：00-22-33

runeStr := []rune("你好 golang")
runeStr[0] = '我'
runeStr[1] = '爱'
fmt.Println(string(runeStr))  // 输出：我爱 golang
```

#### 字符类型 (byte和rune)

Go语言中字符用单引号定义，有两种类型：

**类型说明：**
- `byte`：`uint8`的别名，用于表示ASCII码字符
- `rune`：`int32`的别名，用于表示UTF-8字符（可以处理中文、日文等）
- 遍历字符串时，使用`range`得到的是`rune`类型

**代码示例：**
```go
var bt = 'a'
fmt.Printf("byte bt = %c %v %T \n", bt, bt, bt)
// 输出：byte bt = a 97 int32

str2 := "你好 golang"
// 遍历字符串
for _, avr := range str2 {
	fmt.Printf("%v(%c)", avr, avr)
}
// 输出：20320(你)22909(好)32( )103(g)111(o)108(l)97(a)110(n)103(g)
```

### 3. 类型转换

Go语言是强类型语言，不同类型之间不能直接运算，需要显式转换。

#### 其他类型转String

**方法1：使用fmt.Sprintf**
```go
fmt.Sprintf("%d", int)    // int转string
fmt.Sprintf("%f", float)  // float转string
fmt.Sprintf("%t", bool)   // bool转string
```

**方法2：使用strconv包**
```go
strconv.FormatInt(int64, 10)                    // int转string
strconv.FormatFloat(float, 'f', 2, 64)          // float转string（保留2位小数）
strconv.FormatBool(bool)                        // bool转string
```

**代码示例：**
```go
var ti int = 100
var tf float64 = 23.45
var tb bool = true

// fmt.Sprintf方式
strTi := fmt.Sprintf("%d", ti)
strTf := fmt.Sprintf("%f", tf)
strTb := fmt.Sprintf("%t", tb)
fmt.Printf("strTi=%v %T, strTf=%v %T, strTb=%v %T \n", strTi, strTi, strTf, strTf, strTb, strTb)
// 输出：strTi=100 string, strTf=23.450000 string, strTb=true string

// strconv包方式
strTi1 := strconv.FormatInt(int64(ti), 10)
strTf1 := strconv.FormatFloat(tf, 'f', 2, 64)
strTb1 := strconv.FormatBool(tb)
fmt.Printf("strTi1=%v %T, strTf1=%v %T, strTb1=%v %T \n", strTi1, strTi1, strTf1, strTf1, strTb1, strTb1)
// 输出：strTi1=100 string, strTf1=23.45 string, strTb1=true string
```

#### String转其他类型

使用`strconv`包进行转换：

```go
strconv.Atoi(string)              // string转int
strconv.ParseFloat(string, 64)    // string转float64
strconv.ParseBool(string)         // string转bool
```

**代码示例：**
```go
sn1 := "12"
sn2 := "12.34"
sn3 := "true"

intn, _ := strconv.Atoi(sn1)
floatn, _ := strconv.ParseFloat(sn2, 64)
booln, _ := strconv.ParseBool(sn3)

fmt.Printf("intn=%v %T, floatn=%v %T, booln=%v %T \n", intn, intn, floatn, floatn, booln, booln)
// 输出：intn=12 int, floatn=12.34 float64, booln=true bool
```

#### 数值类型转换

```go
// 不同整型之间转换
int16(int8值)

// 整型转浮点型
float64(int值)

// 浮点型转整型（会截断小数部分）
int(float值)
```

**代码示例：**
```go
var a1 int8 = 10
var a2 int16 = 200
// a1+a2会报错，需要进行类型转换
fmt.Printf("a1+a2 = %v \n", int16(a1)+a2)
// 输出：a1+a2 = 210

a := 10
c0 := float64(a)        // int转float64
var c1 float32 = 23.4
var c2 = float64(c1)    // float32转float64
var c3 = int8(c1)       // float32转int8（会截断）
fmt.Printf("c0 = %v %T, c1 = %v %T, c2 = %v %T, c3 = %v %T \n", c0, c0, c1, c1, c2, c2, c3, c3)
// 输出：c0 = 10 float64, c1 = 23.4 float32, c2 = 23.4 float64, c3 = 23 int8
```

### 4. 值类型与引用类型

Go语言中的数据类型分为值类型和引用类型，它们的赋值行为不同。

#### 值类型

**包含类型：**
- 基本数据类型：`int`、`float`、`bool`、`string`
- 数组
- 结构体

**特点：**
赋值时，是将值复制一份给新的变量，两个变量互不影响。

#### 引用类型

**包含类型：**
- 指针（pointer）
- 切片（slice）
- 映射（map）
- 通道（channel）
- 接口（interface）
- 函数（function）

**特点：**
赋值时，是将引用地址复制一份给新的变量，两个变量指向同一块内存，修改其中一个变量会影响另一个变量。

**代码示例：**
```go
// 值类型示例
a := 10
b := a
b = 20
fmt.Printf("a=%d , b=%d \n", a, b)
// 输出：a=10 , b=20（互不影响）

// 数组值类型
var array1 = [3]int{1, 2, 3}
array2 := array1
array2[0] = 10
fmt.Printf("array1=%v , array2=%v \n", array1, array2)
// 输出：array1=[1 2 3] , array2=[10 2 3]（互不影响）

// 切片引用类型
var array3 = []int{1, 2, 3}
slice1 := array3
slice1[0] = 10
fmt.Printf("array3=%v , slice1=%v \n", array3, slice1)
// 输出：array3=[10 2 3] , slice1=[10 2 3]（互相影响）
```

### 5. 数组

数组是具有相同类型且长度固定的数据集合。

#### 数组定义

**语法格式：**
```go
// 方式1：声明指定长度的数组
var 数组名 [长度]类型

// 方式2：声明并初始化
var 数组名 = [长度]类型{值1, 值2, ...}

// 方式3：自动推断长度
数组名 := [...]类型{值1, 值2, ...}

// 方式4：指定索引初始化
数组名 := [...]类型{索引1: 值1, 索引2: 值2, ...}
```

**数组特点：**
- 长度固定，不能动态扩展
- 数组是值类型，赋值会复制整个数组
- 长度是数组类型的一部分，`[5]int`和`[10]int`是不同类型

**代码示例：**
```go
// 指定长度
array2 := [10]int{1, 2, 3, 4}
fmt.Println("array2: ", array2)
// 输出：array2: [1 2 3 4 0 0 0 0 0 0]

// 自动推断长度
array3 := [...]string{"java", "c++", "python", "php"}
array3[0] = "golang"
fmt.Println("array3: ", array3)
// 输出：array3: [golang c++ python php]

// 指定索引初始化
array4 := [...]string{1: "java", 3: "c++", 0: "python", 2: "php"}
fmt.Println("array4: ", array4)
// 输出：array4: [python java php c++]

// 遍历数组
for k, v := range array4 {
	fmt.Printf("array4[%d]=%s \n", k, v)
}

// 多维数组
var array6 = [...][2]string{
	{"java", "c++"},
	{"python", "php"},
	{"golang", "ruby"},
}
fmt.Println(array6[0][1])  // 输出：c++
```

#### 多维数组

```go
var 数组名 [...][长度]类型 = [...][长度]类型{
	{值1, 值2, ...},
	{值3, 值4, ...},
}
```

---

### 6. 切片 (Slice)

切片是对数组的一个抽象，提供更灵活、更强大的序列操作接口。

#### 切片定义

**语法格式：**
```go
// 方式1：声明nil切片
var slice []type

// 方式2：字面量初始化
slice := []type{值1, 值2, ...}

// 方式3：使用make创建（推荐）
slice := make([]type, length, capacity)
```

**切片特点：**
- 切片是对数组的引用，是引用类型
- 长度可变，可以动态扩展
- 切片的底层是数组，切片本身不存储数据，只是对底层数组的引用
- 切片包含长度（len）和容量（cap）两个属性

**代码示例：**
```go
// 切片的定义
var slice []int
fmt.Println("slice: ", slice, slice == nil)
// 输出：slice: [] true

slice = []int{1, 2, 3, 4, 5}
fmt.Println("slice: ", slice)
// 输出：slice: [1 2 3 4 5]

// 使用make创建
slice3 := make([]int, 5, 10)
fmt.Println("slice3: ", slice3, slice3 == nil)
// 输出：slice3: [0 0 0 0 0] false

// 数组转切片
array := [5]int{10, 20, 30, 40, 50}
slice4 := array[1:4]
fmt.Println("slice4: ", slice4)
fmt.Println("slice4 length: ", len(slice4))
fmt.Println("slice4 capacity: ", cap(slice4))
// 输出：slice4: [20 30 40]
//      slice4 length: 3
//      slice4 capacity: 4
```

#### 长度和容量

- **长度（len）**：切片中元素的个数，使用`len()`函数获取
- **容量（cap）**：切片的底层数组的容量，使用`cap()`函数获取
- 容量 >= 长度

#### 切片操作

**常用操作：**
```go
// 切片操作：[low:high] 包含从low到high-1的元素
slice[low:high]

// 追加元素
append(slice, 元素1, 元素2, ...)

// 合并切片
append(slice1, slice2...)

// 复制切片
copy(dest, src)

// 删除元素（通过重新切片实现）
append(slice[:index], slice[index+1:]...)
```

**代码示例：**
```go
// 动态扩容
var slice5 []int
slice5 = append(slice5, 12)
slice5 = append(slice5, 22, 32, 42)
fmt.Printf("slice5: %v , length: %d , capacity: %d \n", slice5, len(slice5), cap(slice5))
// 输出：slice5: [12 22 32 42] , length: 4 , capacity: 4

// 合并切片
slice6 := []int{1, 2, 3}
slice7 := []int{4, 5, 6}
slice6 = append(slice6, slice7...)
fmt.Printf("slice6: %v , length: %d \n", slice6, len(slice6))
// 输出：slice6: [1 2 3 4 5 6] , length: 6

// 复制切片
slice8 := make([]int, 1, 3)
slice8 = append(slice8, 100, 200)
copy(slice8, slice6)
fmt.Printf("slice8: %v \n", slice8)
// 输出：slice8: [1 2 3]

// 删除元素
slice9 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
index := 4
slice9 = append(slice9[:index], slice9[index+1:]...)
fmt.Printf("slice9: %v \n", slice9)
// 输出：slice9: [1 2 3 4 6 7 8 9]
```

#### 动态扩容

当切片的长度超过容量时，Go会自动分配更大的底层数组（通常是原来的2倍），并将原数据复制到新数组。

---

### 7. Map

Map是Go语言内置的关联数据类型，存储键值对的无序集合。

#### Map定义

**语法格式：**
```go
// 方式1：声明nil map（不能直接使用，需要初始化）
var map1 map[keyType]valueType

// 方式2：使用make创建（推荐）
map1 := make(map[keyType]valueType, 容量)  // 容量可选

// 方式3：字面量初始化
map1 := map[keyType]valueType{
	key1: value1,
	key2: value2,
}
```

**Map特点：**
- Map是引用类型
- 键值对是无序的（遍历顺序随机）
- 键必须是可比较类型（不能是slice、map、function）
- 零值为`nil`，不能直接添加元素

**代码示例：**
```go
// 声明map（nil map）
var map1 map[string]string
if map1 == nil {
	fmt.Println("map1 is null !")
}

// 使用make创建map
map1 = make(map[string]string, 10)
map1["key1"] = "java"
map1["key2"] = "c++"
map1["key3"] = "python"

// 字面量创建map
map3 := map[string]string{
	"one":   "php",
	"two":   "c++",
	"three": "python",
}
```

#### Map操作

**基本操作：**
```go
// 添加/修改元素
map1[key] = value

// 获取元素（ok表示key是否存在）
value, ok := map1[key]

// 删除元素
delete(map1, key)

// 遍历map
for k, v := range map1 {
	// ...
}
```

**代码示例：**
```go
// 删除元素
delete(map3, "two")

// 获取元素（检查key是否存在）
v, ok := map3["one"]
if ok {
	fmt.Println("map3 key one value :", v)
}

// 遍历map
for k, v := range map3 {
	fmt.Printf("map3[%s]=%s \n", k, v)
}
```

#### Map排序

Go的map是无序的，如果需要按顺序遍历，需要将key提取到slice中排序后遍历。

**代码示例：**
```go
map4 := map[int]int{
	1: 10,
	5: 50,
	2: 20,
	6: 60,
}

// 提取key到slice并排序
var keySlice []int
for k := range map4 {
	keySlice = append(keySlice, k)
}
sort.Ints(keySlice)

// 按排序后的key遍历
for _, v := range keySlice {
	fmt.Printf("key = %v , value = %v \n", v, map4[v])
}
```

### 8. 指针

指针是存储另一个变量内存地址的变量，通过指针可以间接访问和修改变量的值。

#### 指针定义

**语法格式：**
```go
var p *type = &变量
p := &变量
```

**指针特点：**
- 指针存储的是变量的内存地址
- 通过`*p`可以访问和修改指针指向的变量的值
- 使用`&`获取变量的地址，使用`*`解引用指针
```go
var a = 10
var p = &a
fmt.Printf("a = %v type: %T 内存地址: %p\n", a, a, &a)
fmt.Printf("p = %v type: %T 内存地址对应的值：%v \n", p, p, *p)
// 输出：a = 10 type: int 内存地址: 0xc0000140a8
//      p = 0xc0000140a8 type: *int 内存地址对应的值：10

*p = 30  // 通过*p改变对应内存地址的值
fmt.Printf("a = %v \n", a)
// 输出：a = 30
```

#### 指针使用

```go
&变量    // 获取变量的地址
*指针    // 解引用，获取指针指向的值
*指针 = 新值  // 通过指针修改值
```

#### 指针作为函数参数

- **值传递**：传递变量的副本，不会修改原变量
- **指针传递**：传递变量的地址，可以修改原变量

**代码示例：**
```go
x := 5
fn2(x)  // 值传递，不会修改x
fmt.Println("x = ", x)  // 输出：x = 5

fn3(&x)  // 指针传递，会修改x
fmt.Println("x = ", x)  // 输出：x = 20

func fn2(x int) {
	x = 10  // 不会影响外部变量
}

func fn3(x *int) {
	*x = 20  // 修改外部变量
}
```

### 9. 流程控制

Go语言提供了if、for、switch、goto等流程控制语句。

#### if条件判断

**语法格式：**
```go
if 条件表达式 {
	执行语句
} else if 条件表达式 {
	执行语句
} else {
	执行语句
}
// if可以包含初始化语句
if 初始化语句; 条件表达式 {
	执行语句
}
```

**代码示例：**
```go
flage := true
if flage {
	fmt.Println("flage is true")
} else {
	fmt.Println("flage is false")
}

// if包含初始化语句
if age := 20; age >= 18 {
	fmt.Println("adult")
} else {
	fmt.Println("teenager")
}
```

#### for循环

**语法格式：**
```go
// 标准for循环
for 初始化语句; 条件表达式; 后置语句 {
	执行语句
}
// 类似while循环
for 条件表达式 {
	执行语句
}
// 无限循环
for {
	执行语句
}
// range遍历
for 索引, 值 := range 数组/切片/字符串/map {
	执行语句
}
```

**代码示例：**
```go
// 标准for循环
var sum = 0
for i := 1; i <= 100; i++ {
	sum += i
}
fmt.Println("1 + 2 + ... + 100 = ", sum)
// 输出：1 + 2 + ... + 100 = 5050

// range遍历字符串
str := "你好golang"
for k, v := range str {
	fmt.Printf("k = %d v = %c \n", k, v)
}

// range遍历切片
arr := []string{"hello", "world", "golang"}
for _, v := range arr {
	fmt.Printf(" v = %s \n", v)
}
```

#### switch条件判断

**语法格式：**
```go
switch 表达式 {
case 值1, 值2:  // 可以多个值
	执行语句
case 值3:
	执行语句
default:
	执行语句
}
// switch可以包含初始化语句
switch 初始化语句; 表达式 {
	...
}
```

**代码示例：**
```go
switch sw := "jpg"; sw {
case "jpg":
	fmt.Println("this is a jpg file")
case "png":
	fmt.Println("this is a png file")
default:
	fmt.Println("unknown file type")
}

// 多个值匹配
switch score := "A"; score {
case "A", "B", "C":
	fmt.Println("pass")
case "D", "E":
	fmt.Println("fail")
}
```

#### break和continue

- `break`：终止当前循环
- `continue`：跳过当前循环，继续下一次循环

#### goto语句

**语法格式：**
```go
goto 标签名
...
标签名:
	执行语句
```

**代码示例：**
```go
age := 30
if age > 24 {
	fmt.Println("成年人！")
	goto lable
}
fmt.Println("AAA")
fmt.Println("BBB")
lable:
fmt.Println("CCC")
// 输出：成年人！
//      CCC
```

### 10. 函数

函数是Go语言的基本构建块，支持多种定义和使用方式。

#### 函数定义

**语法格式：**
```go
1. 固定参数
	func 函数名(参数1 类型1, 参数2 类型2) 返回值类型 {
		...
	}
2. 可变参数
	func 函数名(参数 ...类型) 返回值类型 {
		...
	}
	// 可变参数必须放在最后
	func 函数名(固定参数 类型, 可变参数 ...类型) 返回值类型 {
		...
	}
3. 多返回值
	func 函数名(参数) (返回值1类型, 返回值2类型) {
		return 值1, 值2
	}
	// 命名返回值
	func 函数名(参数) (返回值1 类型1, 返回值2 类型2) {
		返回值1 = 值1
		返回值2 = 值2
		return  // 可以省略返回值
	}
```

**代码示例：**
```go
// 固定参数
func sumFx(x int, y int) int {
	sum := x + y
	return sum
}
sum1 := sumFx(12, 3)
fmt.Println("sum1: ", sum1)  // 输出：sum1: 15

// 可变参数
func sumFx2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
sum2 := sumFx2(1, 2, 3, 4)
fmt.Println("sum2: ", sum2)  // 输出：sum2: 10

// 多返回值
func fub2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}
sum, sub := fub2(12, 3)
fmt.Printf("sum = %v , sub = %v \n", sum, sub)
// 输出：sum = 15 , sub = 9
```

#### 函数类型

Go语言中函数是一等公民，可以作为类型使用。

**定义函数类型：**
```go
type 函数类型名 func(参数列表) 返回值列表

// 示例
type calc func(int, int) int
```

#### 函数作为参数
```go
func 函数名(参数 函数类型) {
	函数类型(参数)
}
```

**代码示例：**
```go
// 自定义函数类型
type calcType func(int, int) int

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc(x, y int, ct calcType) int {
	return ct(x, y)
}

var a = 10
var b = 4
s1 := calc(a, b, sum)  // 传入sum函数
s2 := calc(a, b, sub)  // 传入sub函数
fmt.Printf("a+b=%v, a-b=%v \n", s1, s2)
// 输出：a+b=14, a-b=6

// 匿名函数作为参数
s3 := calc(a, b, func(x, y int) int {
	return x * y
})
fmt.Printf("a*b=%v \n", s3)
// 输出：a*b=40
```

#### 函数作为返回值
```go
func 函数名() 函数类型 {
	return func(参数) 返回值 {
		...
	}
}
```

**代码示例：**
```go
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

str := "+"
var d calcType = do(str)
s4 := d(a, b)
fmt.Printf("a%vb=%v \n", str, s4)
// 输出：a+b=14
```

#### 匿名函数
```go
func(参数)(返回值){
	...
}
// 匿名函数可以立即执行
func(参数)(返回值){
	...
}(参数值)
```

**代码示例：**
```go
// 匿名自执行函数
func() {
	fmt.Println("匿名自执行函数...")
}()

// 匿名函数立即执行并返回结果
s5 := func(x, y int) int {
	return x + y
}(10, 20)
fmt.Println("匿名函数中结果：", s5)
// 输出：匿名函数中结果：30
```

#### 闭包

闭包是指有权访问另一个函数作用域中的变量的函数。

**特点：**
- 创建闭包的常见方式是在一个函数内部创建另外一个函数
- 可以让一个变量常驻内存
- 可以让一个变量不污染全局
- **注意**：由于闭包作用域返回的局部变量资源不会被立即销毁回收，可能占用更多内存，过度使用闭包会导致性能下降

**代码示例：**
```go
func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

fn := adder2()
s6 := fn(10)
s7 := fn(10)
s8 := fn(10)
fmt.Printf("闭包函数中结果：%v, %v, %v \n", s6, s7, s8)
// 输出：闭包函数中结果：20, 30, 40
// 说明：i变量常驻内存，每次调用都会累加
```

#### defer延迟执行

`defer`语句用于延迟函数的执行，直到包含它的函数返回。

**特点：**
- 执行时机：在`return`语句的返回值赋值操作后，RET指令执行前
- 多个`defer`语句按后进先出（LIFO）的顺序执行
- `defer`注册要延迟执行的函数时，该函数所有的参数都需要确定其值

#### panic和recover

- `panic`：可以在任何地方引发异常，程序会立即停止当前函数的执行
- `recover`：只有在`defer`调用的函数中有效，用于捕获`panic`

**代码示例：**
```go
func fn1(x, y int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("fn1 func error: ", err)
		}
	}()
	return x / y  // 如果y=0会panic
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("readFile读取文件失败！")
	}
}

func myFn(filename string) {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员提个醒，错误: ", e)
		}
	}()
	
	err := readFile(filename)
	if err != nil {
		panic(err)
	}
}
```

### 11. make和new的区别

`make`和`new`都是Go语言的内存分配函数，但用途和返回值不同。

#### make

**用途：**
- 用于`slice`、`map`、`channel`的内存创建和初始化
- 返回类型就是这三个类型本身

**语法：**
```go
make(t Type, size ...IntegerType) Type
```

**代码示例：**
```go
// map初始化
m1 := make(map[string]string, 2)
m1["k1"] = "v1"

// 切片初始化
s1 := make([]int, 4, 4)
s1[0] = 1
```

**new：**
```go
用于类型的内存分配
内存对应的值为类型零值
返回的是指向类型的指针
语法：new(Type) *Type
```

**代码示例：**
```go
// 指针初始化
var p *int = new(int)
*p = 100
fmt.Println(*p)  // 输出：100
```

**区别：**
```go
1. make只适用于slice、map、channel，new适用于所有类型
2. make返回引用类型本身，new返回指向类型的指针
3. make分配内存并初始化，new只分配内存（值为零值）
```

### 12. 结构体 (Struct)

结构体是一种聚合数据类型，将零个或多个任意类型的值聚合在一起。

#### 结构体定义

**语法格式：**
```go
type 类型名 struct {
	字段名1 字段类型1
	字段名2 字段类型2
	...
}
```

**代码示例：**
```go
type user struct {
	name   string
	age    int8
	sex    int8
	Hobby  []string
	Score  map[string]int8
	person person //结构体嵌套
}
```

**结构体特点：**
- 结构体实例是相互独立的，不会相互影响
- 结构体的字段类型可以是：基本数据类型、切片、map以及结构体
- 如果结构体的字段类型是指针、slice、map，零值都是`nil`，需要先`make`才能使用

#### 结构体初始化

**初始化方式：**
```go
// 方式1：声明后赋值
var u user
u.name = "damon"

// 方式2：使用new
var u = new(user)
u.name = "damon"

// 方式3：使用&取地址
var u = &user{
	name: "damon",
	age:  18,
}

// 方式4：直接初始化
var u = user{
	name: "damon",
	age:  18,
}
```

**代码示例：**
```go
// 方式1：声明后赋值
var u1 user
u1.name = "damon"
u1.age = 18
u1.Hobby = make([]string, 2, 5)
u1.Hobby[0] = "sleep"
u1.Score = make(map[string]int8)
u1.Score["english"] = 110

// 方式2：使用new
var u2 = new(user)
u2.name = "damon2"
fmt.Printf("u2: %#v  类型：%T \n", u2, u2)
// 输出：u2: &main.user{name:"damon2", age:0, sex:0...}  类型：*main.user

// 方式3：使用&取地址
var u3 = &user{
	name: "damon3",
	age:  20,
	sex:  0,
}

// 方式4：直接初始化
var u4 = user{
	name: "damon4",
	age:  21,
	sex:  0,
}
```

#### 结构体方法

Go语言可以为任意类型定义方法，但接收者类型必须是当前包内定义的类型。

**语法格式：**
```go
// 值接收者
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}

// 指针接收者
func (接收者变量 *接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}
```

**区别：**
- **值接收者**：方法内部对结构体的修改不会影响原结构体
- **指针接收者**：方法内部对结构体的修改会影响原结构体

**代码示例：**
```go
type person struct {
	Name string
	Age  int8
	Sex  int8
}

// 值接收者方法
func (p person) PrintfPersonInfo() {
	fmt.Printf("printf person info: name is %v  age is %v sex is %v\n", p.Name, p.Age, p.Sex)
}

// 指针接收者方法
func (p *person) SetPersonInfo(name string, age int8, sex int8) bool {
	p.Name = name
	p.Age = age
	p.Sex = sex
	return true
}

p1 := person{
	Name: "damon1",
	Age:  18,
}
p1.PrintfPersonInfo()  // 输出：printf person info: name is damon1  age is 18 sex is 0

bl := p1.SetPersonInfo("damon2", 19, 1)
if bl {
	p1.PrintfPersonInfo()  // 输出：printf person info: name is damon2  age is 19 sex is 1
}
```

#### 结构体嵌套

一个结构体可以包含另一个结构体作为字段，通过嵌套可以实现类似继承的效果。

#### 结构体标签 (Tag)

结构体标签用于为结构体字段添加元数据，常用于JSON序列化、数据库映射等。

**格式：**
```go
`标签名:"标签值"`
```

**代码示例：**
```go
type person struct {
	Name string `json:"id"`  // 结构体标签
	Age  int8   `json:"age"`
	Sex  int8
}
```

#### 结构体与JSON

Go语言的标准库`encoding/json`提供了结构体与JSON之间的转换。

**注意：**
- 结构体转JSON时，字段首字母必须大写才能序列化
- 使用`json`标签可以自定义JSON字段名

**转换方法：**
```go
// 结构体转JSON
jsonbyte, err := json.Marshal(结构体实例)
jsonstr := string(jsonbyte)

// JSON转结构体
var 结构体实例 结构体类型
err := json.Unmarshal([]byte(json字符串), &结构体实例)
```

**代码示例：**
```go
// 结构体转JSON
p1 := person{
	Name: "damon1",
	Age:  18,
}
jsonbyte, _ := json.Marshal(p1)
jsonstr := string(jsonbyte)
fmt.Println("jsonstr: ", jsonstr)
// 输出：jsonstr: {"Name":"damon1","Age":18,"Sex":0}

// JSON转结构体
str := `{"Name":"damon1","Age":18,"Sex":0}`
var p2 person
err := json.Unmarshal([]byte(str), &p2)
if err == nil {
	fmt.Printf("p2: %#v \n", p2)
}

// 复杂结构体转换
type class struct {
	Name    string
	Student []student
}

type student struct {
	Id     int8
	Gender string
	Name   string
}

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
```

### 13. 继承/组合

Go语言没有传统面向对象的继承，通过结构体嵌套实现类似继承的效果。

#### 匿名嵌入（类似继承）

通过匿名嵌入结构体，可以直接访问嵌入结构体的字段和方法。

**代码示例：**
```go
type animal struct {
	Name string
}

func (a animal) run() {
	fmt.Printf("%v is running! \n", a.Name)
}

type dog struct {
	animal  // 匿名嵌入
	age int8
}

func (d dog) rescue() {
	fmt.Printf("%v is rescuing! \n", d.Name)
}

d := dog{
	animal: animal{Name: "le qi"},
	age: 2,
}
d.run()      // 直接调用animal的方法
d.rescue()   // 调用dog的方法
```

#### 具名嵌入（组合）

通过具名嵌入结构体，需要通过字段名访问嵌入结构体的字段和方法。

**代码示例：**
```go
type Cat struct {
	Name string
}

func (this *Cat) Walk() {
	fmt.Println("Animal " + this.Name + " Walk()...")
}

type Tom struct {
	Cat Cat  // 具名嵌入
	Say bool
}

tom := Tom{
	Cat: Cat{Name: "Tom"},
	Say: true,
}
tom.Cat.Name   // 需要通过Cat字段访问
tom.Cat.Walk() // 需要通过Cat字段调用方法
```

#### 方法重写

子结构体可以定义同名方法来覆盖父结构体的方法。

### 14. 接口 (Interface)

接口是一种抽象类型，定义了一组方法的集合，实现了这些方法的类型就实现了该接口。

#### 接口定义

**语法格式：**
```go
type 接口名 interface {
	方法名1(参数列表1) 返回值列表1
	方法名2(参数列表2) 返回值列表2
	...
}
```

**接口特点：**
- 接口是一种数据类型，是抽象的类型
- 接口是一组method的集合
- Go语言中的接口不能包含任何变量
- 接口中所有的方法都没有方法体（只定义规范不实现）
- 接口定义了一个对象的行为规范

#### 接口实现

- 一个类型只要实现了接口中的所有方法，就认为实现了该接口
- Go语言中接口的实现是隐式的，不需要显式声明（Duck Typing）

**代码示例：**
```go
type usber interface {
	Start()
	Stop()
}

type phone struct {
	Name string
}

func (p phone) Start() {
	fmt.Printf("%v  is start \n", p.Name)
}

func (p phone) Stop() {
	fmt.Printf("%v  is stop \n", p.Name)
}

type cacmer struct {
	Name string
}

func (p cacmer) Start() {
	fmt.Printf("%v  is start \n", p.Name)
}

func (p cacmer) Stop() {
	fmt.Printf("%v is stop \n", p.Name)
}
```

#### 空接口

空接口（`interface{}`）是特殊接口，不定义任何方法，可以表示任意类型。

**用途：**
- 可以接收任意类型的值
- 常用于函数参数需要接收多种类型的情况

**代码示例：**
```go
var i interface{}
i = 10
i = "hello"
fmt.Printf("%v %T \n", i, i)
```

#### 多态

通过接口实现多态，不同的类型实现同一个接口，可以用接口类型作为参数，传入不同的实现类型。

**代码示例：**
```go
type computer struct {
	Name string
}

func (c computer) transData(u usber) {
	u.Start()
	fmt.Printf("%v computer is working \n", c.Name)
	u.Stop()
}

p := phone{Name: "iphone"}
c := cacmer{Name: "keda"}
computer := computer{Name: "huawei"}

computer.transData(p)  // 传入phone类型
computer.transData(c)  // 传入cacmer类型
// 输出：
// iphone  is start
// huawei computer is working
// iphone  is stop
// keda  is start
// huawei computer is working
// keda is stop
```

### 15. 包 (Package)

包是多个Go源码的集合，用于组织代码和实现代码复用。

#### 包的概念

包是多个Go源码的集合，Go提供了很多内置包，如：`fmt`、`strconv`、`strings`、`sort`、`errors`、`time`、`encoding/json`、`os`、`io`等。

#### 包的分类

**1. 系统内置包：**
- Golang语言提供的内置包，引入后可以直接使用
- 如：`fmt`、`strconv`、`strings`、`sort`、`errors`、`time`、`encoding/json`、`os`、`io`等

**2. 自定义包：**
- 开发者自己写的包
- 导入路径从`go.mod`的module名称开始
- 示例：`import "go_study/jiegouti"`

**3. 第三方包：**
- 属于自定义包的一种，需要下载安装到本地才可以使用
- 包仓库：https://pkg.go.dev/

#### 包管理

Go语言使用`go mod`进行包管理：

```bash
# 初始化项目
go mod init 项目名称

# 下载项目缺少的依赖
go mod tidy

# 下载包
go get github.com/package/path
go install github.com/package/path@latest
go mod tidy
```

#### 包的导入

```go
import "包路径"           // 正常导入
import _ "包路径"        // 匿名导入（只执行init函数）
import 别名 "包路径"      // 别名导入
import . "包路径"        // 点导入（可以直接使用包中的函数，不需要包名）
```

#### 常用内置包

Go语言提供了丰富的标准库，以下是一些常用的内置包：

- **fmt**：格式化输入输出
- **strconv**：字符串与基本数据类型转换
- **strings**：字符串操作
- **sort**：排序算法
- **errors**：错误处理
- **time**：时间处理
- **encoding/json**：JSON编解码
- **os**：操作系统接口
- **io**：I/O操作
- **unsafe**：不安全的操作（如获取变量大小）

这些包都是Go语言标准库的一部分，无需下载即可使用。

---

### 16. 错误处理

Go语言使用`error`类型来表示错误，这是一种显式的错误处理机制。

#### 错误处理机制

- Go语言使用`error`类型来表示错误
- 函数通常返回两个值：结果值和`error`
- 通过检查`error`是否为`nil`来判断是否发生错误

#### 错误创建

```go
errors.New("错误信息")
```

**代码示例：**
```go
// 模拟一个读取文件的方法
func readFile(fileName string) error {
	fmt.Println("filename: ", fileName)
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("readFile读取文件失败！")
	}
}
```

#### 错误处理

```go
err := 可能出错的函数()
if err != nil {
	// 处理错误
}
```

#### panic和recover

`panic`和`recover`用于处理程序运行时的异常情况。

**panic：**
- 可以在任何地方引发异常
- 会立即停止当前函数的执行
- 如果没有`recover`捕获，程序会终止

**recover：**
- 只有在`defer`调用的函数中有效
- 用于捕获`panic`，使程序可以继续执行

**使用方式：**
```go
panic("错误信息")
panic(err)

defer func() {
	if err := recover(); err != nil {
		// 处理panic
		fmt.Println("错误:", err)
	}
}()
```

**代码示例：**
```go
// 使用recover捕获panic
func fn1(x, y int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("fn1 func error: ", err)
		}
	}()
	return x / y  // 如果y=0会panic
}

// 使用panic和recover处理错误
func myFn(filename string) {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员提个醒，错误: ", e)
		}
	}()
	
	err := readFile(filename)
	if err != nil {
		panic(err)
	}
}

myFn("xxx.go")
// 输出：filename:  xxx.go
//      给管理员提个醒，错误:  readFile读取文件失败！
```

### 17. 时间处理

Go语言的`time`包提供了时间处理相关功能。

#### 时间格式化

Go语言使用特定的时间作为格式化模板：`2006-01-02 15:04:05`，这个时间是Go语言的诞生时间。

**代码示例：**
```go
// 当前时间
timeObj := time.Now()

// 格式化
strTime1 := timeObj.Format("2006-01-02 03:04:05")  // 12小时制
strTime2 := timeObj.Format("2006-01-02 15:04:05")  // 24小时制
fmt.Println("time: ", strTime1, strTime2)
```

#### 时间戳

```go
// 获取当前时间戳（秒）
unixTime := timeObj.Unix()

// 时间戳转时间
timeObj := time.Unix(unixTime, 0)

// 日期字符串转时间戳
timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 12:00:00", time.Local)
unixTime := timeObj.Unix()
```

**代码示例：**
```go
// 当前时间戳
unixTime := timeObj.Unix()
fmt.Println("unixTime: ", unixTime)

// 时间戳转时间
timeObj1 := time.Unix(unixTime, 0)
strTime3 := timeObj1.Format("2006-01-02 15:04:05")
fmt.Println("strTime3: ", strTime3)

// 日期字符串转时间戳
timeObj2, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime3, time.Local)
fmt.Println("time to unix", timeObj2.Unix())
```

#### 时间操作

```go
// 时间加减
timeAdd := timeObj.Add(time.Hour)                    // 加1小时
timeSub := timeObj.Add(-time.Hour)                   // 减1小时
timeAdd := timeObj.Add(time.Duration(2) * time.Hour) // 加2小时
```

**代码示例：**
```go
timeAdd := timeObj2.Add(time.Hour)
fmt.Println("time add ", time.Hour, timeAdd)
```

#### 定时器和睡眠

**定时器（Ticker）：**
```go
ticker := time.NewTicker(time.Second)
for t := range ticker.C {
	// 每秒执行一次
	// 停止：ticker.Stop()
}
```

**睡眠：**
```go
time.Sleep(time.Second)  // 睡眠1秒
```

**代码示例：**
```go
// 定时器
n := 5
ticker := time.NewTicker(time.Second)
for t := range ticker.C {
	n--
	if n == 0 {
		ticker.Stop()  // 终止定时器
		break
	}
	fmt.Println("ticker: ", t)
}

// 睡眠
fmt.Println("开始休眠...")
time.Sleep(time.Second * 3)
fmt.Println("休眠结束...")
```

### 18. 排序算法

Go语言标准库提供了`sort`包，包含了多种排序算法的实现，可以直接使用。

#### sort包常用函数

```go
sort.Ints([]int)          // int切片排序
sort.Strings([]string)    // string切片排序
sort.Float64s([]float64)  // float64切片排序
```

**代码示例：**
```go
numSlice := []int{5, 2, 6, 3, 1, 4}
sort.Ints(numSlice)
fmt.Println("sorted numSlice: ", numSlice)
// 输出：sorted numSlice: [1 2 3 4 5 6]
```

#### 自定义排序

实现`sort.Interface`接口可以自定义排序规则：

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

#### 常见排序算法

冒泡排序、选择排序、插入排序、归并排序、快速排序等。Go标准库已经实现了高效的排序算法，建议直接使用标准库。

**代码示例（冒泡排序）：**
```go
numSlice := []int{5, 2, 6, 3, 1, 4}
for i := 0; i < len(numSlice)-1; i++ {
	for j := i + 1; j < len(numSlice); j++ {
		if numSlice[i] > numSlice[j] {
			temp := numSlice[i]
			numSlice[i] = numSlice[j]
			numSlice[j] = temp
		}
	}
	fmt.Println("sorted numSlice: ", numSlice)
}
```


## 三、学习要点总结

1. **变量和常量**：掌握var、:=、const、iota的使用
2. **数据类型**：理解各种数据类型的特点和转换方式
3. **数组和切片**：理解值类型和引用类型的区别
4. **Map**：掌握map的创建、操作和遍历
5. **指针**：理解指针的概念和使用场景
6. **结构体**：掌握结构体的定义、方法和嵌套
7. **接口**：理解接口的实现和多态
8. **函数**：掌握函数类型、匿名函数、闭包、defer的使用
9. **错误处理**：掌握error、panic、recover的使用
10. **包管理**：了解go mod的使用和第三方包的引入
