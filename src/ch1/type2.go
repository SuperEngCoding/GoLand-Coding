package main

import (
	"fmt"
	"reflect"
)

//在结构体成员嵌入时使用别名

// 定义商标结构

type Brand struct {
}

// 为商标结构添加Show()方法

func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand

type FakeBrand = Brand

// 定义车辆结构

type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	// a.Show() 编译器将发生报错： ambiguous selector a.Show 在调用 Show() 方法时，因为两个类型都有 Show() 方法，会发生歧义，证明 FakeBrand 的本质确实是 Brand 类型。
	a.FakeBrand.Show()
	// 取a的类型反射对象
	ta := reflect.TypeOf(a)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}
	//FieldName: FakeBrand, FieldType: Brand
	//FieldName: Brand, FieldType: Brand
}
