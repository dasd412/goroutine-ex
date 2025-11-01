package main

import "testing"

func TestMutex(t *testing.T) {
	var mu SpinMutex
	counter := 0

	//버퍼 없는 채널이라 블로킹됨
	done := make(chan bool)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			done <- true //송신 (send): 채널에 값을 넣기 (신호 보내기)
		}()
	}

	// 모든 고루틴 완료 대기
	for i := 0; i < 100; i++ {
		//수신 (receive): 채널에서 값을 꺼내기
		<-done // 값을 버림 (동기화 목적, 신호 올 때까지 블로킹)
	}

	println("Counter:", counter)
}
