package main

import "fmt"

// 函数 返回单个值
// (参数 类型,参数 类型) 返回值(没有则不写) {
func func1(variableA string, variableB int) int {
	//形参可以不使用
	//打印变量
	//fmt.Println(variableA)
	//fmt.Println(variableB)
	return 0
}

// 函数 返回多个值无形参名
func func2(a int, b string) (int, string) {
	return 1, "hello"
}

//函数 返回多个值有形参名 不同返回值类型
func func3(a int, b string) (c int, d string) {
	c = 1
	d = "hello"
	return
}

// 函数 返回多个值有形参名 相同返回值类型
func func4(a int, b string) (c, d int) {

	//c = 1
	//d = 2
	//直接 return 默认返回会查找c d变量的值
	//return
	//这种也可以
	//return 3,5
	// 如果没有定义c d直接return 的话c,d 是默认值0
	return

}
func main() {
	//调用方法
	res := func1("hello", 1)
	fmt.Println(res)

	//调用方法
	res1, res2 := func2(1, "hello")
	fmt.Println(res1, res2)

	//调用方法
	res3, res4 := func4(1, "hello")
	fmt.Println(res3, res4)
}
