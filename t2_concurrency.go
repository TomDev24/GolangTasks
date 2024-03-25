package main

import (
	"fmt"
	"sync"
)

func powerOfTwo(num int){
	fmt.Printf("%d ", num * num)
}

// Solution using WaitGroups
func main(){
	var wg sync.WaitGroup
	nums := []int{2,4,6,8,10}

	wg.Add(len(nums))
	for _, n := range nums {
		go func(num int){
			defer wg.Done()
			powerOfTwo(num)
		}(n)
	}

	wg.Wait()
	fmt.Println()
}


