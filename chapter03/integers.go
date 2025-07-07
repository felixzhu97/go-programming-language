// Integers 演示整数类型和操作
package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 整数类型
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 42

	fmt.Printf("有符号整数:\n")
	fmt.Printf("int8: %d (范围: %d 到 %d)\n", i8, math.MinInt8, math.MaxInt8)
	fmt.Printf("int16: %d (范围: %d 到 %d)\n", i16, math.MinInt16, math.MaxInt16)
	fmt.Printf("int32: %d (范围: %d 到 %d)\n", i32, math.MinInt32, math.MaxInt32)
	fmt.Printf("int64: %d (范围: %d 到 %d)\n", i64, math.MinInt64, math.MaxInt64)
	fmt.Printf("int: %d\n", i)

	// 2. 无符号整数类型
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 42

	fmt.Printf("\n无符号整数:\n")
	fmt.Printf("uint8: %d (范围: 0 到 %d)\n", u8, math.MaxUint8)
	fmt.Printf("uint16: %d (范围: 0 到 %d)\n", u16, math.MaxUint16)
	fmt.Printf("uint32: %d (范围: 0 到 %d)\n", u32, math.MaxUint32)
	fmt.Printf("uint64: %d (范围: 0 到 %d)\n", u64, uint64(math.MaxUint64))
	fmt.Printf("uint: %d\n", u)

	// 3. 字节和符文类型
	var b byte = 'A'      // byte 是 uint8 的别名
	var r rune = '中'      // rune 是 int32 的别名
	var ptr uintptr = 0x1000 // 指针大小的无符号整数

	fmt.Printf("\n特殊类型:\n")
	fmt.Printf("byte: %d ('%c')\n", b, b)
	fmt.Printf("rune: %d ('%c')\n", r, r)
	fmt.Printf("uintptr: %d (0x%x)\n", ptr, ptr)

	// 4. 整数运算
	a, b := 10, 3
	fmt.Printf("\n整数运算 (a=%d, b=%d):\n", a, b)
	fmt.Printf("加法: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("减法: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("乘法: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("除法: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("取模: %d %% %d = %d\n", a, b, a%b)

	// 5. 位运算
	x, y := 12, 10 // 二进制: 1100, 1010
	fmt.Printf("\n位运算 (x=%d, y=%d):\n", x, y)
	fmt.Printf("按位与: %d & %d = %d\n", x, y, x&y)
	fmt.Printf("按位或: %d | %d = %d\n", x, y, x|y)
	fmt.Printf("按位异或: %d ^ %d = %d\n", x, y, x^y)
	fmt.Printf("按位取反: ^%d = %d\n", x, ^x)
	fmt.Printf("左移: %d << 1 = %d\n", x, x<<1)
	fmt.Printf("右移: %d >> 1 = %d\n", x, x>>1)

	// 6. 数字字面量
	fmt.Printf("\n数字字面量:\n")
	fmt.Printf("十进制: %d\n", 42)
	fmt.Printf("八进制: %o (十进制: %d)\n", 052, 052)
	fmt.Printf("十六进制: %x (十进制: %d)\n", 0x2a, 0x2a)
	fmt.Printf("二进制: %b (十进制: %d)\n", 0b101010, 0b101010)

	// 7. 溢出演示
	fmt.Printf("\n溢出演示:\n")
	var max8 int8 = 127
	fmt.Printf("int8最大值: %d\n", max8)
	fmt.Printf("溢出后: %d\n", max8+1) // 会发生溢出

	// 8. 类型转换
	fmt.Printf("\n类型转换:\n")
	var big int64 = 1000000
	var small int8 = int8(big % 128)
	fmt.Printf("int64 %d 转换为 int8: %d\n", big, small)
} 