package main

import (
	"fmt"
)

var justString string

func createHugeString(n int) string {
	return string(make([]rune, n))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string( []rune(v)[0:100] )
}


func main() {
	// Here is the main problem in two lines of code. String is array of bytes
	// so when we slicing it, we can corrupt multibyte characters
	txt := "hello, 世界"
	fmt.Println(txt, txt[:8])

	fmt.Println(txt, string( []rune(txt)[:8] ))
	someFunc()
}
