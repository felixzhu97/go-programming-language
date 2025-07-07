package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

// Calculator 计算器结构体，用于演示测试
type Calculator struct {
	history []string
}

// NewCalculator 创建一个新的计算器
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

// Add 加法运算
func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.history = append(c.history, fmt.Sprintf("%d + %d = %d", a, b, result))
	return result
}

// Subtract 减法运算
func (c *Calculator) Subtract(a, b int) int {
	result := a - b
	c.history = append(c.history, fmt.Sprintf("%d - %d = %d", a, b, result))
	return result
}

// Multiply 乘法运算
func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.history = append(c.history, fmt.Sprintf("%d * %d = %d", a, b, result))
	return result
}

// Divide 除法运算
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	c.history = append(c.history, fmt.Sprintf("%d / %d = %d", a, b, result))
	return result, nil
}

// GetHistory 获取计算历史
func (c *Calculator) GetHistory() []string {
	return c.history
}

// Clear 清空计算历史
func (c *Calculator) Clear() {
	c.history = make([]string, 0)
}

// 数学工具函数

// IsPrime 判断是否为质数
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci 计算斐波那契数列第n项
func Fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Factorial 计算阶乘
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// 字符串工具函数

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 判断是否为回文
func IsPalindrome(s string) bool {
	// 转换为小写并去除非字母数字字符
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == ReverseString(cleaned)
}

// CountWords 计算单词数量
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// 排序工具函数

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)
	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort 快速排序
func QuickSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// BinarySearch 二分查找
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 慢速函数，用于基准测试
func SlowFunction(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += i * j
		}
	}
	return sum
}

// FastFunction 快速版本
func FastFunction(n int) int {
	if n == 0 {
		return 0
	}
	sum1 := n * (n - 1) / 2
	return sum1 * sum1
}

// 数据结构：Stack
type Stack struct {
	items []int
}

// Push 压栈
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 获取栈大小
func (s *Stack) Size() int {
	return len(s.items)
}

// 用于示例测试的函数
func Greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// 用于演示表格驱动测试的函数
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 用于演示并发测试的函数
func ProcessConcurrently(data []int, workers int) []int {
	if len(data) == 0 {
		return data
	}
	
	result := make([]int, len(data))
	copy(result, data)
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println("《Go程序设计语言》第11章：测试")
	fmt.Println("这个文件包含了用于测试的各种函数和数据结构")
	fmt.Println("请运行 'go test' 来执行测试")
	
	// 演示一些功能
	calc := NewCalculator()
	result := calc.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	fmt.Printf("7 是质数吗？%t\n", IsPrime(7))
	fmt.Printf("斐波那契数列第10项：%d\n", Fibonacci(10))
	
	factorial, _ := Factorial(5)
	fmt.Printf("5! = %d\n", factorial)
} 