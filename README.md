# Golang语言学习

## 一、语法

### 数据类型

```
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
```



### 值类型与引用类型

```
值类型：
	基本数据类型（int、float、bool、string）、数组、结构体
	引用类型：指针、切片、map、chan、interface、函数

值类型赋值时，是将值复制一份给新的变量，两个变量互不影响
引用类型赋值时，是将引用地址复制一份给新的变量，两个变量指向同一块内存，修改其中一个变量会影响另一个变量
```



### 流程控制

```
if条件判断格式：
			if 条件表达式 {
				执行语句
			} else if 条件表达式 {
				执行语句
			} else {
				执行语句
			}

for循环格式：
			for 初始化语句; 条件表达式; 后置语句 {
				执行语句
			}
			

switch条件判断格式：
			switch 表达式 {
			case 值1:
				执行语句
			case 值2:
				执行语句
			default:
				执行语句
			}

break/continue：
			break用于终止当前循环
			continue跳出当前循环继续下次循环

goto格式：
			goto 标签名
			****
			标签名:
```



###  函数

```
函数定义格式:
		1、固定参数
		func 函数名(参数)(返回值){
			....
		}

		2、可变参数
		func 函数名(...)(返回值){
			....
		}
```



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
				在go语言函数中return语句底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前，defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
			panic:
				可以在任何地方引发，但recover只有在defer调用的函数中有效