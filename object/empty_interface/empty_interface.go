package main

import "fmt"

//interface{}表示任意类型
func myFun(arg interface{}) {
	fmt.Println("myFun is called")
	fmt.Println(arg)
	//区分interface{}引用的底层类型,类型断言
	//注意arg必须是 interface{} 类型才可用

	//类型断言,判断arg是否是string类型
	value, ok := arg.(string)
	if ok {
		fmt.Println("string")
		fmt.Println(value)
	} else {
		fmt.Println("not string")
	}
}

type Book struct {
	name string
}

func main() {
	book := Book{
		name: "book",
	}
	myFun(book)
	myFun(100)
	myFun(3.12)
	myFun("你好")
}
