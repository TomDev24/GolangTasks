package main

import (
	"sync"
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func main() {
	nWorkers := 5
	wg := sync.WaitGroup{}
	m := sync.RWMutex{} // so that we can read from map concurrently
	dict := map[int]string{}

	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func(i int){
			defer wg.Done()
			m.Lock()
			dict[i] = RandStringBytes(10)
			m.Unlock()
		}(i)
	}
	wg.Wait()

	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func(i int){
			defer wg.Done()
			m.RLock()
			fmt.Println(i, dict[i]) // writing to STDOUT should be also locked, but we drop it for study purpose
			m.RUnlock()
		}(i)
	}

	wg.Wait()
}

// Also, there is sync.Map https://pkg.go.dev/sync#Map
// but i am lazy)
