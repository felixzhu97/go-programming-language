// Package mypackage 提供了一些数学计算功能。
//
// 这是一个示例包，演示了Go语言包的基本结构和文档规范。
//
// 示例用法：
//    result := mypackage.Add(3, 4)
//    fmt.Println(result) // 输出: 7
package mypackage

import (
	"errors"
	"fmt"
)

// 包级别常量
const (
	// MaxInt 表示支持的最大整数值
	MaxInt = 1000000
	// MinInt 表示支持的最小整数值
	MinInt = -1000000
)

// 包级别变量
var (
	// DefaultPrecision 默认精度
	DefaultPrecision = 2
	// operations 记录操作次数（私有变量）
	operations int
)

// 包初始化函数
func init() {
	fmt.Println("mypackage: 包初始化完成")
	operations = 0
}

// Add 计算两个整数的和。
//
// 这个函数演示了基本的加法运算，包括输入验证和错误处理。
//
// 参数：
//   a - 第一个整数
//   b - 第二个整数
//
// 返回值：
//   int - 两个整数的和
//   error - 如果发生错误返回错误信息
//
// 示例：
//    result, err := mypackage.Add(3, 4)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(result) // 输出: 7
func Add(a, b int) (int, error) {
	// 输入验证
	if a < MinInt || a > MaxInt {
		return 0, errors.New("参数a超出范围")
	}
	if b < MinInt || b > MaxInt {
		return 0, errors.New("参数b超出范围")
	}
	
	// 溢出检查
	if a > 0 && b > 0 && a > MaxInt-b {
		return 0, errors.New("计算结果溢出")
	}
	if a < 0 && b < 0 && a < MinInt-b {
		return 0, errors.New("计算结果溢出")
	}
	
	operations++
	return a + b, nil
}

// Multiply 计算两个整数的乘积。
//
// 这个函数演示了乘法运算，包括溢出检查。
//
// 参数：
//   a - 第一个整数
//   b - 第二个整数
//
// 返回值：
//   int - 两个整数的乘积
//   error - 如果发生错误返回错误信息
func Multiply(a, b int) (int, error) {
	// 输入验证
	if a < MinInt || a > MaxInt {
		return 0, errors.New("参数a超出范围")
	}
	if b < MinInt || b > MaxInt {
		return 0, errors.New("参数b超出范围")
	}
	
	// 特殊情况
	if a == 0 || b == 0 {
		operations++
		return 0, nil
	}
	
	// 溢出检查
	if a > 0 && b > 0 && a > MaxInt/b {
		return 0, errors.New("计算结果溢出")
	}
	if a < 0 && b < 0 && a < MaxInt/b {
		return 0, errors.New("计算结果溢出")
	}
	if (a > 0 && b < 0 && b < MinInt/a) || (a < 0 && b > 0 && a < MinInt/b) {
		return 0, errors.New("计算结果溢出")
	}
	
	operations++
	return a * b, nil
}

// GetOperationCount 返回已执行的操作次数。
//
// 这个函数演示了访问包级别私有变量的方法。
//
// 返回值：
//   int - 已执行的操作次数
func GetOperationCount() int {
	return operations
}

// ResetOperationCount 重置操作次数计数器。
//
// 这个函数演示了修改包级别私有变量的方法。
func ResetOperationCount() {
	operations = 0
}

// Calculator 计算器结构体，演示了面向对象的设计。
type Calculator struct {
	// Name 计算器名称
	Name string
	// precision 精度（私有字段）
	precision int
}

// NewCalculator 创建一个新的计算器实例。
//
// 这是一个构造函数，演示了Go语言中创建对象的常用模式。
//
// 参数：
//   name - 计算器名称
//
// 返回值：
//   *Calculator - 计算器实例指针
func NewCalculator(name string) *Calculator {
	return &Calculator{
		Name:      name,
		precision: DefaultPrecision,
	}
}

// SetPrecision 设置计算器的精度。
//
// 这个方法演示了结构体方法的定义和使用。
//
// 参数：
//   precision - 新的精度值
func (c *Calculator) SetPrecision(precision int) {
	c.precision = precision
}

// GetPrecision 获取计算器的精度。
//
// 返回值：
//   int - 当前精度值
func (c *Calculator) GetPrecision() int {
	return c.precision
}

// Calculate 执行计算操作。
//
// 这个方法演示了结构体方法如何使用包级别函数。
//
// 参数：
//   a - 第一个操作数
//   b - 第二个操作数
//   operation - 操作类型 ("add" 或 "multiply")
//
// 返回值：
//   int - 计算结果
//   error - 如果发生错误返回错误信息
func (c *Calculator) Calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "add":
		return Add(a, b)
	case "multiply":
		return Multiply(a, b)
	default:
		return 0, errors.New("不支持的操作")
	}
}

// String 返回计算器的字符串表示。
//
// 这个方法实现了fmt.Stringer接口，演示了接口的使用。
//
// 返回值：
//   string - 计算器的字符串表示
func (c *Calculator) String() string {
	return fmt.Sprintf("Calculator{Name: %s, Precision: %d}", c.Name, c.precision)
}

// 私有辅助函数，演示了包内部的实现细节
func validateInput(value int) bool {
	return value >= MinInt && value <= MaxInt
} 