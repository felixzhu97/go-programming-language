// Arrays 演示数组类型和操作
package main

import "fmt"

func main() {
	// 1. 数组声明和初始化
	fmt.Printf("数组声明和初始化:\n")
	var a [5]int // 零值数组
	fmt.Printf("零值数组: %v\n", a)
	
	var b = [5]int{1, 2, 3, 4, 5} // 完整初始化
	fmt.Printf("完整初始化: %v\n", b)
	
	c := [...]int{1, 2, 3} // 自动推断长度
	fmt.Printf("自动推断长度: %v\n", c)
	
	d := [5]int{1: 10, 3: 30} // 索引初始化
	fmt.Printf("索引初始化: %v\n", d)

	// 2. 数组长度和容量
	fmt.Printf("\n数组长度和容量:\n")
	fmt.Printf("数组b的长度: %d\n", len(b))
	fmt.Printf("数组b的容量: %d\n", cap(b))

	// 3. 数组访问和修改
	fmt.Printf("\n数组访问和修改:\n")
	fmt.Printf("b[0] = %d\n", b[0])
	b[0] = 100
	fmt.Printf("修改后 b[0] = %d\n", b[0])
	fmt.Printf("修改后的数组: %v\n", b)

	// 4. 数组遍历
	fmt.Printf("\n数组遍历:\n")
	// 方法1：传统for循环
	fmt.Printf("传统for循环: ")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()
	
	// 方法2：range循环
	fmt.Printf("range循环: ")
	for i, v := range b {
		fmt.Printf("[%d]:%d ", i, v)
	}
	fmt.Println()
	
	// 方法3：只要值
	fmt.Printf("只要值: ")
	for _, v := range b {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 5. 数组比较
	fmt.Printf("\n数组比较:\n")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)

	// 6. 多维数组
	fmt.Printf("\n多维数组:\n")
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}
	fmt.Printf("3x3矩阵:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 7. 数组的值传递
	fmt.Printf("\n数组的值传递:\n")
	original := [3]int{1, 2, 3}
	fmt.Printf("原数组: %v\n", original)
	modifyArray(original)
	fmt.Printf("函数调用后: %v\n", original)
	
	// 8. 数组指针
	fmt.Printf("\n数组指针:\n")
	fmt.Printf("修改前: %v\n", original)
	modifyArrayPtr(&original)
	fmt.Printf("修改后: %v\n", original)

	// 9. 数组切片
	fmt.Printf("\n数组切片:\n")
	full := [5]int{1, 2, 3, 4, 5}
	slice := full[1:4]
	fmt.Printf("原数组: %v\n", full)
	fmt.Printf("切片[1:4]: %v\n", slice)
	
	// 10. 数组的实际应用
	fmt.Printf("\n数组的实际应用:\n")
	// 计算数组元素总和
	numbers := [5]int{10, 20, 30, 40, 50}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("数组元素总和: %d\n", sum)
	
	// 查找最大值
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("最大值: %d\n", max)
}

// 值传递：不会修改原数组
func modifyArray(arr [3]int) {
	arr[0] = 100
	fmt.Printf("函数内修改: %v\n", arr)
}

// 指针传递：会修改原数组
func modifyArrayPtr(arr *[3]int) {
	arr[0] = 100
	fmt.Printf("函数内修改: %v\n", *arr)
} 