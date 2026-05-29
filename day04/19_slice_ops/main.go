package main

import "fmt"

func main() {
	// [슬라이스 심화 연산: 복사, 삭제, 2D 슬라이스]

	// [1. copy: 슬라이스 깊은 복사(Deep Copy)]
	// 단순 대입(=)은 같은 메모리 주소를 공유하는 얕은 복사(Shallow Copy)가 됩니다.
	// 한쪽을 수정하면 다른 쪽도 같이 바뀌는 문제가 생깁니다.
	fmt.Println("--- 1. 얕은 복사(Shallow) vs 깊은 복사(Deep) ---")

	original := []int{1, 2, 3}
	shallow := original    // 얕은 복사: 같은 메모리를 공유
	shallow[0] = 999
	fmt.Println("   [얕은 복사] original 수정됨:", original) // 999, 2, 3으로 바뀜!

	original2 := []int{1, 2, 3}
	deep := make([]int, len(original2))
	copy(deep, original2)  // 깊은 복사: 완전히 독립된 메모리에 복사
	deep[0] = 999
	fmt.Println("   [깊은 복사] original2 유지됨:", original2) // 1, 2, 3 그대로

	// [2. 슬라이스에서 특정 인덱스 원소 삭제]
	// Go에는 내장 삭제 함수가 없으며, append와 슬라이싱을 조합하여 삭제합니다.
	fmt.Println("\n--- 2. 특정 인덱스 원소 삭제 ---")
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("   삭제 전:", items)

	deleteIndex := 2 // "c" 삭제
	// items[:2]와 items[3:]을 이어 붙입니다. (...)은 슬라이스를 펼쳐서 전달
	items = append(items[:deleteIndex], items[deleteIndex+1:]...)
	fmt.Println("   인덱스 2 삭제 후:", items) // a, b, d, e

	// [3. 2차원 슬라이스 (슬라이스 안에 슬라이스)]
	// 표(Grid) 형태의 데이터를 저장할 때 활용합니다.
	fmt.Println("\n--- 3. 2차원 슬라이스 ---")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("   행렬 출력:")
	for rowIndex, row := range matrix {
		fmt.Printf("   행[%d]: %v\n", rowIndex, row)
	}
	fmt.Println("   matrix[1][2] (2행 3열):", matrix[1][2]) // 6
}
