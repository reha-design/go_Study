package main

import "fmt"

func main() {
	// [맵 심화 활용: 빈도수 계산, 중첩 맵]

	// [1. 맵으로 빈도수(Frequency) 계산하기]
	// 맵의 가장 대표적인 실전 활용 패턴입니다.
	// 존재하지 않는 키를 조회하면 기본값(0)을 반환하는 특성을 활용합니다.
	fmt.Println("--- 1. 단어 빈도수 계산 ---")
	words := []string{"go", "is", "fast", "go", "is", "simple", "go"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++ // 키가 없으면 0에서 시작해서 1씩 증가
	}
	fmt.Println("   단어들:", words)
	fmt.Println("   빈도수:", wordCount)

	// [2. 맵을 값으로 갖는 중첩 맵]
	// 더 복잡한 데이터 구조를 표현할 때 사용합니다.
	fmt.Println("\n--- 2. 중첩 맵 (학생별 과목 점수) ---")
	scores := map[string]map[string]int{
		"김철수": {"수학": 90, "영어": 85, "과학": 92},
		"이영희": {"수학": 78, "영어": 95, "과학": 88},
	}
	for student, subjects := range scores {
		fmt.Printf("   [%s] 성적:\n", student)
		for subject, score := range subjects {
			fmt.Printf("      %s: %d점\n", subject, score)
		}
	}

	// [3. 슬라이스와 맵의 조합]
	// 맵의 값으로 슬라이스를 사용하여 그룹화할 수 있습니다.
	fmt.Println("\n--- 3. 맵 + 슬라이스 조합 (반별 학생 명단) ---")
	classRoom := make(map[string][]string)
	classRoom["A반"] = append(classRoom["A반"], "김철수", "이영희")
	classRoom["B반"] = append(classRoom["B반"], "박민준", "최수연")
	classRoom["A반"] = append(classRoom["A반"], "정도현") // A반에 추가 학생

	for class, students := range classRoom {
		fmt.Printf("   %s: %v\n", class, students)
	}
}
