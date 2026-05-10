package main

import "fmt"

func main() {
	// 1. 상수 선언 (const)
	// 상수는 선언과 동시에 값을 할당해야 하며, 이후에 값을 변경할 수 없습니다.
	const pi float64 = 3.141592
	const title = "Go Study" // 타입 추론 가능

	fmt.Println("1. 상수 출력:", pi, title)

	// 2. 값 변경 시도 (에러 발생)
	// pi = 3.14 // 주석을 해제하면 'cannot assign to pi' 에러 발생
	
	// 3. 그룹 상수 선언 (많이 사용됨)
	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
	)
	fmt.Println("2. 그룹 상수:", Sunday, Monday, Tuesday)

	// 4. iota 키워드
	// iota를 사용하면 0부터 1씩 자동으로 증가하는 상수를 만들 수 있습니다.
	const (
		Apple = iota  // 0
		Banana        // 1
		Grape         // 2
	)
	fmt.Println("3. iota 상수:", Apple, Banana, Grape)

	// 5. 타입이 없는 상수 (Untyped Constant)
	// 상수는 필요한 시점에 타입이 결정되는 유연함을 가지고 있습니다.
	const maxCount = 100
	var currentCount int = maxCount
	var currentPrice float64 = maxCount // int형 변수에도, float64형 변수에도 대입 가능!

	fmt.Println("4. 유연한 상수:", currentCount, currentPrice)
}
