package main

import "fmt"

//今天做题看了一下可选参数和默认参数
//下面我大概写一下我自己的认知吧
//就是默认参数就是不设定的时候这个参数有默认值，这个参数在函数内部是要使用的
//可选参数的认识：我认为就是多个默认参数选几个用用
//因为在Go里面没有这个概念，那么我们就需要用另一种方式来实现这个东西
//看代码

const (
	Sname = "小王"
	Ssex  = "男"
	Sage  = 10
)

type Select struct {
	name string
	sex  string
	age  int
}

type Option struct {
}

type option func(s *Select)

func (op *Option) Name(name string) option {
	return func(s *Select) {
		s.name = name
	}
}

func (op *Option) Sex(sex string) option {
	return func(s *Select) {
		s.sex = sex
	}
}

func (op *Option) Age(age int) option {
	return func(s *Select) {
		s.age = age
	}
}

func InitSel(options ...option) Select {
	//一般情况下我们要给结构体赋值就是实例化一个对象,给属性赋值，我们希望有些值是一直保持不变的
	//不想每次都是传递参数来赋值->默认参数
	//声明一个常量表，指定值然后调用
	Sel := &Select{
		name: Sname,
		sex:  Ssex,
		age:  Sage,
	}
	//但是有的时候我们又不想用默认的值怎么般呢？这个时候你有又会想到直接传递参数然后赋值上去，那默认参数的配置的意义何在
	//不就是为了减少参数便于读懂
	//为不同的可选参数定义不同函数可以使得api更清晰和理解,然后我们就会想起了自定义类型了,到时候要选的时候直接把函数传进去
	//然后遍历运行一下就行
	for _, value := range options {
		value(Sel)
	}
	return *Sel

}

func main() {
	//这个时候在像能不能归类呢？就是我直接把这个几函数归类到一个可选参数列表然后像选的时候直接掉用然后初始化
	//这个时候我想到了一个好用的方法就是直接写成是一个空结构体的方法，归类到这个空结构体，然后我们只需要实例化一个对象之后就可以直接调用对应的方法
	Op := new(Option)
	//这个就是可选参数的修改
	SetSex := Op.Name("xiaoli")
	//调用函数进行修改
	Sel := InitSel(SetSex)
	fmt.Println(Sel)
}
