// Constants 演示常量的声明和使用
package main

import (
	"fmt"
	"math"
	"time"
)

// 1. 包级常量
const (
	Pi      = 3.14159265358979323846
	E       = 2.71828182845904523536
	Version = "1.0.0"
	Debug   = true
)

// 2. 无类型常量
const (
	Big   = 1e100
	Small = Big / 1e99
)

// 3. iota 枚举器
const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// 4. 复杂的 iota 使用
const (
	_  = iota             // 忽略第一个值
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30 = 1073741824
	TB                    // 1 << 40 = 1099511627776
)

// 5. 标志位
const (
	FlagNone   = 0
	FlagRead   = 1 << iota // 1
	FlagWrite              // 2
	FlagExecute            // 4
)

// 6. 类型化常量
const (
	Timeout time.Duration = 30 * time.Second
	MaxRetries            = 3
)

func main() {
	// 1. 基本常量使用
	fmt.Printf("基本常量:\n")
	fmt.Printf("Pi: %.15f\n", Pi)
	fmt.Printf("E: %.15f\n", E)
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Debug: %t\n", Debug)

	// 2. 无类型常量的灵活性
	fmt.Printf("\n无类型常量:\n")
	fmt.Printf("Big: %g\n", Big)
	fmt.Printf("Small: %g\n", Small)
	// 无类型常量可以赋值给任何兼容类型
	var f32 float32 = Small
	var f64 float64 = Small
	fmt.Printf("Small 作为 float32: %g\n", f32)
	fmt.Printf("Small 作为 float64: %g\n", f64)

	// 3. iota 枚举
	fmt.Printf("\n星期枚举:\n")
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tuesday: %d\n", Tuesday)
	fmt.Printf("Wednesday: %d\n", Wednesday)
	fmt.Printf("Thursday: %d\n", Thursday)
	fmt.Printf("Friday: %d\n", Friday)
	fmt.Printf("Saturday: %d\n", Saturday)
	fmt.Printf("Sunday: %d\n", Sunday)

	// 4. 字节单位
	fmt.Printf("\n字节单位:\n")
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)

	// 5. 标志位操作
	fmt.Printf("\n标志位操作:\n")
	fmt.Printf("FlagRead: %d\n", FlagRead)
	fmt.Printf("FlagWrite: %d\n", FlagWrite)
	fmt.Printf("FlagExecute: %d\n", FlagExecute)

	// 组合标志
	permissions := FlagRead | FlagWrite
	fmt.Printf("Read+Write: %d\n", permissions)
	
	// 检查标志
	fmt.Printf("Has read permission: %t\n", permissions&FlagRead != 0)
	fmt.Printf("Has write permission: %t\n", permissions&FlagWrite != 0)
	fmt.Printf("Has execute permission: %t\n", permissions&FlagExecute != 0)

	// 6. 类型化常量
	fmt.Printf("\n类型化常量:\n")
	fmt.Printf("Timeout: %v\n", Timeout)
	fmt.Printf("MaxRetries: %d\n", MaxRetries)

	// 7. 常量表达式
	fmt.Printf("\n常量表达式:\n")
	const (
		width  = 100
		height = 200
		area   = width * height
	)
	fmt.Printf("Area: %d\n", area)

	// 8. 字符串常量
	fmt.Printf("\n字符串常量:\n")
	const greeting = "Hello, " + "World!"
	fmt.Printf("Greeting: %s\n", greeting)

	// 9. 复数常量
	fmt.Printf("\n复数常量:\n")
	const (
		real = 3.0
		imag = 4.0
		complex_num = real + imag*1i
	)
	fmt.Printf("Complex number: %v\n", complex_num)
	fmt.Printf("Real part: %.1f\n", real)
	fmt.Printf("Imaginary part: %.1f\n", imag)

	// 10. 预声明常量
	fmt.Printf("\n预声明常量:\n")
	fmt.Printf("true: %t\n", true)
	fmt.Printf("false: %t\n", false)
	fmt.Printf("iota: %d\n", iota) // 在函数中，iota 的值为 0

	// 11. 常量的类型推断
	fmt.Printf("\n常量的类型推断:\n")
	const untyped = 42
	var i int = untyped
	var f float64 = untyped
	var c complex128 = untyped
	fmt.Printf("作为 int: %d\n", i)
	fmt.Printf("作为 float64: %.1f\n", f)
	fmt.Printf("作为 complex128: %v\n", c)

	// 12. 常量在编译时计算
	fmt.Printf("\n编译时计算:\n")
	const calculated = 1 + 2*3 + 4*5
	fmt.Printf("1 + 2*3 + 4*5 = %d\n", calculated)

	// 13. 数学常量
	fmt.Printf("\n数学常量:\n")
	fmt.Printf("math.Pi: %.15f\n", math.Pi)
	fmt.Printf("math.E: %.15f\n", math.E)
	fmt.Printf("math.Sqrt2: %.15f\n", math.Sqrt2)
	fmt.Printf("math.Phi: %.15f\n", math.Phi)

	// 14. 实际应用：状态机
	fmt.Printf("\n状态机示例:\n")
	const (
		StateIdle = iota
		StateRunning
		StatePaused
		StateStopped
	)
	
	state := StateRunning
	switch state {
	case StateIdle:
		fmt.Println("系统空闲")
	case StateRunning:
		fmt.Println("系统运行中")
	case StatePaused:
		fmt.Println("系统暂停")
	case StateStopped:
		fmt.Println("系统停止")
	}
} 