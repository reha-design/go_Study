package main

import (
	"fmt"
	"sync"
	"time"
)

// [고루틴(Goroutine): Go 언어의 가볍고 빠른 동시성 실행 단위]
// 일반적인 스레드(Thread)보다 훨씬 적은 메모리를 사용하며 생성 속도가 매우 빠릅니다.
// 함수 호출 앞에 'go' 키워드만 붙이면 새로운 고루틴이 생성되어 비동기적으로 실행됩니다.

func printNumbers(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("[%s] 번호: %d\n", name, i)
		time.Sleep(100 * time.Millisecond) // 0.1초 대기
	}
}

func main() {
	fmt.Println("--- 1. 고루틴 기본 실행 ---")
	
	// 일반적인 함수 호출 (동기적 실행: 끝날 때까지 기다림)
	// printNumbers("동기 함수") 

	// go 키워드를 사용하여 고루틴 생성 (비동기적 실행: 기다리지 않고 다음 줄로 넘어감)
	go printNumbers("고루틴 A")
	go printNumbers("고루틴 B")

	// [주의] 메인 고루틴(main 함수)이 종료되면 다른 모든 고루틴도 즉시 강제 종료됩니다.
	// 따라서 다른 고루틴들이 작업을 마칠 수 있도록 메인 고루틴을 잠시 대기시켜야 합니다.
	time.Sleep(400 * time.Millisecond) // 0.4초 대기
	fmt.Println("메인 함수 종료 (고루틴 기본)")


	fmt.Println("\n--- 2. sync.WaitGroup을 이용한 우아한 대기 ---")
	// time.Sleep()으로 대기 시간을 하드코딩하는 것은 매우 비효율적이고 위험합니다.
	// sync.WaitGroup을 사용하면 모든 고루틴이 끝날 때까지 정확하게 기다릴 수 있습니다.

	var wg sync.WaitGroup // WaitGroup 생성

	// 고루틴 3개를 실행할 것이므로 WaitGroup 카운터를 3으로 설정
	wg.Add(3)

	go func() {
		defer wg.Done() // 작업 완료 시 카운터 감소 (반드시 실행되도록 defer 사용)
		fmt.Println("작업 1 완료")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("작업 2 완료")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("작업 3 완료")
	}()

	fmt.Println("모든 작업이 끝날 때까지 대기 중...")
	wg.Wait() // 카운터가 0이 될 때까지 메인 고루틴 대기
	fmt.Println("모든 고루틴 작업 완료!")
}
