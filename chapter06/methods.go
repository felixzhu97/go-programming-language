// Methods 演示方法的定义和使用
package main

import (
	"fmt"
	"math"
)

// 1. 基本类型的方法
type Counter int

func (c *Counter) Increment() {
	*c++
}

func (c *Counter) Decrement() {
	*c--
}

func (c Counter) Value() int {
	return int(c)
}

// 2. 结构体方法
type Point struct {
	X, Y float64
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt((p.X-other.X)*(p.X-other.X) + (p.Y-other.Y)*(p.Y-other.Y))
}

func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%.2f, %.2f)", p.X, p.Y)
}

// 3. 银行账户示例
type Account struct {
	number  string
	balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("存款金额必须大于0")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("取款金额必须大于0")
	}
	if amount > a.balance {
		return fmt.Errorf("余额不足")
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a Account) Number() string {
	return a.number
}

// 4. 几何图形示例
type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Center: %s, Radius: %.2f}", c.Center, c.Radius)
}

// 5. 方法值和方法表达式
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d years old", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

// 6. 嵌入类型的方法
type Employee struct {
	Person
	ID     int
	Salary float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee{ID: %d, Name: %s, Age: %d, Salary: %.2f}", 
		e.ID, e.Name, e.Age, e.Salary)
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary *= (1 + percent/100)
}

func main() {
	// 1. 基本类型的方法
	fmt.Printf("基本类型的方法:\n")
	var c Counter = 5
	fmt.Printf("初始值: %d\n", c.Value())
	c.Increment()
	fmt.Printf("递增后: %d\n", c.Value())
	c.Decrement()
	fmt.Printf("递减后: %d\n", c.Value())

	// 2. 结构体方法
	fmt.Printf("\n结构体方法:\n")
	p1 := Point{3, 4}
	p2 := Point{0, 0}
	fmt.Printf("点1: %s\n", p1)
	fmt.Printf("点2: %s\n", p2)
	fmt.Printf("两点距离: %.2f\n", p1.Distance(p2))

	fmt.Printf("缩放前: %s\n", p1)
	p1.Scale(2)
	fmt.Printf("缩放后: %s\n", p1)

	// 3. 指针接收者 vs 值接收者
	fmt.Printf("\n指针接收者 vs 值接收者:\n")
	p3 := Point{1, 1}
	fmt.Printf("原点: %s\n", p3)
	
	// 值接收者：不会修改原值
	p3_copy := p3
	fmt.Printf("值方法调用前: %s\n", p3_copy)
	distance := p3_copy.Distance(Point{0, 0})
	fmt.Printf("距离: %.2f\n", distance)
	
	// 指针接收者：会修改原值
	fmt.Printf("指针方法调用前: %s\n", p3)
	p3.Scale(3)
	fmt.Printf("指针方法调用后: %s\n", p3)

	// 4. 银行账户示例
	fmt.Printf("\n银行账户示例:\n")
	account := Account{number: "123456", balance: 1000}
	fmt.Printf("账户: %s, 余额: %.2f\n", account.Number(), account.Balance())
	
	err := account.Deposit(500)
	if err != nil {
		fmt.Printf("存款错误: %v\n", err)
	} else {
		fmt.Printf("存款后余额: %.2f\n", account.Balance())
	}
	
	err = account.Withdraw(200)
	if err != nil {
		fmt.Printf("取款错误: %v\n", err)
	} else {
		fmt.Printf("取款后余额: %.2f\n", account.Balance())
	}
	
	err = account.Withdraw(2000)
	if err != nil {
		fmt.Printf("取款错误: %v\n", err)
	}

	// 5. 几何图形示例
	fmt.Printf("\n几何图形示例:\n")
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}
	fmt.Printf("圆形: %s\n", circle)
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Perimeter())
	
	circle.Scale(2)
	fmt.Printf("缩放后: %s\n", circle)
	fmt.Printf("新面积: %.2f\n", circle.Area())

	// 6. 方法值
	fmt.Printf("\n方法值:\n")
	person := Person{Name: "Alice", Age: 25}
	greetMethod := person.Greet // 方法值
	fmt.Printf("方法值调用: %s\n", greetMethod())
	
	// 方法表达式
	greetExpr := Person.Greet // 方法表达式
	fmt.Printf("方法表达式调用: %s\n", greetExpr(person))

	// 7. 嵌入类型的方法
	fmt.Printf("\n嵌入类型的方法:\n")
	emp := Employee{
		Person: Person{Name: "Bob", Age: 30},
		ID:     1001,
		Salary: 5000,
	}
	fmt.Printf("员工信息: %s\n", emp)
	fmt.Printf("员工问候: %s\n", emp.Greet()) // 继承的方法
	
	emp.GiveRaise(10)
	fmt.Printf("加薪后: %s\n", emp)

	// 8. 方法的可见性
	fmt.Printf("\n方法的可见性:\n")
	fmt.Printf("大写开头的方法（如Greet）是公开的\n")
	fmt.Printf("小写开头的方法（如greet）是私有的\n")

	// 9. nil接收者
	fmt.Printf("\nnil接收者:\n")
	var nilAccount *Account
	// 注意：调用nil接收者的方法会panic，实际代码中需要检查
	fmt.Printf("nil账户: %v\n", nilAccount)
	
	// 10. 方法的实际应用
	fmt.Printf("\n方法的实际应用:\n")
	calculator := &Calculator{}
	calculator.Add(10)
	calculator.Subtract(3)
	calculator.Multiply(2)
	calculator.Divide(2)
	fmt.Printf("计算结果: %.2f\n", calculator.Result())
	fmt.Printf("计算历史: %v\n", calculator.History())
}

// 计算器示例
type Calculator struct {
	result  float64
	history []string
}

func (c *Calculator) Add(value float64) {
	c.result += value
	c.history = append(c.history, fmt.Sprintf("+ %.2f", value))
}

func (c *Calculator) Subtract(value float64) {
	c.result -= value
	c.history = append(c.history, fmt.Sprintf("- %.2f", value))
}

func (c *Calculator) Multiply(value float64) {
	c.result *= value
	c.history = append(c.history, fmt.Sprintf("* %.2f", value))
}

func (c *Calculator) Divide(value float64) {
	if value != 0 {
		c.result /= value
		c.history = append(c.history, fmt.Sprintf("/ %.2f", value))
	}
}

func (c Calculator) Result() float64 {
	return c.result
}

func (c Calculator) History() []string {
	return c.history
} 