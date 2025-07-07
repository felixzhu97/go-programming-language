// 演示包的使用
package main

import (
	"fmt"
	"go-programming-language/chapter02/packages/tempconv"
)

func main() {
	// 1. 使用包中的类型
	var c tempconv.Celsius = 20
	var f tempconv.Fahrenheit = 68
	var k tempconv.Kelvin = 293.15

	fmt.Printf("温度值: %s, %s, %s\n", c, f, k)

	// 2. 使用包中的常量
	fmt.Printf("绝对零度: %s\n", tempconv.AbsoluteZeroC)
	fmt.Printf("水的冰点: %s\n", tempconv.FreezingC)
	fmt.Printf("水的沸点: %s\n", tempconv.BoilingC)

	// 3. 使用包中的函数
	fmt.Printf("20°C = %s\n", tempconv.CToF(c))
	fmt.Printf("68°F = %s\n", tempconv.FToC(f))
	fmt.Printf("20°C = %s\n", tempconv.CToK(c))
	fmt.Printf("293.15K = %s\n", tempconv.KToC(k))

	// 4. 类型转换
	fmt.Printf("显式类型转换: %s = %s\n", c, tempconv.Celsius(f))

	// 5. 包的别名
	fmt.Printf("使用包的完整路径访问: %s\n", tempconv.CToF(25))
} 