# Go 강의 보고서

## 1. 강의 목표 및 성취 수준

본 강의는 Go 언어에서 복수의 데이터를 효율적으로 저장하고 탐색하기 위한 핵심 자료 구조인 배열(Array), 슬라이스(Slice), 맵(Map)의 개념과 동작 원리를 학습하고, 각 자료 구조의 특성과 실무 활용 패턴을 직접 코드로 작성·실행해 보는 것을 목표로 진행하였습니다.

수강생들은 강의와 실습을 통해 다음 내용을 수행할 수 있도록 학습하였습니다.

1. 배열의 고정 크기 특성을 이해하고 선언, 초기화, 인덱스 접근, for 및 range 기반 순회를 수행할 수 있다.
2. 슬라이스의 동적 크기 특성을 이해하고 make, 리터럴 선언, append를 통한 원소 추가, 슬라이싱을 활용할 수 있다.
3. len과 cap의 차이를 이해하고 append 시 내부 메모리 재할당이 발생하는 원리를 설명할 수 있다.
4. 얕은 복사(Shallow Copy)와 깊은 복사(Deep Copy)의 차이를 이해하고, copy 함수를 사용해 독립된 슬라이스를 생성할 수 있다.
5. 맵의 키-값 쌍 구조를 이해하고 선언, 조회, 추가, 삭제, range 순회를 실습할 수 있다.
6. ok 패턴을 사용하여 맵에 키가 존재하는지 안전하게 확인하고, 중첩 맵 및 맵과 슬라이스의 조합을 설계할 수 있다.

---

## 2. 실습 및 제출 산출물 현황

강의 과정에서는 각 주제별 실습 코드를 작성하고 실행 결과를 확인하였습니다. 수강생들은 아래 구조의 예제 코드를 직접 작성하며 Go의 주요 복합 자료 구조의 동작 방식을 검증하였습니다.

```text
day04/
├── 17_array/
│   └── main.go           (배열 선언, 인덱스 접근, for/range 순회 실습)
├── 18_slice/
│   └── main.go           (슬라이스 선언, append, 슬라이싱, len/cap 비교 실습)
├── 19_slice_ops/
│   └── main.go           (얕은·깊은 복사, 원소 삭제, 2차원 슬라이스 실습)
├── 20_map/
│   └── main.go           (맵 선언, 조회, ok 패턴, delete, range 순회 실습)
└── 21_map_ops/
    └── main.go           (빈도수 계산, 중첩 맵, 맵과 슬라이스 조합 실습)
```

### 각 챕터별 상세 실습 소스코드

#### ① `17_array/main.go`

```go
package main

import "fmt"

func main() {
	// [배열(Array): 같은 타입의 데이터를 정해진 개수만큼 담는 바구니 묶음]
	// 배열은 크기가 고정되어 있습니다. 한 번 선언하면 늘리거나 줄일 수 없습니다.

	// [1. 배열 선언]
	var scores [5]int // 정수 5개짜리 배열, 초기값은 모두 0
	fmt.Println("1. 초기 배열 (기본값 0으로 자동 채워짐):", scores)

	// [2. 인덱스로 값 대입]
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
	colors := [...]string{"빨강", "파랑", "초록", "노랑"}
	fmt.Printf("4. 크기 자동 추론 (총 %d개): %v\n", len(colors), colors)

	// [5. for문으로 배열 순회]
	fmt.Println("\n5. for문으로 배열 순회:")
	total := 0
	for i := 0; i < len(scores); i++ {
		fmt.Printf("   scores[%d] = %d\n", i, scores[i])
		total += scores[i]
	}
	fmt.Printf("   합계: %d, 평균: %.2f\n", total, float64(total)/float64(len(scores)))

	// [6. range를 활용한 순회]
	fmt.Println("\n6. range로 순회 (인덱스와 값 동시 접근):")
	for index, value := range fruits {
		fmt.Printf("   fruits[%d] = %s\n", index, value)
	}
}
```

