package main

import (
	"fmt"
	"testing"
)

func main() {
	//greeter := gophersGreeter{"Hello", "gophers"}
	//greeter.greet("Hello", "gophers")
	//var sss float32 =  3.4
	//var c1 complex64 = 5 + 10i
	//fmt.Printf("The value is: %v", c1)
	//fmt.Printf("%1.3e",sss)

	//c := complex(5, 10)
	//fmt.Printf("实部：%v",real(c))
	//fmt.Printf("虚部: %v",imag(c))
	//const hardEight = (1 << 100) >> 97
	//fmt.Print(hardEight)

	//traversalString()

	//var str = [5]string{3: "hello world", 4: "tom"}
	//
	//var arr1 = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	//
	//s := [...]struct {
	//	id   int
	//	name string
	//}{
	//	{1, "miss"},
	//	{2, "hdh"},
	//}
	//
	//fmt.Println(str, s, arr1)

	//var arr1 [5]int
	//printArr(&arr1)
	//fmt.Println(arr1)
	////arr2 := [...]int{2, 4, 6, 8, 10}
	//arr2 := arr1
	//printArr(&arr2)
	//fmt.Println(arr2)

	//transString()

	//rand.Seed(time.Now().Unix())
	//var b [5]int
	//for i := 0; i < len(b); i++ {
	//	b[i] = rand.Intn(1000)
	//	fmt.Println(b[i])
	//}
	//sum := sumArr(b)
	//fmt.Printf("sum=%d", sum)

	//b := [5]int{1, 3, 5, 8, 7}
	//myTest(b, 8)

	//p := new([]int)
	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Printf("arr : %p\n", &arr)
	////var s6 []int
	//// 前包后不包
	//s6 := arr[:3]
	//fmt.Printf("s6 : %p\n", &s6)
	//s7 := append(s6, 6)
	//fmt.Printf("s7 : %p\n", &s7)
	//s7[2] = 4
	////s6 = arr[:len(arr)-1]
	//fmt.Println(arr)
	//fmt.Println(s6)
	//fmt.Println(s7)
	//fmt.Printf("%p\n", &arr)
	//fmt.Printf("%p\n", &s6)
	//fmt.Printf("%p\n", &s7)

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
	//
	//var a = [...]int{1, 2, 4, 5, 6}
	////b := a[:] //相当于 var slice_b []int = arr[:] （引用类型不改变内存地址） 切片在原数组后追加元素
	//b := a //var b[]int = var a[]int  (值类型直接拷贝,内存地址变化)
	//b[1] = 10000
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p\n", &b)
	//fmt.Printf("%p\n", &b)

	//s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	//fmt.Println(s1, len(s1), cap(s1))
	//v := make([]int, 2, 4)
	//fmt.Println(len(v), cap(v))

	//s := []int{0, 1, 2, 4}
	//p := &s[3] // *int, 获取底层数组元素指针。
	//fmt.Println(*p)
	//*p += 100
	//fmt.Println(s)

	//d := [5]struct {
	//	x int
	//}{
	//	{2},
	//}
	//s := d
	//d[1].x = 10
	//s[2].x = 20
	//fmt.Println(d)

	//var a = []int{1, 2, 3}
	//fmt.Printf("slice a : %v\n", a)
	//var b = []int{4, 5, 6}
	//fmt.Printf("slice b : %v\n", b)
	//c := append(a, b...)
	//fmt.Printf("slice c : %v\n", c)

	//s1 := make([]int, 1, 5)
	//fmt.Println(s1)
	//fmt.Printf("%p\n", &s1)
	//
	//s1 = append(s1, 1)
	////fmt.Printf("%p\n", &s2)
	//
	//fmt.Println(s1)

	//s1 := make([]int, 1, 5)
	//fmt.Println(s1)
	//fmt.Printf("%p\n", &s1)
	//
	//s2 := append(s1, 1)
	//fmt.Println(s2)
	//fmt.Printf("%p\n", &s2)

	//data := [...]int{0, 1, 2, 3, 4, 10: 0}
	//s := data[:2:3]
	//fmt.Println(s)
	//s = append(s, 100, 200)      // 一次 append 两个值，超出 s.cap 限制。
	//fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	//fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

	//var a = []int{1, 3, 4, 5}
	//fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	//b := a[1:2]
	//fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	//fmt.Println(&b[0])
	//c := b
	//fmt.Println(&c[0])
	//fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))

	//str := "Hello world"
	//s := []byte(str) //中文字符需要用[]rune(str)
	//s[6] = 'G'
	//fmt.Println(string(s))
	//s1 := s[:8]
	//s1[1] = 'K'
	//fmt.Println(string(s1))
	//fmt.Println(string(s))
	//s = append(s1, '!')
	//str = string(s)
	//fmt.Println(str)

	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//d1 := slice[6:8]
	//fmt.Println(d1, len(d1), cap(d1))
	//d2 := slice[:6:9]
	//fmt.Println(d2, len(d2), cap(d2))
	//
	////数组或者切片转字符串
	//s := strings.Replace(strings.Trim(fmt.Sprint(d1), "[]"), " ", ",", -1)
	//
	//fmt.Println(s)
	//
	//fmt.Println(fmt.Sprint(d1))
	//
	//fmt.Println(strings.Trim(fmt.Sprint(d1), "[ ]"))
	//
	//fmt.Println(d1)

	//arrayA := []int{100, 200}
	//fmt.Printf("%+v\n", arrayA)
	//testArrayPoint(&arrayA) // 1.传数组指针
	//arrayB := arrayA[:]
	//fmt.Printf("%+v", arrayA)
	//testArrayPoint(&arrayB) // 2.传切片
	//fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)

	//s := make([]byte, 200)
	////fmt.Printf("%p\n", &s)
	////fmt.Printf("%p\n", &s[0])
	////fmt.Println(&s[0])
	////ptr := unsafe.Pointer(&s[0])
	//ptr := &s[0]
	//fmt.Println(ptr)
	//fmt.Printf("%p", ptr)


}

func add(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return   // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func testArrayPoint(x *[]int) {
	//fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}

func sumArr(a [5]int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func printArr(arr *[5]int) {
	(*arr)[0] = 10
	fmt.Printf("%p\n", &arr)
}

type gophersGreeter struct {
	how string
	who string
}

func (greeter gophersGreeter) greet(how string, who string) {
	fmt.Printf("%s %s!", greeter.how, greeter.who)
	fmt.Printf("%+v", greeter.how)
}

// 遍历字符串
func traversalString() {
	s := "博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func transString() {
	s1 := "hello"
	bytes1 := []byte(s1)
	bytes1[0] = 'H'
	fmt.Println(string(bytes1))
}

func m() {
	a := "O"
	fmt.Print(a)
}
