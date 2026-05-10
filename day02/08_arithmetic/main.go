package main

import "fmt"

func main() {
	// 1. 기본 산술 연산자: +, -, *, /, %
	// Go의 산술 연산자는 C/Java와 동일합니다.
	a, b := 17, 5

	fmt.Println("1. 기본 산술 연산자 (a=17, b=5):")
	fmt.Printf("   a + b = %d\n", a+b)  // 22
	fmt.Printf("   a - b = %d\n", a-b)  // 12
	fmt.Printf("   a * b = %d\n", a*b)  // 85
	fmt.Printf("   a / b = %d\n", a/b)  // 3  ← 정수 나눗셈: 소수점 버림
	fmt.Printf("   a %% b = %d\n", a%b) // 2  ← 나머지 (Modulo)

	// 2. 정수 나눗셈 vs 실수 나눗셈
	// 정수끼리 나누면 소수점이 버려집니다.
	// 실수 결과가 필요하면 float64로 변환 후 나눠야 합니다.
	fmt.Println("\n2. 정수 나눗셈 vs 실수 나눗셈:")
	fmt.Printf("   17 / 5        = %d   (정수 나눗셈, 소수 버림)\n", a/b)
	fmt.Printf("   17.0 / 5.0    = %f (실수 나눗셈)\n", 17.0/5.0)
	fmt.Printf("   float64(a)/float64(b) = %f (타입 변환 후 나눗셈)\n", float64(a)/float64(b))

	// 3. 나머지 연산자 (%) — 정수 전용
	// 부호는 피제수(앞 숫자)를 따라갑니다.
	fmt.Println("\n3. 나머지(%) 부호 규칙:")
	fmt.Printf("    7 %% 3 = %d\n", 7%3)   //  1
	fmt.Printf("   -7 %% 3 = %d\n", -7%3)  // -1 ← 부호는 -7(피제수) 기준
	fmt.Printf("    7 %% -3 = %d\n", 7%-3) //  1 ← 부호는 7(피제수) 기준

	// 4. 복합 대입 연산자: +=, -=, *=, /=, %=
	// 변수에 연산 결과를 바로 저장합니다.
	fmt.Println("\n4. 복합 대입 연산자:")
	x := 10
	x += 3
	fmt.Printf("   x += 3  → x = %d\n", x) // 13
	x -= 2
	fmt.Printf("   x -= 2  → x = %d\n", x) // 11
	x *= 4
	fmt.Printf("   x *= 4  → x = %d\n", x) // 44
	x /= 5
	fmt.Printf("   x /= 5  → x = %d\n", x) // 8
	x %= 3
	fmt.Printf("   x %%= 3 → x = %d\n", x) // 2

	// 5. 증감 연산자: ++, --
	// Go에서 ++, --는 "문장(statement)"입니다.
	// C/Java처럼 식(expression)으로 사용할 수 없습니다.
	//   y := x++  ← Go에서는 컴파일 에러!
	fmt.Println("\n5. 증감 연산자 (++ / --):")
	n := 5
	n++ // n = n + 1
	fmt.Printf("   n++ → n = %d\n", n) // 6
	n-- // n = n - 1
	fmt.Printf("   n-- → n = %d\n", n) // 5
	// 전위(++n)는 Go에서 지원하지 않습니다.

	// 6. 연산자 우선순위
	// *, /, % 가 +, - 보다 먼저 계산됩니다. (수학과 동일)
	fmt.Println("\n6. 연산자 우선순위:")
	fmt.Printf("   2 + 3*4    = %d  (3*4 먼저)\n", 2+3*4)      // 14
	fmt.Printf("   (2+3)*4   = %d  (괄호 먼저)\n", (2+3)*4)    // 20
	fmt.Printf("   10/2 + 3  = %d  (10/2 먼저)\n", 10/2+3)     // 8
	fmt.Printf("   10%%3 + 1  = %d  (10%%3 먼저)\n", 10%3+1)   // 2
}
