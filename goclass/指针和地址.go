package main

//指针与变量 每一个变量的声明都会产生一个地址
/*

func main() {
	var a int
	//声明一个变量
	a = 123
	fmt.Println(a)
	//声明一个指针
	var b *int
	//另指针等于变量a的原始地址
	b = &a
	//改变当前地址的内容
	*b = 999
	//查看是否被改变
	fmt.Println(a, b, *b)
	fmt.Println(a == *b)
}
*/
//
/*数组指针
func main() {
	var arr [5]string
	arr = [5]string{"1,2,3,4,5"}
	fmt.Println(arr)
	var arrp *[5]string
	arrp = &arr
	fmt.Println(arr, arrp)
}
*/
/*指针数组，把一些地址放进一个数组内，可以通过指针来进行取值和变更
func main() {
	var arrpp [5]*string
	var str1 = "str1"
	var str2 = "str2"
	var str3 = "str3"
	var str4 = "str4"
	arrpp = [5]*string{&str1, &str2, &str3, &str4}
	*arrpp[2] = "999"
	for k, v := range arrpp {
		fmt.Println(k, v)
	}

	fmt.Println(str3)
}
*/
//指针传参 创建一个指针方法，用此方法来改变地址的参数，达到了改变原有变量的参数。各类函数都适用（map，数组，字符串
/*
func main() {
	var str1 = "我定义了"
	point(&str1)
	fmt.Println(str1)
}
func point(p1 *string) {
	*p1 = "wobianle"
}
*/
//隐式定义指针
func main() {
	var str1 = "yinshidengdai"
	//:=赋值
	p := &str1

	*p = "bianle"
}
