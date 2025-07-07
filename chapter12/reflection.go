package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 示例结构体
type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email" validate:"email"`
	private string // 私有字段
}

// 为Person定义一些方法
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", p.Name, p.Age, p.Email)
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) GetPrivate() string {
	return p.private
}

func (p *Person) SetPrivate(value string) {
	p.private = value
}

// 示例1：基本反射操作
func basicReflectionExample() {
	fmt.Println("=== 基本反射示例 ===")
	
	// 创建一个Person实例
	person := Person{
		Name:    "Alice",
		Age:     25,
		Email:   "alice@example.com",
		private: "secret",
	}
	
	// 获取类型信息
	t := reflect.TypeOf(person)
	fmt.Printf("类型：%s\n", t)
	fmt.Printf("类型名称：%s\n", t.Name())
	fmt.Printf("包路径：%s\n", t.PkgPath())
	fmt.Printf("类型种类：%s\n", t.Kind())
	
	// 获取值信息
	v := reflect.ValueOf(person)
	fmt.Printf("值：%v\n", v)
	fmt.Printf("值类型：%s\n", v.Type())
	fmt.Printf("值种类：%s\n", v.Kind())
	fmt.Printf("是否有效：%t\n", v.IsValid())
	fmt.Printf("是否为零值：%t\n", v.IsZero())
}

// 示例2：结构体字段反射
func structFieldReflectionExample() {
	fmt.Println("\n=== 结构体字段反射示例 ===")
	
	person := Person{
		Name:    "Bob",
		Age:     30,
		Email:   "bob@example.com",
		private: "confidential",
	}
	
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	
	fmt.Printf("结构体有 %d 个字段：\n", t.NumField())
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		fmt.Printf("字段 %d:\n", i)
		fmt.Printf("  名称：%s\n", field.Name)
		fmt.Printf("  类型：%s\n", field.Type)
		fmt.Printf("  标签：%s\n", field.Tag)
		fmt.Printf("  偏移量：%d\n", field.Offset)
		fmt.Printf("  是否导出：%t\n", field.IsExported())
		
		// 获取标签信息
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("  JSON标签：%s\n", jsonTag)
		}
		if validateTag := field.Tag.Get("validate"); validateTag != "" {
			fmt.Printf("  验证标签：%s\n", validateTag)
		}
		
		// 获取字段值
		if value.CanInterface() {
			fmt.Printf("  值：%v\n", value.Interface())
		} else {
			fmt.Printf("  值：<不可访问>\n")
		}
		fmt.Println()
	}
}

// 示例3：方法反射
func methodReflectionExample() {
	fmt.Println("=== 方法反射示例 ===")
	
	person := Person{
		Name:  "Charlie",
		Age:   28,
		Email: "charlie@example.com",
	}
	
	// 值类型的方法
	v := reflect.ValueOf(person)
	t := reflect.TypeOf(person)
	
	fmt.Printf("值类型有 %d 个方法：\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("方法 %d: %s\n", i, method.Name)
		fmt.Printf("  类型：%s\n", method.Type)
		fmt.Printf("  函数：%v\n", method.Func)
		
		// 调用方法
		if method.Name == "String" {
			result := v.Method(i).Call(nil)
			fmt.Printf("  调用结果：%s\n", result[0].String())
		}
		if method.Name == "Greet" {
			result := v.Method(i).Call(nil)
			fmt.Printf("  调用结果：%s\n", result[0].String())
		}
	}
	
	// 指针类型的方法
	fmt.Println("\n指针类型的方法：")
	pv := reflect.ValueOf(&person)
	pt := reflect.TypeOf(&person)
	
	fmt.Printf("指针类型有 %d 个方法：\n", pt.NumMethod())
	for i := 0; i < pt.NumMethod(); i++ {
		method := pt.Method(i)
		fmt.Printf("方法 %d: %s\n", i, method.Name)
		
		// 调用SetAge方法
		if method.Name == "SetAge" {
			args := []reflect.Value{reflect.ValueOf(35)}
			pv.Method(i).Call(args)
			fmt.Printf("  设置年龄为35后：%s\n", person.String())
		}
	}
}

// 示例4：动态创建和修改值
func dynamicValueExample() {
	fmt.Println("\n=== 动态创建和修改值示例 ===")
	
	// 创建一个Person指针
	personType := reflect.TypeOf(Person{})
	personValue := reflect.New(personType)
	person := personValue.Interface().(*Person)
	
	fmt.Printf("创建的Person：%s\n", person.String())
	
	// 通过反射设置字段值
	elem := personValue.Elem()
	nameField := elem.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("David")
	}
	
	ageField := elem.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(32)
	}
	
	emailField := elem.FieldByName("Email")
	if emailField.IsValid() && emailField.CanSet() {
		emailField.SetString("david@example.com")
	}
	
	fmt.Printf("设置字段后的Person：%s\n", person.String())
	
	// 通过反射调用方法
	method := personValue.MethodByName("Greet")
	if method.IsValid() {
		result := method.Call(nil)
		fmt.Printf("调用Greet方法：%s\n", result[0].String())
	}
}

