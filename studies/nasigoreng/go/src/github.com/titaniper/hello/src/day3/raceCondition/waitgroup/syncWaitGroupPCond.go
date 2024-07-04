package main

import (
	"fmt"
	"sync"
	"time"
)

const maxConcurrentRequests = 3

var mu sync.Mutex
var cond = sync.NewCond(&mu)
var currentRequests int

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	for currentRequests >= maxConcurrentRequests {
		cond.Wait() // 조건이 만족되지 않으면 기다림
	}
	currentRequests++
	mu.Unlock()

	// 작업 수행
	fmt.Printf("Worker %d is doing work\n", id)
	time.Sleep(1 * time.Second) // 작업을 시뮬레이션

	mu.Lock()
	currentRequests--
	cond.Signal() // 기다리고 있는 고루틴 하나를 깨움
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
}
