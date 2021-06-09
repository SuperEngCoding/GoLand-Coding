package main

import "fmt"

// 与多数语言不同的是，Go语言中的循环语句只支持 for 关键字，而不支持 while 和 do-while 结构

func main() {
	sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	// Go语言还进一步考虑到无限循环的场景，让开发者不用写无聊的 for(;;){}和do{} while(1);，而直接简化为如下的写法：
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)
	// Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环

	//	for j := 0; j < 5; j++ {
	//		for i := 0; i < 10; i++ {
	//			if i > 5 {
	//				break JLoop
	//			}
	//			fmt.Println(i)
	//		}
	//	}
	//JLoop:
	//	// ...

	// for中的初始语句 -- 开始循环时的执行的语句
	//初始语句是在第一次循环前执行的语句，一般使用初始语句执行变量初始化，如果变量在此处被声明，其作用域将被局限在这个 for 的范围内。
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	//这段代码将 step 放在 for 的前面进行初始化，for 中没有初始语句，此时 step 的作用域就比在初始语句中声明 step 要大。

}
