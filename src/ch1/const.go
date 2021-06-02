package main

import (
	"fmt"
	"math"
)

// Go语言常量和Const关键字
// Go语言的常量是使用Const关键字进行定义用来存储不会改变的变量。常量在编译器时被创建及时在函数内部
// 并且只能时布尔类型，数字类型，和字符串类型。
// 由于在编译器确定变量的值，定义常量的表达式必须能被编译器求值的常量表达式。

func main() {
	// 定义常量的格式为 const name [type] = value  其中的type可以省去不写
	const str = "我是个常量" // 常量在本文件中未被使用编译器也不会报错，可能被其他文件进行使用

	// 编译期间可以确定
	const c1 = 2 / 3
	// 编译不通过  因为在编译期间不能确定c2的值
	//const c2 = getNumber()

	// 定义多个常量
	const (
		a1 = 1
		a2 = 2
		a3 = 3
	)

	const (
		a = 1
		b // 在这里b默认是没有赋值的 默认和上面保持一直
		c = 2
		d //d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	// iota常量生成器 - 在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 0 1 2 3 4 5 6

	// 无类型常量 - 通过延迟明确常量的具体类型，不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。
	const Pi64 float64 = math.Pi
	fmt.Println(Pi64)

}

func getNumber() (c int) {
	return 1
}
