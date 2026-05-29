# Go 강의 보고서

## 1. 강의 목표 및 성취 수준

본 강의는 Go 언어의 가장 강력한 특징인 동시성(Concurrency) 프로그래밍의 핵심 요소인 고루틴(Goroutine)과 채널(Channel), 그리고 데이터 동기화를 위한 `sync` 패키지의 원리를 이해하고 직접 구현해 보는 것을 목표로 진행하였습니다.

수강생들은 강의와 실습을 통해 다음 내용을 수행할 수 있도록 학습하였습니다.

1. `go` 키워드를 사용하여 가볍고 빠른 비동기 실행 단위인 고루틴을 생성하고 실행 흐름을 분기할 수 있다.
2. `sync.WaitGroup`을 활용하여 메인 함수가 모든 고루틴의 작업 완료를 기다리도록 동기화할 수 있다.
3. 채널(Channel)을 생성하여 고루틴 간에 데이터를 안전하게 송수신할 수 있으며, 언버퍼드(Unbuffered)와 버퍼드(Buffered) 채널의 통신 방식 차이를 이해할 수 있다.
4. 채널 닫기(`close`)와 `for range` 문을 조합하여 데이터 스트림을 안전하게 순회하고 종료할 수 있다.
5. `select` 문을 사용하여 다중 채널 통신을 제어하고, `time.After`를 활용하여 타임아웃 처리를 구현할 수 있다.
6. 다수의 고루틴이 공유 데이터에 동시 접근할 때 발생하는 경쟁 상태(Race Condition)를 이해하고, `sync.Mutex`를 사용하여 임계 구역(Critical Section)을 보호할 수 있다.

---

## 2. 실습 및 제출 산출물 현황

강의 과정에서는 각 주제별 실습 코드를 작성하고 실행 결과를 확인하였습니다. 수강생들은 아래 구조의 예제 코드를 직접 작성하며 Go의 동시성 프로그래밍 모델을 검증하였습니다.

```text
day06/
├── 26_goroutine/
│   └── main.go           (고루틴 생성, 비동기 실행, sync.WaitGroup 대기 실습)
├── 27_channel/
│   └── main.go           (채널 송수신, 버퍼드/언버퍼드 채널, 단방향 채널 실습)
├── 28_channel_sync/
│   └── main.go           (채널 닫기, ok 확인, for range 채널 순회 실습)
├── 29_select/
│   └── main.go           (select 다중 채널 모니터링, 타임아웃 처리 실습)
└── 30_mutex/
    └── main.go           (경쟁 상태 재현, sync.Mutex를 이용한 안전한 동시 접근 실습)
```

### 각 챕터별 상세 실습 소스코드

#### ① `26_goroutine/main.go`

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("[%s] 번호: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("--- 1. 고루틴 기본 실행 ---")
	go printNumbers("고루틴 A")
	go printNumbers("고루틴 B")
	
	time.Sleep(400 * time.Millisecond)
	fmt.Println("메인 함수 종료 (고루틴 기본)")

	fmt.Println("\n--- 2. sync.WaitGroup을 이용한 우아한 대기 ---")
	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); fmt.Println("작업 1 완료") }()
	go func() { defer wg.Done(); fmt.Println("작업 2 완료") }()
	go func() { defer wg.Done(); fmt.Println("작업 3 완료") }()

	fmt.Println("모든 작업이 끝날 때까지 대기 중...")
	wg.Wait()
	fmt.Println("모든 고루틴 작업 완료!")
}
```

#### ② `27_channel/main.go`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- 1. 기본 채널 (Unbuffered Channel) ---")
	ch := make(chan int) 

	go func() {
		fmt.Println("고루틴: 채널에 42를 보냅니다.")
		ch <- 42
	}()

	time.Sleep(1 * time.Second)
	val := <-ch
	fmt.Printf("메인: 채널로부터 %d를 받았습니다.\n", val)

	fmt.Println("\n--- 2. 버퍼가 있는 채널 (Buffered Channel) ---")
	bufCh := make(chan string, 2)
	bufCh <- "첫 번째 메시지"
	bufCh <- "두 번째 메시지"
	fmt.Println("채널에서 수신 1:", <-bufCh)
	fmt.Println("채널에서 수신 2:", <-bufCh)

	fmt.Println("\n--- 3. 단방향 채널 ---")
	sendCh := make(chan int)
	go sendData(sendCh)
	receiveData(sendCh)
}

func sendData(ch chan<- int) { ch <- 99 }
func receiveData(ch <-chan int) { fmt.Println("수신된 값 =", <-ch) }
```

