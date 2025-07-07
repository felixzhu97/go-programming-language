// Booleans 演示布尔类型和逻辑运算
package main

import "fmt"

func main() {
	// 1. 布尔类型基础
	var b1 bool = true
	var b2 bool = false
	var b3 bool // 零值为false
	
	fmt.Printf("布尔类型:\n")
	fmt.Printf("b1: %t\n", b1)
	fmt.Printf("b2: %t\n", b2)
	fmt.Printf("b3 (零值): %t\n", b3)

	// 2. 逻辑运算符
	fmt.Printf("\n逻辑运算符:\n")
	fmt.Printf("true && true = %t\n", true && true)
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("false && true = %t\n", false && true)
	fmt.Printf("false && false = %t\n", false && false)
	
	fmt.Printf("true || true = %t\n", true || true)
	fmt.Printf("true || false = %t\n", true || false)
	fmt.Printf("false || true = %t\n", false || true)
	fmt.Printf("false || false = %t\n", false || false)
	
	fmt.Printf("!true = %t\n", !true)
	fmt.Printf("!false = %t\n", !false)

	// 3. 短路求值
	fmt.Printf("\n短路求值:\n")
	x, y := 5, 0
	
	// && 短路：如果左边为false，右边不会被计算
	if x > 10 && y/x > 1 {
		fmt.Println("不会执行")
	} else {
		fmt.Println("&& 短路：避免了除零错误")
	}
	
	// || 短路：如果左边为true，右边不会被计算
	if x > 0 || y/x > 1 {
		fmt.Println("|| 短路：避免了除零错误")
	}

	// 4. 比较运算符
	fmt.Printf("\n比较运算符:\n")
	a, b := 10, 20
	fmt.Printf("%d == %d: %t\n", a, b, a == b)
	fmt.Printf("%d != %d: %t\n", a, b, a != b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)

	// 5. 字符串比较
	fmt.Printf("\n字符串比较:\n")
	str1, str2 := "apple", "banana"
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str2, str1 == str2)
	fmt.Printf("\"%s\" < \"%s\": %t\n", str1, str2, str1 < str2)
	fmt.Printf("\"%s\" > \"%s\": %t\n", str1, str2, str1 > str2)

	// 6. 条件语句
	fmt.Printf("\n条件语句:\n")
	age := 18
	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	// 7. 复合条件
	fmt.Printf("\n复合条件:\n")
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 8. 布尔值作为函数参数和返回值
	fmt.Printf("\n布尔值作为函数参数和返回值:\n")
	fmt.Printf("isEven(4): %t\n", isEven(4))
	fmt.Printf("isEven(5): %t\n", isEven(5))
	fmt.Printf("isPositive(-3): %t\n", isPositive(-3))
	fmt.Printf("isPositive(3): %t\n", isPositive(3))

	// 9. 三元运算符的替代
	fmt.Printf("\n三元运算符的替代:\n")
	// Go没有三元运算符，但可以用函数实现
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Printf("max(10, 20): %d\n", max(10, 20))

	// 10. 布尔值转换
	fmt.Printf("\n布尔值转换:\n")
	fmt.Printf("boolToInt(true): %d\n", boolToInt(true))
	fmt.Printf("boolToInt(false): %d\n", boolToInt(false))
	fmt.Printf("intToBool(0): %t\n", intToBool(0))
	fmt.Printf("intToBool(1): %t\n", intToBool(1))
	fmt.Printf("intToBool(-1): %t\n", intToBool(-1))
}

// 判断是否为偶数
func isEven(n int) bool {
	return n%2 == 0
}

// 判断是否为正数
func isPositive(n int) bool {
	return n > 0
}

// 布尔值转整数
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 整数转布尔值
func intToBool(i int) bool {
	return i != 0
} 