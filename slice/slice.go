package main

import "fmt"

func main() {

	//1. 定义切片,给切片赋值
	//slice :=[]int{1,2}

	//2. 定义动态切片,但是没有给切片分配空间,注意判断切片是否为nil,否则直接赋值会报错
	//var slice []int
	//此时会报错:因为没有分配空间,所以不能赋值
	//slice[0] = 1

	//3. 定义动态切片,给切片分配空间, 1代表只给了一个位置, 下标从0开始,如果分配空间数量小于实际数量,进行赋值操作会报错
	//var slice = make([]int, 1)

	//4. 定义动态切片  , 第二个参数代表给了一个位置
	// 第二个参数分配空间数量  设置的几len就是几,数组中的值默认值被填充为0. 如果设置为0,则不分配空间,只是创建一个空切片cap=1
	// 第三个参数代表cap容量,不写默认为设置的分配空间数值
	//slice := make([]int, 3, 5)
	slice := make([]int, 0)
	//赋值 如果第二个参数没有分配空间,则会报错
	//slice[0] = 3

	if slice != nil {
		fmt.Printf("slice空间长度len=%d ,cap容量=%d  %v\n", len(slice), cap(slice), slice)
	} else {
		fmt.Println("slice is nil")
	}

	//切片追加元素,动态增加空间,如果定义了cap容量, 调用append时会动态增加设置的cap容量
	slice = append(slice, 2)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)

	if slice != nil {
		fmt.Printf("slice空间长度len=%d ,cap容量=%d  %v\n", len(slice), cap(slice), slice)
	} else {
		fmt.Println("slice is nil")
	}

	//截取数组 左闭右开
	//[下标:(第几个元素=下标+1)]
	// 取下标为1到2的元素
	s1 := slice[1:3]
	fmt.Printf("s1空间长度len=%d ,cap容量=%d  %v\n", len(s1), cap(s1), s1)
	//取下标为0到1的元素
	s2 := slice[0:2]
	fmt.Printf("s2空间长度len=%d ,cap容量=%d  %v\n", len(s2), cap(s2), s2)
	//从开始到第几个元素
	s3 := slice[:2]
	fmt.Printf("s3空间长度len=%d ,cap容量=%d  %v\n", len(s3), cap(s3), s3)
	//从下标到结束
	s4 := slice[2:]
	fmt.Printf("s4空间长度len=%d ,cap容量=%d  %v\n", len(s4), cap(s4), s4)

	//切片后的新切片,进行赋值,会修改原切片的值
	s4[0] = 100
	fmt.Printf("s4空间长度len=%d ,cap容量=%d  %v\n", len(s4), cap(s4), s4)

	//切片的深拷贝
	s5 := make([]int, len(s4))
	copy(s5, s4)
	fmt.Printf("s5空间长度len=%d ,cap容量=%d  %v\n", len(s5), cap(s5), s5)
	//此时修改s5的值,不会修改s4的值
	s5[0] = 200
	fmt.Printf("s4空间长度len=%d ,cap容量=%d  %v\n", len(s4), cap(s4), s4)
	fmt.Printf("s5空间长度len=%d ,cap容量=%d  %v\n", len(s5), cap(s5), s5)

}
