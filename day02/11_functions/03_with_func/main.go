package main

import "fmt"

// [함수 정의: 계산기 한 번만 만들어두기]
// 이렇게 한 번 만들어두면, 몇 명의 학생이 오든 이 기계에 집어넣기만 하면 됩니다.
func printReport(name string, math, eng, hist int) {
	total := math + eng + hist
	avg := float64(total) / 3.0
	fmt.Printf("%s -> 총점: %d, 평균: %.2f\n", name, total, avg)
}

func main() {
	fmt.Println("=== 함수를 사용하는 경우 (자동화의 힘) ===")

	// [장점 1: 가독성]
	// 이전 예제와 비교해보세요. 메인 코드가 한눈에 쏙 들어올 만큼 깨끗해졌습니다!
	// 무엇을 하는지(성적 출력)가 명확히 보입니다.
	printReport("학생1", 85, 92, 78)
	printReport("학생2", 70, 80, 95)
	printReport("학생3", 100, 88, 92)

	// [장점 2: 유지보수]
	// 만약 "평균 소수점을 1자리만 보여달라"는 요청이 오면?
	// 우리는 위쪽의 printReport 함수 안에서 %.2f를 %.1f로 '딱 한 군데'만 바꾸면 됩니다.
	// 그러면 학생이 100명이든 1,000명이든 한 번에 수정이 완료됩니다!
	fmt.Println("\n[결론] 함수는 '중복을 줄이고, 수정을 쉽게 하며, 코드를 읽기 좋게' 만들어줍니다.")
}
