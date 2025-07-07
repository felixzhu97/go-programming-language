// Slices 演示切片类型和操作
package main

import "fmt"

func main() {
	// 1. 切片声明和初始化
	fmt.Printf("切片声明和初始化:\n")
	var s1 []int // 零值切片
	fmt.Printf("零值切片: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	
	s2 := []int{1, 2, 3, 4, 5} // 直接初始化
	fmt.Printf("直接初始化: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	
	s3 := make([]int, 5) // 使用make创建
	fmt.Printf("make创建: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	
	s4 := make([]int, 3, 5) // 指定长度和容量
	fmt.Printf("指定容量: %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	// 2. 从数组创建切片
	fmt.Printf("\n从数组创建切片:\n")
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("切片[1:4]: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 3. 切片操作
	fmt.Printf("\n切片操作:\n")
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", s)
	fmt.Printf("s[1:4]: %v\n", s[1:4])
	fmt.Printf("s[:3]: %v\n", s[:3])
	fmt.Printf("s[2:]: %v\n", s[2:])
	fmt.Printf("s[:]: %v\n", s[:])

	// 4. 切片扩容
	fmt.Printf("\n切片扩容:\n")
	var nums []int
	fmt.Printf("初始: %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))
	
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("添加%d: len=%d, cap=%d\n", i, len(nums), cap(nums))
	}

	// 5. append操作
	fmt.Printf("\nappend操作:\n")
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	
	// 追加单个元素
	a = append(a, 4)
	fmt.Printf("append(a, 4): %v\n", a)
	
	// 追加多个元素
	a = append(a, 5, 6, 7)
	fmt.Printf("append(a, 5, 6, 7): %v\n", a)
	
	// 追加另一个切片
	a = append(a, b...)
	fmt.Printf("append(a, b...): %v\n", a)

	// 6. copy操作
	fmt.Printf("\ncopy操作:\n")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	n := copy(dst, src)
	fmt.Printf("源切片: %v\n", src)
	fmt.Printf("目标切片: %v\n", dst)
	fmt.Printf("复制了 %d 个元素\n", n)

	// 7. 切片的引用性质
	fmt.Printf("\n切片的引用性质:\n")
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:4]
	slice2 := original[2:5]
	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("切片1[1:4]: %v\n", slice1)
	fmt.Printf("切片2[2:5]: %v\n", slice2)
	
	slice1[1] = 100 // 修改slice1
	fmt.Printf("修改slice1[1]=100后:\n")
	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("切片1: %v\n", slice1)
	fmt.Printf("切片2: %v\n", slice2)

	// 8. 切片遍历
	fmt.Printf("\n切片遍历:\n")
	data := []string{"apple", "banana", "cherry"}
	for i, v := range data {
		fmt.Printf("索引%d: %s\n", i, v)
	}

	// 9. 切片作为函数参数
	fmt.Printf("\n切片作为函数参数:\n")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("函数调用后: %v\n", numbers)

	// 10. 删除切片元素
	fmt.Printf("\n删除切片元素:\n")
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("原切片: %v\n", items)
	
	// 删除索引2的元素
	i := 2
	items = append(items[:i], items[i+1:]...)
	fmt.Printf("删除索引2: %v\n", items)

	// 11. 切片排序
	fmt.Printf("\n切片排序:\n")
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("排序前: %v\n", unsorted)
	
	// 手动冒泡排序
	for i := 0; i < len(unsorted); i++ {
		for j := 0; j < len(unsorted)-1-i; j++ {
			if unsorted[j] > unsorted[j+1] {
				unsorted[j], unsorted[j+1] = unsorted[j+1], unsorted[j]
			}
		}
	}
	fmt.Printf("排序后: %v\n", unsorted)

	// 12. 二维切片
	fmt.Printf("\n二维切片:\n")
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
		for j := range matrix[i] {
			matrix[i][j] = i*4 + j + 1
		}
	}
	fmt.Printf("二维切片:\n")
	for i := range matrix {
		fmt.Printf("%v\n", matrix[i])
	}

	// 13. 切片的nil检查
	fmt.Printf("\n切片的nil检查:\n")
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nil切片: %v, len=%d, cap=%d, is nil: %t\n", 
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("空切片: %v, len=%d, cap=%d, is nil: %t\n", 
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
}

// 修改切片内容
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
} 