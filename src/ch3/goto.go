package main

import "fmt"

// Go语言中 goto 语句通过标签进行代码间的无条件跳转，同时 goto 语句在快速跳出循环、
//避免重复退出上也有一定的帮助，使用 goto 语句能简化一些代码的实现过程。

func main() {
	//  使用goto退出多层循环

	// 传统的编写方式
	//var breakAgain bool
	////外循环
	//for i := 0; i < 10; i++ {
	//	//内循环
	//	for y := 0; y < 10; y++ {
	//		if y == 2 {
	//			breakAgain = true
	//			// 退出内循环
	//			break
	//
	//		}
	//	}
	//	if breakAgain {
	//		// 退出外循环
	//		break
	//	}
	//}
	//fmt.Println("done")

	// 使用goto 进行优化

	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				//使用 goto 语句跳转到指明的标签处，标签在第 42 行定义。
				goto breakHere
			}
		}
	}
	return
	//定义 breakHere 标签。
breakHere:
	fmt.Println("done")
}
