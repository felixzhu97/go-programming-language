// Interfaces 演示接口的定义和使用
package main

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

// 1. 基本接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. 实现接口的类型
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 3. 接口组合
type Drawable interface {
	Draw()
}

type Movable interface {
	Move(x, y float64)
}

type GameObject interface {
	Shape
	Drawable
	Movable
}

// 4. 实现GameObject接口
type Player struct {
	Name   string
	X, Y   float64
	Radius float64
}

func (p Player) Area() float64 {
	return math.Pi * p.Radius * p.Radius
}

func (p Player) Perimeter() float64 {
	return 2 * math.Pi * p.Radius
}

func (p Player) Draw() {
	fmt.Printf("绘制玩家 %s 在位置 (%.2f, %.2f)\n", p.Name, p.X, p.Y)
}

func (p *Player) Move(x, y float64) {
	p.X += x
	p.Y += y
}

// 5. 空接口
func printAnything(value interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", value, value)
}

// 6. 类型断言
func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case bool:
		fmt.Printf("布尔值: %t\n", v)
	case Shape:
		fmt.Printf("图形，面积: %.2f\n", v.Area())
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 7. Writer接口示例
type Logger struct {
	prefix string
}

func (l Logger) Write(data []byte) (int, error) {
	output := fmt.Sprintf("[%s] %s", l.prefix, string(data))
	fmt.Print(output)
	return len(output), nil
}

// 8. Stringer接口
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d岁)", p.Name, p.Age)
}

// 9. sort.Interface示例
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// 10. 错误接口
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.Code, e.Message)
}

func riskyOperation() error {
	return CustomError{Code: 500, Message: "内部服务器错误"}
}

func main() {
	// 1. 基本接口使用
	fmt.Printf("基本接口使用:\n")
	var shapes []Shape = []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
	}
	
	for i, shape := range shapes {
		fmt.Printf("图形 %d: 面积=%.2f, 周长=%.2f\n", 
			i+1, shape.Area(), shape.Perimeter())
	}

	// 2. 接口的多态性
	fmt.Printf("\n接口的多态性:\n")
	printShapeInfo := func(s Shape) {
		fmt.Printf("图形信息 - 面积: %.2f, 周长: %.2f, 类型: %T\n", 
			s.Area(), s.Perimeter(), s)
	}
	
	rect := Rectangle{Width: 4, Height: 6}
	circle := Circle{Radius: 2.5}
	printShapeInfo(rect)
	printShapeInfo(circle)

	// 3. 接口组合
	fmt.Printf("\n接口组合:\n")
	player := &Player{Name: "Alice", X: 0, Y: 0, Radius: 1}
	player.Draw()
	player.Move(10, 5)
	player.Draw()
	fmt.Printf("玩家面积: %.2f\n", player.Area())

	// 4. 空接口
	fmt.Printf("\n空接口:\n")
	printAnything(42)
	printAnything("Hello")
	printAnything(true)
	printAnything(rect)

	// 5. 类型断言
	fmt.Printf("\n类型断言:\n")
	var value interface{} = "Hello, World!"
	if str, ok := value.(string); ok {
		fmt.Printf("字符串长度: %d\n", len(str))
	}

	// 安全类型断言
	if num, ok := value.(int); ok {
		fmt.Printf("整数: %d\n", num)
	} else {
		fmt.Printf("不是整数类型\n")
	}

	// 6. 类型开关
	fmt.Printf("\n类型开关:\n")
	values := []interface{}{42, "hello", true, Rectangle{3, 4}, 3.14}
	for _, v := range values {
		processValue(v)
	}

	// 7. Writer接口
	fmt.Printf("\nWriter接口:\n")
	logger := Logger{prefix: "INFO"}
	io.WriteString(&logger, "这是一条日志消息\n")

	// 8. Stringer接口
	fmt.Printf("\nStringer接口:\n")
	person := Person{Name: "Bob", Age: 30}
	fmt.Printf("Person: %s\n", person) // 自动调用String()方法

	// 9. sort.Interface
	fmt.Printf("\nsort.Interface:\n")
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}
	fmt.Printf("排序前: %v\n", people)
	sort.Sort(ByAge(people))
	fmt.Printf("按年龄排序后: %v\n", people)

	// 10. 错误接口
	fmt.Printf("\n错误接口:\n")
	if err := riskyOperation(); err != nil {
		fmt.Printf("操作失败: %v\n", err)
		if customErr, ok := err.(CustomError); ok {
			fmt.Printf("错误代码: %d\n", customErr.Code)
		}
	}

	// 11. 接口的零值
	fmt.Printf("\n接口的零值:\n")
	var shape Shape
	fmt.Printf("零值接口: %v\n", shape)
	if shape == nil {
		fmt.Printf("接口为nil\n")
	}

	// 12. 接口的动态类型和动态值
	fmt.Printf("\n接口的动态类型和动态值:\n")
	var w io.Writer
	fmt.Printf("空接口: %T, %v\n", w, w)
	
	w = &strings.Builder{}
	fmt.Printf("赋值后: %T, %v\n", w, w)

	// 13. 接口的实际应用
	fmt.Printf("\n接口的实际应用:\n")
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}
	
	for _, animal := range animals {
		animal.Speak()
		animal.Move()
	}
}

// 动物接口示例
type Animal interface {
	Speak()
	Move()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Printf("%s 说: 汪汪!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s 跑来跑去\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("%s 说: 喵喵!\n", c.Name)
}

func (c Cat) Move() {
	fmt.Printf("%s 悄悄地走\n", c.Name)
}

type Bird struct {
	Name string
}

func (b Bird) Speak() {
	fmt.Printf("%s 说: 啾啾!\n", b.Name)
}

func (b Bird) Move() {
	fmt.Printf("%s 飞来飞去\n", b.Name)
} 