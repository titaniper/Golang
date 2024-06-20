package src

import (
	"container/list"
	"fmt"
)

/*
*

	list.New(): 새로운 링크드 리스트를 생성합니다.
	PushBack(value interface{}): 리스트의 끝에 요소를 추가합니다.
	PushFront(value interface{}): 리스트의 앞에 요소를 추가합니다.
	Remove(element *Element): 리스트에서 특정 요소를 제거합니다.
	Front() *Element: 리스트의 첫 번째 요소를 반환합니다.
	Back() *Element: 리스트의 마지막 요소를 반환합니다.
	Next() *Element: 다음 요소를 반환합니다.
	Prev() *Element: 이전 요소를 반환합니다.
*/
func List() {
	// 새로운 링크드 리스트 생성
	myList := list.New()

	// 리스트에 요소 추가
	myList.PushBack("First")
	myList.PushBack("Second")
	myList.PushBack("Third")

	// 리스트 요소 순회 및 출력
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 리스트에서 요소 제거
	element := myList.Front() // 첫 번째 요소를 가리킴
	myList.Remove(element)

	// 요소 제거 후 리스트 출력
	fmt.Println("After removal:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
