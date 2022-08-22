package main

import "fmt"

//定义接口
type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	name string
}

func (d Dog) Eat() {
	fmt.Println("eat")
}

func (d Dog) Run() {
	fmt.Println("run")
}

func (d Dog) Sleep() {
	fmt.Println("sleep")
}

//猫
type Cat struct {
	name string
}

func (c Cat) Eat() {
	fmt.Println("eat")
}

func (c Cat) Run() {
	fmt.Println("run")
}

func main() {

	var a Animal
	a = Dog{"dog"}
	a.Eat()
	a.Run()

	//
}
