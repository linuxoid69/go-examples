package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1)

	start := time.Now()

	for i := 1; i < 1000; i++ {
		ch <- i
		go worker(ch)
	}

	elapsed := time.Since(start)

	fmt.Printf("Цикл выполнен за %s\n", elapsed)
}

func worker(ch chan int) {
	var mu sync.Mutex
	mu.Lock()

	fmt.Println(<-ch)

	mu.Unlock()
}