// 示例5：处理切片和映射
func sliceMapReflectionExample() {
	fmt.Println("\n=== 切片和映射反射示例 ===")
	
	// 切片反射
	numbers := []int{1, 2, 3, 4, 5}
	sliceValue := reflect.ValueOf(numbers)
	
	fmt.Printf("切片类型：%s\n", sliceValue.Type())
	fmt.Printf("切片长度：%d\n", sliceValue.Len())
	fmt.Printf("切片容量：%d\n", sliceValue.Cap())
	
	// 遍历切片
	fmt.Print("切片元素：")
	for i := 0; i < sliceValue.Len(); i++ {
		fmt.Printf("%d ", sliceValue.Index(i).Int())
	}
	fmt.Println()
	
	// 动态创建切片
	stringSliceType := reflect.TypeOf([]string{})
	stringSlice := reflect.MakeSlice(stringSliceType, 0, 5)
	
	// 添加元素
	stringSlice = reflect.Append(stringSlice, reflect.ValueOf("hello"))
	stringSlice = reflect.Append(stringSlice, reflect.ValueOf("world"))
	
	fmt.Printf("动态创建的切片：%v\n", stringSlice.Interface())
	
	// 映射反射
	userMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Charlie": 28,
	}
	
	mapValue := reflect.ValueOf(userMap)
	fmt.Printf("映射类型：%s\n", mapValue.Type())
	fmt.Printf("映射长度：%d\n", mapValue.Len())
	
	// 遍历映射
	fmt.Println("映射内容：")
	for _, key := range mapValue.MapKeys() {
		value := mapValue.MapIndex(key)
		fmt.Printf("  %s: %d\n", key.String(), value.Int())
	}
	
	// 动态创建映射
	mapType := reflect.TypeOf(map[string]int{})
	newMap := reflect.MakeMap(mapType)
	
	// 设置值
	newMap.SetMapIndex(reflect.ValueOf("Eve"), reflect.ValueOf(22))
	newMap.SetMapIndex(reflect.ValueOf("Frank"), reflect.ValueOf(35))
	
	fmt.Printf("动态创建的映射：%v\n", newMap.Interface())
}

// 示例6：接口和类型断言
func interfaceReflectionExample() {
	fmt.Println("\n=== 接口反射示例 ===")
	
	var data interface{} = 42
	
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	
	fmt.Printf("接口值：%v\n", data)
	fmt.Printf("动态类型：%s\n", t)
	fmt.Printf("动态值：%v\n", v)
	
	// 类型检查
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("这是一个整数：%d\n", v.Int())
	case reflect.String:
		fmt.Printf("这是一个字符串：%s\n", v.String())
	case reflect.Struct:
		fmt.Printf("这是一个结构体：%s\n", t.Name())
	}
	
	// 动态类型转换
	if v.Type().ConvertibleTo(reflect.TypeOf(float64(0))) {
		converted := v.Convert(reflect.TypeOf(float64(0)))
		fmt.Printf("转换为float64：%f\n", converted.Float())
	}
}

// 示例7：反射实现通用函数
func genericFunctionExample() {
	fmt.Println("\n=== 通用函数示例 ===")
	
	// 通用的打印函数
	printValue := func(v interface{}) {
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)
		
		fmt.Printf("值：%v，类型：%s，种类：%s\n", v, rt, rt.Kind())
		
		switch rt.Kind() {
		case reflect.Struct:
			fmt.Printf("结构体有%d个字段：\n", rt.NumField())
			for i := 0; i < rt.NumField(); i++ {
				field := rt.Field(i)
				value := rv.Field(i)
				fmt.Printf("  %s: %v\n", field.Name, value.Interface())
			}
		case reflect.Slice:
			fmt.Printf("切片有%d个元素：\n", rv.Len())
			for i := 0; i < rv.Len(); i++ {
				fmt.Printf("  [%d]: %v\n", i, rv.Index(i).Interface())
			}
		case reflect.Map:
			fmt.Printf("映射有%d个键值对：\n", rv.Len())
			for _, key := range rv.MapKeys() {
				value := rv.MapIndex(key)
				fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
			}
		}
	}
	
	// 测试不同类型的值
	person := Person{Name: "Eve", Age: 22, Email: "eve@example.com"}
	numbers := []int{1, 2, 3}
	scores := map[string]int{"Math": 90, "English": 85}
	
	printValue(person)
	printValue(numbers)
	printValue(scores)
}

