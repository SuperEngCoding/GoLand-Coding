package main

import (
	"fmt"
	"net"
)

// Go语言的变量声明 var 变量名  变量类型

func main() {
	/**变量的声明*/
	//01 标准格式: var name type
	var age int
	//02 批量格式
	var (
		a string
		b int
		c bool
	)
	//03 简洁的格式 名字 := 表达式  可以不用指定类型 通过复制会预判该变量的类型
	d := 1
	i, j := 0, 1

	/** 变量的初始化 */
	//01 标准格式 var 变量名 类型 = 表达式
	var name string = "小明" // 这里从idea提示中可以看出string是可以除去的
	//02 编译器推导类型的格式
	var address = "安徽"
	//03 短变量声明并初始化
	sex := "男" //由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。
	//address := "阜阳" 这里会编译不通过
	// 短变量声明并初始化 多个  从func Dial(network, address string) (Conn, error) 方法可知返回值两个
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	fmt.Println(age)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(sex)

	fmt.Println(conn)
	fmt.Println(err)

}
