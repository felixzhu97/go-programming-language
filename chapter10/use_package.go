package main

import (
	"fmt"
	"log"
	
	// 导入本地包
	"./mypackage"
)

func main() {
	fmt.Println("=== 使用自定义包示例 ===")
	
	// 使用包级别函数
	result, err := mypackage.Add(10, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("10 + 20 = %d\n", result)
	
	// 使用包级别常量
	fmt.Printf("支持的最大值: %d\n", mypackage.MaxInt)
	fmt.Printf("支持的最小值: %d\n", mypackage.MinInt)
	
	// 使用包级别变量
	fmt.Printf("默认精度: %d\n", mypackage.DefaultPrecision)
	
	// 测试多个操作
	multiply, err := mypackage.Multiply(5, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("5 * 6 = %d\n", multiply)
	
	// 查看操作次数
	fmt.Printf("已执行操作次数: %d\n", mypackage.GetOperationCount())
	
	// 使用结构体
	fmt.Println("\n=== 使用Calculator结构体 ===")
	calc := mypackage.NewCalculator("My Calculator")
	fmt.Printf("计算器: %s\n", calc)
	
	// 设置精度
	calc.SetPrecision(5)
	fmt.Printf("设置精度后: %s\n", calc)
	
	// 使用计算器进行计算
	calcResult, err := calc.Calculate(8, 9, "add")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("计算器计算 8 + 9 = %d\n", calcResult)
	
	calcResult, err = calc.Calculate(7, 8, "multiply")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("计算器计算 7 * 8 = %d\n", calcResult)
	
	// 最终操作次数
	fmt.Printf("最终操作次数: %d\n", mypackage.GetOperationCount())
	
	// 重置计数器
	mypackage.ResetOperationCount()
	fmt.Printf("重置后操作次数: %d\n", mypackage.GetOperationCount())
} 