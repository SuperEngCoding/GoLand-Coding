package main

import "fmt"

//主要讲述一下匿名变量 - 匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符
//匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

func main() {
	// 只需要getData的第一个返回值
	a, _ := getData()
	// 只需要getData的第二个返回值
	_, b := getData()
	fmt.Println(a)
	fmt.Println(b)
}

func getData() (a int, b int) {
	return 200, 100
}
