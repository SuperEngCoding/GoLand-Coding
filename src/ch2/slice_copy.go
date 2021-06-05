package main

import "fmt"

// Go语言的内置copy()可以将一个数组切片复制到另外一个数组切片中，如果加入的二个数组大小不一致。就会按照其中较小的数组切片的元素进行复制。
// copy() 函数使用格式如下:
// copy(destSlice,srcSlice[]T)int - 其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice），
//目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。\

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//for _, i := range slice2 {
	//	fmt.Print(i)   //123
	//}
	fmt.Println("===")
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	for _, i2 := range slice1 {
		fmt.Print(i2) // 54345
	}

	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := range srcData {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
