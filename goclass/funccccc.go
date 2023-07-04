package main

/*
func  方法 函数
函数内不允许创建函数
函数名首字母大写，可以被其他途径引用

//普函数，函数内不允许创建函数

	func main() {
		r1, r2 := a(0, "data2的数据")
		fmt.Println(r1, r2)
	}

	func a(data1 int, data2 string) (ret1 int, ret2 string) {
		ret1 = data1
		ret2 = data2
		return
	}
*/
//传入不定项参数进入到函数内部
/*
func main() {
	om(8778, "1", "2", "3", "4", "5")
}

func om(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
	//挨个取出不定项参数
	for k, v := range data2 {
		fmt.Println(k, v)
	}
}

func a(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return
}
*/
//自执行函数

/*
	func main() {
		(func() {
			fmt.Println("我是自执行")
		})()
	}
*/

/*
func main() {
	a()(6)
}

// 闭包函数，一个函数接受一个函数然后再返回这个函数
func a() func(num int) {
	return func(num int) {
		//调用
		fmt.Println("闭包函数调用", num)
	}
}
*/
