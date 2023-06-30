package main

import "fmt"

// 死循环
/*func main() {
	z1 := 0
	for {
		z1++
		fmt.Println(z1)
		if z1 > 5 {
			break
		}

	}
}
*/

/* for循环1 先判断在加加
func main() {
	for b := 0; b < 5; b++ {
		fmt.Println(b)
	}
}
*/
/*for循环2 先++后判断
func main() {
	a := 0
	for a < 5 {
		a++
		fmt.Println(a)
	}
}
*/

/*
// 跳转语句break可与多层循环嵌套跳出/continue跳过本次循环
func main() {
	a := 0
	for a < 9 {
		a++
		if a == 5 {
			continue
		}
		fmt.Println(a)
	}
}
*/
//goto 跳转到某个循环

func main() {
A:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 5 {
				break A
				goto B
			}
		}
	}
B:
	fmt.Println("我来b了")
}
