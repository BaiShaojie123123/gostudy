package lib1

import "fmt"

//当前lib1包提供的api 注意开头大写
func Lib1test() {
	fmt.Println("lib1test")
}

func init() {
	fmt.Println("lib1")
}
