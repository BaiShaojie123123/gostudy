package main

import "fmt"

//接口定义 interface 本质是指针
//实现接口要实现接口的所有方法
type Animal interface {
	Sleep()
	GetColor() string
}

//定义具体的类,继承接口不需要写到struct中
type Dog struct {
	name string
}

//实现接口睡觉 *Dog 才可以修改Dog的属性
func (d *Dog) Sleep() {
	fmt.Println("dog sleep")
}

//实现接口获取颜色 *Dog 才可以修改Dog的属性
func (d *Dog) GetColor() string {
	return "black"
}

type Cat struct {
	name string
}

func (c *Cat) Sleep() {
	fmt.Println("cat sleep")
}

func (c *Cat) GetColor() string {
	return "black"
}

//类似依赖注入,规定形参必须是接口类型的指针,所以传进来的变量必须实现接口的所有方法,并且传递的是地址
func showAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
}

func main() {

	//第一种方式
	//定义一个接口类型的变量,指针类型,这个变量的类必须实现接口的所有方法
	//var dogVariable Animal
	////实例化一个类,&Dog 传递的是Dog的地址才可以实例化接口
	//dogVariable = &Dog{
	//	name: "dog",
	//}

	//第二种方式
	//定义一个接口类型的变量,指针类型,这个变量的类必须实现接口的所有方法
	//var dogVariable Animal = &Dog{
	//	name: "dog",
	//}

	//dogVariable.Sleep()
	//fmt.Println(dogVariable.GetColor())

	//第三种方式
	dog := Dog{
		name: "dog",
	}
	//形参变量是接口类型的,接口本质是指针,所以要传地址.
	//形参变量是接口类型的,传递的变量必须是实现了接口的所有方法
	showAnimal(&dog)

}
