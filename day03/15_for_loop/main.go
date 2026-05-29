package main

import "fmt"

func main() {
	// [for문: 컴퓨터가 제일 잘하는 반복 작업 자동화]
	// Go 언어에는 다른 언어에 있는 while문이 없으며, 오직 'for' 키워드 하나로 모든 반복을 해결합니다.

	// [1. 기본 for문]
	// for 초기문; 조건식; 증감식 { ... }
	fmt.Println("--- 1. 기본 1부터 5까지 반복 ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// [2. 조건식만 있는 for문 (while 형태)]
	// 초기문과 증감식을 생략하고 조건만 남겨 다른 언어의 while문처럼 동작하게 합니다.
	fmt.Println("\n--- 2. 조건만 있는 반복문 ---")
	count := 3
	for count > 0 {
		fmt.Printf("카운트다운: %d\n", count)
		count--
	}

	// [3. break와 continue 제어]
	// break: 반복문을 그 자리에서 완전히 폭파(종료)하고 탈출합니다.
	// continue: 아래 코드를 무시하고 다음 바퀴(반복 주기)로 즉시 넘어갑니다.
	fmt.Println("\n--- 3. break & continue ---")
	fmt.Println("   [1부터 10 중 짝수만 출력하되, 8을 만나면 탈출]")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // 홀수면 아래 출력 건너뛰고 다음 바퀴로 점프
		}
		if i == 8 {
			fmt.Println(" 8을 만났으니 중단합니다!")
			break // 반복 완전 탈출
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
