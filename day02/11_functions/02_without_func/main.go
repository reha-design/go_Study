package main

import "fmt"

func main() {
	fmt.Println("=== 함수를 사용하지 않는 경우 (불편함 체험) ===")

	// [상황] 3명의 학생 성적을 계산해야 합니다.
	
	// 학생 1 성적 처리
	math1, eng1, hist1 := 85, 92, 78
	total1 := math1 + eng1 + hist1
	avg1 := float64(total1) / 3.0 // <-- 이 계산식을 기억하세요!
	fmt.Printf("학생1 -> 총점: %d, 평균: %.2f\n", total1, avg1)

	// 학생 2 성적 처리
	// 똑같은 계산 코드를 또 씁니다. 복사해서 붙여넣기(Ctrl+C, Ctrl+V)를 하게 되죠.
	math2, eng2, hist2 := 70, 80, 95
	total2 := math2 + eng2 + hist2
	avg2 := float64(total2) / 3.0
	fmt.Printf("학생2 -> 총점: %d, 평균: %.2f\n", total2, avg2)

	// 학생 3 성적 처리
	math3, eng3, hist3 := 100, 88, 92
	total3 := math3 + eng3 + hist3
	avg3 := float64(total3) / 3.0
	fmt.Printf("학생3 -> 총점: %d, 평균: %.2f\n", total3, avg3)

	// [생각해볼 점]
	// 만약 평균 계산법이 "3.0으로 나누는 것"에서 "가산점을 주는 방식"으로 바뀐다면?
	// 우리는 위 코드 3곳을 모두 찾아가서 일일이 수정해야 합니다. 
	// 학생이 100명이라면? 100곳을 수정하다가 분명히 한두 군데는 오타가 날 거예요.
	fmt.Println("\n[결론] 똑같은 일을 반복하는 코드는 '함수'라는 기계로 만들어버리는 것이 훨씬 편합니다.")
}
