package main

import "fmt"

func main() {
	// [배열(Array): 같은 타입의 데이터를 정해진 개수만큼 담는 바구니 묶음]
	// 배열은 크기가 고정되어 있습니다. 한 번 선언하면 늘리거나 줄일 수 없습니다.

	// [1. 배열 선언]
	// var 변수명 [크기]타입
	var scores [5]int // 정수 5개짜리 배열, 초기값은 모두 0
	fmt.Println("1. 초기 배열 (기본값 0으로 자동 채워짐):", scores)

	// [2. 인덱스로 값 대입]
	// 배열의 각 칸은 0번부터 시작하는 '인덱스(Index)'로 접근합니다.
	scores[0] = 85
	scores[1] = 92
	scores[2] = 78
	scores[3] = 95
	scores[4] = 88
	fmt.Println("2. 값 대입 후:", scores)

	// [3. 배열 리터럴 (선언과 동시에 값 지정)]
	fruits := [3]string{"사과", "바나나", "포도"}
	fmt.Println("3. 선언과 동시에 초기화:", fruits)

	// [4. ... 을 사용한 크기 자동 추론]
	// 초기값의 개수를 보고 컴파일러가 배열 크기를 알아서 결정합니다.
	colors := [...]string{"빨강", "파랑", "초록", "노랑"}
	fmt.Printf("4. 크기 자동 추론 (총 %d개): %v\n", len(colors), colors)

	// [5. for문으로 배열 순회]
	// len(배열): 배열의 원소 개수를 반환합니다.
	fmt.Println("\n5. for문으로 배열 순회:")
	total := 0
	for i := 0; i < len(scores); i++ {
		fmt.Printf("   scores[%d] = %d\n", i, scores[i])
		total += scores[i]
	}
	fmt.Printf("   합계: %d, 평균: %.2f\n", total, float64(total)/float64(len(scores)))

	// [6. range를 활용한 순회 (더 Go다운 방법)]
	// for 인덱스, 값 := range 배열 { ... }
	// 인덱스가 필요 없으면 _로 무시할 수 있습니다.
	fmt.Println("\n6. range로 순회 (인덱스와 값 동시 접근):")
	for index, value := range fruits {
		fmt.Printf("   fruits[%d] = %s\n", index, value)
	}
}
