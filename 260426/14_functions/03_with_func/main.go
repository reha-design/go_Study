package main

import "fmt"

// 성적 리포트를 출력하는 함수를 딱 한 번만 정의합니다.
func printReport(name string, math, eng, hist int) {
	total := math + eng + hist
	avg := float64(total) / 3.0
	fmt.Printf("%s -> 총점: %d, 평균: %.2f\n", name, total, avg)
}

func main() {
	fmt.Println("=== 함수를 사용하는 경우 (재사용성) ===")

	// 이제 학생 정보만 넘겨주면 끝! 코드가 아주 간결해집니다.
	printReport("학생1", 85, 92, 78)
	printReport("학생2", 70, 80, 95)
	printReport("학생3", 100, 88, 92)

	fmt.Println("\n[장점] 계산 방식이 바뀌어도 printReport 함수 한 곳만 수정하면 모든 학생에게 바로 적용됩니다.")
}
