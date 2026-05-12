package main

import "fmt"

// [함수 정의: 특별한 기능을 하는 기계 만들기]
// func 함수이름(재료1, 재료2...) (결과물1, 결과물2...) { ... }
// math, eng, history라는 3가지 점수를 받아서 total(합계)과 avg(평균)라는 2가지 결과를 돌려줍니다.
func getGrade(math, eng, history int) (total int, avg float64) {
	total = math + eng + history
	avg = float64(total) / 3.0
	// return: 결과물을 바깥으로 내보내고 함수를 끝냅니다.
	// 미리 total, avg라고 이름을 지어놨기 때문에 그냥 return만 써도 됩니다.
	return 
}

func main() {
	// 1. 함수에 넣을 재료들을 준비합니다.
	mathScore := 85
	engScore := 92
	historyScore := 78

	// 2. 함수 호출: 기계를 가동시키고 결과물을 받습니다.
	// Go 언어의 멋진 점은 결과물을 한 번에 여러 개(total, avg) 받을 수 있다는 것입니다!
	total, avg := getGrade(mathScore, engScore, historyScore)

	// 3. 결과 출력
	fmt.Println("--- 성적 결과 리포트 ---")
	fmt.Printf("총점: %d점\n", total)
	fmt.Printf("평균: %.2f점\n", avg)
}
