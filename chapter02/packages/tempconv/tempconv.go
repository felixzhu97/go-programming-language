// Package tempconv 提供摄氏度和华氏度的温度转换功能
package tempconv

import "fmt"

// Celsius 摄氏度类型
type Celsius float64

// Fahrenheit 华氏度类型
type Fahrenheit float64

// 温度常量
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// String 方法让Celsius满足fmt.Stringer接口
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

// String 方法让Fahrenheit满足fmt.Stringer接口
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", float64(f))
}

// CToF 将摄氏度转换为华氏度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 将华氏度转换为摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK 将摄氏度转换为开氏度
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KToC 将开氏度转换为摄氏度
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// Kelvin 开氏度类型
type Kelvin float64

// String 方法让Kelvin满足fmt.Stringer接口
func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", float64(k))
} 