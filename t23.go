package main

import (
	"fmt"
	"errors"
)

func Delete1[T any](s []T, i int) ([]T, error) {
	if !(i >= 0 && i < len(s)) {
		return nil, errors.New("wrong index")
	}
	copy(s[i:], s[i+1:])
	return s[:len(s)-1], nil
}

func Delete2[T any](s []T, i int) ([]T, error) {
	if !(i >= 0 && i < len(s)) {
		return nil, errors.New("wrong index")
	}
	return append(s[:i], s[i+1:]...), nil
}

// without saving order
func Delete3[T any](s []T, i int) ([]T, error) {
	if !(i >= 0 && i < len(s)) {
		return nil, errors.New("wrong index")
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1], nil
}

func main() {
	s1 := []int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(s1[len(s1):], "empty")
	fmt.Println(s1, 9)
	s1, _ = Delete1(s1, 9)
	fmt.Println(s1)

	fmt.Println(s1, 4)
	s1,_ = Delete1(s1, 4)
	fmt.Println(s1)

	// add more tests here!
}
