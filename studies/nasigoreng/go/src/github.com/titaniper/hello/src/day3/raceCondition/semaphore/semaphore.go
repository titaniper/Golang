package main

import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

const maxConcurrentRequests = 3

var sem = semaphore.NewWeighted(maxConcurrentRequests)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := sem.Acquire(nil, 1); err != nil {
		return
	}
	defer sem.Release(1)

	// 작업 수행
	fmt.Printf("Worker %d is doing work\n", id)
	time.Sleep(1 * time.Second) // 작업을 시뮬레이션
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
