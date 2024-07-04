package main

//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//const maxConcurrentRequests = 3
//
//var mu sync.Mutex
//var currentRequests int
//
//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// 임계 구역에 들어가기 전에 잠금
//	mu.Lock()
//	if currentRequests >= maxConcurrentRequests {
//		mu.Unlock() // 조건이 만족되지 않으면 잠금을 해제하고 리턴
//		return
//	}
//	currentRequests++
//	mu.Unlock()
//
//	// 작업 수행
//	fmt.Printf("Worker %d is doing work\n", id)
//	time.Sleep(1 * time.Second) // 작업을 시뮬레이션
//
//	// 작업 완료 후 임계 구역에 들어가기 전에 잠금
//	mu.Lock()
//	currentRequests--
//	mu.Unlock()
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	for i := 1; i <= 10; i++ {
//		wg.Add(1)
//		go worker(i, &wg)
//	}
//
//	wg.Wait()
//	fmt.Println("All workers finished.")
//}
