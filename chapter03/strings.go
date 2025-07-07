// Strings 演示字符串类型和操作
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1. 字符串基础
	var s1 string = "Hello, World!"
	var s2 string = "你好，世界！"
	var s3 string // 零值为空字符串
	
	fmt.Printf("字符串基础:\n")
	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("s2: %s\n", s2)
	fmt.Printf("s3 (零值): %q\n", s3)
	fmt.Printf("s3 是否为空: %t\n", s3 == "")

	// 2. 字符串长度
	fmt.Printf("\n字符串长度:\n")
	fmt.Printf("len(\"%s\"): %d 字节\n", s1, len(s1))
	fmt.Printf("len(\"%s\"): %d 字节\n", s2, len(s2))
	fmt.Printf("utf8.RuneCountInString(\"%s\"): %d 符文\n", s2, utf8.RuneCountInString(s2))

	// 3. 字符串索引和切片
	fmt.Printf("\n字符串索引和切片:\n")
	hello := "Hello"
	fmt.Printf("hello[0]: %c (ASCII: %d)\n", hello[0], hello[0])
	fmt.Printf("hello[1:3]: %s\n", hello[1:3])
	fmt.Printf("hello[:2]: %s\n", hello[:2])
	fmt.Printf("hello[2:]: %s\n", hello[2:])

	// 4. 字符串拼接
	fmt.Printf("\n字符串拼接:\n")
	first := "Hello"
	second := "World"
	fmt.Printf("使用+: %s\n", first+" "+second)
	fmt.Printf("使用fmt.Sprintf: %s\n", fmt.Sprintf("%s %s", first, second))
	fmt.Printf("使用strings.Join: %s\n", strings.Join([]string{first, second}, " "))

	// 5. 字符串比较
	fmt.Printf("\n字符串比较:\n")
	fmt.Printf("\"abc\" == \"abc\": %t\n", "abc" == "abc")
	fmt.Printf("\"abc\" < \"def\": %t\n", "abc" < "def")
	fmt.Printf("\"ABC\" < \"abc\": %t\n", "ABC" < "abc")
	fmt.Printf("strings.EqualFold(\"ABC\", \"abc\"): %t\n", strings.EqualFold("ABC", "abc"))

	// 6. 字符串常用函数
	fmt.Printf("\n字符串常用函数:\n")
	text := "  Hello, Go Programming!  "
	fmt.Printf("原字符串: %q\n", text)
	fmt.Printf("Trim: %q\n", strings.Trim(text, " "))
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(text))
	fmt.Printf("ToUpper: %q\n", strings.ToUpper(text))
	fmt.Printf("ToLower: %q\n", strings.ToLower(text))
	fmt.Printf("Replace: %q\n", strings.Replace(text, "Go", "Golang", -1))
	fmt.Printf("Contains: %t\n", strings.Contains(text, "Go"))
	fmt.Printf("HasPrefix: %t\n", strings.HasPrefix(text, "  Hello"))
	fmt.Printf("HasSuffix: %t\n", strings.HasSuffix(text, "!  "))

	// 7. 字符串分割和连接
	fmt.Printf("\n字符串分割和连接:\n")
	csv := "apple,banana,orange"
	parts := strings.Split(csv, ",")
	fmt.Printf("Split: %v\n", parts)
	fmt.Printf("Join: %s\n", strings.Join(parts, " | "))
	
	// 8. 字符串查找
	fmt.Printf("\n字符串查找:\n")
	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("Index of 'fox': %d\n", strings.Index(sentence, "fox"))
	fmt.Printf("LastIndex of 'the': %d\n", strings.LastIndex(sentence, "the"))
	fmt.Printf("Count of 'the': %d\n", strings.Count(sentence, "the"))

	// 9. 字符串和字节切片转换
	fmt.Printf("\n字符串和字节切片转换:\n")
	str := "Hello, 世界"
	bytes := []byte(str)
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("从字节切片恢复: %s\n", string(bytes))

	// 10. 符文(rune)处理
	fmt.Printf("\n符文(rune)处理:\n")
	chinese := "你好世界"
	fmt.Printf("字符串: %s\n", chinese)
	fmt.Printf("字节长度: %d\n", len(chinese))
	fmt.Printf("符文长度: %d\n", utf8.RuneCountInString(chinese))
	
	// 遍历符文
	fmt.Printf("遍历符文:\n")
	for i, r := range chinese {
		fmt.Printf("位置 %d: %c (Unicode: U+%04X)\n", i, r, r)
	}

	// 11. 字符串字面量
	fmt.Printf("\n字符串字面量:\n")
	fmt.Printf("普通字符串: %s\n", "Hello\nWorld")
	fmt.Printf("原始字符串: %s\n", `Hello\nWorld`)
	fmt.Printf("多行原始字符串:\n%s\n", `第一行
第二行
第三行`)

	// 12. 字符串格式化
	fmt.Printf("\n字符串格式化:\n")
	name := "Alice"
	age := 25
	fmt.Printf("我的名字是%s，今年%d岁\n", name, age)
	fmt.Printf("使用占位符: %[1]s今年%[2]d岁，%[1]s很开心\n", name, age)
	fmt.Printf("十六进制: %x\n", "Hello")
	fmt.Printf("引用格式: %q\n", "Hello\nWorld")
} 