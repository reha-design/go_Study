package main

import (
	"fmt"
	"time"
)

// [채널 닫기와 range 순회]
// 더 이상 채널로 보낼 데이터가 없음을 수신자에게 알리기 위해 채널을 닫습니다(close).
// 수신자는 for range 문을 사용하여 채널이 닫힐 때까지 지속적으로 데이터를 받을 수 있습니다.

func main() {
	fmt.Println("--- 1. 채널 닫기 (close) ---")
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	// 데이터 전송 완료 후 채널을 닫음
	// 주의: 닫힌 채널에 데이터를 보내려고 하면 패닉(Panic)이 발생합니다.
	close(ch) 

	// 닫힌 채널에서도 이미 전송된 데이터는 정상적으로 읽을 수 있습니다.
	val1, ok1 := <-ch
	fmt.Printf("값: %d, 정상수신(ok): %t\n", val1, ok1)
	
	val2, ok2 := <-ch
	fmt.Printf("값: %d, 정상수신(ok): %t\n", val2, ok2)
	
	val3, ok3 := <-ch
	fmt.Printf("값: %d, 정상수신(ok): %t\n", val3, ok3)

	// 채널이 비어있고 닫혀있다면, 해당 타입의 기본값과 함께 ok=false가 반환됩니다.
	val4, ok4 := <-ch
	fmt.Printf("닫힌 채널 수신 - 값: %d, 정상수신(ok): %t\n", val4, ok4)


	fmt.Println("\n--- 2. for range를 이용한 채널 순회 ---")
	queue := make(chan string, 5)

	// 고루틴에서 데이터를 생성하여 채널로 전송
	go func() {
		for i := 1; i <= 3; i++ {
			queue <- fmt.Sprintf("작업 %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		// 모든 데이터 전송 후 반드시 채널을 닫아야 합니다.
		// 그렇지 않으면 메인 함수의 for range가 데이터가 오기를 영원히 기다리며 데드락에 빠집니다.
		close(queue)
		fmt.Println("데이터 생성 완료 및 채널 닫음")
	}()

	// 채널이 닫힐 때까지 자동으로 데이터를 수신합니다.
	for msg := range queue {
		fmt.Println("수신:", msg)
	}
	fmt.Println("모든 작업 수신 완료")
}
