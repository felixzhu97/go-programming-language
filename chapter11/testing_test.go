package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试函数必须以Test开头，并接收 *testing.T 参数

// TestCalculator 测试计算器基本功能
func TestCalculator(t *testing.T) {
	calc := NewCalculator()
	
	// 测试加法
	result := calc.Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
	
	// 测试减法
	result = calc.Subtract(5, 2)
	if result != 3 {
		t.Errorf("Subtract(5, 2) = %d; want 3", result)
	}
	
	// 测试乘法
	result = calc.Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply(3, 4) = %d; want 12", result)
	}
	
	// 测试除法
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}
	
	// 测试除零错误
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}
	
	// 测试历史记录
	history := calc.GetHistory()
	if len(history) != 4 {
		t.Errorf("History length = %d; want 4", len(history))
	}
}

// TestIsPrime 测试质数判断
func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{17, true},
		{25, false},
		{29, true},
	}
	
	for _, test := range tests {
		result := IsPrime(test.input)
		if result != test.expected {
			t.Errorf("IsPrime(%d) = %t; want %t", test.input, result, test.expected)
		}
	}
}

// TestFibonacci 测试斐波那契数列
func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{10, 55},
	}
	
	for _, test := range tests {
		result := Fibonacci(test.input)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

// TestFactorial 测试阶乘计算
func TestFactorial(t *testing.T) {
	tests := []struct {
		input       int
		expected    int
		expectError bool
	}{
		{0, 1, false},
		{1, 1, false},
		{2, 2, false},
		{3, 6, false},
		{4, 24, false},
		{5, 120, false},
		{-1, 0, true},
		{-5, 0, true},
	}
	
	for _, test := range tests {
		result, err := Factorial(test.input)
		
		if test.expectError {
			if err == nil {
				t.Errorf("Factorial(%d) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Factorial(%d) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Factorial(%d) = %d; want %d", test.input, result, test.expected)
			}
		}
	}
}

// TestReverseString 测试字符串反转
func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"hello", "olleh"},
		{"Go测试", "试测oG"},
		{"12345", "54321"},
	}
	
	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// TestIsPalindrome 测试回文判断
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abc", false},
		{"A man a plan a canal Panama", true},
		{"race a car", false},
		{"hello", false},
		{"level", true},
		{"noon", true},
	}
	
	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %t; want %t", test.input, result, test.expected)
		}
	}
}

// TestBubbleSort 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}
	
	for _, test := range tests {
		result := BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestQuickSort 测试快速排序
func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}
	
	for _, test := range tests {
		result := QuickSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("QuickSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestBinarySearch 测试二分查找
func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{5, 4},
		{10, 9},
		{3, 2},
		{11, -1},
		{0, -1},
	}
	
	for _, test := range tests {
		result := BinarySearch(arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", arr, test.target, result, test.expected)
		}
	}
}

// TestStack 测试栈数据结构
func TestStack(t *testing.T) {
	stack := &Stack{}
	
	// 测试空栈
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}
	
	if stack.Size() != 0 {
		t.Errorf("Empty stack size = %d; want 0", stack.Size())
	}
	
	// 测试出栈错误
	_, err := stack.Pop()
	if err == nil {
		t.Error("Pop from empty stack should return error")
	}
	
	// 测试压栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	if stack.IsEmpty() {
		t.Error("Stack with items should not be empty")
	}
	
	if stack.Size() != 3 {
		t.Errorf("Stack size = %d; want 3", stack.Size())
	}
	
	// 测试出栈
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop returned error: %v", err)
	}
	if item != 3 {
		t.Errorf("Pop() = %d; want 3", item)
	}
	
	if stack.Size() != 2 {
		t.Errorf("Stack size after pop = %d; want 2", stack.Size())
	}
}

// TestMax 演示表格驱动测试
func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 5},
		{"negative numbers", -2, -5, -2},
		{"mixed numbers", -3, 2, 2},
		{"equal numbers", 4, 4, 4},
		{"zero and positive", 0, 5, 5},
		{"zero and negative", 0, -3, 0},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Max(test.a, test.b)
			if result != test.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

