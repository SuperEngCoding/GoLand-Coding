package main

import (
	"container/list"
	"fmt"
)

// 列表是一种非连续的存储容器。由多个节点通过一些变量记录彼此之间的关系。列表实现的方法由多种。例如链表和双向链表。
// Go语言中列表是使用container/list包来实现，内部实现原理是双链表，列表能够高效的进行任意位置的删除添加操作

func main() {
	// 初始化列表有两种方法 New()函数 和 var关键字声明
	// 1) 通过container/list包中的New()函数初始化List 格式为: 变量名 := list.New()
	// 2) 通过var关键字声明list 格式:  var 变量名 listList
	// 需要注意点: 列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，
	//这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

	//01  在list中添加元素
	// 双链表支持从队列头和尾插入元素。 对应的方法为 PushFront 和 PushBack。
	// 提示: 这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，
	//这种方法可以让删除更加效率化，同时也是双链表特性之一。
	//l := list.New()
	//front := l.PushFront("first")
	//l.PushBack(30)
	//在 front 点之后插入元素，front 点由其他插入函数提供
	//l.InsertAfter("two",front)
	//InsertBefore(v interface {}, mark * Element) *Element	  在 mark 点之前插入元素，mark 点由其他插入函数提供
	// PushBackList(other *List)	添加 other 列表元素到尾部 - 添加列表
	//PushFrontList(other *List)	添加 other 列表元素到头部  - 添加列表

	//02 从列表中删除元素
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 移除 element 变量对应的元素。
	l.Remove(element)

	// 03 遍历列表
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
