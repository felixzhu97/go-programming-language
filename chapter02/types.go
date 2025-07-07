// Types 演示类型声明和使用
package main

import "fmt"

// 1. 基本类型别名
type Celsius float64
type Fahrenheit float64

// 2. 结构体类型
type Point struct {
	X, Y int
}

// 3. 函数类型
type Calculator func(int, int) int

// 4. 方法定义
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", float64(f))
}

// 5. 类型转换函数
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 6. 计算器实现
func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }

func main() {
	// 1. 使用自定义类型
	var c Celsius = 20
	var f Fahrenheit = 68
	fmt.Printf("温度: %s, %s\n", c, f)

	// 2. 类型转换
	fmt.Printf("20°C = %s\n", CToF(c))
	fmt.Printf("68°F = %s\n", FToC(f))

	// 3. 结构体使用
	p1 := Point{1, 2}
	p2 := Point{X: 3, Y: 4}
	fmt.Printf("点: p1=%v, p2=%v\n", p1, p2)

	// 4. 函数类型使用
	var calc Calculator
	calc = add
	fmt.Printf("加法: %d\n", calc(3, 4))
	calc = mul
	fmt.Printf("乘法: %d\n", calc(3, 4))

	// 5. 类型断言和类型检查
	var x interface{} = 42
	if i, ok := x.(int); ok {
		fmt.Printf("x 是 int 类型，值为: %d\n", i)
	}

	// 6. 类型比较
	var c1 Celsius = 25
	var c2 Celsius = 25
	fmt.Printf("c1 == c2: %t\n", c1 == c2)

	// 注意：不同类型不能直接比较
	// fmt.Printf("c == f: %t\n", c == Celsius(f)) // 需要类型转换
} 