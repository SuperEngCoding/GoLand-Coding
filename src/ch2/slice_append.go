package main

import "fmt"

// Go语言append()为切片添加元素

func main() {
	var a []int
	// 追加一个元素
	a = append(a, 1)
	// 追加多个元素
	a = append(a, 1, 2, 3, 4)
	// 追加一个切片
	a = append(a, []int{5, 6, 7}...)
	for _, i2 := range a {
		fmt.Println(i2)
	}

	// append()为切片动态的添加元素。如果空间不足的话切片就会扩容，扩容的规律是按照2倍的方式进行的

	//  len: 1  cap: 1 pointer: 0xc00001a0d0
	//len: 2  cap: 2 pointer: 0xc00008e020
	//len: 3  cap: 4 pointer: 0xc000014120
	//len: 4  cap: 4 pointer: 0xc000014120
	//len: 5  cap: 8 pointer: 0xc000018100
	//len: 6  cap: 8 pointer: 0xc000018100
	//len: 7  cap: 8 pointer: 0xc000018100
	//len: 8  cap: 8 pointer: 0xc000018100
	//len: 9  cap: 16 pointer: 0xc000092000
	//len: 10  cap: 16 pointer: 0xc000092000

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		//打印输出切片的长度、容量和指针变化，使用函数 len() 查看切片拥有的元素个数，使用函数 cap() 查看切片的容量情况。
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	//除了在切片的尾部追加，我们还可以在切片的开头添加元素：
	var b = []int{1, 2, 3}
	b = append([]int{0}, b...)          // 在开头添加1个元素
	b = append([]int{-3, -2, -1}, b...) // 在开头添加1个切片

	// 因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素
	//var c []int
	//c = append(c[:i], append([]int{x}, c[i:]...)...) // 在第i个位置插入x
	//c = append(c[:i], append([]int{1,2,3}, c[i:]...)...) // 在第i个位置插入切片

}
