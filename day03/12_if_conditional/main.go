package main

import "fmt"

func main() {
	// [1. 기본 if문]
	// 조건이 참(true)일 때만 블록 내부의 코드를 실행합니다.
	score := 85

	fmt.Println("--- 1. 기본 if-else 구조 ---")
	if score >= 90 {
		fmt.Println("합격 등급: A")
	} else if score >= 80 {
		// 바로 앞 조건이 거짓이고, 현재 조건이 참일 때 실행됩니다.
		fmt.Println("합격 등급: B")
	} else {
		// 위의 모든 조건이 거짓일 때 최종 실행됩니다.
		fmt.Println("합격 등급: C")
	}

	// [2. if 초기화 문법]
	// Go 언어에서는 if 조건식 직전에 짧은 변수 선언 및 초기화를 수행할 수 있습니다.
	// 이 변수는 해당 if-else 블록 내부에서만 살아있고 블록이 끝나면 소멸합니다. (메모리 절약에 효과적)
	fmt.Println("\n--- 2. if 초기화 문법 ---")
	if tempScore := 95; tempScore >= 90 {
		fmt.Printf("임시 점수 %d는 A등급입니다.\n", tempScore)
	}
	// 이곳에서 tempScore 변수를 출력하려고 하면 컴파일 에러가 발생합니다.
	// fmt.Println(tempScore) // 에러!
}
