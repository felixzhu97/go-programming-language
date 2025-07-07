// Maps 演示映射类型和操作
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. 映射声明和初始化
	fmt.Printf("映射声明和初始化:\n")
	var m1 map[string]int // 零值映射
	fmt.Printf("零值映射: %v, len=%d, is nil: %t\n", m1, len(m1), m1 == nil)
	
	m2 := make(map[string]int) // 使用make创建
	fmt.Printf("make创建: %v, len=%d\n", m2, len(m2))
	
	m3 := map[string]int{ // 字面量初始化
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}
	fmt.Printf("字面量初始化: %v, len=%d\n", m3, len(m3))

	// 2. 映射的基本操作
	fmt.Printf("\n映射的基本操作:\n")
	scores := make(map[string]int)
	
	// 添加键值对
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("添加后: %v\n", scores)
	
	// 访问值
	fmt.Printf("Alice的分数: %d\n", scores["Alice"])
	
	// 修改值
	scores["Alice"] = 98
	fmt.Printf("修改后Alice的分数: %d\n", scores["Alice"])
	
	// 删除键值对
	delete(scores, "Bob")
	fmt.Printf("删除Bob后: %v\n", scores)

	// 3. 零值测试
	fmt.Printf("\n零值测试:\n")
	value := scores["NonExistent"]
	fmt.Printf("不存在的键的值: %d\n", value)
	
	// 使用ok惯用法检查键是否存在
	if value, ok := scores["Alice"]; ok {
		fmt.Printf("Alice存在，分数: %d\n", value)
	}
	
	if value, ok := scores["David"]; ok {
		fmt.Printf("David存在，分数: %d\n", value)
	} else {
		fmt.Printf("David不存在\n")
	}

	// 4. 映射遍历
	fmt.Printf("\n映射遍历:\n")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
	
	// 只遍历键
	fmt.Printf("只遍历键: ")
	for name := range scores {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// 5. 映射的键类型
	fmt.Printf("\n映射的键类型:\n")
	// 字符串键
	strMap := map[string]int{"one": 1, "two": 2}
	fmt.Printf("字符串键: %v\n", strMap)
	
	// 整数键
	intMap := map[int]string{1: "one", 2: "two"}
	fmt.Printf("整数键: %v\n", intMap)
	
	// 结构体键
	type Point struct{ x, y int }
	pointMap := map[Point]string{
		{0, 0}: "origin",
		{1, 1}: "diagonal",
	}
	fmt.Printf("结构体键: %v\n", pointMap)

	// 6. 映射的值类型
	fmt.Printf("\n映射的值类型:\n")
	// 切片值
	sliceMap := map[string][]int{
		"evens": {2, 4, 6, 8},
		"odds":  {1, 3, 5, 7},
	}
	fmt.Printf("切片值: %v\n", sliceMap)
	
	// 映射值
	nestedMap := map[string]map[string]int{
		"math": {"Alice": 95, "Bob": 87},
		"english": {"Alice": 92, "Bob": 89},
	}
	fmt.Printf("嵌套映射: %v\n", nestedMap)

	// 7. 映射排序
	fmt.Printf("\n映射排序:\n")
	ages := map[string]int{
		"Alice": 25, "Bob": 30, "Charlie": 20, "David": 35,
	}
	fmt.Printf("原映射: %v\n", ages)
	
	// 按键排序
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Printf("按键排序: ")
	for _, name := range names {
		fmt.Printf("%s:%d ", name, ages[name])
	}
	fmt.Println()

	// 8. 映射作为集合
	fmt.Printf("\n映射作为集合:\n")
	set := make(map[string]bool)
	words := []string{"apple", "banana", "apple", "cherry", "banana"}
	
	// 添加到集合
	for _, word := range words {
		set[word] = true
	}
	fmt.Printf("原数组: %v\n", words)
	fmt.Printf("去重后: ")
	for word := range set {
		fmt.Printf("%s ", word)
	}
	fmt.Println()

	// 9. 映射的引用性质
	fmt.Printf("\n映射的引用性质:\n")
	original := map[string]int{"a": 1, "b": 2}
	copy := original
	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("复制映射: %v\n", copy)
	
	copy["c"] = 3
	fmt.Printf("修改复制映射后:\n")
	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("复制映射: %v\n", copy)

	// 10. 映射作为函数参数
	fmt.Printf("\n映射作为函数参数:\n")
	data := map[string]int{"x": 10, "y": 20}
	fmt.Printf("调用前: %v\n", data)
	modifyMap(data)
	fmt.Printf("调用后: %v\n", data)

	// 11. 映射的实际应用
	fmt.Printf("\n映射的实际应用:\n")
	// 统计字符频率
	text := "hello world"
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	fmt.Printf("字符频率统计:\n")
	for char, count := range freq {
		fmt.Printf("'%c': %d\n", char, count)
	}

	// 12. 映射的容量
	fmt.Printf("\n映射的容量:\n")
	// 映射没有cap函数，但可以指定初始容量
	bigMap := make(map[int]int, 1000)
	for i := 0; i < 5; i++ {
		bigMap[i] = i * i
	}
	fmt.Printf("大映射: %v, len=%d\n", bigMap, len(bigMap))

	// 13. 映射的线程安全
	fmt.Printf("\n映射的线程安全:\n")
	fmt.Printf("注意：映射不是线程安全的，并发访问需要同步\n")
	
	// 14. 映射的性能
	fmt.Printf("\n映射的性能:\n")
	fmt.Printf("查找、插入、删除的平均时间复杂度: O(1)\n")
	fmt.Printf("最坏情况时间复杂度: O(n)\n")
}

// 修改映射内容
func modifyMap(m map[string]int) {
	m["z"] = 30
} 