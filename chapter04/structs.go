// Structs 演示结构体类型和操作
package main

import (
	"fmt"
	"time"
)

// 1. 基本结构体定义
type Person struct {
	Name string
	Age  int
	City string
}

// 2. 带标签的结构体
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Salary   float64 `json:"salary"`
	IsActive bool   `json:"is_active"`
}

// 3. 嵌套结构体
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

type Customer struct {
	Person
	Address Address
	Orders  []Order
}

type Order struct {
	ID     int
	Amount float64
	Date   time.Time
}

// 4. 匿名结构体
var config = struct {
	Host string
	Port int
	SSL  bool
}{
	Host: "localhost",
	Port: 8080,
	SSL:  false,
}

// 5. 结构体方法
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, City: %s}", 
		p.Name, p.Age, p.City)
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s from %s", p.Name, p.City)
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary *= (1 + percent/100)
}

func (e Employee) GetAnnualSalary() float64 {
	return e.Salary * 12
}

func main() {
	// 1. 结构体声明和初始化
	fmt.Printf("结构体声明和初始化:\n")
	var p1 Person // 零值结构体
	fmt.Printf("零值结构体: %+v\n", p1)
	
	p2 := Person{
		Name: "Alice",
		Age:  25,
		City: "New York",
	}
	fmt.Printf("字面量初始化: %+v\n", p2)
	
	p3 := Person{"Bob", 30, "London"} // 按顺序初始化
	fmt.Printf("按顺序初始化: %+v\n", p3)

	// 2. 结构体字段访问
	fmt.Printf("\n结构体字段访问:\n")
	fmt.Printf("p2.Name: %s\n", p2.Name)
	fmt.Printf("p2.Age: %d\n", p2.Age)
	
	// 修改字段
	p2.Age = 26
	fmt.Printf("修改后 p2.Age: %d\n", p2.Age)

	// 3. 结构体指针
	fmt.Printf("\n结构体指针:\n")
	pp := &p2
	fmt.Printf("指针: %p\n", pp)
	fmt.Printf("通过指针访问: %s\n", pp.Name)
	
	// 通过指针修改
	pp.Name = "Alice Smith"
	fmt.Printf("通过指针修改: %s\n", p2.Name)

	// 4. 结构体比较
	fmt.Printf("\n结构体比较:\n")
	p4 := Person{"Charlie", 35, "Paris"}
	p5 := Person{"Charlie", 35, "Paris"}
	fmt.Printf("p4 == p5: %t\n", p4 == p5)
	
	p6 := Person{"Charlie", 36, "Paris"}
	fmt.Printf("p4 == p6: %t\n", p4 == p6)

	// 5. 结构体方法调用
	fmt.Printf("\n结构体方法调用:\n")
	fmt.Printf("p2.String(): %s\n", p2.String())
	fmt.Printf("p2.Greet(): %s\n", p2.Greet())

	// 6. 员工结构体示例
	fmt.Printf("\n员工结构体示例:\n")
	emp := Employee{
		ID:       1001,
		Name:     "John Doe",
		Email:    "john@example.com",
		Salary:   5000.0,
		IsActive: true,
	}
	fmt.Printf("员工信息: %+v\n", emp)
	fmt.Printf("年薪: %.2f\n", emp.GetAnnualSalary())
	
	emp.GiveRaise(10)
	fmt.Printf("加薪10%%后月薪: %.2f\n", emp.Salary)

	// 7. 嵌套结构体
	fmt.Printf("\n嵌套结构体:\n")
	customer := Customer{
		Person: Person{
			Name: "Jane Smith",
			Age:  28,
			City: "Seattle",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Seattle",
			Country: "USA",
			ZipCode: "98101",
		},
		Orders: []Order{
			{ID: 1, Amount: 99.99, Date: time.Now()},
			{ID: 2, Amount: 149.99, Date: time.Now().AddDate(0, 0, -1)},
		},
	}
	fmt.Printf("客户信息: %+v\n", customer)
	fmt.Printf("客户姓名: %s\n", customer.Name) // 嵌入字段
	fmt.Printf("客户地址: %s, %s\n", customer.Address.Street, customer.Address.City)
	fmt.Printf("订单数量: %d\n", len(customer.Orders))

	// 8. 匿名结构体
	fmt.Printf("\n匿名结构体:\n")
	point := struct {
		X, Y int
	}{10, 20}
	fmt.Printf("点坐标: %+v\n", point)
	fmt.Printf("配置: %+v\n", config)

	// 9. 结构体切片
	fmt.Printf("\n结构体切片:\n")
	people := []Person{
		{"Alice", 25, "NYC"},
		{"Bob", 30, "LA"},
		{"Charlie", 35, "Chicago"},
	}
	fmt.Printf("人员列表:\n")
	for i, person := range people {
		fmt.Printf("%d: %s\n", i+1, person.Greet())
	}

	// 10. 结构体映射
	fmt.Printf("\n结构体映射:\n")
	employees := map[int]Employee{
		1001: {ID: 1001, Name: "John", Email: "john@example.com", Salary: 5000},
		1002: {ID: 1002, Name: "Jane", Email: "jane@example.com", Salary: 5500},
	}
	
	for id, emp := range employees {
		fmt.Printf("员工ID %d: %s, 薪水: %.2f\n", id, emp.Name, emp.Salary)
	}

	// 11. 结构体作为函数参数
	fmt.Printf("\n结构体作为函数参数:\n")
	fmt.Printf("值传递前: %+v\n", p2)
	modifyPersonByValue(p2)
	fmt.Printf("值传递后: %+v\n", p2)
	
	fmt.Printf("指针传递前: %+v\n", p2)
	modifyPersonByPointer(&p2)
	fmt.Printf("指针传递后: %+v\n", p2)

	// 12. 结构体的零值
	fmt.Printf("\n结构体的零值:\n")
	var empty Employee
	fmt.Printf("空员工: %+v\n", empty)
	fmt.Printf("空员工姓名为空: %t\n", empty.Name == "")

	// 13. 结构体字段的可见性
	fmt.Printf("\n结构体字段的可见性:\n")
	fmt.Printf("大写字段名（如Name）是公开的，可以被其他包访问\n")
	fmt.Printf("小写字段名（如age）是私有的，只能在同一包内访问\n")

	// 14. 结构体的实际应用
	fmt.Printf("\n结构体的实际应用:\n")
	// 创建一个简单的数据库记录
	users := []struct {
		ID       int
		Username string
		Email    string
		Created  time.Time
	}{
		{1, "alice", "alice@example.com", time.Now()},
		{2, "bob", "bob@example.com", time.Now().AddDate(0, 0, -1)},
	}
	
	fmt.Printf("用户记录:\n")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", 
			user.ID, user.Username, user.Email)
	}
}

// 值传递：不会修改原结构体
func modifyPersonByValue(p Person) {
	p.Age = 100
	fmt.Printf("函数内修改: %+v\n", p)
}

// 指针传递：会修改原结构体
func modifyPersonByPointer(p *Person) {
	p.Age = 100
	fmt.Printf("函数内修改: %+v\n", *p)
} 