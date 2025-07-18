# 第 4 章：复合类型

这一章介绍了 Go 语言的复合类型，包括数组、切片、映射和结构体。这些类型允许程序员组合基本类型来创建更复杂的数据结构。

## 示例程序

### 1. 数组

- `arrays.go` - 演示数组声明、初始化、操作和多维数组等

### 2. 切片

- `slices.go` - 演示切片的创建、操作、扩容和与数组的关系等

### 3. 映射

- `maps.go` - 演示映射的创建、操作、遍历和实际应用等

### 4. 结构体

- `structs.go` - 演示结构体定义、方法、嵌套和实际应用等

## 运行方式

```bash
# 运行各个示例
go run arrays.go
go run slices.go
go run maps.go
go run structs.go
```

## 类型对比

| 类型   | 大小 | 元素类型 | 键类型     | 零值       | 可比较 |
| ------ | ---- | -------- | ---------- | ---------- | ------ |
| 数组   | 固定 | 任意     | 整数索引   | 零值数组   | 是     |
| 切片   | 可变 | 任意     | 整数索引   | nil        | 否     |
| 映射   | 可变 | 任意     | 可比较类型 | nil        | 否     |
| 结构体 | 固定 | 任意     | 字段名     | 零值结构体 | 是\*   |

\*结构体可比较当且仅当所有字段都可比较

## 关键概念

### 1. 数组 (Arrays)

- **固定大小**: 声明时必须指定大小
- **值类型**: 赋值和传参时会复制整个数组
- **同质性**: 所有元素必须是相同类型
- **索引访问**: 使用`arr[index]`访问元素
- **比较**: 相同类型的数组可以使用`==`比较

### 2. 切片 (Slices)

- **动态数组**: 大小可以在运行时改变
- **引用类型**: 多个切片可以共享底层数组
- **三要素**: 指针、长度(len)、容量(cap)
- **nil 切片**: 零值为 nil，长度和容量都是 0
- **切片操作**: `[low:high]`、`append()`、`copy()`

### 3. 映射 (Maps)

- **键值对**: 存储键值对的无序集合
- **引用类型**: 零值为 nil
- **动态大小**: 可以动态添加和删除元素
- **哈希表**: 底层实现为哈希表，平均 O(1)查找
- **键的要求**: 键类型必须可比较

### 4. 结构体 (Structs)

- **字段集合**: 将相关数据组合在一起
- **值类型**: 赋值时复制所有字段
- **字段访问**: 使用`.`操作符访问字段
- **方法**: 可以为结构体定义方法
- **嵌套**: 支持结构体嵌套和匿名字段

## 最佳实践

### 数组最佳实践

1. **使用场景**: 大小已知且固定的情况
2. **性能**: 访问速度快，内存局部性好
3. **传参**: 大数组传参使用指针避免复制

### 切片最佳实践

1. **容量管理**: 预分配容量避免频繁扩容
2. **nil 检查**: 使用前检查切片是否为 nil
3. **内存泄漏**: 注意长期引用大切片的小片段

### 映射最佳实践

1. **零值检查**: 使用`value, ok := map[key]`检查键是否存在
2. **初始化**: 使用`make()`创建映射
3. **并发安全**: 映射不是线程安全的，需要同步

### 结构体最佳实践

1. **字段顺序**: 考虑内存对齐，将相似大小的字段放在一起
2. **指针 vs 值**: 大结构体使用指针，小结构体使用值
3. **标签**: 使用结构体标签进行序列化配置

## 内存模型

### 切片内存模型

```
切片头:
[指针] -> [底层数组]
[长度]
[容量]
```

### 映射内存模型

```
映射头:
[哈希表指针] -> [桶数组]
[元素数量]
[其他元数据]
```

## 常用操作

### 切片操作

```go
// 创建
s := make([]int, len, cap)
s := []int{1, 2, 3}

// 追加
s = append(s, 4, 5, 6)
s = append(s, otherSlice...)

// 复制
copy(dst, src)

// 删除元素
s = append(s[:i], s[i+1:]...)
```

### 映射操作

```go
// 创建
m := make(map[string]int)
m := map[string]int{"key": value}

// 检查键
if value, ok := m[key]; ok {
    // 键存在
}

// 删除
delete(m, key)

// 遍历
for key, value := range m {
    // 处理键值对
}
```

## 性能考虑

### 时间复杂度

| 操作 | 数组 | 切片     | 映射     | 结构体 |
| ---- | ---- | -------- | -------- | ------ |
| 访问 | O(1) | O(1)     | O(1)平均 | O(1)   |
| 插入 | -    | O(1)摊销 | O(1)平均 | -      |
| 删除 | -    | O(n)     | O(1)平均 | -      |
| 搜索 | O(n) | O(n)     | O(1)平均 | O(1)   |

### 空间复杂度

- **数组**: O(n)，n 为元素数量
- **切片**: O(n)，额外的切片头开销
- **映射**: O(n)，额外的哈希表开销
- **结构体**: 所有字段大小之和

## 应用场景

### 数组适用场景

- 固定大小的数据集合
- 需要高性能的数值计算
- 嵌入式系统或内存受限环境

### 切片适用场景

- 动态数据集合
- 函数参数传递
- 数据流处理

### 映射适用场景

- 键值对存储
- 缓存和索引
- 配置管理

### 结构体适用场景

- 数据建模
- API 响应/请求
- 数据库记录

## 练习建议

1. 实现矩阵运算（使用二维切片）
2. 创建一个简单的缓存系统（使用映射）
3. 设计用户管理系统（使用结构体）
4. 比较数组和切片的性能差异
5. 实现栈和队列数据结构
6. 练习结构体的嵌套和方法定义
