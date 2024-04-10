package main

import (
	"fmt"
	"sync"
	"log"
	"os"
)

func powerOfTwoLog(num int, logger *log.Logger){
	//fmt.Printf("%d ", num * num) // We cannot use fmt lib, cause its not safe in concurrent environment
	logger.Printf("%d", num * num) // logger is safe to use, cause it locks mutex when printing
}

func powerOfTwoMutex(num int, m *sync.Mutex){
	m.Lock()
	defer m.Unlock()
	fmt.Printf("%d\n", num * num)
}


// Solution using logger
func useLogger() {
	logger := log.New(os.Stdout, "", 0) // using STDOUT instead of STDERR
	wg := sync.WaitGroup{}
	nums := []int{2,4,6,8,10}

	wg.Add(len(nums))
	for _, n := range nums {
		go func(num int){
			defer wg.Done()
			powerOfTwoLog(num, logger)
		}(n)
	}
	wg.Wait()
}

func useMutex() {
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	nums := []int{2,4,6,8,10}

	wg.Add(len(nums))
	for _, n := range nums {
		go func(num int){
			defer wg.Done()
			powerOfTwoMutex(num, &m) // mutexes must not be copied,
		}(n)
	}
	wg.Wait()
}

func main(){
	fmt.Println("Using logger:")
	useLogger()
	fmt.Println("Using mutex:")
	useMutex()
}

