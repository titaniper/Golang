package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
		counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
