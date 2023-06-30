package main

import "fmt"

func main() {
	//[3]为数组长度+元素类型+元素1，元素2...
	a := [3]int{4, 5, 6}
	//[...]自动补充数组的长度
	b := [...]int{23, 7657, 9, 76856, 454, 86786}
	var d = new([10]int)
	d[4] = 9
	fmt.Println(a, b, d)
	//数组的李宇和数组的循环
	zoom := [...]string{"老虎", "狮子", "长颈鹿", "梅花鹿"}
	for i := 0; i < len(zoom); i++ {
		fmt.Println(zoom[i] + "跑")
	}
	//打印出数组内元素对应的坐标
	for i, v := range zoom {
		fmt.Println(i, v)
	}
	// len (zoom)=cap(zoom) 长度和容积 可以快速定位数组内的元素
	//多维数组，
	er := [3][3][1]int{
		{{1}, {2}, {3}},
		{{2}, {3}, {4}},
		{{3}, {4}, {5}},
	}
	fmt.Println(er)
}
