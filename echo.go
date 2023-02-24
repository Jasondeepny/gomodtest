package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, se string
	for i := 1; i < len(os.Args); i++ {
		s += se + os.Args[i]
		se = ""
	}
	fmt.Println(s)
}

func main1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func init() {
	fmt.Println("imp-int")
}

func Print() {
	fmt.Println("Hello!")
}