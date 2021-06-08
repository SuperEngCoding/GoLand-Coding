package main

import (
	"errors"
	"fmt"
)

// 空值/零值
// 在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""
//而指针、切片、映射、通道、函数和接口的零值则是 nil。

func main() {
	// 01 nil 标识符是不能比较的
	//fmt.Println(nil =nil)  - 编译报错

	// 02 nil不是关键字或保留字
	// nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量，比如下面这样：
	// 虽然上面的声明语句可以通过编译，但是并不提倡这么做。
	var nil = errors.New("my god")
	fmt.Println(nil)

	// 03  nil没有默认类型
	fmt.Printf("%T \n", nil) //  *errors.errorString

	// 04  不同类型的nil指针是一样的
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr) // 0x0
	fmt.Printf("%p", num)   // 0x0

	// 不同类型的nil是不能比较的
	//var m map[int]string
	//var ptr *int
	//fmt.Printf(m == ptr)
	/* 需要注意的点 */
	// nil 是 map、slice、pointer、channel、func、interface 的零值
	// 不同类型的 nil 值占用的内存大小可能是不一样的
	// 两个相同类型的 nil 值也可能无法比较
	// 不同类型的 nil 值占用的内存大小可能是不一样的

}
