package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"reflect"
//	"runtime"
//)

//func main() {

//slice := []int{10, 20, 30, 40}
//fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//newSlice := append(slice, 50)
//fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//newSlice[1] += 10
//fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

//array := [4]int{10, 20, 30, 40}
//fmt.Printf("Before array = %v, Pointer = %p, len = %d, cap = %d\n", array, &array, len(array), cap(array))
//slice := array[0:2]           //10,20
//newSlice := append(slice, 50) // 10,20,50
//fmt.Printf("After array = %v, Pointer = %p, len = %d, cap = %d\n", array, &array, len(array), cap(array))
//fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//fmt.Printf("Before newS = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//newSlice[1] += 10
//fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//fmt.Printf("After newS = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//fmt.Printf("After array = %v\n", array)

//var a = [...]int{1, 2, 4, 5, 6}
////b := a[:] //相当于 var slice_b []int = arr[:] （引用类型不改变内存地址） 切片在原数组后追加元素
//b := a //var b[]int = var a[]int  (值类型直接拷贝,内存地址变化)
//b[1] = 10000
//fmt.Println(a)
//fmt.Println(b)
//fmt.Printf("%p\n", &a)
//fmt.Printf("%p\n", &b)
//fmt.Printf("%p\n", &b)

//s := []int{1, 2, 3, 4}
//fmt.Printf("type of s:%T", s)

//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//s1 := arr[2:6] //2 3 4 5
//fmt.Println(s1[3:6])
//fmt.Printf("Before array = %v, Pointer = %p, len = %d, cap = %d\n", s1, &s1, len(s1), cap(s1))
//s2 := s1[3:5]
//fmt.Println(s2)

//printFile("abc.text")
//
//q, r := div(9, 4)
//fmt.Println(q, r)

//fmt.Println(apply(func(a, b int) int {
//	return int(math.Pow(float64(a), float64(b)))
//}, 3, 4))

//a, b := 3, 4
////swap(&a, &b)
////fmt.Println(a, b)
//swapp(a, b)
//fmt.Println(a, b)
//}
//
//func swapp(a, b int) {
//	a, b = b, a
//	fmt.Println(a, b)
//}
//
//func swap(a, b *int) {
//	*a, *b = *b, *a
//}
//
//func apply(f func(x int, y int) int, a, b int) int {
//	pointer := reflect.ValueOf(f).Pointer()
//	opName := runtime.FuncForPC(pointer).Name()
//	fmt.Printf("Calling func %s with args + (%d + %d) \n", opName, a, b)
//	return f(a, b)
//}
//
//func div(a, b int) (q, r int) {
//	q = a / b
//	r = a % b
//	return q, r
//}
//
//func printFile(n string) {
//	filename, err := os.Open(n)
//	if err != nil {
//		panic(err)
//	}
//	newScanner := bufio.NewScanner(filename)
//
//	for newScanner.Scan() {
//		fmt.Println(newScanner.Text())
//	}
//}