#### ② `18_slice/main.go`

```go
package main

import "fmt"

func main() {
	// [슬라이스(Slice): 크기가 유동적인 배열]

	// [1. 슬라이스 선언 - make 함수 사용]
	s1 := make([]int, 5)
	fmt.Println("1. make로 만든 슬라이스:", s1)

	// [2. 슬라이스 리터럴]
	s2 := []string{"Go", "Python", "Java"}
	fmt.Println("2. 슬라이스 리터럴:", s2)

	// [3. append: 슬라이스에 원소 추가하기]
	s2 = append(s2, "Kotlin")
	s2 = append(s2, "Swift", "Rust")
	fmt.Println("3. append로 원소 추가 후:", s2)

	// [4. 슬라이싱: 슬라이스의 일부를 잘라 새 슬라이스 만들기]
	fmt.Println("\n4. 슬라이싱:")
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("   원본:", nums)
	fmt.Println("   nums[1:3] =", nums[1:3])
	fmt.Println("   nums[:3]  =", nums[:3])
	fmt.Println("   nums[2:]  =", nums[2:])

	// [5. len(길이)과 cap(용량)]
	fmt.Println("\n5. len과 cap:")
	s3 := make([]int, 3, 10)
	fmt.Printf("   len=%d, cap=%d, 내용=%v\n", len(s3), cap(s3), s3)

	var autoGrow []int
	for i := 0; i < 5; i++ {
		autoGrow = append(autoGrow, i)
		fmt.Printf("   append 후: len=%d, cap=%d\n", len(autoGrow), cap(autoGrow))
	}
}
```

#### ③ `19_slice_ops/main.go`

```go
package main

import "fmt"

func main() {
	// [슬라이스 심화 연산: 복사, 삭제, 2D 슬라이스]

	// [1. 얕은 복사(Shallow) vs 깊은 복사(Deep)]
	fmt.Println("--- 1. 얕은 복사 vs 깊은 복사 ---")
	original := []int{1, 2, 3}
	shallow := original
	shallow[0] = 999
	fmt.Println("   [얕은 복사] original 수정됨:", original)

	original2 := []int{1, 2, 3}
	deep := make([]int, len(original2))
	copy(deep, original2)
	deep[0] = 999
	fmt.Println("   [깊은 복사] original2 유지됨:", original2)

	// [2. 특정 인덱스 원소 삭제]
	fmt.Println("\n--- 2. 특정 인덱스 원소 삭제 ---")
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("   삭제 전:", items)
	deleteIndex := 2
	items = append(items[:deleteIndex], items[deleteIndex+1:]...)
	fmt.Println("   인덱스 2 삭제 후:", items)

	// [3. 2차원 슬라이스]
	fmt.Println("\n--- 3. 2차원 슬라이스 ---")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for rowIndex, row := range matrix {
		fmt.Printf("   행[%d]: %v\n", rowIndex, row)
	}
	fmt.Println("   matrix[1][2] (2행 3열):", matrix[1][2])
}
```

#### ④ `20_map/main.go`

```go
package main

import "fmt"

func main() {
	// [맵(Map): 키(Key)로 값(Value)을 즉시 찾는 사전(Dictionary) 구조]

	// [1. 맵 선언 - make 함수 사용]
	studentAge := make(map[string]int)
	studentAge["김철수"] = 20
	studentAge["이영희"] = 22
	studentAge["박민준"] = 21
	fmt.Println("1. 학생 나이 맵:", studentAge)

	// [2. 맵 리터럴]
	capitals := map[string]string{
		"한국": "서울",
		"일본": "도쿄",
		"미국": "워싱턴 D.C.",
	}
	fmt.Println("2. 수도 맵:", capitals)

	// [3. 키로 값 조회]
	fmt.Println("\n3. 값 조회:", capitals["한국"])

	// [4. ok 패턴으로 키 존재 여부 확인]
	fmt.Println("\n4. ok 패턴으로 존재 여부 확인:")
	if age, ok := studentAge["김철수"]; ok {
		fmt.Printf("   김철수의 나이: %d세\n", age)
	}
	if _, ok := studentAge["홍길동"]; !ok {
		fmt.Println("   홍길동은 맵에 존재하지 않습니다.")
	}

	// [5. delete로 요소 삭제]
	fmt.Println("\n5. 요소 삭제:")
	delete(studentAge, "박민준")
	fmt.Println("   박민준 삭제 후:", studentAge)

	// [6. range를 활용한 맵 순회]
	fmt.Println("\n6. range로 맵 순회:")
	for key, value := range capitals {
		fmt.Printf("   %s의 수도 → %s\n", key, value)
	}
}
```

