package main

import "fmt"

// 多维切片  格式: var sliceName [][]...[]sliceType
// 其中，sliceName 为切片的名字，sliceType为切片的类型，每个[ ]代表着一个维度，切片有几个维度就需要几个[]

func main() {
	////声明一个二维切片
	//var slice [][]int
	////为二维切片赋值
	//slice = [][]int{{10}, {100, 200}}

	// 声明一个二维整型切片并赋值
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)
	fmt.Println(slice) // [[10 20] [100 200]]
}
