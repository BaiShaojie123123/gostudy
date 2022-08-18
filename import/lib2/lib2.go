package lib2

import "fmt"

//当前lib2包提供的api 注意开头大写
func Lib2test() {
	fmt.Println("lib2test")
}

func init() {
	fmt.Println("lib2")
}
