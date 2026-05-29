package main

import (
	"fmt"
	"time"
)

// [채널(Channel): 고루틴 간에 데이터를 주고받는 안전한 파이프(통로)]
// 고루틴은 서로 독립적으로 실행되므로, 데이터를 안전하게 교환할 방법이 필요합니다.
// 채널을 사용하면 데이터 통신뿐만 아니라 고루틴 간의 동기화(타이밍 맞추기)도 가능합니다.
// 형태: make(chan 데이터타입)

func main() {
	fmt.Println("--- 1. 기본 채널 (Unbuffered Channel) ---")
	// 버퍼가 없는 채널은 보내는 쪽과 받는 쪽이 동시에 준비되어야만 데이터가 전달됩니다. (동기식)
	ch := make(chan int) 

	// 데이터를 보내는 고루틴
	go func() {
		fmt.Println("고루틴: 채널에 42를 보냅니다.")
		ch <- 42 // 채널에 데이터 전송 (수신자가 받을 때까지 여기서 대기)
		fmt.Println("고루틴: 전송 완료.")
	}()

	fmt.Println("메인: 1초 대기 후 채널에서 데이터를 받습니다.")
	time.Sleep(1 * time.Second)
	
	val := <-ch // 채널에서 데이터 수신
	fmt.Printf("메인: 채널로부터 %d를 받았습니다.\n", val)


	fmt.Println("\n--- 2. 버퍼가 있는 채널 (Buffered Channel) ---")
	// 버퍼 크기를 지정하면, 버퍼가 가득 찰 때까지는 수신자가 없어도 전송자가 대기하지 않습니다. (비동기식)
	bufCh := make(chan string, 2) // 크기가 2인 문자열 채널 생성

	// 고루틴 없이도 메인 함수에서 바로 전송 가능 (버퍼 여유가 있으므로 대기하지 않음)
	bufCh <- "첫 번째 메시지"
	bufCh <- "두 번째 메시지"
	// bufCh <- "세 번째 메시지" // 이 줄을 주석 해제하면 버퍼가 꽉 차서 데드락(Deadlock) 발생!

	fmt.Println("채널에서 수신 1:", <-bufCh)
	fmt.Println("채널에서 수신 2:", <-bufCh)


	fmt.Println("\n--- 3. 단방향 채널 (함수 매개변수로 사용 시) ---")
	// 함수의 매개변수에서 채널의 방향을 제한하여 안전성을 높일 수 있습니다.

	sendCh := make(chan int)

	// 전송 전용 고루틴 실행
	go sendData(sendCh)
	
	// 수신 전용 함수 실행
	receiveData(sendCh)
}

// 전송 전용 채널: chan<- int (보내기만 가능)
func sendData(ch chan<- int) {
	fmt.Println("sendData: 99 전송")
	ch <- 99
}

// 수신 전용 채널: <-chan int (받기만 가능)
func receiveData(ch <-chan int) {
	val := <-ch
	fmt.Println("receiveData: 수신된 값 =", val)
}
