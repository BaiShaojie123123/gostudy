package main

import "fmt"

func func1() {
	fmt.Println("func1")
}
func func2() {
	fmt.Println("func2")
}
func main() {
	//defer 是延迟执行的意思，defer 函数会在函数返回前执行，defer 函数可以用于释放资源，比如文件句柄，网络连接，数据库连接等。
	//压栈，函数执行完毕后，会按照逆序的顺序，依次执行defer函数。
	defer func1() //后执行
	defer func2() //前执行

	//return 在defer之前执行

	fmt.Println("main")
}
