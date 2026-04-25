package main

import "fmt"

func main() {
	// 테스트용 데이터
	var i = 123
	var f = 3.141592
	var s = "Go Language"
	var b = true

	fmt.Println("--- 1. 기본 서식 문자 ---")
	fmt.Printf("값 그대로(%%v): %v\n", i)
	fmt.Printf("타입 출력(%%T): %T\n", i)
	fmt.Printf("값과 타입 같이: %v(%T)\n", s, s)

	fmt.Println("\n--- 2. 정수(Integer) ---")
	fmt.Printf("10진수(%%d): %d\n", i)
	fmt.Printf(" 2진수(%%b): %b\n", i)
	fmt.Printf(" 8진수(%%o): %o\n", i)
	fmt.Printf("16진수(%%x): %x\n", i)
	fmt.Printf("16진수(%%X): %X\n", i)

	fmt.Println("\n--- 3. 실수(Float) ---")
	fmt.Printf("기본 출력(%%f): %f\n", f)
	fmt.Printf("소수점 2자리(%%.2f): %.2f\n", f)
	fmt.Printf("지수 표기(%%e): %e\n", f)

	fmt.Println("\n--- 4. 문자열 및 불리언 ---")
	fmt.Printf("문자열(%%s): %s\n", s)
	fmt.Printf("따옴표 포함(%%q): %q\n", s)
	fmt.Printf("불리언(%%t): %t\n", b)

	fmt.Println("\n--- 5. 정렬 및 패딩 ---")
	fmt.Printf("|%5d| (5칸 확보, 우측 정렬)\n", i)
	fmt.Printf("|%-5d| (5칸 확보, 좌측 정렬)\n", i)
	fmt.Printf("|%05d| (5칸 확보, 빈곳은 0으로)\n", i)

	fmt.Println("\n--- 6. 특수 문자 (Escape Characters) ---")
	fmt.Printf("줄바꿈: \\n\n")
	fmt.Printf("탭 문자: \t <-- 탭\n")
	fmt.Printf("큰따옴표: \"Hello\"\n")
	fmt.Printf("백슬래시: \\ \n")
}
