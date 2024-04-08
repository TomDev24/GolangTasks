package main

import (
	"fmt"
	"strings"
)

func rotateArr(s string) string {
	b := strings.Builder{}
	words := strings.Fields(s)

	for i := len(words)-1; i >= 0; i-- {
		b.WriteString(words[i])
		b.WriteString(" ")
	}

	return b.String()
}

func main() {
	t1 := "hey there world"
	t2 := "1 2 3 4"

	fmt.Println(t1, "->", rotateArr(t1))
	fmt.Println(t2, "->", rotateArr(t2))
}
