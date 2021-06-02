package main

import "fmt"

//Go语言数组详解

// 数组定义的格式如下所示:
// var 数组名 [元素数量]Type

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 1
	a[2] = 2
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// // 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 打印索引和元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，
	//同时也可以使用数组字面值语法，用一组值来初始化数组：

	var b [3]int = [3]int{1, 2, 3}
	for i, i2 := range b {
		fmt.Printf("%d %d\n", i, i2)
	}

	//在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	//，因此，上面数组 q 的定义可以简化为：
	c := [...]int{2, 3, 4, 4, 4, 4, 4}
	for i, i2 := range c {
		fmt.Printf("%d %d\n", i, i2)
	}

	//比较两个数组是否相等

	d := [2]int{1, 2}
	e := [2]int{1, 3}
	f := [2]int{1, 3}
	fmt.Println(d == e) //false
	fmt.Println(f == e) //true

}
