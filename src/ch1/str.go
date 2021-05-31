package main

import "fmt"

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

	// 计算字符串长度 len()
	fmt.Println(len(name)) //38

}
