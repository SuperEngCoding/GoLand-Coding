package main

import "fmt"

// map 删除和清空
// map 删除 使用delete

func main() {
	scene := make(map[string]int)
	scene = map[string]int{"route": 66, "brazil": 4, "china": 960}
	delete(scene, "route")
	//value: 960
	//value: 4
	for _, i := range scene {
		fmt.Printf("value: %d \n", i)
	}

	// 清空 map 中的所有元素
	// Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map
	//的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

}
