package gotest

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	fmt.Println(i)
	for i > -1 {
		result = append(result, s[:i])
		fmt.Println(result)
		fmt.Println(len(sep))
		s = s[i+len(sep):]
		fmt.Println(s)
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}