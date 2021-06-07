package main

import "fmt"

// map 是一种特殊的数据结构，一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值）
//所以这个结构也称为关联数组或字典，这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。
// 格式为: var mapName map[keyType]valueType
// 在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。

func main() {
	// 定义map 类型 的变量  mapLit
	var mapLit map[string]int
	// 定义map类型的 mapAssigned
	var mapAssigned map[string]int
	// {key:value,key:value} 来初始化map的值
	mapLit = map[string]int{"one": 1, "two": 2}
	//可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，
	//会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	// make 被用来分配引用类型的内存： map, slice, channel
	// new 被用来分配除了引用类型的所有其他类型的内存： int, string, array等
	// mapCreated := make(map[string]float32) 等价于 mapCreated := map[string]float32{}
	mapCreated := make(map[string]float32)

	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	// Map literal at "one" is: 1
	// Map created at "key2" is: 3.141590
	// Map assigned at "two" is: 3
	// Map literal at "ten" is: 0
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	// map 的容量  map和数组不同是根据key-value动态的伸缩。因此它不存在最大的容量的限制。当然我们也可以指明初始容量，减少扩容的过程。
	// 格式:  make(map[keyType]valueType,cap)
	map2 := make(map[string]float32, 10)
	fmt.Println(map2)

	// 用切片作为map的值
	// 既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？
	// 例如，当我们要处理 unix 机器上的所有进程，以父进程（pid 为整形）作为 key，
	// 所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将 value 定义为 []int 类型或者其他类型的切片，
	// 就可以优雅的解决这个问题，示例代码如下所示：
	//mp1 := make(map[int][]int)
	//mp2 := make(map[int]*[]int)

}
