package main

import (
	"os"
	"fmt"
	"strings"
)

func isUnique(str string) bool{
	dict := map[rune]rune{}
	s := strings.ToLower(str)

	for _, l := range s {
		_, exist := dict[l]
		if exist {
			return false
		}
		dict[l] = l
	}
	return true
}

func main(){
	if len(os.Args) == 2 {
		fmt.Println(isUnique(os.Args[1]))
		return
	}
	fmt.Println("---TESTS---")
	fmt.Println("abcd", isUnique("abcd"))
	fmt.Println("abCdefAaf", isUnique("abCdefAaf"))
	fmt.Println("aabcd", isUnique("aabcd"))
	fmt.Println("ABCDa", isUnique("ABCDa"))
}
