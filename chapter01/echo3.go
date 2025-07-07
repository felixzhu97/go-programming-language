// Echo3 打印命令行参数 (使用strings.Join)
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
} 