package main

import (
	"fmt"
	"strings"
)

func rotate(s string) string {
	b := strings.Builder{}
	runes := []rune(s)

	for i := len(runes)-1; i >=0 ; i-- {
		b.WriteRune(runes[i])
	}

	return b.String()
}

func main() {
	t1 := "ab日cd"
	t2 := "hello"

	fmt.Println(t1, "->", rotate(t1))
	fmt.Println(t2, "->", rotate(t2))
}
