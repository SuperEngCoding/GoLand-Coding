package main

import "fmt"

func main() {
	//Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	//跨越 case 的 fallthrough——兼容C语言的 case 设计
	// 在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，
	// 但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，

	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

}
