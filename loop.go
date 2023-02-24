package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func convertString() {
	s := "这个是个字符串"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Print("-------------")

}
