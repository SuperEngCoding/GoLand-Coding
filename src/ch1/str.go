package main

import (
	"fmt"
	"unicode/utf8"
)

//字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，
//更深入地讲，字符串是字节的定长数组。

func main() {
	// 使用""来定义字符串 字符串可以使用转义符实现换行缩进等效果 常见的转义符如下所示
	// \n 换行  \r 回车符 \t tab键  \u Unicode字符  \\ 反斜杠本身
	str := "Java\nGo"
	fmt.Println(str)
	// 字符串拼接使用+
	msg := "一个消息:" + "好消息"
	fmt.Println(msg)
	// 字符串拼接也可以使用 +=
	s := "夏天"
	s += "秋天"
	fmt.Println(s)

	// 定义多行字符串
	var name = `第一行
	第二行
	第三行
    /n`
	fmt.Println(name)

	// ASCII 字符串长度使用 len() 函数。
	// Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
	// 计算字符串长度 len() 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度。
	str1 := "abcd efg"
	str2 := "中国"
	fmt.Println(len(str1)) //8
	// 中国为什么是6呢而不是2呢
	//这里的差异是由于 Go 语言的字符串都以 UTF-8 格式保存，
	//每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。
	fmt.Println(len(str2)) //6
	str3 := "中国，niubi"
	//UTF-8 包提供的 RuneCountInString() 函数，统计 Uncode 字符数量。
	fmt.Println(utf8.RuneCountInString(str2)) //2
	fmt.Println(utf8.RuneCountInString(str3)) //8

	// 字符串遍历 有二种遍历方法
	// 01 遍历每个ASCII字符  - ASCII 字符串遍历直接使用下标。
	//如果出现中文就会出现乱码
	str4 := "中国,China"
	for i := 0; i < len(str4); i++ {
		fmt.Printf("ascii: %c %d\n", str4[i], str4[i])
	}

	// 按Unicode字符遍历字符串 - Unicode 字符串遍历用 for range 出现中文也不会乱码
	for _, s := range str4 {
		fmt.Printf("unicode: %c %d\n", s, s)
	}

}
