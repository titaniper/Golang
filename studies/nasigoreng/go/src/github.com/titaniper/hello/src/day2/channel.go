package day2

import (
	"fmt"
	"time"
)

/**
채널은 고루틴 간에 데이터를 주고받는 동기화 메커니즘입니다.
- 채널은 타입을 가지며, 해당 타입의 데이터만 전송할 수 있습니다.
*/

func Channel() {
	//test := make(string)
	// 채널 생성
	messages := make(chan string)

	// 고루틴에서 메시지 전송
	go func() {
		time.Sleep(2 * time.Second)
		messages <- "ping"
	}()

	// 채널에서 메시지 수신
	msg := <-messages
	fmt.Println(msg)

	// 버퍼링된 채널: 버퍼를 가진 채널을 만들 수 있으며, 버퍼가 꽉 차지 않으면 송신이 블로킹되지 않습니다.
	//buffered := make(chan int, 10)

	// select 문: 여러 채널 연산을 대기할 때 사용됩니다.
	//select {
	//case msg1 := <-chan1:
	//	fmt.Println("Received", msg1)
	//case msg2 := <-chan2:
	//	fmt.Println("Received", msg2)
	//case chan3 <- "hello":
	//	fmt.Println("Sent message")
	//default:
	//	fmt.Println("No communication")
	//}

	// 채널 닫기: 채널을 닫으면 더 이상 전송할 수 없습니다.
	close(messages)
}
