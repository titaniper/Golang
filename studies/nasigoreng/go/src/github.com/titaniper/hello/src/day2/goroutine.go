package day2

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴은 Go에서 경량 스레드로, 비동기 작업을 수행하는데 사용됩니다.
/**
-  고루틴은 매우 적은 메모리 오버헤드로 동작합니다.
   - 고루틴은 운영체제 스레드보다 훨씬 가볍습니다. 이는 다음과 같은 이유들 때문입니다:
		- 작은 초기 스택 크기: 고루틴의 초기 스택 크기는 몇 KB로 매우 작습니다. 필요에 따라 자동으로 크기가 증가합니다.
 		- 고루틴의 스택 크기: 초기 스택 크기가 작고 필요에 따라 자동으로 증가합니다.
		- 고루틴 스케줄러: Go 런타임에 내장된 고루틴 스케줄러는 수많은 고루틴을 소수의 운영체제 스레드에 효율적으로 매핑합니다.
- 고루틴 스택 크기 출력
 	- 고루틴의 스택 크기는 직접 출력할 수 없지만, 디버깅 시 스택 크기를 확인할 수 있습니다.
- 고루틴 동기화 문제
	- 고루틴 간의 동기화 문제를 해결하는 방법에는 sync 패키지와 채널을 사용하는 방법이 있습니다.


*/

func printNumbers(c string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(c, i)
	}
}

func Goroutine() {
	go printNumbers("routine")  // 고루틴 시작
	defer printNumbers("defer") // 고루틴 시작

	// 메인 함수의 코드
	fmt.Println("Hello from main function")
	time.Sleep(6 * time.Second) // 메인 함수 종료 지연
}

func Sync() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// 채널 이용한 동기화
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func SyncChannel() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func GoroutineTest1() {
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	//time.Sleep(time.Second * 3)
}