#### ③ `28_channel_sync/main.go`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- 1. 채널 닫기 (close) ---")
	ch := make(chan int, 3)
	ch <- 10; ch <- 20; ch <- 30
	close(ch) 

	val1, ok1 := <-ch
	fmt.Printf("값: %d, 정상수신(ok): %t\n", val1, ok1)
	val4, ok4 := <-ch // 빈 채널
	fmt.Printf("닫힌 채널 수신 - 값: %d, 정상수신(ok): %t\n", val4, ok4)

	fmt.Println("\n--- 2. for range를 이용한 채널 순회 ---")
	queue := make(chan string, 5)

	go func() {
		for i := 1; i <= 3; i++ {
			queue <- fmt.Sprintf("작업 %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(queue)
	}()

	for msg := range queue {
		fmt.Println("수신:", msg)
	}
	fmt.Println("모든 작업 수신 완료")
}
```

#### ④ `29_select/main.go`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { time.Sleep(500 * time.Millisecond); ch1 <- "채널 1에서 온 메시지" }()
	go func() { time.Sleep(1 * time.Second); ch2 <- "채널 2에서 온 메시지" }()

	fmt.Println("--- 다중 채널 수신 대기 (select) ---")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("수신:", msg1)
		case msg2 := <-ch2:
			fmt.Println("수신:", msg2)
		}
	}

	fmt.Println("\n--- 타임아웃(Timeout) 처리 ---")
	slowCh := make(chan string)
	
	go func() { time.Sleep(3 * time.Second); slowCh <- "작업 완료 데이터" }()

	select {
	case result := <-slowCh:
		fmt.Println("성공:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("타임아웃 발생! 작업을 취소합니다.")
	}
}
```

#### ⑤ `30_mutex/main.go`

```go
package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock() 
	a.balance += amount
	fmt.Printf("입금: %d (현재 잔액: %d)\n", amount, a.balance)
}

func main() {
	fmt.Println("--- Mutex를 이용한 안전한 데이터 동시 접근 ---")
	account := BankAccount{balance: 0}
	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go account.Deposit(1000, &wg)
	}
	wg.Wait()
	fmt.Printf("\n모든 입금 완료. 최종 잔액: %d\n", account.balance)
}
```

---

## 3. 핵심 교육 내용 요약

### 고루틴(Goroutine)과 동기화 대기

OS 스레드보다 훨씬 가벼운 Go만의 동시성 실행 단위인 고루틴의 동작 원리를 학습하였습니다. 단순히 함수 호출 앞에 `go` 키워드를 붙이는 것만으로 독립적인 비동기 실행이 가능함을 확인하였습니다. 

그러나 메인 고루틴이 종료되면 파생된 고루틴도 강제 종료된다는 점을 실습으로 체험하였고, 이를 해결하기 위해 `time.Sleep` 같은 불안전한 방식 대신 `sync.WaitGroup`을 활용하여 모든 작업이 끝날 때까지 정확하게 동기화(대기)하는 권장 패턴을 습득하였습니다.

### 채널(Channel)을 통한 통신

"메모리를 공유하여 통신하지 말고, 통신하여 메모리를 공유하라"는 Go의 철학을 구현하는 핵심 도구인 채널을 실습하였습니다.

- Unbuffered 채널: 버퍼가 없어 송신자와 수신자가 동시에 준비되어야만(랑데부) 데이터가 전달되는 동기적 특성
- Buffered 채널: 큐(Queue)처럼 작동하여 지정된 용량만큼 대기 없이 전송 가능한 비동기적 특성

또한, 함수 매개변수에 `chan<-` (송신 전용), `<-chan` (수신 전용)을 명시하여 채널의 방향을 강제함으로써 컴파일 단계에서 논리적 오류를 방지하는 방법도 학습하였습니다.

### 채널 닫기 및 다중 제어

데이터의 전송 완료를 알리는 `close(ch)`와 이를 수신 측에서 `v, ok := <-ch`로 감지하는 패턴을 실습하였습니다. 닫힌 채널에서 `for range` 루프가 자동으로 종료되는 우아한 스트림 처리 방식을 체득하였습니다.

더불어, 여러 채널의 이벤트를 동시에 대기하고 처리할 수 있는 `select` 문을 학습하였습니다. 이를 응용하여 `time.After` 채널과 결합, 네트워크 프로그래밍 등에서 필수적인 타임아웃(Timeout) 패턴을 효과적으로 구현하는 방식을 익혔습니다.

### 경쟁 상태(Race Condition)와 Mutex 보호

복수의 고루틴이 동일한 메모리(변수)에 동시 접근해 변경할 때 발생하는 치명적인 데이터 유실 문제(경쟁 상태)를 시뮬레이션하였습니다.

이를 방지하기 위해 `sync.Mutex`를 도입하여 임계 구역(Critical Section) 진입 시 `Lock()`, 종료 시 `Unlock()`을 수행하여 상호 배제(Mutual Exclusion)를 보장하는 전통적이고 강력한 동기화 기법을 완벽히 소화하였습니다.

---

## 4. 강의 평가 및 교육적 성과

이번 6일차 강의에서는 Go 언어 도입의 가장 큰 목적이라 할 수 있는 '강력하고 쉬운 동시성 처리'를 마스터하였습니다.

수강생들은 고루틴이라는 극히 가벼운 스레드를 수십, 수백 개씩 손쉽게 생성해보며 기존 언어들에서 겪던 멀티스레딩 프로그래밍의 복잡성과 성능 제약에 대한 부담을 크게 덜어냈습니다. 특히 `sync.WaitGroup`을 통한 스레드 풀링 및 대기 제어는 실무 백엔드 개발에서 즉시 사용할 수 있는 훌륭한 무기가 되었습니다.

또한 스레드 간 데이터 공유의 어려움을 채널이라는 파이프라인 개념으로 대체하여, 생산자-소비자(Producer-Consumer) 패턴과 같은 복잡한 비동기 아키텍처를 안전하고 직관적으로 설계할 수 있게 되었습니다. 더 나아가 `select`를 이용한 이벤트 루프 및 타임아웃 처리를 실습함으로써 고성능 서버 프로그래밍의 기초를 튼튼히 다졌습니다.

종합적으로 볼 때, 수강생들은 Go 언어의 코어 문법부터 동시성 프로그래밍까지 실무에 필요한 모든 핵심 기능을 섭렵하였습니다. 이제 7일차에서는 지금까지 배운 지식을 모두 통합하여 패키지(Package)와 모듈(Module)을 관리하고, 간단한 웹 서버나 CLI 도구를 제작해 보는 미니 프로젝트 과정으로 매끄럽게 진입할 준비를 100% 완료하였습니다.
