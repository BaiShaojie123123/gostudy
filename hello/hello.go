package main

import (
	"fmt"
	"time"
)

//主函数
func main() {
	fmt.Println("Hello, world.")
	//定义一个数组
	var arr [5]int
	//赋值
	arr[0] = 1
	arr[1] = 2

	//定义变量num
	var num int = 0

	//遍历数组
	for i := 0; i < len(arr); i++ {
		//判断是否大于1
		if arr[i] > 1 {
			//赋值
			num = arr[i]
		}
		time.Sleep(time.Minute)
	}

	//输出
	fmt.Println(num)
}
