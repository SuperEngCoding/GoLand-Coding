package main

import (
	"fmt"
	"math"
)

// 浮点类型
// Go语言提供了两种精度的浮点数 float32 和 float64

func main() {
	// 通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，
	//并且 float32 能精确表示的正整数并不是很大
	var f float32 = 16777216
	fmt.Println(f == f+1) //true 有误差
	//浮点数在声明的时候可以只写整数部分或者小数部分，像下面这样：
	e := .89 // 0.89
	h := 1.  // 1
	fmt.Println(e)
	fmt.Println(h)
	//用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数
	fmt.Printf("%f\n", math.Pi)   // 3.141593
	fmt.Printf("%.2f\n", math.Pi) // 3.14

}
