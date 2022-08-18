package main

import "fmt"

/*
四种变量的声明方式
*/

//生命全局变量 方法123可以 4不可以
var gAge int = 10

func main() {
	//1.var 变量名 类型 = 值
	var a int = 10
	//打印变量a的类型
	fmt.Printf("a的类型是%T\n", a)

	//2.var 变量名 类型 默认值整形是0，字符串是空字符串，布尔是false
	var b int
	fmt.Printf("b的类型是%T\n", b)
	//3.var 变量名  = 值 省去类型
	var c = "hello"
	fmt.Printf("c的类型是%T\n", c)
	//4.  变量名 := 值 省去类型 常用 只能用在函数体内
	d := "world"
	fmt.Printf("d的类型是%T\n", d)

	//打印全局变量gAge
	fmt.Println(gAge)

	//生命多个变量 aa,bb,cc
	var aa, bb, cc int = 1, 2, 3
	fmt.Println(aa, bb, cc)

	//生命多个变量 aa1,bb1,cc1
	var (
		aa1 = 1
		bb1 = 2
		cc1 = 3
	)
	fmt.Println(aa1, bb1, cc1)

}
