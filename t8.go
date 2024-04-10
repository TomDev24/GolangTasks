package main

import (
	"fmt"
	"strconv"
	"os"
	"log"
)

func setBit(num int64, i int) int64 {
	return num | (1 << i)
}

func resetBit(num int64, i int) int64 {
	return num & ^(1 << i)
}

func setFlag(num int64, i int) int64 {
	return num ^ (1 << i)
}

func main(){
	// Accept first two parameters of the program as Number and Index
	if len(os.Args) == 3 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("incorrect input")
		}
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("incorrect input")
		}

		res := setFlag(int64(num), i)
		fmt.Println(strconv.FormatInt(res, 2))
		os.Exit(0)
	}

	// If no input provided, runing default tests
	var num int64 = 1 << 5
	fmt.Println("Number:\t", strconv.FormatInt(num, 2))

	// toggle i = 2
	num = setFlag(num, 2)
	fmt.Println("toggle i=2  :", strconv.FormatInt(num, 2))
	// toggle i = 3
	num = setFlag(num, 3)
	fmt.Println("toggle i=3  :", strconv.FormatInt(num, 2))
	// toggle i = 2
	num = setFlag(num, 2)
	fmt.Println("toggle i=2:", strconv.FormatInt(num, 2))
	// toggle i = 0
	num = setFlag(num, 0)
	fmt.Println("toggle i=0  :", strconv.FormatInt(num, 2))

	// reset i = 0
	num = setFlag(num, 0)
	fmt.Println("reset i=0:\t", strconv.FormatInt(num, 2))
	// reset i = 3
	num = resetBit(num, 3)
	fmt.Println("reset i=3:", strconv.FormatInt(num, 2))

}
