package main

import (
	"fmt"
	"sync"
)

// [경쟁 상태(Race Condition)와 Mutex]
// 여러 고루틴이 동시에 같은 메모리(변수)에 접근하여 값을 변경하려고 할 때 발생하는 문제를 '경쟁 상태'라고 합니다.
// 이를 방지하기 위해 sync.Mutex (Mutual Exclusion, 상호 배제)를 사용하여 한 번에 하나의 고루틴만 접근하도록 보호합니다.

// 은행 계좌를 표현하는 구조체
type BankAccount struct {
	balance int
	mu      sync.Mutex // 데이터를 보호할 Mutex
}

// 입금 메서드
func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 1. 임계 구역(Critical Section) 진입 전에 Lock()을 걸어 다른 고루틴의 접근을 차단합니다.
	a.mu.Lock()
	
	// 2. 현재 작업이 끝나면 반드시 Unlock()을 호출하여 차단을 해제해야 합니다.
	defer a.mu.Unlock() 

	// 공유 자원에 접근하여 값 변경
	currentBalance := a.balance
	// 일부러 아주 짧은 지연을 주어 경쟁 상태를 유발하기 쉽게 만들 수도 있습니다.
	a.balance = currentBalance + amount
	
	fmt.Printf("입금: %d (현재 잔액: %d)\n", amount, a.balance)
}

func main() {
	fmt.Println("--- Mutex를 이용한 안전한 데이터 동시 접근 ---")
	
	account := BankAccount{balance: 0}
	var wg sync.WaitGroup

	// 10개의 고루틴이 동시에 1000원씩 입금을 시도합니다.
	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go account.Deposit(1000, &wg)
	}

	wg.Wait()
	
	// Mutex를 사용하지 않았다면 결과가 10000보다 작게 나올 확률이 높습니다. (데이터 유실 발생)
	// Mutex로 보호했기 때문에 항상 정확히 10000이 출력됩니다.
	fmt.Printf("\n모든 입금 완료. 최종 잔액: %d\n", account.balance)
}
