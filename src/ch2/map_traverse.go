package main

import (
	"fmt"
	"sort"
)

// 遍历Map

func main() {
	//scene := make(map[string]int)
	//scene = map[string]int{"one": 1, "two": 2}
	//for s, i := range scene {
	//	fmt.Printf("key:,%s,value:%d \n", s, i)
	//}

	// 注意：遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果。

	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList) //[brazil china route]

}
