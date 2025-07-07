# 第 1 章：入门

这一章包含了 Go 语言的入门示例，展示了 Go 语言的基本特性和常用模式。

## 示例程序

### 1. Hello World

- `hello.go` - 最简单的 Go 程序

### 2. 命令行参数处理

- `echo1.go` - 使用传统 for 循环处理命令行参数
- `echo2.go` - 使用 for range 循环处理命令行参数
- `echo3.go` - 使用 strings.Join 处理命令行参数

### 3. 查找重复行

- `dup1.go` - 从标准输入查找重复行
- `dup2.go` - 从文件或标准输入查找重复行
- `dup3.go` - 一次性读取整个文件查找重复行

### 4. 动画 GIF

- `lissajous.go` - 生成利萨如图形动画 GIF

### 5. 网络编程

- `fetch.go` - 获取 URL 内容
- `fetchall.go` - 并发获取多个 URL
- `server1.go` - 最小的 Web 服务器
- `server2.go` - 带计数器的 Web 服务器
- `server3.go` - 显示请求详细信息的 Web 服务器

## 运行方式

### 命令行程序

```bash
# Hello World
go run hello.go

# 命令行参数
go run echo1.go Hello World
go run echo2.go Hello World
go run echo3.go Hello World

# 查找重复行
go run dup1.go < input.txt
go run dup2.go file1.txt file2.txt
go run dup3.go file1.txt file2.txt

# 动画GIF
go run lissajous.go > out.gif

# 网络获取
go run fetch.go https://golang.org
go run fetchall.go https://golang.org https://godoc.org https://golang.org/help/
```

### Web 服务器

```bash
# 启动服务器
go run server1.go
go run server2.go
go run server3.go

# 访问地址
# http://localhost:8000/
# http://localhost:8000/count (只对server2.go有效)
```

## 学习要点

1. **Go 程序结构**：包声明、导入、函数定义
2. **命令行参数**：os.Args 的使用
3. **映射（map）**：键值对数据结构
4. **文件 I/O**：读取文件和标准输入
5. **并发编程**：goroutine 和 channel 的基本使用
6. **网络编程**：HTTP 客户端和服务器
7. **错误处理**：Go 语言的错误处理模式
