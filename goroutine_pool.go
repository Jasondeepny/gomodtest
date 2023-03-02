package main

//
//import (
//	"fmt"
//	"math/rand"
//)
//
//func main()  {
//	jobChan := make(chan *Job,128)
//	resultChan := make(chan *Result,128)
//
//	createPool(64,jobChan,resultChan)
//
//	go func(resultChan chan *Result) {
//		for result := range resultChan {
//			fmt.Printf("job id is %#v,job random is %#v, result is %#v\n",
//				result.job.id, result.job.randNum, result.sum)
//		}
//	}(resultChan)
//
//	var num int
//	for {
//		num++
//		if num > 1000 {
//			break
//		}
//		job := &Job{
//			id:      num,
//			randNum: rand.Int(),
//		}
//		jobChan <- job
//	}
//}
//
//func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
//	for i := 0; i < num; i++ {
//		go func(jobChan chan *Job, resultChan chan *Result) {
//			for job := range jobChan {
//				// 随机数接过来
//				rNum := job.randNum
//				// 随机数每一位相加
//				// 定义返回值
//				var sum int
//				for rNum != 0 {
//					tmp := rNum % 10
//					sum += tmp
//					rNum /= 10
//				}
//				// 想要的结果是Result
//				r := &Result{
//					job: job,
//					sum: sum,
//				}
//				//运算结果扔到管道
//				resultChan <- r
//			}
//		}(jobChan, resultChan)
//	}
//}
//
//
//type Job struct {
//	// id
//	id int
//	// 需要计算的随机数
//	randNum int
//}
//
//type Result struct {
//	// 这里必须传对象实例
//	job *Job
//	// 求和
//	sum int
//}
