package main

import "fmt"

//定义一个类(结构体) 人类
//如果首字母大写，则表示其他包可见
type Human struct {
	//定义类属性 如果首字母大写，则表示其他包可见,如果首字母小写，则表示该属性对外不可见
	//名字
	Name string
	//年龄
	Age int
}

//定义一个超人类 继承人类
type SuperMan struct {
	//继承人类属性,不管是否大写，都可以继承
	Human
	//超能力
	SuperPower string
}

//获取名字方法
func (h *Human) GetName() string {
	return h.Name
}

//重新定义获取名字方法
func (h *SuperMan) GetName() string {
	return h.Name + " " + h.SuperPower
}

//定义吃方法
func (h *SuperMan) Eat() string {
	return "吃饭"
}

func main() {
	//定义一个人类,创建类的实例 初始化属性的值,可以不用全部初始化
	var humanVariable Human = Human{
		Name: "宋江",
		Age:  10,
	}
	fmt.Println(humanVariable)

	//定义一个超人类,是人类的子类,创建类的实例 初始化属性的值,可以不用全部初始化
	var superManVariable SuperMan = SuperMan{
		//继承人类属性,不管是否大写，都可以继承,方法也可以继承
		Human: Human{
			Name: "宋江",
		},
		SuperPower: "会飞",
	}
	//获取超人名字,调用父类的方法,子类如果有重写方法，则调用子类的方法
	fmt.Println(superManVariable.GetName())

	//另一种定义方式
	var superManVariable2 SuperMan
	//设定父类属性
	//superManVariable2.Human.Name = "张飞"
	//简写方式
	superManVariable2.Name = "张飞"
	//设定父类属性age
	superManVariable2.Age = 20
	//设定子类属性
	superManVariable2.SuperPower = "会遁地"
	fmt.Println(superManVariable2.GetName())

}
