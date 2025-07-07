// Conversions 演示类型转换
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. 数值类型转换
	fmt.Printf("数值类型转换:\n")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	fmt.Printf("int: %d -> float64: %.1f\n", i, f)
	fmt.Printf("int: %d -> uint: %d\n", i, u)

	// 2. 浮点数转整数（截断）
	fmt.Printf("\n浮点数转整数:\n")
	var pi float64 = 3.14159
	var truncated int = int(pi)
	fmt.Printf("float64: %.5f -> int: %d\n", pi, truncated)

	// 3. 字符串转数值
	fmt.Printf("\n字符串转数值:\n")
	str := "123"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("字符串 \"%s\" -> int: %d\n", str, num)
	}

	strFloat := "3.14159"
	if fnum, err := strconv.ParseFloat(strFloat, 64); err == nil {
		fmt.Printf("字符串 \"%s\" -> float64: %.5f\n", strFloat, fnum)
	}

	strBool := "true"
	if bnum, err := strconv.ParseBool(strBool); err == nil {
		fmt.Printf("字符串 \"%s\" -> bool: %t\n", strBool, bnum)
	}

	// 4. 数值转字符串
	fmt.Printf("\n数值转字符串:\n")
	num := 42
	str1 := strconv.Itoa(num)
	fmt.Printf("int %d -> 字符串: \"%s\"\n", num, str1)

	fnum := 3.14159
	str2 := strconv.FormatFloat(fnum, 'f', 2, 64)
	fmt.Printf("float64 %.5f -> 字符串: \"%s\"\n", fnum, str2)

	bnum := true
	str3 := strconv.FormatBool(bnum)
	fmt.Printf("bool %t -> 字符串: \"%s\"\n", bnum, str3)

	// 5. 字符和数值转换
	fmt.Printf("\n字符和数值转换:\n")
	var ch rune = 'A'
	var chNum int = int(ch)
	fmt.Printf("字符 '%c' -> int: %d\n", ch, chNum)

	var num2 int = 65
	var ch2 rune = rune(num2)
	fmt.Printf("int %d -> 字符: '%c'\n", num2, ch2)

	// 6. 字节切片和字符串转换
	fmt.Printf("\n字节切片和字符串转换:\n")
	str4 := "Hello, 世界"
	bytes := []byte(str4)
	fmt.Printf("字符串: \"%s\"\n", str4)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("恢复字符串: \"%s\"\n", string(bytes))

	// 7. 不同进制的转换
	fmt.Printf("\n不同进制的转换:\n")
	decimal := 255
	binary := strconv.FormatInt(int64(decimal), 2)
	octal := strconv.FormatInt(int64(decimal), 8)
	hex := strconv.FormatInt(int64(decimal), 16)
	fmt.Printf("十进制 %d:\n", decimal)
	fmt.Printf("  二进制: %s\n", binary)
	fmt.Printf("  八进制: %s\n", octal)
	fmt.Printf("  十六进制: %s\n", hex)

	// 从不同进制解析
	if binNum, err := strconv.ParseInt("11111111", 2, 64); err == nil {
		fmt.Printf("二进制 \"11111111\" -> 十进制: %d\n", binNum)
	}

	// 8. 类型断言（接口转换）
	fmt.Printf("\n类型断言:\n")
	var x interface{} = 42
	if i, ok := x.(int); ok {
		fmt.Printf("interface{} 值 %v 是 int 类型: %d\n", x, i)
	}

	var y interface{} = "hello"
	if s, ok := y.(string); ok {
		fmt.Printf("interface{} 值 %v 是 string 类型: \"%s\"\n", y, s)
	}

	// 9. 自定义类型转换
	fmt.Printf("\n自定义类型转换:\n")
	type Celsius float64
	type Fahrenheit float64

	var c Celsius = 20.0
	var f Fahrenheit = Fahrenheit(c*9/5 + 32)
	fmt.Printf("%.1f°C = %.1f°F\n", float64(c), float64(f))

	// 10. 错误处理
	fmt.Printf("\n错误处理:\n")
	invalidStr := "abc"
	if _, err := strconv.Atoi(invalidStr); err != nil {
		fmt.Printf("转换失败: \"%s\" 不是有效的数字\n", invalidStr)
		fmt.Printf("错误信息: %v\n", err)
	}

	// 11. 安全的类型转换
	fmt.Printf("\n安全的类型转换:\n")
	var bigNum int64 = 1000000000000
	if bigNum <= int64(^uint(0)>>1) { // 检查是否在 int 范围内
		var safeInt int = int(bigNum)
		fmt.Printf("安全转换: int64 %d -> int %d\n", bigNum, safeInt)
	} else {
		fmt.Printf("数值 %d 超出 int 范围\n", bigNum)
	}

	// 12. 复数类型转换
	fmt.Printf("\n复数类型转换:\n")
	var c1 complex64 = 3 + 4i
	var c2 complex128 = complex128(c1)
	fmt.Printf("complex64: %v -> complex128: %v\n", c1, c2)

	// 提取实部和虚部
	real := real(c2)
	imag := imag(c2)
	fmt.Printf("实部: %.1f, 虚部: %.1f\n", real, imag)

	// 13. 指针类型转换
	fmt.Printf("\n指针类型转换:\n")
	var val int = 42
	var ptr *int = &val
	var uintPtr uintptr = uintptr(ptr)
	fmt.Printf("指针地址: %p\n", ptr)
	fmt.Printf("转换为 uintptr: %d (0x%x)\n", uintPtr, uintPtr)

	// 14. 切片类型转换
	fmt.Printf("\n切片类型转换:\n")
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", intSlice)
	
	// 注意：只能转换底层类型相同的切片
	type IntSlice []int
	customSlice := IntSlice(intSlice)
	fmt.Printf("自定义切片类型: %v\n", customSlice)
} 