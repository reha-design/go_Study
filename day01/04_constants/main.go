package main

import "fmt"

func main() {
	fmt.Println("=== 1. 변수(var)와 상수(const)의 차이점 ===")
	// [변수 vs 상수]
	// 변수(var): 값을 언제든지 바꿀 수 있는 그릇 (예: 나이, 오늘의 기온)
	// 상수(const): 한 번 정하면 절대 바꿀 수 없는 고정된 값 (예: 파이(3.14), 생년월일)
	var changeableVar = "나는 변수야 (변경 가능)"
	const unchangeableConst = "나는 상수야 (변경 불가)"

	fmt.Println("초기 상태 ->", changeableVar, "|", unchangeableConst)

	// 변수는 값을 덮어쓸 수 있습니다.
	changeableVar = "변수 값 변경 완료!"
	fmt.Println("변경 후   ->", changeableVar, "|", unchangeableConst)

	// 상수는 값 변경을 시도하면 빨간 줄(컴파일 에러)이 뜹니다.
	// 실수로 값이 바뀌면 안 되는 중요한 데이터를 지키기 위해 사용합니다.
	// unchangeableConst = "상수 값 변경 시도!" 

	fmt.Println("\n=== 2. 그룹 상수 선언 ===")
	// 괄호를 사용해 연관된 상수를 한 번에 묶어서 선언할 수 있어 코드가 깔끔해집니다.
	const (
		Sunday  = 0
		Monday  = 1
		Tuesday = 2
	)
	fmt.Println("요일별 번호:", Sunday, Monday, Tuesday)

	fmt.Println("\n=== 3. iota 키워드 (자동 번호표) ===")
	// iota는 '아이오타'라고 읽으며, 0부터 시작해서 아래로 갈수록 1씩 자동으로 늘어납니다.
	// 일일이 0, 1, 2... 적기 귀찮을 때 컴퓨터에게 "너가 알아서 번호 좀 매겨줘"라고 하는 것과 같습니다.
	const (
		Apple  = iota // 0이 자동으로 들어갑니다.
		Banana        // 1이 자동으로 들어갑니다.
		Grape         // 2가 자동으로 들어갑니다.
	)
	fmt.Printf("자동 번호 - 사과: %d, 바나나: %d, 포도: %d\n", Apple, Banana, Grape)

	fmt.Println("\n=== 4. 타입이 없는 상수의 유연함 ===")
	// 타입을 정하지 않은 상수는 '만능 열쇠'와 같습니다.
	// 나중에 이 값을 대입하는 변수의 성격(정수/실수 등)에 맞춰서 유연하게 변신합니다.
	const maxCount = 100

	var currentCount int = maxCount     // 정수(int) 변수에 들어가면 정수가 됨
	var currentPrice float64 = maxCount // 실수(float64) 변수에 들어가면 실수가 됨

	fmt.Printf("정수 변수에 담긴 100: %v (타입: %T)\n", currentCount, currentCount)
	fmt.Printf("실수 변수에 담긴 100: %v (타입: %T)\n", currentPrice, currentPrice)
}
