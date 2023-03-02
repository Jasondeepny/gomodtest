package main

import "fmt"

//
//import (
//	"fmt"
//)

//func main() {
// 合起来写
//go func() {
//	i := 0
//	for {
//		i++
//		fmt.Printf("new goroutine: i = %d\n", i)
//		time.Sleep(time.Second)
//	}
//}()
//i := 0
//for {
//	i++
//	fmt.Printf("main goroutine: i = %d\n", i)
//	time.Sleep(time.Second)
//	if i == 2 {
//		break
//	}
//}
//
//var wg sync.WaitGroup
//wg.Add(10)
//for i := 0; i < 10; i++ {
//	go func(i int) {
//		defer wg.Done()
//		fmt.Println("goroutine",i)
//	}(i)
//}
//wg.Wait()
//
//for i := 0; i < 10; i++ {
//	wg.Add(1) // 启动一个goroutine就登记+1
//	go hello(i)
//}
//wg.Wait()

//go func(s string) {
//	for i := 0; i < 2; i++ {
//		fmt.Println(s)
//	}
//}("world")
//// 主协程
//runtime.Gosched()
//for i := 0; i < 2; i++ {
//	// 切一下，再次分配任务
//	fmt.Println("hello")
//}

//go func() {
//	defer fmt.Println("A.defer")
//	func() {
//		defer fmt.Println("B.defer")
//		// 结束协程
//		runtime.Goexit()
//		defer fmt.Println("C.defer")
//		fmt.Println("B")
//	}()
//	fmt.Println("A")
//}()
//for {
//}
//ch := make(chan int)
//go recv(ch) // 启用goroutine从通道接收值
//ch <- 10
//fmt.Println("发送成功")

//ch1 := make(chan int)
//ch2 := make(chan int)
// 开启goroutine将0~100的数发送到ch1中
//count(ch1)
//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//go func() {
//	for {
//		i, ok := <-ch1 // 通道关闭后再取值ok=false
//		if !ok {
//			break
//		}
//		ch2 <- i * i
//	}
//	close(ch2)
//}()
//// 在主goroutine中从ch2中接收值打印
////s := 0
////for i := range ch2 { // 通道关闭后会退出for range循环
////	fmt.Println(i)
////	s++
////}
////fmt.Println("数量",s)
//fmt.Println("接收成功")

//go counter(ch1)
//go squarer(ch2, ch1)
//printer(ch2)
//}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func count(ch1 chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
}
