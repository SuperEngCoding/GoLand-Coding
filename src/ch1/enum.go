package main

import "fmt"

//Go语言模拟枚举 - 为什么叫模拟枚举呢 因为目前Go语言是没有枚举类型
// 是通过const  和 iota 模拟枚举
// 声明芯片类型

type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

func main() {

	type Weapon int //将 int 定义为 Weapon 类型，就像枚举类型的本质是一个 int 类型一样
	const (
		Arrow Weapon = iota // 开始生成枚举值, 默认为0
		Shrike
		SniperRifle
		Rifle
		Blower
	)
	// 输出所有枚举值
	fmt.Println(Arrow, Shrike, SniperRifle, Rifle, Blower) //0 1 2 3 4
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower // 4
	fmt.Println(weapon)

	//iota 不仅可以生成每次增加 1 的枚举值。还可以利用 iota 来做一些强大的枚举常量值生成器
	const (
		FlagNone = 1 << iota //第 2 行中 iota 使用了一个移位操作，每次将上一次的值左移一位（二进制位），以得出每一位的常量值。
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue) //2 4 8
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue) //10 100 1000 将枚举值按二进制格式输出，可以清晰地看到每一位的变化。

	// 将枚举值转换为字符串

	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
}

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
