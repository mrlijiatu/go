package main

import (
	"fmt"
	"reflect"
)

// assertion 把一个接口类型指定为他的原始类型
// 创建一个名为user的结构体
type User struct {
	Name string
	Age  int
	Sex  bool
}

// 创建一个名为student的结构体
type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名字叫做", name)
}

//func main() {
//	u := User{
//		Name: "malaoban",
//		Age:  18,
//		Sex:  true,
//	}
//	check(u)
//}
//
//func check(v interface{}) {
//	v.(User).SayName("malaoban")
//}

//把一个接口方面的断言成他的原始类型

func main() {
	u := User{
		Name: "malaoban",
		Age:  18,
		Sex:  true,
	}
	s := Student{"三年二班", u}

	check(&s)
}

//func check(v interface{}) {
//	switch v.(type) {
//	case User:
//		fmt.Println("我是user")
//	case Student:
//		fmt.Println("我是student")
//	}
//}

//反射reflect
//reflect.valueof(判断接口中的数据值)			普通反射
//reflect.typeof(动态获取接口中的值的类型)		stuct反射
//reflect.typeof（）.kind（） 用来判断类型		匿名货嵌入字段的反射
//reflect.valueof.findle（） 用来获取值		判断传入的类型是否是我们想要的类型
//reflect.fieldbyindex（[]int{0,1}）获取层级	通过反射修改内同
//reflect.valueof（）。elem获取原始数据并操作	通过反射调用方法

//	func check(inter interface{}) {
//		t := reflect.TypeOf(inter) //获取接口中的数据值
//		//v := reflect.ValueOf(inter) //获取接口中数据类型
//		ty := t.Kind() //判断返回类型
//		if ty == reflect.Struct {
//			fmt.Println("struct")
//			if ty == reflect.Int16 {
//				fmt.Println("int")
//				if ty == reflect.Bool {
//					fmt.Println("bool")
//				}
//			}
//		}

// 使用反射改变原始数据
// func check(inter interface{}) {
//
//		v := reflect.ValueOf(inter) //获取接口中数据类型
//		e := v.Elem()
//		e.FieldByName("Class").SetString("四年二班")
//		fmt.Println(inter)
//	}
//
// 调取原始数据上层的方法
func check(inter interface{}) {
	v := reflect.ValueOf(inter) //获取接口中数据类型
	m := v.Method(0)            //获取数据上层方法的地址
	fmt.Println(m)

	m.Call([]reflect.Value{reflect.ValueOf("马老板")}) //改变原始数据上层的方法
}

//按照位置取对应的值
//fmt.Println(t, v)
//fmt.Println(v.FieldByName("Class"))      //通过名称取值
//fmt.Println(v.FieldByIndex([]int{1, 1})) //通过层级取值，第一项的第一个值，取匿名函数
//for i := 0; i < t.NumField(); i++ {
//	fmt.Println(v.Field(i))
//}
