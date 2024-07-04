package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//var counter int
//
//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 0; i < 1000; i++ {
//		counter++
//		fmt.Println("gogo count:", id)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go worker(i, &wg)
//	}
//
//	wg.Wait()
//	fmt.Println("Final counter value:", counter)
//}
