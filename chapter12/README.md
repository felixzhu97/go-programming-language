# 第 12 章：反射

本章介绍了 Go 语言的反射机制，包括类型信息获取、值操作、方法调用以及反射的实际应用。

## 项目结构

```
chapter12/
├── README.md        # 本文档
└── reflection.go    # 反射示例代码
```

## 主要内容

### 1. 反射基础

#### reflect.Type 和 reflect.Value

- **reflect.Type**：表示类型信息
- **reflect.Value**：表示值信息
- **获取方式**：
  - `reflect.TypeOf(v)` - 获取类型
  - `reflect.ValueOf(v)` - 获取值

#### 基本操作

```go
t := reflect.TypeOf(person)
v := reflect.ValueOf(person)

fmt.Printf("类型：%s\n", t)
fmt.Printf("值：%v\n", v)
fmt.Printf("种类：%s\n", t.Kind())
```

### 2. 结构体反射

#### 字段信息获取

- **字段数量**：`t.NumField()`
- **字段信息**：`t.Field(i)`
- **字段值**：`v.Field(i)`
- **标签信息**：`field.Tag.Get("json")`

#### 字段属性

- **名称**：`field.Name`
- **类型**：`field.Type`
- **标签**：`field.Tag`
- **偏移量**：`field.Offset`
- **是否导出**：`field.IsExported()`

### 3. 方法反射

#### 方法信息获取

- **方法数量**：`t.NumMethod()`
- **方法信息**：`t.Method(i)`
- **方法调用**：`v.Method(i).Call(args)`

#### 方法调用

```go
method := v.MethodByName("Greet")
if method.IsValid() {
    result := method.Call(nil)
    fmt.Println(result[0].String())
}
```

### 4. 动态操作

#### 值的创建和修改

- **创建新值**：`reflect.New(t)`
- **设置字段值**：`field.SetString("value")`
- **检查可设置性**：`field.CanSet()`

#### 集合操作

- **切片**：`reflect.MakeSlice()`, `reflect.Append()`
- **映射**：`reflect.MakeMap()`, `MapIndex()`, `SetMapIndex()`
- **数组**：`reflect.ArrayOf()`, `Index()`

### 5. 接口和类型转换

#### 接口处理

- **类型检查**：`t.Kind()`
- **类型转换**：`v.Convert(targetType)`
- **接口获取**：`v.Interface()`

#### 类型断言

```go
if v.Type().ConvertibleTo(targetType) {
    converted := v.Convert(targetType)
    // 使用转换后的值
}
```

## 运行示例

```bash
# 运行反射示例
go run reflection.go

# 查看输出，了解反射的各种用法
# 包括类型信息、字段访问、方法调用等
```

## 实际应用

### 1. JSON 序列化

```go
func toJSON(v interface{}) string {
    rv := reflect.ValueOf(v)
    rt := reflect.TypeOf(v)

    // 遍历结构体字段
    for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        value := rv.Field(i)

        // 获取JSON标签
        jsonTag := field.Tag.Get("json")
        // 序列化字段值
    }
}
```

### 2. 结构体验证

```go
func validate(v interface{}) bool {
    rv := reflect.ValueOf(v)
    rt := reflect.TypeOf(v)

    for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        value := rv.Field(i)

        // 获取验证标签
        validateTag := field.Tag.Get("validate")
        // 执行验证逻辑
    }
}
```

### 3. 通用函数

```go
func printValue(v interface{}) {
    rv := reflect.ValueOf(v)
    rt := reflect.TypeOf(v)

    switch rt.Kind() {
    case reflect.Struct:
        // 处理结构体
    case reflect.Slice:
        // 处理切片
    case reflect.Map:
        // 处理映射
    }
}
```

## 反射的特性

### 类型种类 (Kind)

- **基本类型**：`Bool`, `Int`, `Float`, `String`
- **复合类型**：`Array`, `Slice`, `Map`, `Struct`
- **引用类型**：`Chan`, `Func`, `Interface`, `Ptr`
- **特殊类型**：`UnsafePointer`

### 值的属性

- **IsValid()**：值是否有效
- **IsZero()**：是否为零值
- **CanInterface()**：是否可以调用 Interface()
- **CanSet()**：是否可以设置
- **CanAddr()**：是否可以取地址

## 性能考虑

### 性能开销

1. **反射比直接访问慢**：通常慢 10-100 倍
2. **类型检查开销**：运行时类型检查
3. **内存分配**：反射操作可能分配内存
4. **方法调用开销**：动态方法调用

### 优化策略

1. **缓存反射信息**：避免重复获取类型信息
2. **预先验证**：在初始化时验证反射操作
3. **减少反射使用**：只在必要时使用反射
4. **使用接口**：优先使用接口而非反射

## 最佳实践

### 何时使用反射

1. **序列化/反序列化**：JSON、XML、Protocol Buffers
2. **框架开发**：Web 框架、ORM、依赖注入
3. **通用算法**：需要处理不同类型的算法
4. **配置解析**：动态配置加载和验证

### 何时避免反射

1. **性能敏感**：对性能要求极高的场景
2. **类型已知**：编译时已知类型的情况
3. **简单操作**：简单的类型转换和操作
4. **API 设计**：公共 API 应该避免暴露反射

### 安全使用反射

1. **错误处理**：检查 IsValid()、CanSet()等
2. **恐慌恢复**：使用 recover 捕获可能的恐慌
3. **类型检查**：在操作前检查类型
4. **边界检查**：检查索引和长度

## 常见陷阱

### 1. 未导出字段

```go
// 无法访问未导出字段
field := v.FieldByName("private")
if !field.CanInterface() {
    // 无法获取未导出字段的值
}
```

### 2. 指针和值的混淆

```go
// 需要区分指针和值
if t.Kind() == reflect.Ptr {
    t = t.Elem()
    v = v.Elem()
}
```

### 3. 方法集差异

```go
// 值类型和指针类型的方法集不同
valueType := reflect.TypeOf(person)
pointerType := reflect.TypeOf(&person)
```

### 4. 零值检查

```go
// 检查值是否有效
if !v.IsValid() {
    // 处理无效值
}
```

## 高级技巧

### 1. 动态创建类型

```go
// 创建结构体类型
fields := []reflect.StructField{
    {Name: "Name", Type: reflect.TypeOf("")},
    {Name: "Age", Type: reflect.TypeOf(0)},
}
structType := reflect.StructOf(fields)
```

### 2. 函数反射

```go
funcValue := reflect.ValueOf(someFunc)
funcType := funcValue.Type()

// 检查函数签名
fmt.Printf("参数个数: %d\n", funcType.NumIn())
fmt.Printf("返回值个数: %d\n", funcType.NumOut())
```

### 3. 通道反射

```go
chanType := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0))
chanValue := reflect.MakeChan(chanType, 10)
```

## 扩展学习

- 深入理解 Go 的类型系统
- 学习更多反射应用场景
- 探索反射的性能优化技巧
- 了解其他语言的反射机制
- 研究 Go 标准库中的反射使用
