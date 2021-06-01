package main

import "fmt"

// Go语言指针 - 两个核心的概念
// 类型指针: 允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
// 切片: 由指向起始元素的原始指针,元素数量和容量组成。

func main() {
	// 指针地址和指针类型
	cat := "猫咪"
	// 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址
	fmt.Println(&cat) //0xc000090030

	//从指针获取指针指向的值 - 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	str := "China NO1 88773"
	ptr := &str
	//打印ptr的数据类型 -string
	fmt.Printf("ptr type: %T \n", ptr)
	//打印ptr的地址
	fmt.Printf("ptr address: %p\n", ptr)
	// 根据地址取出地址指向的值
	value := *ptr
	fmt.Printf("value type %T\n", value)
	fmt.Printf("value  %s\n", value)

	// 使用指针修改值
	//

}
