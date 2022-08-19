package main

import (
	"fmt"
	"reflect"
)

//1. 声明数据类型
//type 变量名 类型
//定义一个类型叫myint 类型是int
type myint int

//type 变量名 类型 []表示数组
//type 变量名 类型 map表示map

//type 变量名 类型 struct表示结构体
//定义结构体类型的名称如果大写，则表示其他包可见
type mystruct struct {
	//定义结构体的字段名称  字段名称 字段类型
	//字段名称如果大写，表示该属性对外可见,如果小写，表示该属性对外不可见
	name string
	age  int
}

//type 变量名 类型 func表示函数

func main() {
	//使用自定义类型
	var a myint = 1
	//等价于
	var b int = 1
	fmt.Println(a)
	fmt.Println(b)

	//使用自定义结构体
	var structVariable mystruct
	//修改结构体的值
	structVariable.name = "张三"
	structVariable.age = 18
	fmt.Println(structVariable)

	//传参打印结构体,值传递
	fmt.Println("值传递修改前", structVariable)
	changeStruct(structVariable)
	fmt.Println("值传递修改后", structVariable)

	//传参打印结构体,引用传递
	fmt.Println("引用传递修改前", structVariable)
	changeStruct2(&structVariable)
	fmt.Println("引用传递修改后", structVariable)

	//遍历结构体
	// 遍历结构体
	//t := reflect.TypeOf(structVariable)
	//v := reflect.ValueOf(structVariable)
	//for k := 0; k < t.NumField(); k++ {
	//	// 判断是否指针类型
	//	if kind := v.Field(k).Kind(); kind == reflect.Ptr {
	//		// 判断是否结构体
	//		if kind := v.Field(k).Elem().Kind(); kind == reflect.Struct {
	//			fmt.Printf("%#v \n", v.Field(k).Interface())
	//		}
	//	}
	//}
	//遍历结构体
	//获取结构体的类型 注意结构体类型不能是指针类型 否则会报错 panic recovered: reflect: NumField of non-struct type
	t := reflect.TypeOf(structVariable)
	//获取结构体的值
	v := reflect.ValueOf(structVariable)
	//获取结构体的字段数量
	k := t.NumField()
	for i := 0; i < k; i++ {
		//获取结构体的字段名称
		fmt.Println(t.Field(i).Name)
		//获取结构体的字段类型
		fmt.Println(t.Field(i).Type)
		//获取结构体的字段值
		fmt.Println(v.Field(i))
	}

}

//结构体传参默认是值传递，不是引用传递
func changeStruct(structVariable mystruct) {
	structVariable.age = 19
}

//引用传递
func changeStruct2(structVariable *mystruct) {
	structVariable.age = 19
}

//func ForeachStruct(obj mystruct) {
//	t := reflect.TypeOf(obj) // 注意，obj不能为指针类型，否则会报：panic recovered: reflect: NumField of non-struct type
//	v := reflect.ValueOf(obj)
//	for k := 0; k < t.NumField(); k++ {
//		fmt.Printf("%s -- %v \n", t.Field(k).Tag, v.Field(k).Interface())
//	}
//}
