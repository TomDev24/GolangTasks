package main

import (
	"sync"
	"fmt"
	"time"
	"context"
)

//1. Close the channel
func byClosing() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
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
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

//2. Send message over special stop channel
func byMessage() {
	stopCh := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(){
		defer wg.Done()
		for {
			select {
				case <-stopCh:
						fmt.Println("Recived message from stop channel. Stopping goroutine")
						return
				default:
					fmt.Println("doing work")
					time.Sleep(1 * time.Second)
			}
		}
	}()
time.Sleep(1 * time.Second)
	stopCh <- true
	wg.Wait()
}

// 3. Stop goroutine using context
func byContext(){
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		for {
			select {
			case <- ctx.Done(): //will happen, when we call cancel()
				fmt.Println("Context is canceled. Stopping goroutine")
				ch <- struct{}{}
				return
			default:
				fmt.Println("doing work")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// close the context in separate gorutine after 5 sec
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	<-ch // waiting for last msg from channel
}

func main() {
	byClosing()
	byMessage()
	byContext()
}
