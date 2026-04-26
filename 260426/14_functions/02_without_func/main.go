package main

import "fmt"

func main() {
	fmt.Println("=== 함수를 사용하지 않는 경우 (코드 중복) ===")

	// 학생 1 성적 처리
	math1, eng1, hist1 := 85, 92, 78
	total1 := math1 + eng1 + hist1
	avg1 := float64(total1) / 3.0
	fmt.Printf("학생1 -> 총점: %d, 평균: %.2f\n", total1, avg1)

	// 학생 2 성적 처리 (똑같은 계산 코드를 또 써야 함)
	math2, eng2, hist2 := 70, 80, 95
	total2 := math2 + eng2 + hist2
	avg2 := float64(total2) / 3.0
	fmt.Printf("학생2 -> 총점: %d, 평균: %.2f\n", total2, avg2)

	// 학생 3 성적 처리 (복사/붙여넣기 반복...)
	math3, eng3, hist3 := 100, 88, 92
	total3 := math3 + eng3 + hist3
	avg3 := float64(total3) / 3.0
	fmt.Printf("학생3 -> 총점: %d, 평균: %.2f\n", total3, avg3)

	fmt.Println("\n[문제점] 만약 평균 계산 방식이 바뀌면(예: 4과목으로 변경), 모든 복사된 코드를 일일이 다 수정해야 합니다.")
}
