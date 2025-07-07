// Pointers 演示指针的使用
package main

import "fmt"

func main() {
	// 1. 基本指针操作
	x := 1
	p := &x                   // p指向x
	fmt.Printf("x=%d\n", x)   // x=1
	fmt.Printf("p=%p\n", p)   // p=0x... (x的地址)
	fmt.Printf("*p=%d\n", *p) // *p=1

	// 2. 通过指针修改值
	*p = 2
	fmt.Printf("修改后 x=%d\n", x) // x=2

	// 3. 零值指针
	var q *int
	fmt.Printf("零值指针: q=%p\n", q) // q=0x0 (nil)

	// 4. new函数
	r := new(int)
	fmt.Printf("new创建的指针: r=%p, *r=%d\n", r, *r) // *r=0

	// 5. 函数参数传递
	fmt.Printf("传值调用前: x=%d\n", x)
	incByValue(x)
	fmt.Printf("传值调用后: x=%d\n", x)

	fmt.Printf("传指针调用前: x=%d\n", x)
	incByPointer(&x)
	fmt.Printf("传指针调用后: x=%d\n", x)

	// 6. 指针比较
	y := 1
	px := &x
	py := &y
	fmt.Printf("px == py: %t\n", px == py)     // false
	fmt.Printf("px == &x: %t\n", px == &x)    // true
	fmt.Printf("*px == *py: %t\n", *px == *py) // 比较值

	// 7. 返回局部变量的地址
	ptr := newInt()
	fmt.Printf("函数返回的指针: %p, 值: %d\n", ptr, *ptr)
}

// 传值调用 - 不会修改原值
func incByValue(x int) {
	x++
}

// 传指针调用 - 会修改原值
func incByPointer(p *int) {
	*p++
}

// 返回局部变量的地址（Go会自动处理）
func newInt() *int {
	x := 42
	return &x
} 