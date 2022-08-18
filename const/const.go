package main

import "fmt"

//iota 按照每行增加

//定义枚举类型间隔1
const (
	//iota可以自动增加数值 第一行默认0
	BEIJING   = iota //iota=0   0
	SHANGHAI         //iota=1	1
	GUANGZHOU        //iota=2	2
)

//定义枚举类型 间隔10
const (
	//iota可以自动增加数值 第一行默认0
	BEIJING_1   = 10 * iota //iota=0*10   0
	SHANGHAI_2              //iota=1*10	10
	GUANGZHOU_3             //iota=2*10	20
)

func main() {
	//常量 status 按照变量定义方法 前面var 改为const 只读不可修改
	const status int = 10
	const type1 = 10
	//打印status 常量
	fmt.Println(status)

}
