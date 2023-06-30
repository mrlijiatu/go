package main

import (
	"fmt"
	"goclass/testpkg"
)

func main() {
	var a string = "7788"
	//关键字 变量名 变量类型 =变量值
	b := "2020"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(testpkg.A)
	fmt.Printf("%T", a)
}

//tips
// 01,引入包时可以在引入的包前面增加一个.如."goclass/testpkg"，这样在使用包里面的东西时就不需要输入包名了
// 02,go语言的小特性，在包里面写的变量名、方法名必须以大写字母开头才可以被引用。小写字母开头的是私有的，不能被引用的
// 03,一个文件夹下只能有一个包，但是可以有多个文件。即package后面的后缀需要一致
/* 04,函数类
int/整数型 有int8包含2的8次方个数字，一半为负数一半为0和正数（-128~127）、int16(-32768~32767)、int32、int64
uint/正整数 有uint8、uint16、uint32、uint62，从2的8次方到2的64次方，个数字从0开始
float/浮点型，表示小数 有float32，float64
string/字符串，需要用双引号引起来
bool/布尔型 true 、false
查看变量数据类型 	fmt.Printf("%T", a)
2023年6月30日14:31:01
*/

//变量名只能以字母开头
/*
死循环：


*/
//int的默认值是0
//切片是数组的一部分，可以通过切片的方式改变数组内元素的值
//切片内可以增加
//问题录入
/*01
问题描述：同个包下的不容文件不能相互调用函数
原因：在 run/debug configurations 设置中run kind(运行类型)选择了File(文件)类型，在运行go程序时，
只会编译指定文件(main.go)，其他文件不会编译，所以main.go不能调用其他文件的函数
解决方案：在run->edit Configuration界面设置run kind(运行类型)为package
package path为main，即src源文件夹下的main包的相对路径(相对于源文件夹src)2023年6月30日11:50:44
*/
