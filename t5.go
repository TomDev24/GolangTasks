package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Printf("Enter time: ")
	fmt.Scan(&n)

	ch := make(chan int)
	wg := sync.WaitGroup{}
	timeout := time.After(time.Duration(n) * time.Second)

	wg.Add(1)
	go func(){
		defer wg.Done()
		for {
			n, ok := <-ch
			switch {
			case !ok:
				fmt.Println("channel is closed. Stopping goroutine")
				return
			case ok:
				fmt.Println(n)
				time.Sleep(1 * time.Second)
			}

		}
	}()

	i := 1
	LOOP:
	for {
		select {
		case ch <- i:
			i++
		case <-timeout:
			close(ch)
			break LOOP
		}
	}
	wg.Wait()
}
