package main

import (
	"fmt"
	"sync"
	"time"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

// 1. using one bufferd channel
func powerOfTwo(num int, out chan<- int){
	out <- num * num
}

func buffChannel(nums []int) {
	result := 0
	wg := sync.WaitGroup{}
	ch := make(chan int, len(nums))

	wg.Add(len(nums))
	for i := range nums{
		go func(index int){
			defer wg.Done()
			powerOfTwo(nums[index], ch)
		}(i)
	}
	wg.Wait()
	close(ch) // closing channel, also sends END message to recivers, so we will be able iterate over it

	for v := range ch {
		result += v
	}
	fmt.Println(result)
}


func main(){
	buffChannel([]int{2,4,6,8,10})
	fmt.Println("Testing speed with 1 million numbers")
	start := time.Now()
	buffChannel(makeRange(1, 1 << 20))
	fmt.Println(time.Since(start))
}
