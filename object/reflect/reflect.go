package main

import (
	"fmt"
	"reflect"
)

//定义一个类(结构体) 英雄
type Hero struct {
	Name   string
	Attack int
}

//获取英雄名字方法
func (h *Hero) GetName() string {
	return h.Name
}

func main() {
	var f float64 = 1.232
	//获取f的反射类型
	fmt.Println(reflect.TypeOf(f))
	//获取f的反射值
	fmt.Println(reflect.ValueOf(f))

	//实例化一个英雄
	hero := Hero{
		Name:   "宋江",
		Attack: 10,
	}

	//获取hero的反射类型
	heroTypeOf := reflect.TypeOf(hero)
	//打印hero的反射类型
	fmt.Println("heroTypeOf:", heroTypeOf)

	//获取hero的反射值
	heroValueOf := reflect.ValueOf(hero)
	//打印hero的反射值
	fmt.Println("heroValueOf:", heroValueOf)

	//获取对象类型的方法集合,循环
	for i := 0; i < heroTypeOf.NumMethod(); i++ {
		m := heroTypeOf.Method(i)
		//打印方法名
		fmt.Println("Method name: ", m.Name)

	}

	//获取对象类型的变量集合,循环
	for i := 0; i < heroTypeOf.NumField(); i++ {
		f := heroTypeOf.Field(i)
		//打印变量名
		fmt.Println("Field name: ", f.Name)
		//打印变量类型
		fmt.Println("Field type: ", f.Type)
		//打印变量值
		fmt.Println("Field value: ", heroValueOf.Field(i).Interface())
	}

}
