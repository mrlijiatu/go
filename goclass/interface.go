package main

//go接口是一类东西的的集合，只要实现这个接口里面的方法，就可以说属于这个接口
import "fmt"

type animal interface {
	Eat()
	Run()
}

type cat struct {
	Name string
	Sex  bool
}

type dog struct {
	Name string
}

func (d dog) Eat() {
	fmt.Println(d.Name, "开始吃")
}

func (d dog) Run() {
	fmt.Println(d.Name, "开始跑")
}

func (c cat) Run() {
	fmt.Println(c.Name, "开始跑")
}

func (c cat) Eat() {
	fmt.Println(c.Name, "开始跑")
}

// 只要实现这个接口的结构体，都可以用这个接口声明
/*func main() {
	var a animal
	//声明一个接口。声明一个符合接口规范的结构体c，让c去实现这个接口，于是就可以通过接口调用c内部的方法
	c := cat{
		Name: "Tom",
		Sex:  false,
	}
	a = c
	a.Run()
	a.Eat()
}
*/
//使用一个泛型 （创建一个func，里面传入一个空接口）

/*func main() {
	myfunc(1)
	myfunc("dasdasdsa")
	myfunc(([]string{"12", "212"}))
}
func myfunc(a interface{}) {
	fmt.Println(a)
}
*/
//实现接口里定义的操作
/*
func main() {
	c := cat{
		Name: "tom",
		Sex:  false,
	}
	myfunc(c)

}
func myfunc(a animal) {
	a.Run()

}
*/
//进行一种解耦合的操作 有一个全局的变量，想让这个变量做固定的事情（在实现需求里去调用别人的方法
var L animal

func main() {
	c := cat{
		Name: "tom",
		Sex:  false,
	}
	myfunc(c)
	L.Run()

}
func myfunc(a animal) {
	L = a

}
