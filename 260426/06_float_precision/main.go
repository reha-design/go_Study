package main

import "fmt"

func main() {
	// 1. 실수 타입 (float32 vs float64)
	// Go에서 실수의 기본 타입은 float64입니다. (더 정밀함)
	var f1 float32 = 0.1234567890123456789
	var f2 float64 = 0.1234567890123456789

	fmt.Println("1. 정밀도 비교:")
	fmt.Printf("   float32: %v\n", f1) // 약 7자리까지 정확
	fmt.Printf("   float64: %v\n", f2) // 약 15자리까지 정확

	// 2. 지수 표기법 (Scientific Notation)
	// 아주 크거나 작은 수를 표현할 때 사용합니다.
	var e1 float64 = 1.23e4  // 1.23 * 10^4 = 12300
	var e2 float64 = 1.23e-2 // 1.23 * 10^-2 = 0.0123

	fmt.Println("\n2. 지수 표기법:")
	fmt.Printf("   1.23e4  : %v\n", e1)
	fmt.Printf("   1.23e-2 : %v\n", e2)

	// 3. 실수 출력 포맷팅
	// 소수점 자리수를 제한하여 출력할 때 자주 사용합니다.
	fmt.Println("\n3. 포맷팅 출력:")
	fmt.Printf("   소수점 2자리까지: %.2f\n", f2)
	fmt.Printf("   전체 10칸, 소수점 3자리: %10.3f\n", f2)

	// 4. 중요: 실수의 정밀도 한계 (오차)
	// 컴퓨터는 이진수를 사용하기 때문에 십진수 실수를 완벽하게 표현하지 못할 때가 있습니다.
	fmt.Println("\n4. 실수의 계산 오차 주의:")
	var a float64 = 0.1
	var b float64 = 0.2
	fmt.Printf("   0.1 + 0.2 = %v\n", a+b) 
	// 결과가 0.3이 아니라 0.30000000000000004가 나옵니다!
}
