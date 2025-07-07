// Scope 演示变量作用域
package main

import "fmt"

// 1. 包级作用域
var globalVar = "全局变量"

// 2. 函数作用域
func main() {
	// 局部变量
	var localVar = "局部变量"
	fmt.Printf("函数内: %s, %s\n", globalVar, localVar)

	// 3. 块作用域
	{
		blockVar := "块变量"
		fmt.Printf("块内: %s, %s, %s\n", globalVar, localVar, blockVar)
		
		// 遮蔽外层变量
		globalVar := "遮蔽的全局变量"
		fmt.Printf("遮蔽后: %s\n", globalVar)
	}
	// blockVar 在这里不可访问
	fmt.Printf("块外: %s, %s\n", globalVar, localVar)

	// 4. if 语句作用域
	if condition := true; condition {
		ifVar := "if变量"
		fmt.Printf("if块内: %s\n", ifVar)
	}
	// condition 和 ifVar 在这里不可访问

	// 5. for 循环作用域
	for i := 0; i < 3; i++ {
		loopVar := fmt.Sprintf("循环变量_%d", i)
		fmt.Printf("循环内: %s\n", loopVar)
	}
	// i 和 loopVar 在这里不可访问

	// 6. 短变量声明的作用域
	x := 1
	fmt.Printf("外层 x: %d\n", x)
	{
		x := 2 // 新的变量，遮蔽外层的x
		fmt.Printf("内层 x: %d\n", x)
	}
	fmt.Printf("外层 x: %d\n", x) // 外层x不变

	// 7. 函数调用和作用域
	testFunctionScope()
	fmt.Printf("main中的全局变量: %s\n", globalVar)
}

// 函数有自己的作用域
func testFunctionScope() {
	// 可以访问包级变量
	fmt.Printf("函数中的全局变量: %s\n", globalVar)
	
	// 局部变量
	localVar := "函数局部变量"
	fmt.Printf("函数局部变量: %s\n", localVar)
	
	// 遮蔽包级变量
	globalVar := "函数内遮蔽的全局变量"
	fmt.Printf("遮蔽后: %s\n", globalVar)
} 