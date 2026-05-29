package main

import "fmt"

func main() {
	// [중첩 반복문: 반복문 안에 또 반복문 집어넣기]
	// 2차원적인 표(Grid) 형식의 데이터를 다루거나 순서쌍을 연산할 때 자주 활용됩니다.
	// 가장 대표적인 실습 예제인 '구구단'을 작성해 봅니다.

	fmt.Println("--- 중첩 반복문 실습: 구구단 출력 ---")

	// 바깥쪽 for문: 2단부터 9단까지 단수를 제어합니다.
	for dan := 2; dan <= 9; dan++ {
		fmt.Printf("=== %d단 ===\n", dan)
		
		// 안쪽 for문: 각 단마다 1부터 9까지 곱해지는 숫자를 제어합니다.
		for i := 1; i <= 9; i++ {
			fmt.Printf("  %d * %d = %d\n", dan, i, dan*i)
		}
		fmt.Println() // 단마다 구분 빈 줄 출력
	}
}
