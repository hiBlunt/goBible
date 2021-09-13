// 尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异
package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	// 比较字符串拼接速度。 使用+和strings.Join拼接2^12次字符串
	length := int64(math.Pow(2, 12))

	t1 := time.Now().UnixNano()
	tmp := ""
	var i int64
	for ; i < length; i++ {
		tmp += "ha"
	}
	t2 := time.Now().UnixNano()
	fmt.Println("+ use time:", t2-t1)

	tmp = ""
	strSlice := make([]string, length)
	//NOTE(lhuang): 遍历切片是，value值只是一个拷贝，无法修改原始值
	for index, _ := range strSlice {
		strSlice[index] = "ha"
	}
	t3 := time.Now().UnixNano()
	tmp = strings.Join(strSlice, "")
	t4 := time.Now().UnixNano()
	fmt.Println("strings.Join use time:", t4-t3)
}

// 结论：strings.Join速度几乎为0，比单纯+运算速度快很多个数量级
