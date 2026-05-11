package main

import "fmt"

func main() {
	fmt.Println("=== 1. 변수(var)와 상수(const)의 차이점 ===")
	// 변수는 값을 바꿀 수 있지만, 상수는 바꿀 수 없습니다.
	var changeableVar = "나는 변수야 (변경 가능)"
	const unchangeableConst = "나는 상수야 (변경 불가)"
	
	fmt.Println("초기 상태 ->", changeableVar, "|", unchangeableConst)

	// 변수 값 변경 (정상 동작)
	changeableVar = "변수 값 변경 완료!"
	fmt.Println("변경 후   ->", changeableVar, "|", unchangeableConst)

	// 상수는 값 변경을 시도하면 에러가 발생합니다.
	// 아래 주석을 해제해서 직접 에러를 확인해보세요!
	// unchangeableConst = "상수 값 변경 시도!" // 에러: cannot assign to unchangeableConst

	fmt.Println("\n=== 2. 그룹 상수 선언 ===")
	// 여러 상수를 한 번에 선언할 수 있습니다.
	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
	)
	fmt.Println("그룹 상수 값:", Sunday, Monday, Tuesday)

	fmt.Println("\n=== 3. iota 키워드 (자동 증가) ===")
	// iota를 사용하면 0부터 1씩 자동으로 증가합니다.
	const (
		Apple = iota  // 0
		Banana        // 1
		Grape         // 2
	)
	fmt.Printf("Apple: %d, Banana: %d, Grape: %d\n", Apple, Banana, Grape)

	fmt.Println("\n=== 4. 타입이 없는 상수 (Untyped Constant)의 유연함 ===")
	// maxCount는 타입이 명시되지 않았습니다. (int, float64 등 어떤 것이든 될 수 있는 상태)
	const maxCount = 100
	
	var currentCount int = maxCount
	var currentPrice float64 = maxCount // 같은 상수(maxCount)를 넣었지만 변수 타입에 맞춰서 들어갑니다!

	// %T를 사용해 변수의 타입을 직접 눈으로 확인해봅시다.
	fmt.Printf("currentCount 값: %-4v | 타입: %T\n", currentCount, currentCount)
	fmt.Printf("currentPrice 값: %-4v | 타입: %T\n", currentPrice, currentPrice)
}
