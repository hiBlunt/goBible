// 修改echo程序输出os.Args[0], 即命令行名字
package main

import (
	"fmt"
	"os"
)

func main() {
	processName := os.Args[0]
	fmt.Println("process name is ", processName)
}
