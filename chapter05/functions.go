// Functions 演示函数的定义和使用
package main

import (
	"fmt"
	"math"
)

// 1. 基本函数定义
func greet(name string) string {
	return "Hello, " + name
}

// 2. 多参数函数
func add(a, b int) int {
	return a + b
}

// 3. 多返回值函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 4. 命名返回值
func minMax(a, b int) (min, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return // 裸返回
}

// 5. 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 6. 递归函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 7. 高阶函数：接收函数作为参数
func apply(f func(int) int, value int) int {
	return f(value)
}

// 8. 高阶函数：返回函数
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 9. 闭包示例
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 1. 基本函数调用
	fmt.Printf("基本函数调用:\n")
	result := greet("Alice")
	fmt.Printf("greet(\"Alice\"): %s\n", result)

	// 2. 多参数函数
	fmt.Printf("\n多参数函数:\n")
	sum_result := add(3, 4)
	fmt.Printf("add(3, 4): %d\n", sum_result)

	// 3. 多返回值函数
	fmt.Printf("\n多返回值函数:\n")
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("divide(10, 3): %.2f\n", quotient)
	}

	// 错误情况
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("divide(10, 0) 错误: %v\n", err)
	}

	// 4. 命名返回值
	fmt.Printf("\n命名返回值:\n")
	min, max := minMax(5, 3)
	fmt.Printf("minMax(5, 3): min=%d, max=%d\n", min, max)

	// 5. 可变参数函数
	fmt.Printf("\n可变参数函数:\n")
	fmt.Printf("sum(): %d\n", sum())
	fmt.Printf("sum(1): %d\n", sum(1))
	fmt.Printf("sum(1, 2, 3, 4, 5): %d\n", sum(1, 2, 3, 4, 5))

	// 传递切片
	numbers := []int{2, 4, 6, 8}
	fmt.Printf("sum(numbers...): %d\n", sum(numbers...))

	// 6. 递归函数
	fmt.Printf("\n递归函数:\n")
	fmt.Printf("factorial(5): %d\n", factorial(5))
	fmt.Printf("fibonacci(10): %d\n", fibonacci(10))

	// 7. 函数作为值
	fmt.Printf("\n函数作为值:\n")
	var fn func(int) int = func(x int) int { return x * x }
	fmt.Printf("平方函数(5): %d\n", fn(5))

	// 8. 高阶函数
	fmt.Printf("\n高阶函数:\n")
	square := func(x int) int { return x * x }
	fmt.Printf("apply(square, 5): %d\n", apply(square, 5))

	double := multiplier(2)
	triple := multiplier(3)
	fmt.Printf("double(5): %d\n", double(5))
	fmt.Printf("triple(5): %d\n", triple(5))

	// 9. 闭包
	fmt.Printf("\n闭包:\n")
	c1 := counter()
	c2 := counter()
	fmt.Printf("c1(): %d\n", c1())
	fmt.Printf("c1(): %d\n", c1())
	fmt.Printf("c2(): %d\n", c2())
	fmt.Printf("c1(): %d\n", c1())

	// 10. 匿名函数
	fmt.Printf("\n匿名函数:\n")
	func() {
		fmt.Println("立即执行的匿名函数")
	}()

	// 11. 函数类型
	fmt.Printf("\n函数类型:\n")
	type BinaryOp func(int, int) int
	var op BinaryOp = add
	fmt.Printf("op(3, 4): %d\n", op(3, 4))

	// 12. defer语句
	fmt.Printf("\ndefer语句:\n")
	deferExample()

	// 13. 函数的实际应用
	fmt.Printf("\n函数的实际应用:\n")
	numbers = []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("原数组: %v\n", numbers)
	sorted := bubbleSort(numbers)
	fmt.Printf("排序后: %v\n", sorted)

	// 使用函数实现过滤
	evens := filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)
}

func deferExample() {
	fmt.Println("开始")
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	fmt.Println("结束")
}

// 冒泡排序函数
func bubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)
	
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// 过滤函数
func filter(arr []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
} 