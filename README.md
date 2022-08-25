1. [打印 hello world ](./hello/hello.go)
2. [变量定义](./variable/variable.go)
3. [常量定义](./const/const.go)
4. [函数](./func/func.go)
5. [import导入包路径问题与init方法调用流程,别名导入](./import/main.go)  
   ![img.png](img.png)  
   ├── lib1  
   │ └── lib1.go  
   ├── lib2  
   │ └── lib2.go  
   └── main.go

6. [指针](./pointer/pointer.go)  
   引用传递,传递指针
7. [defer](./defer/defer.go)
8. [数组与动态数组定义操作](./array_slice/array_slice.go)
9. [数组切片定义与操作](slice/slice.go)
   ![img_1.png](img_1.png)
10. [map创建使用](./map/map.go)
    1. 创建
    2. 添加修改
    3. 删除
    4. 遍历
    5. 查找
    6. 往函数传参是引用传递
11. 面向对象
    1. [声明数据类型结构体](./object/struct/struct.go)
        1. 定义结构体
        2. 修改结构体的值
        2. 往函数传参结构体默认是值传递,引用传递按照定义普通变量方式
        3. [遍历结构体](./object/reflect/reflect.go)
    2. [类的表示与封装](./object/class/class.go)
        1. 类的创建
            1. 对象实例化
            2. 类名大写开头代表其他包可以访问,小写开头代表只有本包可以访问
            2. 类名小写开头代表只有本包可以访问,大写开头代表其他包可以访问
        2. 类方法创建,形参要用指针
    3. [面向对象继承](./object/inherit/inherit.go)
        1. 继承
            1. 子类继承父类
            2. 子类方法重写父类方法
            3. 对象实例化
    4. 面向对象多态接口
        1. [接口](./object/interface/interface.go)
            1. 接口定义
            2. 接口实现
            3. 接口使用 依赖注入,一个方法的形参接收变量,变量类型是一个接口
        2. [万能类型空接口](./object/empty_interface/empty_interface.go)
            1. 类型断言,判断类型
            2. 空接口 形参中类型位置是interface{} 例如: func (arg interface{})  
    5. 反射
       1. pair<type,value>
       2. 断言成功是pair中的type相同
          ![img_2.png](img_2.png)
       3. [反射reflect机制用法](./object/reflect/reflect.go)
          1. 获取一个变量的类型 reflect.TypeOf(变量)
          2. 获取一个变量的值 reflect.ValueOf(变量)
          3. 获取对象类型的方法集合,循环 
          4. 获取对象类型的变量集合,循环
       4. [反射解析结构体标签Tag](./object/reflect_tag/reflect_tag.go)
          1. json标签
12. goroutine 协程
    1. [基础创建](./object/goroutine/goroutine.go) 
    2. [协程参数通信-无缓冲,有阻塞](./object/goroutine_channel/goroutine_channel.go)
               ![img_4.png](img_4.png)
    3. [协程参数通信-有缓冲](./object/goroutine_channel_huanchong/goroutine_channel_huanchong.go)
               ![img_3.png](img_3.png)
13. 方法
    1. 小数加减乘除
    2. 整数加减乘除
    3. 字符串变更
    4. 字符串截取
    5. 字符串拼接
    6. 字符串替换
    7. [json字符串转对象](./object/reflect_tag/reflect_tag.go)
    8. [对象转json字符串](./object/reflect_tag/reflect_tag.go)