#### ⑤ `21_map_ops/main.go`

```go
package main

import "fmt"

func main() {
	// [맵 심화 활용: 빈도수 계산, 중첩 맵]

	// [1. 단어 빈도수 계산]
	fmt.Println("--- 1. 단어 빈도수 계산 ---")
	words := []string{"go", "is", "fast", "go", "is", "simple", "go"}
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println("   단어들:", words)
	fmt.Println("   빈도수:", wordCount)

	// [2. 중첩 맵 (학생별 과목 점수)]
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

	// [3. 맵 + 슬라이스 조합 (반별 학생 명단)]
	fmt.Println("\n--- 3. 맵 + 슬라이스 조합 ---")
	classRoom := make(map[string][]string)
	classRoom["A반"] = append(classRoom["A반"], "김철수", "이영희")
	classRoom["B반"] = append(classRoom["B반"], "박민준", "최수연")
	classRoom["A반"] = append(classRoom["A반"], "정도현")
	for class, students := range classRoom {
		fmt.Printf("   %s: %v\n", class, students)
	}
}
```

---

## 3. 핵심 교육 내용 요약

### 배열의 구조와 고정 크기 특성

Go의 배열은 동일한 타입의 값을 정해진 개수만큼 연속된 메모리 공간에 저장하는 가장 기본적인 자료 구조입니다. 선언 시 크기가 타입의 일부로 결정되므로 `[3]int`와 `[5]int`는 서로 다른 타입으로 취급됩니다.

배열 선언 방식을 다양하게 실습하였습니다.

- `var scores [5]int`: 크기 5의 정수 배열, 초기값 0으로 자동 설정
- `fruits := [3]string{"사과", "바나나", "포도"}`: 리터럴로 선언과 동시에 초기화
- `colors := [...]string{"빨강", "파랑"}`: `...`으로 컴파일러가 크기를 자동 추론

순회는 `for i := 0; i < len(arr); i++` 형태와 `for index, value := range arr` 형태 모두를 실습하였으며, 후자가 더 간결하고 Go다운 방식임을 확인하였습니다.

### 슬라이스의 동적 확장과 내부 구조

슬라이스는 Go 개발에서 배열보다 훨씬 빈번하게 활용되는 자료 구조입니다. 내부적으로 포인터, 길이(len), 용량(cap)의 세 요소로 구성됩니다.

- len(길이): 현재 슬라이스에 실제로 담긴 원소의 수
- cap(용량): 내부 배열이 메모리 재할당 없이 수용 가능한 최대 원소 수

`append` 함수를 통해 원소를 추가할 때, 현재 cap을 초과하면 Go 런타임이 더 큰 내부 배열을 자동으로 새로 할당하고 기존 데이터를 복사합니다. 이 과정에서 cap이 보통 2배 단위로 증가하는 패턴을 실습을 통해 직접 확인하였습니다.

슬라이싱 문법 `s[시작:끝]`은 원본 슬라이스의 메모리를 공유하는 새로운 슬라이스를 반환하므로, 의도치 않은 원본 변경을 주의해야 한다는 점도 함께 학습하였습니다.

### 얕은 복사와 깊은 복사의 차이

