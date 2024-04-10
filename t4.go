package main

import (
	"fmt"
	"sync"
	"os"
	"time"
	"os/signal"
	"syscall"
)

func main() {
	var n int
	fmt.Printf("Num of workers: ")
	fmt.Scan(&n)

	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	ch := make(chan int, n)
	osCh := make(chan os.Signal)

	wg.Add(n)
	signal.Notify(osCh, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < n; i++ {
		go func(id int){
			defer wg.Done()
			for data := range ch {
				m.Lock()
				fmt.Printf("worker: %d, got: %d\n", id, data)
				m.Unlock()
				time.Sleep(3 * time.Second)
			}
		}(i)
	}

	num := 0
	for {
		select {
		case ch <- num:
			num++
		case <- osCh:
			close(ch)
			fmt.Println("Exiting because of OS signal") // not mutex protected, cause we dont care at the end
			wg.Wait()
			os.Exit(0) // exiting with zero, showing there is no error
		}
	}
}
