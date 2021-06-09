package main

// 使用 goto 集中处理错误

//多处错误处理存在代码重复时是非常棘手的，例如：

func main() {
	// 传统的处理方法
	//	err := firstCheckError()
	//	if err != nil {
	//		fmt.Println(err)
	//		exitProcess()
	//		return
	//	}
	//	err = secondCheckError()
	//	if err != nil {
	//		fmt.Println(err)
	//		exitProcess()
	//		return
	//	}
	//	fmt.Println("done")
	//
	//	// 使用goto优化
	//
	//	err := firstCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	err = secondCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	fmt.Println("done")
	//	return
	//onExit:
	//	fmt.Println(err)
	//	exitProcess()

}
