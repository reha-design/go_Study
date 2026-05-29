package main

import (
	"fmt"
	"time"
)

// [select문: 다중 채널 통신 제어]
// 하나의 고루틴이 여러 채널의 상태(수신 가능, 송신 가능)를 동시에 모니터링할 때 사용합니다.
// switch문과 비슷하게 생겼지만, case에는 채널 연산만 올 수 있습니다.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 고루틴 1: 0.5초 후 데이터 전송
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "채널 1에서 온 메시지"
	}()

	// 고루틴 2: 1초 후 데이터 전송
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "채널 2에서 온 메시지"
	}()

	fmt.Println("--- 다중 채널 수신 대기 (select) ---")
	
	// select를 반복하여 여러 채널에서 오는 데이터를 처리
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("수신:", msg1)
		case msg2 := <-ch2:
			fmt.Println("수신:", msg2)
		// default:
			// default가 있으면 채널 데이터가 준비되지 않았을 때 대기하지 않고 즉시 default를 실행합니다.
			// 여기서는 데이터를 기다려야 하므로 default를 생략합니다.
		}
	}


	fmt.Println("\n--- 타임아웃(Timeout) 처리 ---")
	// 서버 프로그래밍에서 특정 시간 내에 응답이 오지 않으면 포기해야 할 때 매우 유용합니다.
	
	slowCh := make(chan string)
	
	go func() {
		time.Sleep(3 * time.Second) // 3초 걸리는 작업 시뮬레이션
		slowCh <- "작업 완료 데이터"
	}()

	select {
	case result := <-slowCh:
		fmt.Println("성공:", result)
	case <-time.After(2 * time.Second): // 2초 타임아웃 설정
		// time.After는 지정된 시간이 지나면 현재 시간을 채널로 보내주는 함수입니다.
		fmt.Println("타임아웃 발생! 작업을 취소합니다.")
	}
}
