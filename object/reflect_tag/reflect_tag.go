package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	//结构体标签就是类似 `json:"name" info:"这是说明信息"` 这样的标签,info,json自定义字段名
	Name string `json:"name"`
	//json标签  year 代表在转换为json中显示的字段名
	Year  int `json:"year"`
	Price int
	//演员数组 字符串数组
	Actors []string
}

func main() {
	m := Movie{
		Name:   "西游记",
		Year:   2019,
		Price:  100,
		Actors: []string{"吴承恩", "白马王子", "猪八戒"},
	}

	//将结构体转换为json格式 对象转json字符串
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("jsonStr:", string(jsonStr))
	//格式化输出
	fmt.Printf("jsonStr:%s: \n", jsonStr)

	//将json格式转换为结构体 字符串转json对象
	//定义一个空的结构体,后面用于接受json转换的结构体
	newMovie := Movie{}
	//将json转换为结构体
	err = json.Unmarshal(jsonStr, &newMovie)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("newMovie:", newMovie)
}
