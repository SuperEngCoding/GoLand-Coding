package main

import "fmt"

// Go中的切片 - 是对数组的一个连续片段的引用。所以切片是一个引用类型。其中这个片段可以是整个数组也可以是数组起始和终止索引标识的子集。
//  注意点: 切片是不包含终止的一个索引
// 切片内部包含: 地址,大小，容量。

func main() {
	// 从数组或者切片中生成新的切片 - 切片默认指向一段连续的内存区域。 可以是数组也可以切片本身。
	// 格式: slice[开始的位置:结束的位置] - 其中slice: 表示目标切片对象  开始位置: 对应目标对象的索引  结束位置: 对应目标对象的结束索引。

	// 从数组中生产切片
	var a = [3]int{1, 2, 3}
	its := a[1:3]       // 下标从1 到3 但是不包括3包含1
	fmt.Println(a, its) //[1 2 3] [2 3]    [2 3] 就是新的一个切片

	//从指定范围中生成切片
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15]) //[11 12 13 14 15]
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:]) //[21 22 23 24 25 26 27 28 29 30]
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2]) //[1 2]

	// 表示原有的切片
	b := []int{1, 2, 3}
	fmt.Println(b[:]) //[1 2 3]
	// 清空切片
	fmt.Println(b[0:0]) //[]

	//声明新的切片 格式: var name[] type

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	//numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。
	fmt.Println(numListEmpty == nil)

	// 使用 make() 函数构造切片  格式: make( []Type, size, cap )
	// 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，
	// 只是能提前分配空间，降低多次分配空间造成的性能问题。
	c := make([]int, 2)
	d := make([]int, 2, 10)
	fmt.Println(c, d)           //[0 0] [0 0]
	fmt.Println(len(c), len(d)) // 2  2
	// 需要注意的点:  使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，
	// 设定开始与结束位置，不会发生内存分配操作。

}
