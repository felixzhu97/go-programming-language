# 第 11 章：测试

本章介绍了 Go 语言的测试框架，包括单元测试、基准测试、示例测试以及测试的最佳实践。

## 项目结构

```
chapter11/
├── README.md       # 本文档
├── testing.go      # 被测试的代码
└── testing_test.go # 测试文件
```

## 主要内容

### 1. 测试类型

#### 单元测试 (Unit Tests)

- **命名规则**：函数名以`Test`开头，接收`*testing.T`参数
- **目的**：验证代码的正确性
- **示例**：`TestCalculator`, `TestIsPrime`, `TestFibonacci`

#### 基准测试 (Benchmark Tests)

- **命名规则**：函数名以`Benchmark`开头，接收`*testing.B`参数
- **目的**：测试代码的性能
- **示例**：`BenchmarkBubbleSort`, `BenchmarkQuickSort`

#### 示例测试 (Example Tests)

- **命名规则**：函数名以`Example`开头
- **目的**：提供可执行的文档和示例
- **特点**：通过`// Output:`注释验证输出

### 2. 测试技术

#### 表格驱动测试 (Table-driven Tests)

- 使用结构体切片定义测试用例
- 便于添加新的测试用例
- 测试逻辑清晰，易于维护

```go
tests := []struct {
    name     string
    input    int
    expected bool
}{
    {"prime", 7, true},
    {"not prime", 8, false},
}
```

#### 子测试 (Sub-tests)

- 使用`t.Run()`创建子测试
- 可以并行执行
- 更好的测试组织和错误报告

#### 并行测试 (Parallel Tests)

- 使用`t.Parallel()`标记并行测试
- 提高测试执行效率
- 需要注意共享资源的使用

### 3. 测试工具和技术

#### 测试断言

- `t.Error()`, `t.Errorf()` - 报告错误但继续执行
- `t.Fatal()`, `t.Fatalf()` - 报告错误并停止执行
- `t.Skip()` - 跳过测试
- `t.Helper()` - 标记辅助函数

#### 测试设置和清理

- `t.Cleanup()` - 注册清理函数
- `testing.Short()` - 检查是否为短测试模式
- `b.ResetTimer()` - 重置基准测试计时器

## 运行测试

```bash
# 运行所有测试
go test

# 运行测试并显示详细信息
go test -v

# 运行特定测试
go test -run TestCalculator

# 运行测试并生成覆盖率报告
go test -cover

# 生成详细的覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# 运行基准测试
go test -bench=.

# 运行基准测试并显示内存分配
go test -bench=. -benchmem

# 运行并行基准测试
go test -bench=BenchmarkFibonacciParallel -cpu=1,2,4

# 运行示例测试
go test -run Example

# 短测试模式（跳过耗时测试）
go test -short
```

## 测试覆盖率

### 覆盖率类型

- **语句覆盖率**：测试执行的语句百分比
- **分支覆盖率**：测试执行的分支百分比
- **函数覆盖率**：测试调用的函数百分比

### 覆盖率命令

```bash
# 基本覆盖率
go test -cover

# 生成覆盖率文件
go test -coverprofile=coverage.out

# 查看覆盖率详情
go tool cover -func=coverage.out

# 生成HTML覆盖率报告
go tool cover -html=coverage.out
```

## 最佳实践

### 测试设计原则

1. **FIRST 原则**：

   - **Fast**：测试应该快速执行
   - **Independent**：测试应该独立，不依赖其他测试
   - **Repeatable**：测试应该可重复执行
   - **Self-validating**：测试应该有明确的通过/失败结果
   - **Timely**：测试应该及时编写

2. **AAA 模式**：

   - **Arrange**：设置测试数据和环境
   - **Act**：执行被测试的代码
   - **Assert**：验证结果

3. **命名规范**：
   - 测试函数名应该清楚地描述测试内容
   - 使用有意义的变量名
   - 子测试名应该描述测试场景

### 测试编写技巧

1. **边界值测试**：测试边界条件和特殊值
2. **错误处理测试**：验证错误情况的处理
3. **并发安全测试**：测试并发访问的安全性
4. **性能回归测试**：使用基准测试防止性能回归

### 测试文件组织

```
package/
├── main.go
├── main_test.go        # 主测试文件
├── helper_test.go      # 测试辅助函数
├── benchmark_test.go   # 基准测试
└── example_test.go     # 示例测试
```

## 高级测试技术

### 模拟和存根 (Mocking and Stubbing)

- 使用接口进行依赖注入
- 创建测试用的模拟对象
- 隔离被测试的代码

### 集成测试

- 测试多个组件的交互
- 使用真实的外部依赖
- 测试完整的用户场景

### 性能测试

- 使用基准测试测量性能
- 分析内存使用情况
- 识别性能瓶颈

## 测试工具

### 内置工具

- `go test` - 测试运行器
- `go tool cover` - 覆盖率分析
- `go tool pprof` - 性能分析

### 第三方工具

- `testify` - 增强的断言库
- `gomock` - 模拟对象生成器
- `ginkgo` - BDD 风格的测试框架

## 常见问题

### 测试失败的原因

1. **逻辑错误**：代码实现不正确
2. **边界条件**：没有考虑特殊情况
3. **并发问题**：竞态条件或死锁
4. **外部依赖**：依赖外部服务或文件

### 测试性能问题

1. **测试过慢**：使用并行测试或减少测试数据
2. **内存泄漏**：检查资源是否正确释放
3. **CPU 密集**：优化算法或使用更高效的数据结构

## 扩展学习

- 学习更多测试模式和技巧
- 了解测试驱动开发（TDD）
- 探索行为驱动开发（BDD）
- 学习持续集成和持续部署
- 深入理解 Go 语言的测试机制
