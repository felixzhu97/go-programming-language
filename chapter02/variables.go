// Variables 演示变量声明和初始化的各种方式
package main

import "fmt"

// 包级变量
var (
	name     = "Go语言"
	version  = "1.21"
	year     = 2023
	isActive = true
)

func main() {
	// 1. 基本变量声明
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("零值: i=%d, f=%g, b=%t, s=%q\n", i, f, b, s)

	// 2. 变量声明并初始化
	var x, y int = 1, 2
	var name, age = "Alice", 25
	fmt.Printf("初始化: x=%d, y=%d, name=%s, age=%d\n", x, y, name, age)

	// 3. 短变量声明
	msg := "Hello, World!"
	count := 42
	fmt.Printf("短声明: msg=%s, count=%d\n", msg, count)

	// 4. 多重赋值
	a, b := 10, 20
	fmt.Printf("多重赋值前: a=%d, b=%d\n", a, b)
	a, b = b, a // 交换
	fmt.Printf("多重赋值后: a=%d, b=%d\n", a, b)

	// 5. 类型推断
	var auto = 3.14159
	fmt.Printf("类型推断: auto=%T, value=%g\n", auto, auto)

	// 6. 包级变量
	fmt.Printf("包级变量: %s %s (%d), 活跃状态: %t\n", name, version, year, isActive)
} 