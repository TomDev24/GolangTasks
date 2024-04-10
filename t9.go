package main

import "fmt"

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    ch := gen([]int{2, 4, 8, 16})
    out := sq(ch)

	for n := range out {
    	fmt.Println(n)
	}
	fmt.Println("-------")
   	// Or even better way to write it
	arr := []int{2, 4}
	for n := range sq(sq(gen(arr))) {
        fmt.Println(n)
    }
}
