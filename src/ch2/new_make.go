package main

// make 和 new的区别
// Go语言中new 和 make 是两个内置函数，主要用来创建并分配类型的内存。
// new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。
// make 只能用于slice,map,channel 的初始化。

func main() {
	// new
	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly
	// allocated zero value of that type.
	// func new(Type) *Type
	// 后续待更新

}
