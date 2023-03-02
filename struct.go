package main

//
//import (
//	"errors"
//	"fmt"
//	"net/http"
//	"os"
//)
//
////func main()  {
////
////	var p1 person
////	p1.name = "小明"
////	p1.age = 18
////	p1.city = "北京"
////	p1.money = 3000.0
////
////	var p2 = new(person)
////	p2.name = ""
////
////	p3 := &person{
////		name: "jj",
////		age: 12,
////		city: "beijing",
////	}
////	p3.name = "沙哈" //p3.name = “博客” 其实在底层是(*p3).name = “博客”
//
////fmt.Printf("p3 = %#v\n",p3)
//
////fmt.Printf("p1 = %v\n",p1)
////fmt.Printf("p1 = %+v\n",p1)
////fmt.Printf("p1 = %#v\n" ,p1)
////
////var user struct{name string; age int}
////user.name = "小虎"
////user.age = 13
////fmt.Printf("user = %+v",user)
//
////p4 := newPerson("xix","shanghai",18)
////fmt.Println(p4.age)
////fmt.Printf("%#v",p4)
//
////p5 := newPerson("hah","nanjing",12)
////p5.SetAge2(23)
////fmt.Println(p5.age)
////fmt.Printf("%#v\n",p5)
////
////p5.setAge(24)
////fmt.Println(p5.age)
////fmt.Printf("%#v\n",p5)
//
////dog := Dog{
////	feet: "21",
////	Animal: &Animal{
////		name: "旺财",
////	},
////}
////dog.move()
////dog.wang()
////fmt.Printf("这个 一个dog 示例 %#v", dog)
//
////class01 := Class{
////	Title: "一年级",
////	Students: []*Student{
////		{
////			ID:     1,
////			Gender: "男",
////			Name:   "小明",
////					},
////		{
////			ID:     2,
////			Gender: "女",
////			Name:   "小丽",
////					},
////	},
////}
////data, err := json.Marshal(class01)
////fmt.Printf("json:%s\n", data)
////fmt.Printf("class one 的信息%+v,学生信息%+v\n",class01,class01.Students[:1])
////for _, student := range class01.Students {
////	if student.Name == "小丽" {
////		fmt.Printf("学生小丽%+v",*student)
////	}
////}
////str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
////c1 := &Class{}
////err = json.Unmarshal([]byte(str), c1)
////if err != nil {
////	fmt.Println("json unmarshal failed!")
////	return
////}
////fmt.Printf("%#v\n", c1)
//
////ce := []student{
////	student{1, "xiaoming", 22},
////	student{2, "xiaozhang", 33},
////}
////fmt.Println(ce)
////demo(&ce)
////fmt.Println(ce)
//
////var c1, c2, c3 chan int
////var i1, i2 int
////select {
////case i1 = <-c1:
////	fmt.Printf("received ", i1, " from c1\n")
////case c2 <- i2:
////	fmt.Printf("sent ", i2, " to c2\n")
////case i3, ok := <-c3: // same as: i3, ok := <-c3
////	if ok {
////		fmt.Printf("received ", i3, " from c3\n")
////	} else {
////		fmt.Printf("c3 is closed\n")
////	}
////default:
////	fmt.Printf("no communication\n")
////}
//
////a := [3]int{0, 1, 2}
////for i, v := range a { // index、value 都是从复制品中取出。
////	fmt.Printf("v值%d\n",v)
////	//if i == 0 { // 在修改前，我们先修改原数组。
////		a[1], a[2] = 999, 999
////		//fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
////	//}
////	a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
////	fmt.Println(a)
////}
////arrays(a)
////fmt.Println(a)
////
////s := []int{1, 2, 3, 4, 5}
////
////for i, v := range s { // 复制 struct slice { pointer, len, cap }。
////	if i == 0 {
////		s = s[:3]  // 对 slice 的修改，不会影响 range。
////		s[2] = 100 // 对底层数据的修改。
////	}
////	println(i, v)
////	fmt.Println(s)
////}
////fmt.Println(s)
//
////c := a()
////c()
////fmt.Println(a()())
////
////tmp := ads(10)
////fmt.Println(tmp(1),tmp(2))
//
////temp01,temp02 := test01(10)
////fmt.Println(temp02(1),temp01(1))
////
////var whatever [5]struct{}
////for i := range whatever {
////	defer func() {
////		fmt.Println(i)
////	}()
////}
////
////var a = 1
////var b = 2
////
////defer func(a int, b int) {
////	fmt.Println(a + b)
////}(a, b)
////a = 2
////fmt.Println("main")
////
////test()
////i, err := foo(2, 0)
////if err != nil {
////	//panic(err)
////}
////fmt.Println(i)
//
////fmt.Println(f())
////fmt.Println(f1())
////fmt.Println(f2())
////fmt.Println(f3())
////test05()
////chanClose()
//
////test12(2,1)
////test11()
////errtest()
//
////Try(func() {
////	panic("this is err")
////}, func(err interface{}) {
////	fmt.Println(err)
////})
////}
//
//func Try(fun func(), handler func(interface{})) {
//	defer func() {
//		if err := recover(); err != nil {
//			handler(err)
//		}
//	}()
//	fun()
//}
//
//func test12(x, y int) {
//	var z int
//	errtest()
//	fmt.Printf("x / y = %d\n", z)
//}
//
//func errtest() {
//	defer func() {
//		//if recover() != nil {
//		fmt.Println(recover())
//		//}
//	}()
//	panic("test panic")
//	//z = x / y
//}
//
//func chanClose() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	ch := make(chan int, 10)
//	close(ch)
//	ch <- 1
//}
//
//func test11() {
//	defer func() {
//		fmt.Println(recover()) //有效
//	}()
//	//defer recover()              //无效！
//	//defer fmt.Println(recover()) //无效！
//	//defer func() {
//	//	println("defer inner")
//	//	fmt.Println(recover()) //无效！
//	//}()
//	defer func() {
//		func() {
//			println("defer inner")
//			fmt.Println(recover())
//		}()
//		panic("test panic")
//	}()
//}
//
//func test10() {
//	defer func() {
//		//println("defer inner")
//		fmt.Println(recover())
//	}()
//
//	//defer func() {
//	//	panic("defer panic")
//	//}()
//
//	panic("test panic")
//}
//
//func test09() {
//	defer func() {
//		if err := recover(); err != nil {
//			println(err.(string)) // 将 interface{} 转型为具体类型。
//		}
//	}()
//	panic("panic error!")
//}
//
//func do02() error {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func() {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close book.txt err %v\n", err)
//			}
//		}()
//	}
//
//	f, err = os.Open("another-book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func() {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close another-book.txt err %v\n", err)
//			}
//		}()
//	}
//
//	return nil
//}
//
//func do01() (err error) {
//	f, err := os.Open("book.text")
//	if err != nil {
//		return err
//	}
//
//	if f != nil {
//		defer func() {
//			if ferr := f.Close(); ferr != nil {
//				err = ferr
//				fmt.Printf("defer close book.txt err %v\n", err)
//			}
//		}()
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func do() error {
//	res, err := http.Get("http://www.google.com")
//	defer res.Body.Close()
//	if err != nil {
//		return err
//	}
//	// ..code...
//	return nil
//}
//
//func test05() {
//	var run func()
//	defer run()
//	fmt.Println("runs")
//}
//
//func f() (result int) {
//	defer func() {
//		result++
//	}()
//	return 0
//}
//
//func f1() (r int) {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}
//
//func f2() (r int) {
//	fmt.Println(r)
//	defer func(r int) {
//		r = r + 5
//		fmt.Println(r)
//	}(r)
//	return 1
//}
//
//func f3() (r int) {
//	//给返回值赋值
//	r = 1
//	func(r int) { //这里改的r是传值传进去的r，不会改变要返回的那个r值
//		r = r + 5
//		fmt.Println(r)
//	}(r)
//	return //空的return
//}
//
//func foo(a, b int) (i int, err error) {
//	defer fmt.Printf("first defer err %v\n", err)
//	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
//	defer func() { fmt.Printf("third defer err %v\n", err) }()
//	if b == 0 {
//		err = errors.New("divided by zero!")
//		return
//	}
//	i = a / b
//	return
//}
//
//func test() {
//	x, y := 10, 20
//
//	defer func(i int) {
//		println("defer:", i, y) // y 闭包引用
//	}(x) // x 被复制
//
//	x += 10
//	y += 100
//	println("x =", x, "y =", y)
//}
//
//func test01(base int) (func(int) int, func(int) int) {
//	add := func(i int) int {
//		return base + i
//	}
//	sub := func(i int) int {
//		return base - i
//	}
//	return add, sub
//}
//
//func ads(base int) func(int) int {
//	return func(i int) int {
//		base += i
//		return base
//	}
//}
//
//func a() func() int {
//	i := 0
//	b := func() int {
//		i++
//		fmt.Println(i)
//		return i
//	}
//	return b
//}
//
//func arrays(a [3]int) {
//	for i, v := range a { // index、value 都是从复制品中取出。
//		fmt.Printf("v值%d\n", v)
//		//if i == 0 { // 在修改前，我们先修改原数组。
//		a[1], a[2] = 999, 999
//		//fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
//		//}
//		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
//		fmt.Println(a)
//	}
//}
//
//type student struct {
//	id   int
//	name string
//	age  int
//}
//
//func demo(ce *[]student) {
//	//切片是引用传递，是可以改变值的
//	//ce = &[]student{
//	//	{id: 1,name: "xi",age: 12},
//	//}
//
//	(*ce)[0].age = 999
//	fmt.Printf("ce %+v\n", *ce)
//	// ce = append(ce, student{3, "xiaowang", 56})
//	// return ce
//}
//
//type Student struct {
//	ID     int
//	Gender string
//	Name   string
//}
//
//// Class 班级
//type Class struct {
//	Title    string
//	Students []*Student
//}
//
//type Animal struct {
//	name string
//}
//
//type Dog struct {
//	feet string
//	*Animal
//}
//
//func (d Dog) wang() {
//	fmt.Printf("这是 dog的方法 %+v\n", d)
//}
//
//func (a Animal) move() {
//	fmt.Printf("这是animal 的方法 %+v\n", a)
//}
//
//func (p *person) setAge(newAge int) {
//	p.age = newAge
//	fmt.Printf("指针传递,%v\n", p.age)
//}
//
//func (p person) SetAge2(newAge int) {
//	p.age = newAge
//	fmt.Printf("值传递,%v\n", p.age)
//}
//
//func Dream(p person) {
//	fmt.Println(p.name)
//}
//
//func (p person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
//}
//
//func newPerson(name, city string, age int) *person {
//	return &person{
//		name: name,
//		city: city,
//		age:  age,
//	}
//}
//
//type person struct {
//	name, city string
//	age        int
//	money      float32
//}
