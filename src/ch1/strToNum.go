package main

import (
	"fmt"
	"strconv"
)

//  字符串和数值类型的相互转换
//在实际开发中我们往往需要对一些常用的数据类型进行转换，如 string、int、int64、float 等数据类型之间的转换，
//Go语言中的 strconv 包为我们提供了字符串和基本数据类型之间的转换功能。
//strconv 包中常用的函数包括 Atoi()、Itia()、parse 系列函数、format 系列函数、append 系列函数等，下面就来分别介绍一下

func main() {
	//string 与 int 类型之间的转换
	//Itoa()：整型转字符串
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v \n", str, str) //type:string value:"100"

	//Atoi()：字符串转整型
	str1 := "110"
	str2 := "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1) //type:int value:110
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！", str2) //s100 转换失败！
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}

}