// 基准测试函数必须以Benchmark开头，并接收 *testing.B 参数

// BenchmarkBubbleSort 基准测试冒泡排序
func BenchmarkBubbleSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	
	b.ResetTimer() // 重置计时器，排除准备时间
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}

// BenchmarkQuickSort 基准测试快速排序
func BenchmarkQuickSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}

// BenchmarkSlowFunction 基准测试慢速函数
func BenchmarkSlowFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowFunction(100)
	}
}

// BenchmarkFastFunction 基准测试快速函数
func BenchmarkFastFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFunction(100)
	}
}

// BenchmarkFibonacci 基准测试斐波那契数列
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

// 并行基准测试
func BenchmarkFibonacciParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(20)
		}
	})
}

// 示例测试函数必须以Example开头

// ExampleGreet 示例测试问候函数
func ExampleGreet() {
	fmt.Println(Greet("World"))
	// Output: Hello, World!
}

// ExampleGreet_empty 示例测试空名称
func ExampleGreet_empty() {
	fmt.Println(Greet(""))
	// Output: Hello, World!
}

// ExampleCalculator_Add 示例测试计算器加法
func ExampleCalculator_Add() {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// ExampleIsPrime 示例测试质数判断
func ExampleIsPrime() {
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(8))
	// Output:
	// true
	// false
}

// ExampleReverseString 示例测试字符串反转
func ExampleReverseString() {
	fmt.Println(ReverseString("hello"))
	// Output: olleh
}

// ExampleBubbleSort 示例测试冒泡排序
func ExampleBubbleSort() {
	data := []int{3, 1, 4, 1, 5}
	sorted := BubbleSort(data)
	fmt.Println(sorted)
	// Output: [1 1 3 4 5]
}

// 子测试示例
func TestMathFunctions(t *testing.T) {
	t.Run("IsPrime", func(t *testing.T) {
		if !IsPrime(7) {
			t.Error("7 should be prime")
		}
		if IsPrime(8) {
			t.Error("8 should not be prime")
		}
	})
	
	t.Run("Fibonacci", func(t *testing.T) {
		if Fibonacci(5) != 5 {
			t.Error("Fibonacci(5) should be 5")
		}
		if Fibonacci(10) != 55 {
			t.Error("Fibonacci(10) should be 55")
		}
	})
	
	t.Run("Factorial", func(t *testing.T) {
		result, err := Factorial(5)
		if err != nil {
			t.Errorf("Factorial(5) returned error: %v", err)
		}
		if result != 120 {
			t.Errorf("Factorial(5) = %d; want 120", result)
		}
	})
}

// 测试辅助函数
func TestHelperFunction(t *testing.T) {
	t.Helper() // 标记这是一个辅助函数
	
	// 这里可以包含一些通用的测试逻辑
	calc := NewCalculator()
	if calc == nil {
		t.Fatal("NewCalculator() returned nil")
	}
}

// 测试清理函数
func TestWithCleanup(t *testing.T) {
	// 设置测试环境
	calc := NewCalculator()
	
	// 注册清理函数
	t.Cleanup(func() {
		calc.Clear()
	})
	
	// 执行测试
	calc.Add(1, 2)
	calc.Add(3, 4)
	
	history := calc.GetHistory()
	if len(history) != 2 {
		t.Errorf("History length = %d; want 2", len(history))
	}
	
	// 清理函数会在测试结束时自动调用
}

// 跳过测试示例
func TestSkipExample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	
	// 这里是耗时的测试代码
	// 只有在非 short 模式下才会执行
}

// 并行测试示例
func TestParallelExample(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"small", 10},
		{"medium", 100},
		{"large", 1000},
	}
	
	for _, test := range tests {
		test := test // 捕获循环变量
		t.Run(test.name, func(t *testing.T) {
			t.Parallel() // 标记为并行测试
			
			result := Fibonacci(test.n)
			if result < 0 {
				t.Errorf("Fibonacci(%d) returned negative number: %d", test.n, result)
			}
		})
	}
} 