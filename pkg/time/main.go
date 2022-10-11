package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个时间对象
	now := time.Now()
	fmt.Println(now)

	// 格式化成给人看的时间
	// strNow := now.Format("2022-10-11 23:29:05")
	// golang 的时间格式化必须是固定的 2006-01-02 15:04:05
	strNow := now.Format("2006-01-02 15:04:05")
	fmt.Println(strNow)

	// 时间戳 从1970-01-01 00:00:00 到现在的秒数或者毫秒数
	fmt.Println(now.Unix())      // 秒
	fmt.Println(now.UnixMilli()) // 毫秒

	// 格式化字符串转时间对象
	loc, err := time.LoadLocation("Asia/Shanghai") // 设置时区
	if err != nil {
		fmt.Println(err)
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", strNow, loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)

	// 获取当前时间，以及时间拆分
	now2 := time.Now()
	year := now2.Year()
	month := now2.Month()
	day := now2.Day()
	hour := now2.Hour()
	minute := now2.Minute()
	second := now2.Second()
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second) // 02d 表示两位数，不足两位前面补0

	// 时间戳
	now3 := time.Now()
	timestamp1 := now3.Unix()      // 秒
	timestamp2 := now3.UnixMilli() // 毫秒
	timestamp3 := now3.UnixNano()  // 纳秒
	fmt.Println(timestamp1, timestamp2, timestamp3)
	// 时间戳转时间对象
	timeObj2 := time.Unix(timestamp1, 0)
	fmt.Println(timeObj2)
}
