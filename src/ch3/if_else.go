package main

import "fmt"

// 在Go语言中，关键字 if 是用于测试某个条件（布尔型或逻辑型）的语句，如果该条件成立，
// 则会执行 if 后由大括号{}括起来的代码块，否则就忽略该代码块继续执行后续的代码。
// 格式 ：
/*
if condition {
	//do something
}

if condition1 {
    // do something
} else if condition2 {
    // do something else
}else {
    // catch-all or default
}

*/

func main() {
	// 普通的写法
	var ten int = 11
	if ten > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}

	// 特殊写法
	// Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
	//if err := Connect(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
