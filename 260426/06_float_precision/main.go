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

	// 5. 오차를 피하는 방법들
	fmt.Println("\n5. 오차를 피하는 방법들:")

	// 방법 1: 반올림 출력 (출력 단계에서 소수점 자리수를 제한)
	// - 단순히 보여줄 때만 맞춰주는 방법. 내부 값은 여전히 오차가 있습니다.
	fmt.Printf("   [방법1] %.1f + %.1f = %.1f  (출력 반올림)\n", a, b, a+b)

	// 방법 2: epsilon(허용 오차) 비교 (math 패키지 없이)
	// - math.Abs 대신 직접 절댓값을 계산합니다.
	// - diff < 0 이면 부호를 뒤집어 양수로 만듭니다.
	const epsilon = 1e-9 // 허용 오차 범위 (10^-9)
	expected := 0.3
	diff := (a + b) - expected
	if diff < 0 {
		diff = -diff // 직접 절댓값 계산 (math.Abs 없이)
	}
	if diff < epsilon {
		fmt.Println("   [방법2] 0.1 + 0.2 ≈ 0.3  (epsilon 비교: true)")
	} else {
		fmt.Println("   [방법2] 0.1 + 0.2 ≠ 0.3  (epsilon 비교: false)")
	}

	// 방법 3: 정수 변환 후 연산 (가장 확실한 방법, math 패키지 없이)
	// - 10^n을 곱해 정수로 만든 뒤 계산합니다.
	// - math.Round 대신 양수에서 0.5를 더한 뒤 int64로 잘라내면 반올림 효과가 납니다.
	//   int64(1.0 + 0.5) = int64(1.5) = 1  ✓
	ai := int64(a*10 + 0.5) // 0.1 → int64(1.5) → 1
	bi := int64(b*10 + 0.5) // 0.2 → int64(2.5) → 2
	result := float64(ai+bi) / 10 // (1 + 2) / 10 = 0.3
	fmt.Printf("   [방법3] 0.1 + 0.2 = %v  (정수 변환 후 연산)\n", result)
}
