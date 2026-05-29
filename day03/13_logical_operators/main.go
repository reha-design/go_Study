package main

import "fmt"

func main() {
	// [논리 연산자: 조건들을 엮어서 더 스마트한 판단하기]
	// 비교 연산의 결과들을 묶어 종합적인 참/거짓을 판별합니다.
	
	age := 25
	hasLicense := true

	fmt.Println("--- 1. 논리 AND (&&) ---")
	// && : "그리고" (양쪽 조건이 모두 참이어야만 최종 참이 됨)
	if age >= 20 && hasLicense {
		fmt.Println("운전이 가능합니다.")
	} else {
		fmt.Println("운전이 불가능합니다.")
	}

	fmt.Println("\n--- 2. 논리 OR (||) ---")
	// || : "또는" (두 조건 중 단 하나라도 참이면 최종 참이 됨)
	isWeekend := true
	isHoliday := false
	if isWeekend || isHoliday {
		fmt.Println("학교나 회사에 가지 않고 쉽니다!")
	} else {
		fmt.Println("일과를 진행합니다.")
	}

	fmt.Println("\n--- 3. 논리 NOT (!) ---")
	// ! : "아님" (참을 거짓으로, 거짓을 참으로 뒤집음)
	isRainy := false
	if !isRainy {
		fmt.Println("날씨가 맑으니 우산이 필요 없습니다.")
	}
}
