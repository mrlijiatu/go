package main

import "fmt"

// 结构体 map 数组 都是复杂的数据类型，用结构需要先声明,声明之后可以和其他数据类型一样使用，也有指针 值
type maxinfeng struct {
	Name   string
	age    int
	Sex    bool
	hobbys []string
	myhome home
}

// 给结构体声明一个方法
// 结构体作为参数写在前面，方法名 +入参+出餐+具体方法
func (m *maxinfeng) song(name string) (restr string) {
	restr = "真不错"
	//使用 fmt.Printf 打印后会换行，使用fmt.Print()打印不会换行，fmt.Printf()格式化打印
	fmt.Printf("%v唱了一首%v,观众觉得%v", m.Name, name, restr)
	return restr
}

type home struct {
	p string
}

// 声明1
func main() {
	//声明结构体里面的内容1
	//var mx maxinfeng
	//mx.name = "马老板"
	//mx.age = 18
	//mx.Sex = true
	//mx.hobbys = []string{"拳击", "游泳", "唱歌"}
	//fmt.Println(mx)
	//声明结构体内的内容2（隐式声明
	mx := maxinfeng{
		Name:   "会唱歌的马老板",
		age:    18,
		Sex:    true,
		hobbys: []string{"拳击", "游泳", "唱歌"},
	}

	re := mx.song("《我会等》")
	fmt.Println("\n", re)
	//mxFunc(mx)
	//
	//声明一个空的结构体 new出来的结构体是一个地址，可以直接进行赋值
	//mx := new(maxinfeng)
	//mx.name = "马老板"
	//fmt.Println(mx)

	//结构体的指针与值修改，创建一个指针，将指针指向结构体的地址，通过指针修改结构体内的参数。go可以通过内存地址修改数据类型的值，仅限于复杂的数据类型（map 数组 结构体
	//var mxp *maxinfeng
	//mxp = &mx
	//mxp.age = 19
	//fmt.Println(mx)
}

//结构体可以引入到函数中取
//func mxFunc(qm maxinfeng) {
//	fmt.Println(qm)
//}
//
