package timeUtil

import (
	"fmt"
	"time"
)

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
//var timeLayoutStr = 2006/01/02 03:04:05 //合法, 格式可以改变
//var timeLayoutStr = 2019/01/02 15:04:05 //不合法, 时间必须是2016年1月2号这个时间

func StringToTime(ts string) time.Time {
	parse, _ := time.Parse(timeLayoutStr, ts)
	return parse
}

func TimeToString(ts time.Time) string {
	return ts.Format(timeLayoutStr)
}

func testFormat() {
	t := time.Now() //当前时间
	t.Unix()        //时间戳

	ts := t.Format(timeLayoutStr) //time转string
	fmt.Println(ts)
	st, _ := time.Parse(timeLayoutStr, ts) //string转time
	fmt.Println(st)

	//在go中, 可以格式化一个带前后缀的时间字符串
	prefixTStr := "PREFIX-- 2019-01-01 -TEST- 10:31:12 --SUFFIX"       //带前后缀的时间字符串
	preTimeLayoutStr := "PREFIX-- 2006-01-02 -TEST- 15:04:05 --SUFFIX" //需要转换的时间格式, 格式和前后缀需要一致, 这种写法的限制很大, 但一些特殊场景可以用到
	prefixTime, _ := time.Parse(preTimeLayoutStr, prefixTStr)
	fmt.Println(prefixTime)

	//时间加减 time.ParseDuration()
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	at, _ := time.ParseDuration("2h") //2个小时后的时间, 负数就是之前的时间
	fmt.Println((t.Add(at)).Format(timeLayoutStr))

	//两个时间差
	sub := t.Sub(prefixTime)
	fmt.Println(sub.Seconds()) //秒,  sub.Minutes()分钟,  sub.Hours()小时...

}
