package main

import "fmt"

//定义一个类(结构体) 英雄
//如果首字母大写，则表示其他包可见
type Hero struct {
	//定义类属性 如果首字母大写，则表示其他包可见,如果首字母小写，则表示该属性对外不可见

	//名字
	Name string
	//攻击力
	Attack int
	//防御力
	Def int
	//生命值
	Hp int
	//等级
	Level int
}

//获取英雄名字方法

func (h *Hero) GetName() string {
	return h.Name
}

//设置英雄名字方法
// (h表示形参名 *Hero表示形参类型要用指针接收,否则方法中不能修改形参的值))
func (h *Hero) SetName(name string) {
	h.Name = name
}

func main() {
	//定义一个英雄,创建类的实例 不初始化属性的值
	var hero Hero
	//设置英雄名字
	hero.SetName("宋江")
	//获取英雄名字
	fmt.Println(hero.GetName())

	//定义一个英雄,创建类的实例 初始化属性的值,可以不用全部初始化
	var hero1 Hero = Hero{
		Name:   "宋江",
		Attack: 10,
	}
	fmt.Println(hero1)

}
