// 修改echo程序，输出参数的索引和值，每行一个
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0:] 会产生一个新的slice，所以index又从0开始
	for index, para := range os.Args[1:] {
		fmt.Println("index: ", index+1, ", para:", para)
	}
}
