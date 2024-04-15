package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Println(time.Since(start))
}
