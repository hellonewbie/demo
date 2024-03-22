package main

import (
	"fmt"
	"reflect"
)

//今天上课的时候摸鱼看了一下Go的泛型，简要写几个例子吧，大概记录一下实现的过程
//写的很垃，大家不要介意就是我的个人笔记罢了

//用反射+接口来判断类型，个人觉得性能上可能有较大的差距，个人还是不建议这样用的

func Range3(arr interface{}) {
	//reflect.TypeOf()获得任意值的对象
	typeOfA := reflect.TypeOf(arr)
	fmt.Println(typeOfA, typeOfA.Kind())
	//然后对不不同的值遍历又要判断
	if typeOfA.String() == "[]int" {
		//感觉还是逃不过断言
		for _, v := range arr.([]int) {
			fmt.Println(v)
		}
	}
	if typeOfA.String() == "[]string" {
		for _, v := range arr.([]string) {
			fmt.Println(v)
		}
	}
}

//断言的方式完成泛型

func Range2(arr interface{}) {
	if _, ok := arr.([]string); ok {
		for _, v := range arr.([]string) {
			fmt.Println(v)
		}
	}
	if _, ok := arr.([]int); ok {
		for _, v := range arr.([]int) {
			fmt.Println(v)
		}
	}

}

//直接使用泛型，好处就是会自动匹配类型
//any可以点进去看一下其实本质就是interface{}，可以接收任何类型的数据
//里面放comparable,表示的是Go里面所有内置的可比较类型：`int、uint、float、bool、struct和指针等等

func Range[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

type MySlice[T int | string] []T

func (S MySlice[T]) Sum() T {
	var sum T
	for _, v := range S {
		sum += v
	}
	return sum
}

func Add[T int | string | float32](a, b T) T {
	return a + b
}

func main() {
	INT := []int{1, 2}
	STRING := []string{"xiaoyingtongxue", "xiaolintongxue"}
	//参数传进去之后会自动确认类型
	Range(INT)
	Range(STRING)
	Range2(INT)
	Range2(STRING)
	Range3(INT)
	Range3(STRING)

	//自定义泛型类型，[T ]里面的是类型约束
	//type MySlice[T int | string] []T
	////泛型类型必须实例化为具体类型
	//var a MySlice[int] = []int{1, 2, 3}
	//var b MySlice[string] = []string{"xiaoyingtongxue", "xiaolint"}
	//fmt.Println(a, b)

	type MyMap[KEY int | string, VALUE int | string] map[KEY]VALUE
	var V1 MyMap[int, string] = map[int]string{
		1: "xiaoyingtongxue",
		2: "xiaolintongxue",
	}
	var V2 MyMap[string, string] = map[string]string{
		"1": "xiaoyingtongxue",
		"2": "xiaolintongxue",
	}
	fmt.Println(V1, V2)

	var myslice MySlice[int] = []int{1, 2, 3, 4, 5}
	fmt.Println(myslice.Sum())
	//可以不加泛型约束，会自动转成下面这种，就是一种语法糖
	fmt.Println(Add(1, 2))
	fmt.Println(Add[int](1, 2))
	fmt.Println(Add[string]("1", "2"))
}
