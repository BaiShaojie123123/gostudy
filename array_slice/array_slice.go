package main

import "fmt"

//数组传参
//固定数组不建议, 如果数组长度不对，会报错
//传参在函数内部修改数组的值，函数外部不能看到修改后的值
//func test(arr [5]int) {
//	for index, value := range arr {
//		fmt.Println(index, value)
//	}
//}

//遍历数组1  引用传递,内部改变外部也会改变
func test(arr []int) {

	//遍历数组
	//index是下标，value是值,写了下标就可以取值，没有下标就只能取下标,不能取值
	//for 中的变量定义了必须使用，否则报错
	for index, value := range arr {
		fmt.Println(index, value)
	}

	//修改数组中的值,注意传进来的数组不能是空数组
	arr[0] = 100
}

//遍历数组2 引用传递,内部改变外部也会改变
func test3(arr []int) {
	//_表示忽略变量,可以使用_不会报错,不写就是遍历数组的下标,所以想取值必须写上下标或者用 _ 代替
	for _, value := range arr {
		fmt.Println(value)
	}
}

//遍历数组3 引用传递,内部改变外部也会改变
func test2(arr []int) {
	//遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	//	定义数组
	//	1.固定大小的数组 默认值为0
	var arr [10]int
	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr)
	//遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//2.固定大小给默认值
	myarr := [5]int{1, 2, 3, 4, 5}

	//查看数组类型
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", myarr)

	//数组传参
	//test(myarr)

	//3. 动态大小的数组,切片slice []中没有数字
	var myarr2 = []int{1, 2, 3, 4, 5}
	//空切片不能被赋值!!! 需要判断是否是nil
	var myarr3 = []int{}
	myarr5 := []int{1, 2, 3, 4, 5}
	myarr6 := []int{}

	//打印 myarr2
	//fmt.Println(myarr2)
	//fmt.Println(myarr3)

	//遍历数组
	test3(myarr2)
	test3(myarr3)
	//修改数组中的值,注意传进来的数组不能是空数组
	test(myarr2)
	fmt.Println(myarr2)

	fmt.Println(myarr5)
	fmt.Println(myarr6)

}