슬라이스는 참조 타입이므로 단순 대입(`=`)을 하면 같은 내부 배열을 가리키는 얕은 복사(Shallow Copy)가 발생합니다. 이 경우 하나를 수정하면 다른 쪽에도 영향을 미칩니다.

완전히 독립된 복사본이 필요할 때는 `copy(목적지, 출처)` 함수를 사용하여 깊은 복사(Deep Copy)를 수행해야 합니다.

```
얕은 복사: shallow := original    → 같은 메모리 공유 → 변경 시 원본도 변경
깊은 복사: copy(deep, original)   → 독립된 메모리    → 변경해도 원본 보존
```

### 맵의 키-값 구조와 활용

맵은 해시 테이블 기반의 키-값 쌍 자료 구조로, 임의의 키를 통해 값을 O(1) 수준의 빠른 속도로 조회할 수 있습니다.

- 선언: `make(map[string]int)` 또는 맵 리터럴 방식
- 조회: `value := m[key]`
- 삭제: `delete(m, key)`
- 순회: `for key, value := range m { ... }` (순서 무작위)

맵에서 존재하지 않는 키를 조회할 경우 해당 타입의 제로값이 반환되므로, 반드시 ok 패턴을 통해 키의 존재 여부를 먼저 확인해야 한다는 원칙을 실습하였습니다.

```go
if value, ok := m[key]; ok {
    // 키가 존재할 때만 실행
}
```

### 자료 구조 조합 패턴

실무에서는 단일 자료 구조보다 조합 패턴이 더 빈번하게 사용됩니다. 이번 강의에서는 두 가지 대표적인 조합 패턴을 실습하였습니다.

- 중첩 맵: `map[string]map[string]int` 형태로 2계층 이상의 분류 구조를 표현
- 맵과 슬라이스 조합: `map[string][]string` 형태로 그룹별 목록 데이터를 관리

또한 `wordCount[word]++` 패턴을 활용한 빈도수 계산은 맵의 제로값 특성을 실용적으로 응용한 대표 기법으로, 실제 데이터 분석과 알고리즘 문제에서 광범위하게 활용됨을 학습하였습니다.

---

## 4. 강의 평가 및 교육적 성과

이번 4일차 강의에서는 Go 언어의 핵심 복합 자료 구조인 배열, 슬라이스, 맵을 중심으로 실무 코드 설계에 필수적인 데이터 저장과 탐색 역량을 집중적으로 수련하였습니다.

수강생들은 고정 크기의 배열과 동적으로 성장하는 슬라이스 사이의 메모리 구조적 차이를 직접 len/cap 출력으로 관찰하면서, 단순히 "슬라이스가 배열보다 편리하다"는 수준을 넘어 내부 재할당 원리까지 체득하였습니다. 특히 얕은 복사와 깊은 복사 간의 동작 차이는 참조 타입 특성에 기인한 것으로, 초보자가 슬라이스를 처음 사용할 때 가장 많이 겪는 버그의 원인이자 해결책임을 실코드 비교를 통해 명확히 파악하였습니다.

맵 학습에서는 ok 패턴을 통한 안전한 키 조회 방식과 순회 순서의 비결정성이라는 두 가지 Go 맵의 고유 특성을 실습하였습니다. 이를 통해 맵을 단순한 조회 도구로만 인식하지 않고, 빈도수 계산이나 그룹 분류와 같은 알고리즘적 사고를 데이터 구조와 결합하는 실전 역량을 갖출 수 있었습니다.

종합적으로 볼 때, 수강생들은 배열·슬라이스·맵이라는 Go 언어의 3대 컬렉션 자료 구조를 능숙하게 활용할 수 있는 기초를 완성하였으며, 5일차에 다룰 구조체(Struct)와 포인터(Pointer), 그리고 Go 언어의 인터페이스(Interface) 기초를 이해하기 위한 준비를 충실히 마쳤습니다.
