// Assignments 演示各种赋值操作
package main

import "fmt"

func main() {
	// 1. 基本赋值
	var x int
	x = 1
	fmt.Printf("基本赋值: x=%d\n", x)

	// 2. 多重赋值
	var a, b int
	a, b = 1, 2
	fmt.Printf("多重赋值: a=%d, b=%d\n", a, b)

	// 3. 交换变量
	fmt.Printf("交换前: a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后: a=%d, b=%d\n", a, b)

	// 4. 函数返回多个值
	quotient, remainder := divide(10, 3)
	fmt.Printf("除法结果: 商=%d, 余数=%d\n", quotient, remainder)

	// 5. 丢弃不需要的返回值
	q, _ := divide(15, 4)
	fmt.Printf("只要商: %d\n", q)

	// 6. 复合赋值操作符
	x = 5
	fmt.Printf("初始值: x=%d\n", x)
	x += 3
	fmt.Printf("x += 3: x=%d\n", x)
	x -= 2
	fmt.Printf("x -= 2: x=%d\n", x)
	x *= 2
	fmt.Printf("x *= 2: x=%d\n", x)
	x /= 3
	fmt.Printf("x /= 3: x=%d\n", x)

	// 7. 自增自减
	fmt.Printf("自增前: x=%d\n", x)
	x++
	fmt.Printf("x++: x=%d\n", x)
	x--
	fmt.Printf("x--: x=%d\n", x)

	// 8. 指针赋值
	var p *int
	p = &x
	fmt.Printf("指针赋值: p=%p, *p=%d\n", p, *p)
	*p = 100
	fmt.Printf("通过指针修改: x=%d\n", x)

	// 9. 类型推断赋值
	message := "Hello"
	number := 42
	pi := 3.14159
	fmt.Printf("类型推断: message=%s, number=%d, pi=%f\n", message, number, pi)
}

// 返回两个值的函数
func divide(a, b int) (int, int) {
	return a / b, a % b
} 