package main

import "fmt"

func main() {
	// 테스트용 데이터들
	var i = 123
	var f = 3.141592
	var s = "Go Language"
	var b = true

	// [1. 가장 많이 쓰는 '만능' 서식]
	fmt.Println("--- 1. 자주 쓰는 서식 문자 ---")
	// %v: value의 약자. 타입을 몰라도 알아서 적절한 형식으로 보여줍니다. (가장 많이 씀!)
	fmt.Println("값 그대로 출력 (%v):", i)
	// %T: Type의 약자. 이 변수의 바구니 종류가 무엇인지 보여줍니다.
	fmt.Println("데이터 타입 출력 (%T):", i)
	
	// %v와 %T만 알아도 웬만한 출력은 다 해결할 수 있습니다.
	fmt.Printf("값: %v, 타입: %T\n", s, s)

	// [2. 정수(Integer)를 위한 서식]
	fmt.Println("\n--- 2. 정수 (숫자) 전용 ---")
	// %d: decimal(10진수)의 약자. 우리가 평소 쓰는 숫자를 출력합니다.
	fmt.Printf("10진수(%%d): %d\n", i)
	// %b: binary(2진수)의 약자. 컴퓨터가 이해하는 0과 1로 보여줍니다.
	fmt.Printf(" 2진수(%%b): %b\n", i)
	// %x: hex(16진수)의 약자. 메모리 주소 등을 표현할 때 씁니다.
	fmt.Printf("16진수(%%x): %x\n", i)

	// [3. 실수(Float)를 위한 서식]
	fmt.Println("\n--- 3. 실수 (소수점) 전용 ---")
	// %f: float의 약자. 소수점 아래 6자리까지 기본으로 보여줍니다.
	fmt.Printf("기본 출력(%%f): %f\n", f)
	// %.2f: 마침표(.) 뒤의 숫자로 소수점 몇 자리까지 보여줄지 정합니다. (중요!)
	fmt.Printf("소수점 2자리(%%.2f): %.2f\n", f)

	// [4. 문자열 및 불리언 서식]
	fmt.Println("\n--- 4. 문자열 및 참/거짓 ---")
	// %s: string의 약자. 글자 그대로 출력합니다.
	fmt.Printf("문자열(%%s): %s\n", s)
	// %q: quote의 약자. 따옴표를 포함해서 출력해줍니다. (구분하기 좋음)
	fmt.Printf("따옴표 포함(%%q): %q\n", s)
	// %t: true/false의 약자. 불리언 값을 출력합니다.
	fmt.Printf("불리언(%%t): %t\n", b)

	// [5. 줄 맞춤 기능]
	fmt.Println("\n--- 5. 정렬 및 패딩 ---")
	fmt.Printf("|%5d| (5칸 공간 확보 후 우측 정렬)\n", i)
	fmt.Printf("|%-5d| (5칸 공간 확보 후 좌측 정렬)\n", i)
	fmt.Printf("|%05d| (5칸 공간 확보 후 빈 곳은 0으로 채우기)\n", i)

	// [6. 특수 문자 (제어 문자)]
	fmt.Println("\n--- 6. 특수 기능 (이스케이프 문자) ---")
	fmt.Printf("줄바꿈: \\n 은 엔터키를 친 것과 같습니다.\n")
	fmt.Printf("탭 문자: \\t 는 \t 탭키를 누른 것과 같습니다.\n")
	fmt.Printf("큰따옴표: \\\" 는 글자 안에서 \"따옴표\"를 쓰고 싶을 때 씁니다.\n")
}
