
# 복습
세마포어(semaphore): 개수


// - Sync Package
// - 하나의 값에 동시에 접근할 때 → 동시성 문제
// - Mutex (≒ Lock?)
// https://worthpreading.tistory.com/90
// 뮤텍스: 한 쓰레드, 프로세스에 의해 소유될 수 있는 Key🔑를 기반으로 한 상호배제기법
// 세마포어: Signaling mechanism. 현재 공유자원에 접근할 수 있는 쓰레드, 프로세스의 수를 나타내는 값을 두어 상호배제를 달성하는 기법
// 출처: https://worthpreading.tistory.com/90 [Worth spreading:티스토리]