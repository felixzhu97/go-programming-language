// Floats 演示浮点数类型和操作
package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 浮点数类型
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	
	fmt.Printf("浮点数类型:\n")
	fmt.Printf("float32: %.10f\n", f32)
	fmt.Printf("float64: %.16f\n", f64)
	fmt.Printf("float32范围: %g 到 %g\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64范围: %g 到 %g\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	// 2. 科学计数法
	fmt.Printf("\n科学计数法:\n")
	fmt.Printf("1.23e4 = %g\n", 1.23e4)
	fmt.Printf("1.23e-4 = %g\n", 1.23e-4)
	fmt.Printf("1.23E4 = %g\n", 1.23E4)

	// 3. 浮点数运算
	a, b := 10.5, 3.2
	fmt.Printf("\n浮点数运算 (a=%.1f, b=%.1f):\n", a, b)
	fmt.Printf("加法: %.1f + %.1f = %.1f\n", a, b, a+b)
	fmt.Printf("减法: %.1f - %.1f = %.1f\n", a, b, a-b)
	fmt.Printf("乘法: %.1f * %.1f = %.1f\n", a, b, a*b)
	fmt.Printf("除法: %.1f / %.1f = %.1f\n", a, b, a/b)

	// 4. 特殊值
	fmt.Printf("\n特殊值:\n")
	fmt.Printf("正无穷: %g\n", math.Inf(1))
	fmt.Printf("负无穷: %g\n", math.Inf(-1))
	fmt.Printf("NaN: %g\n", math.NaN())
	fmt.Printf("是否为NaN: %t\n", math.IsNaN(math.NaN()))
	fmt.Printf("是否为无穷: %t\n", math.IsInf(math.Inf(1), 1))

	// 5. 精度问题演示
	fmt.Printf("\n精度问题:\n")
	fmt.Printf("0.1 + 0.2 = %.17f\n", 0.1+0.2)
	fmt.Printf("0.3 = %.17f\n", 0.3)
	fmt.Printf("相等比较: 0.1 + 0.2 == 0.3: %t\n", 0.1+0.2 == 0.3)
	
	// 正确的浮点数比较
	const epsilon = 1e-9
	diff := math.Abs((0.1 + 0.2) - 0.3)
	fmt.Printf("使用epsilon比较: |差值| < %g: %t\n", epsilon, diff < epsilon)

	// 6. 数学函数
	fmt.Printf("\n数学函数:\n")
	x := 2.5
	fmt.Printf("sqrt(%.1f) = %.3f\n", x, math.Sqrt(x))
	fmt.Printf("pow(%.1f, 2) = %.3f\n", x, math.Pow(x, 2))
	fmt.Printf("sin(π/4) = %.3f\n", math.Sin(math.Pi/4))
	fmt.Printf("cos(π/4) = %.3f\n", math.Cos(math.Pi/4))
	fmt.Printf("log(e) = %.3f\n", math.Log(math.E))
	fmt.Printf("log10(100) = %.3f\n", math.Log10(100))

	// 7. 取整函数
	fmt.Printf("\n取整函数:\n")
	val := 3.7
	fmt.Printf("原值: %.1f\n", val)
	fmt.Printf("Floor: %.1f\n", math.Floor(val))
	fmt.Printf("Ceil: %.1f\n", math.Ceil(val))
	fmt.Printf("Round: %.1f\n", math.Round(val))
	fmt.Printf("Trunc: %.1f\n", math.Trunc(val))

	// 8. 类型转换
	fmt.Printf("\n类型转换:\n")
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("int %d 转换为 float64: %.1f\n", i, f)
	
	var f2 float64 = 3.14159
	var i2 int = int(f2)
	fmt.Printf("float64 %.5f 转换为 int: %d\n", f2, i2)

	// 9. 格式化输出
	fmt.Printf("\n格式化输出:\n")
	pi := math.Pi
	fmt.Printf("默认: %g\n", pi)
	fmt.Printf("固定小数点: %.2f\n", pi)
	fmt.Printf("科学计数法: %.2e\n", pi)
	fmt.Printf("更紧凑的科学计数法: %.2g\n", pi)
} 