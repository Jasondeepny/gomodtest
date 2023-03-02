package gotest

//
//import (
//	"reflect"
//	"strings"
//	"testing"
//)
//
//func BenchmarkSplit(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Split("枯藤老树昏鸦", "老")
//	}
//}
//
//func TestSplitSlice(t *testing.T) {
//	type test struct {
//		input, sep string
//		want       []string
//	}
//
//	tests := []test{
//		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		{input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//	}
//
//	for _, tc := range tests {
//		split := Split(tc.input, tc.sep)
//		if !reflect.DeepEqual(split, tc.want) {
//			t.Errorf("excepted:%v, got:%v", tc.want, split)
//		}
//	}
//}
//
//func TestSplitMap(t *testing.T) {
//	type test struct { // 定义test结构体
//		input string
//		sep   string
//		want  []string
//	}
//	tests := map[string]test{ // 测试用例使用map存储
//		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//	}
//	for name, tc := range tests {
//		got := Split(tc.input, tc.sep)
//		if reflect.DeepEqual(got, tc.want) {
//			t.Logf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
//		}
//	}
//}
//
//func TestSplitSubTest(t *testing.T) {
//	type test struct { // 定义test结构体
//		input string
//		sep   string
//		want  []string
//	}
//	tests := map[string]test{ // 测试用例使用map存储
//		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//	}
//	teardownTestCase := setupTestCase(t)
//	defer teardownTestCase(t)
//
//	//for name, tc := range tests {
//	//	t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
//	//		got := Split(tc.input, tc.sep)
//	//		if !reflect.DeepEqual(got, tc.want) {
//	//			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//	//		}
//	//	})
//	//}
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			teardownSubCase := setupSubTest(t)
//			defer teardownSubCase(t)
//			got := Split(tc.input, tc.sep)
//			if !reflect.DeepEqual(got, tc.want) {
//				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//			}
//		})
//	}
//}
//
//// 测试集的Setup与Teardown
//func setupTestCase(t *testing.T) func(t *testing.T) {
//	t.Log("如有需要在此执行:测试之前的setup")
//	return func(t *testing.T) {
//		t.Log("如有需要在此执行:测试之后的teardown")
//	}
//}
//
//// 子测试的Setup与Teardown
//func setupSubTest(t *testing.T) func(t *testing.T) {
//	t.Log("如有需要在此执行:子测试之前的setup")
//	return func(t *testing.T) {
//		t.Log("如有需要在此执行:子测试之后的teardown")
//	}
//}
//
//func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Split("a:b:c", ":")         // 程序输出的结果
//	want := []string{"a", "b", "c"}    // 期望的结果
//	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//}
//
//func TestMoreSplit(t *testing.T) {
//	got := Split("abcd", "bc")
//	want := []string{"a", "d"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("excepted:%v, got:%v", want, got)
//	}
//}
//
//func Split(s, sep string) (result []string) {
//	result = make([]string, 0, strings.Count(s, sep)+1)
//	i := strings.Index(s, sep)
//	//fmt.Println(i)
//	for i > -1 {
//		result = append(result, s[:i])
//		//fmt.Println(result)
//		//fmt.Println(len(sep))
//		s = s[i+len(sep):]
//		//fmt.Println(s)
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}
