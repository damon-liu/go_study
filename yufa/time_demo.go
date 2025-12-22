package yufa

import (
	"fmt"
	"time"
)

func init() {
	// timeDemo()
}

func timeDemo() {
	// 时间格式化
	timeObj := time.Now()
	strTime1 := timeObj.Format("2006-01-02 03:04:05")
	strTime2 := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println("time: ", strTime1, strTime2)

	// 当前时间戳（毫秒）
	unixTime := timeObj.Unix()
	timeObj1 := time.Unix(unixTime, 0)
	strTime3 := timeObj1.Format("2006-01-02 15:04:05")
	fmt.Println("unixTime: ", unixTime, strTime3)

	// 日期字符串转成时间戳
	var timeObj2, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime3, time.Local)
	fmt.Println("time to unix", timeObj2.Unix())

	// 时间操作函数
	timeAdd := timeObj2.Add(time.Hour)
	fmt.Println("time add ", time.Hour, timeAdd)

	// 定时器
	n := 5
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		n--
		if n == 0 {
			// 终止定时器
			ticker.Stop()
			break
		}
		fmt.Println("ticker: ", t)
	}

	fmt.Println("开始休眠...")
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("休眠结束...")

}
