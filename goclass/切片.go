package main

import "fmt"

// 切片是数组的一部分，可以通过切片的方式改变数组内的元素
/*func main() {
	a := []int{1, 2, 3}
	cl := a[0:1] //左开右闭
	cl[0] = 5
	fmt.Println(cl)
	fmt.Println(a)
}
*/
//取切片的方式
//[:]取数组里面所有的
//[:5]从第0位取到第4位
//[5:]从第4位取到最后一位
//[3:5]从第2位取到第三位
//切片append
/*
func main() {
	a := [3]int{0, 1, 2}
	cl := a[1:] //左开右闭
	cl = append(cl, 9)
	cl[0] = 4
	//先append在改变切片的数值不会改变原数组
	//先改变原切片数值，再append原数组数值会被改变，在append之后再对其操作就不会改变原数组
	fmt.Println(a, cl)
}
*/
//切片的len cap的值不固定
//可以直接创造切片，不对数组的长度赋值即可
//切片copy把第二个切片里的东西拷贝到第一个切片进行覆盖
//打印一个make的切片会显示默认值，默认值为0
//aa :=make([]int,4)
func main() {
	a := [3]int{0, 1, 2}
	cl := a[:]
	cl1 := a[2:]
	cl = append(cl, 9)
	copy(cl, cl1)
	//先append在改变切片的数值不会改变原数组
	//先改变原切片数值，再append原数组数值会被改变，在append之后再对其操作就不会改变原数组
	aa := make([]int, 4)
	fmt.Println(cl, cl1, aa)
}
