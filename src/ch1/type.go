package main

import (
	"fmt"
	"time"
)

//Go语言Type关键字(类型别名)

// 类型别名是 Go 1.9 版本添加的新功能，主要用于解决代码升级、迁移中存在的类型兼容性问题。

func main() {
	// 区分类型别名和类型定义
	// 定义类型别名的写法为:   type TypeAlias = Type
	// TypeAlias 只是 Type 的别名，本质上 TypeAlias 与 Type 是同一个类型，
	// 就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
	// 类型别名与类型定义表面上看只有一个等号的差异，那么它们之间实际的区别有哪些呢？下面通过一段代码来理解
	// 将NewInt定义为int类型
	type NewInt int //对应的类型为NewInt
	// 将int取一个别名叫IntAlias
	type IntAlias = int // 对应的类型还是int
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a) //a type: main.NewInt
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	////IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	fmt.Printf("a2 type: %T\n", a2) //a2 type: int

}

// 非本地类型不能定义方法
// 定义time.Duration的别名为MyDuration

type MyDuration = time.Duration

// 为MyDuration添加一个函数
// cannot define new methods on non-local type time.Duration 编译报错
// type MyDuration time.Duration 这样定义就不会编译报错了
//func (m MyDuration) EasySet(a string) {
//}
