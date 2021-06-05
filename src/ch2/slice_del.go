package main

import "fmt"

// Go语言从切片中删除元素

func main() {
	var a = []int{1, 2, 3}
	//a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素
	//fmt.Print(a) //[2 3]

	a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[N:]...) // 删除开头N个元素
	fmt.Print(a) //[2 3]

	// 还可以用 copy() 函数来删除开头的元素：

	a = a[:copy(a, a[1:])] // 删除开头1个元素
	//a = a[:copy(a, a[N:])] // 删除开头N个元素

	//从中间位置删除
	//a = []int{1, 2, 3, ...}
	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	// 从尾部删除

	a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N] // 删除尾部N个元素

	//删除切片指定位置的元素。

	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:]) //[a b] [d e]
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq) // [a b d e]

}
