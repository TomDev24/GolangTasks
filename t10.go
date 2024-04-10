package main

import (
	"fmt"
)

func main() {
	dict := map[int][]float32{}
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, v := range temps {
		key := int((v / 10)) * 10
		dict[key] = append(dict[key], v)
	}

	for key, value := range dict {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
