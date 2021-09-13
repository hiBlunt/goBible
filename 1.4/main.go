// 修改dup2程序，输出出现重复行的文件的名称

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dip2: %v\n", err)
			}
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					fmt.Printf("%s include repeated line\n", arg)
				}
				break
			}
			counts = make(map[string]int)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意：忽略input.Err()中可能的错误
}