// 示例8：JSON序列化（简化版）
func jsonSerializeExample() {
	fmt.Println("\n=== JSON序列化示例 ===")
	
	person := Person{
		Name:  "Frank",
		Age:   33,
		Email: "frank@example.com",
	}
	
	// 简化的JSON序列化
	jsonString := toJSON(person)
	fmt.Printf("JSON序列化结果：%s\n", jsonString)
}

// 简化的JSON序列化函数
func toJSON(v interface{}) string {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	
	if rt.Kind() != reflect.Struct {
		return fmt.Sprintf("\"%v\"", v)
	}
	
	var parts []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)
		
		if !field.IsExported() {
			continue
		}
		
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}
		
		fieldName := field.Name
		if jsonTag != "" {
			fieldName = jsonTag
		}
		
		var valueStr string
		switch value.Kind() {
		case reflect.String:
			valueStr = fmt.Sprintf("\"%s\"", value.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueStr = fmt.Sprintf("%d", value.Int())
		case reflect.Float32, reflect.Float64:
			valueStr = fmt.Sprintf("%f", value.Float())
		case reflect.Bool:
			valueStr = fmt.Sprintf("%t", value.Bool())
		default:
			valueStr = fmt.Sprintf("\"%v\"", value.Interface())
		}
		
		parts = append(parts, fmt.Sprintf("\"%s\":%s", fieldName, valueStr))
	}
	
	return "{" + strings.Join(parts, ",") + "}"
}

// 示例9：结构体验证
func structValidationExample() {
	fmt.Println("\n=== 结构体验证示例 ===")
	
	person1 := Person{Name: "Grace", Age: 25, Email: "grace@example.com"}
	person2 := Person{Name: "", Age: -5, Email: "invalid-email"}
	
	fmt.Printf("验证person1：%t\n", validate(person1))
	fmt.Printf("验证person2：%t\n", validate(person2))
}

// 简化的验证函数
func validate(v interface{}) bool {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	
	if rt.Kind() != reflect.Struct {
		return true
	}
	
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)
		
		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			if !validateField(value, rule) {
				fmt.Printf("字段 %s 验证失败，规则：%s\n", field.Name, rule)
				return false
			}
		}
	}
	
	return true
}

// 简化的字段验证函数
func validateField(value reflect.Value, rule string) bool {
	switch rule {
	case "required":
		return !value.IsZero()
	case "email":
		if value.Kind() != reflect.String {
			return false
		}
		email := value.String()
		return strings.Contains(email, "@")
	}
	
	// 处理min和max规则
	if strings.HasPrefix(rule, "min=") {
		minStr := strings.TrimPrefix(rule, "min=")
		min, err := strconv.Atoi(minStr)
		if err != nil {
			return false
		}
		
		if value.Kind() == reflect.Int {
			return value.Int() >= int64(min)
		}
	}
	
	if strings.HasPrefix(rule, "max=") {
		maxStr := strings.TrimPrefix(rule, "max=")
		max, err := strconv.Atoi(maxStr)
		if err != nil {
			return false
		}
		
		if value.Kind() == reflect.Int {
			return value.Int() <= int64(max)
		}
	}
	
	return true
}

// 示例10：反射的性能考虑
func performanceExample() {
	fmt.Println("\n=== 性能考虑示例 ===")
	
	person := Person{Name: "Henry", Age: 40, Email: "henry@example.com"}
	
	// 直接访问
	fmt.Printf("直接访问：%s\n", person.Name)
	
	// 通过反射访问
	v := reflect.ValueOf(person)
	nameField := v.FieldByName("Name")
	fmt.Printf("反射访问：%s\n", nameField.String())
	
	fmt.Println("\n反射的优缺点：")
	fmt.Println("优点：")
	fmt.Println("  - 运行时类型检查")
	fmt.Println("  - 通用编程")
	fmt.Println("  - 序列化/反序列化")
	fmt.Println("  - 框架开发")
	fmt.Println("缺点：")
	fmt.Println("  - 性能开销")
	fmt.Println("  - 失去编译时类型检查")
	fmt.Println("  - 代码可读性降低")
	fmt.Println("  - 运行时错误风险")
}

func main() {
	fmt.Println("《Go程序设计语言》第12章：反射")
	fmt.Println("===================================")
	
	// 运行所有示例
	basicReflectionExample()
	structFieldReflectionExample()
	methodReflectionExample()
	dynamicValueExample()
	sliceMapReflectionExample()
	interfaceReflectionExample()
	genericFunctionExample()
	jsonSerializeExample()
	structValidationExample()
	performanceExample()
	
	fmt.Println("\n===================================")
	fmt.Println("第12章示例运行完成!")
} 