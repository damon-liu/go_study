package yufa

import (
	"errors"
	"fmt"
)

func init() {
	// errorHandler()
}

func errorHandler() {

	// 处理异常
	fn1(10, 0)

	myFn("xxx.go")
}

func fn1(x, y int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("fn1 func error: ", err)
		}
	}()
	return x / y
}

// 模拟一个读取文件的方法
func readFile(fileName string) error {
	fmt.Println("filename: ", fileName)
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
