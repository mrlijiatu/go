package main

import "fmt"

// goroutine 是一个函数，它能让函数进行异步执行,可以通过断点来看异步执行的过程
/*func main() {
	go run()
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)

	}
}

func run() {
	fmt.Println("我跑起来了")
}
*/
//携程管理器在sync包里面
/*
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go run(&wg)
	wg.Wait()
}

func run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}
*/

/*
channel 是goroutine之间通讯的桥梁
chan分为五种
（可读可取c ：=make（chan int

	可读 var readChan <-chan int =c
	可取 var readChan chan<- int =c
	有缓冲 c:=make(chan int,5)
	无缓冲 c：=make（chan int)
*/
//有缓冲区，可以将信息存到缓冲区，然后在取出来
//func main() {
//	c1 := make(chan int, 1)
//	c1 <- 1
//	fmt.Println(<-c1)
//}
//无缓冲区

//	func main() {
//		c1 := make(chan int)
//		go func() {
//			for i := 0; i < 10; i++ {
//				c1 <- i
//			}
//
//		}()
//
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c1)
//		}
//
// }
// 可读，可取
//func main() {
//	c1 := make(chan int, 5)
//	var readc <-chan int = c1
//	var write chan<- int = c1
//	write <- 1
//	fmt.Println(<-readc)
//
//}

//channo可以执行关闭命令

//func main() {
//	c1 := make(chan int, 5)
//	c1 <- 1
//	c1 <- 2
//	c1 <- 3
//	c1 <- 4
//	c1 <- 5
//	close(c1)
//	//可以 通过for range来取channo里面的数值
//
//	for i := range c1 {
//		fmt.Println(i)
//	}
//
//}

//selecrt
//
//func main() {
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	ch3 := make(chan int, 1)
//	select {
//	case <-ch1:
//		fmt.Println("ch1")
//	case <-ch2:
//		fmt.Println("ch2")
//	case <-ch3:
//		fmt.Println("ch2")
//	default:
//		fmt.Println("xx")
//
//	}
//}

// chanon和gorunter实现内部通信
func main() {
	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c
	go setchan(writec)
	getchan(readc)
}
func setchan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		writec <- i
		fmt.Printf("我在个set函数内")
	}
}
func getchan(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("我在个get函数内，我是%d\n", <-readc)
	}
}
