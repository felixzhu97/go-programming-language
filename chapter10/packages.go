package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// 示例1：包的基本概念
func packageBasicsExample() {
	fmt.Println("=== 包的基本概念 ===")
	
	// 包名和导入路径
	fmt.Println("1. 包名：通常是导入路径的最后一个元素")
	fmt.Println("2. 导入路径：唯一标识包的字符串")
	fmt.Println("3. 包声明：每个源文件都必须以package声明开始")
	
	// 包的可见性规则
	fmt.Println("\n可见性规则：")
	fmt.Println("- 大写字母开头的标识符是导出的（public）")
	fmt.Println("- 小写字母开头的标识符是未导出的（private）")
}

// 示例2：包的初始化顺序
// 包级别变量
var (
	initialized = initFunction()
	counter     = 0
)

func initFunction() string {
	fmt.Println("包级别变量初始化")
	return "initialized"
}

// init函数 - 包初始化时自动调用
func init() {
	fmt.Println("init函数被调用")
	counter = 100
}

func packageInitExample() {
	fmt.Println("\n=== 包初始化示例 ===")
	fmt.Printf("初始化的变量: %s\n", initialized)
	fmt.Printf("计数器值: %d\n", counter)
	
	fmt.Println("\n包初始化顺序：")
	fmt.Println("1. 导入的包先被初始化")
	fmt.Println("2. 包级别变量按依赖顺序初始化")
	fmt.Println("3. init函数被调用")
	fmt.Println("4. main函数被调用（如果存在）")
}

// 示例3：空白导入（blank import）
// 这通常用于触发包的初始化而不直接使用包
func blankImportExample() {
	fmt.Println("\n=== 空白导入示例 ===")
	
	fmt.Println("空白导入的用途：")
	fmt.Println("1. 触发包的初始化")
	fmt.Println("2. 注册驱动或插件")
	fmt.Println("3. 调试和测试")
	
	// 示例：_ "database/sql/driver"
	fmt.Println("\n常见的空白导入：")
	fmt.Println("import _ \"database/sql/driver\"")
	fmt.Println("import _ \"net/http/pprof\"")
}

// 示例4：包别名
func packageAliasExample() {
	fmt.Println("\n=== 包别名示例 ===")
	
	fmt.Println("包别名的用途：")
	fmt.Println("1. 解决包名冲突")
	fmt.Println("2. 使用更短的名称")
	fmt.Println("3. 提高代码可读性")
	
	fmt.Println("\n包别名语法：")
	fmt.Println("import alias \"package/path\"")
	fmt.Println("import . \"package/path\"  // 导入到当前命名空间")
	fmt.Println("import _ \"package/path\"  // 空白导入")
}

// 示例5：包的文档
/*
Package main 演示了Go语言包和工具的使用。

这是一个多行注释的示例，用于包级别的文档。
包文档应该简洁地描述包的用途和功能。

示例用法：

    go run packages.go

更多信息请参考：
    https://golang.org/doc/
*/

func packageDocExample() {
	fmt.Println("\n=== 包文档示例 ===")
	
	fmt.Println("包文档规范：")
	fmt.Println("1. 包文档应该在package声明之前")
	fmt.Println("2. 以'Package 包名'开头")
	fmt.Println("3. 简洁描述包的功能")
	fmt.Println("4. 提供使用示例")
	
	fmt.Println("\n函数文档规范：")
	fmt.Println("1. 在函数声明之前")
	fmt.Println("2. 以函数名开头")
	fmt.Println("3. 解释函数的功能、参数和返回值")
}

// AddNumbers 计算两个整数的和。
//
// 这是一个示例函数，演示了函数文档的写法。
// 参数 a 和 b 是要相加的两个整数。
// 返回值是两个整数的和。
//
// 示例：
//    result := AddNumbers(3, 4)
//    fmt.Println(result) // 输出: 7
func AddNumbers(a, b int) int {
	return a + b
}

// 示例6：go build 和 go install
func buildToolsExample() {
	fmt.Println("\n=== 构建工具示例 ===")
	
	fmt.Println("常用的go命令：")
	fmt.Println("1. go build  - 编译包和依赖")
	fmt.Println("2. go install - 编译并安装包")
	fmt.Println("3. go run    - 编译并运行")
	fmt.Println("4. go test   - 运行测试")
	fmt.Println("5. go mod    - 模块管理")
	fmt.Println("6. go get    - 下载并安装包")
	fmt.Println("7. go fmt    - 格式化代码")
	fmt.Println("8. go vet    - 代码静态分析")
}

// 示例7：查询包信息
func packageQueryExample() {
	fmt.Println("\n=== 查询包信息示例 ===")
	
	fmt.Println("查询包信息的方法：")
	fmt.Println("1. go list - 列出包信息")
	fmt.Println("2. go doc  - 查看包文档")
	fmt.Println("3. godoc   - 启动文档服务器")
	
	fmt.Println("\n常用的go list命令：")
	fmt.Println("go list -m all           # 列出所有模块")
	fmt.Println("go list ./...            # 列出当前模块的所有包")
	fmt.Println("go list -f '{{.Dir}}' fmt # 列出fmt包的目录")
}

// 示例8：简单的AST解析
func astExample() {
	fmt.Println("\n=== AST解析示例 ===")
	
	// 解析简单的Go代码
	src := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`
	
	// 创建文件集
	fset := token.NewFileSet()
	
	// 解析源代码
	node, err := parser.ParseFile(fset, "hello.go", src, parser.ParseComments)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}
	
	// 遍历AST
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			fmt.Printf("找到函数: %s\n", x.Name.Name)
		case *ast.CallExpr:
			if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
				if pkg, ok := fun.X.(*ast.Ident); ok {
					fmt.Printf("找到函数调用: %s.%s\n", pkg.Name, fun.Sel.Name)
				}
			}
		}
		return true
	})
}

// 示例9：工作区操作
func workspaceExample() {
	fmt.Println("\n=== 工作区操作示例 ===")
	
	// 获取当前工作目录
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取工作目录失败: %v\n", err)
		return
	}
	fmt.Printf("当前工作目录: %s\n", pwd)
	
	// 查找Go文件
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".go") {
			fmt.Printf("找到Go文件: %s\n", path)
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}
}

// 示例10：环境变量
func environmentExample() {
	fmt.Println("\n=== 环境变量示例 ===")
	
	// 常用的Go环境变量
	envVars := []string{
		"GOPATH",
		"GOROOT",
		"GOOS",
		"GOARCH",
		"CGO_ENABLED",
		"GO111MODULE",
	}
	
	for _, env := range envVars {
		value := os.Getenv(env)
		if value != "" {
			fmt.Printf("%s: %s\n", env, value)
		} else {
			fmt.Printf("%s: (未设置)\n", env)
		}
	}
}

func main() {
	fmt.Println("《Go程序设计语言》第10章：包和工具")
	fmt.Println("====================================")
	
	// 运行所有示例
	packageBasicsExample()
	packageInitExample()
	blankImportExample()
	packageAliasExample()
	packageDocExample()
	
	// 演示文档化函数
	result := AddNumbers(3, 4)
	fmt.Printf("\nAddNumbers(3, 4) = %d\n", result)
	
	buildToolsExample()
	packageQueryExample()
	astExample()
	workspaceExample()
	environmentExample()
	
	fmt.Println("\n====================================")
	fmt.Println("第10章示例运行完成!")
} 