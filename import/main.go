package main

//导入lib1包 lib2bao
import (
	//gopath全路径才可以导入,导入后必须使用包名,否则
	//"gostudy/import/lib1"
	//"gostudy/import/lib2"

	//"gostudy/import/lib2"

	//别名导入,可以导入包但不使用包,这时包的init方法还是会执行
	//_ "gostudy/import/lib2"

	//别名 定义了以后必须使用
	mylib1 "gostudy/import/lib1"

	//导入包下所有方法不建议使用!直接使用包中的方法
	. "gostudy/import/lib2"
	//还可以用mod方式导入
)

func main() {
	//普通调用 "gostudy/import/lib2"
	//	lib2.Lib2test()

	//别名  mylib1 "gostudy/import/lib1"
	mylib1.Lib1test()

	//直接调用包下的方法  导入 . "gostudy/import/lib2"
	Lib2test()
}
