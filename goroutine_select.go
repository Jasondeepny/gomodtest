package main

//import (
//	"fmt"
//	"strconv"
//	"time"
//)

//import (
//	"fmt"
//	"strconv"
//	"time"
//)

//func main()  {
//output1 := make(chan string)
//output2 := make(chan string)
//// 跑2个子协程，写数据
//go test1(output1)
//go test2(output2)
//// 用select监控
//select {
//case value := <-output1:
//	fmt.Println("s1=", value)
//case value := <-output2:
//	fmt.Println("s2=", value)
//}

//	output1 := make(chan string, 10)
//	// 子协程写数据
//	go write(output1)
//	// 取数据
//	for s := range output1 {
//		fmt.Println("res:", s)
//		time.Sleep(time.Second)
//	}
//}

//func write(ch chan string) {
//	for {
//		select {
//		// 写数据
//		case ch <- "hello":
//			fmt.Println("write",ch)
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//}
//
//func test1(ch chan string) {
//	//time.Sleep(time.Second * 5)
//	//ch <- "test1"
//	for i := 0; i < 100; i++ {
//		ch <- strconv.Itoa(i)
//	}
//}
//func test2(ch chan string) {
//	time.Sleep(time.Second * 2)
//	ch <- "test2"
//}
