package main

import "fmt"

//	map的声明方式
func main() {

	//1. 无初始值	常用
	//变量名:=make(map[key类型]value类型)
	m2 := make(map[string]string)
	m2["name"] = "李四"
	m2["age"] = "20"
	fmt.Println(m2)

	//2.常用 带初始值
	//变量名 := map[key类型]value类型{
	// 	key1:value1,
	// 	key2:value2,
	//}
	m3 := map[string]string{
		"name": "王五",
		"age":  "22",
	}
	fmt.Println(m3)

	//3.无初始值  不常用
	//var  变量名 map类型[key是string]value是int
	var m map[string]string
	if m == nil {
		fmt.Println("map 是空")
	}
	//在使用之前必须make
	m = make(map[string]string)
	m["name"] = "张三"
	m["age"] = "18"
	fmt.Println(m)

	//遍历map
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	//key不使用的时候，可以使用_
	for _, v := range m2 {
		fmt.Println(v)
	}

	//删除map中的元素
	delete(m2, "name")
	//允许删除不存在的key
	delete(m2, "nam")
	for k, v := range m2 {
		fmt.Println("删除后key =", k)
		fmt.Println("删除后value =", v)
	}
	//map传参是引用传递,修改会影响到原来的map
	fmt.Println("修改前", m2)
	changeMap(m2)
	fmt.Println("修改后", m2)

}

//map作为函数的参数,传递的是map的地址,在函数中修改map的值,会影响到原来的map
func changeMap(mapvariable map[string]string) {
	//修改map的值
	mapvariable["age"] = "19"

}
