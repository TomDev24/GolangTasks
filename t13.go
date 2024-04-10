package main

import (
	"fmt"
)

func main() {
	a := 42
	b := -1
	a, b = b, a
	fmt.Println(a, b)
}
