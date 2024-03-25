package main

import (
	"fmt"
	"time"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func powerOfTwo(num int, out chan<- int){
	out <- num * num
}

func main(){
	start := time.Now()
	nums := []int{2,4,6,8,10}
	//nums := makeRange(1, 1000000)
	result := 0
	c := make(chan int, len(nums))
	defer close(c)


	for i := range nums{
		go powerOfTwo(nums[i], c)
	}
	for _ = range nums{
		result += <-c
	}
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
