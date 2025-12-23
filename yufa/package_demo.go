package yufa

import (
	"fmt"
	"go_study/jiegouti"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func init() {
	// packageDemo()
}

func packageDemo() {
	/*
		包是多个Go源码的集合，Go为我们提供了很多内置包，如：fmt、strconv、strings、sort、errors、time、encoding/json、os、io等

		分为三种：
			1、系统内置包
			2、自定义包
			3、第三方包

		系统内置包：Golang语言未我们提供的内之宝，引入后可以直接使用，如：fmt、strconv、strings、sort、errors、time、encoding/json、os、io等

		自定义包：开发者自己写的包

		第三方包：属于自定义包的一种，需要下载安装到本地才可以使用
		包仓库：https:pkg.go.dev/
		初始化项目：go init 项目名称
		下载项目缺少的依赖：go mod tidy
		下载包：go get github.com/shopspring/decimal 或 go install github.com/shopspring/decimal@latest 或 go mod tidy
	*/

	// 本地包引入
	jiegouti.ExtendsDemno()

	// 第三方包
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println("price:  ", price)

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	fmt.Println("gjson: ", value.String())
}
