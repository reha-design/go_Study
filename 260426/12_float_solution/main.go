package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- 방법 1: 모든 수를 정수로 변환 (Fixed-point) ---")
	// 예: 1.23달러 + 0.12달러를 계산할 때
	// 1.23 -> 123 (센트 단위 정수)
	// 0.12 -> 12  (센트 단위 정수)
	var price1 int = 123
	var price2 int = 12
	sum := price1 + price2

	fmt.Printf("정수 계산 결과: %d 센트\n", sum)
	fmt.Printf("실제 달러 환산: $%.2f\n", float64(sum)/100.0)

	fmt.Println("\n--- 방법 2: 아주 작은 값(Epsilon) 비교 ---")
	var f1 float64 = 0.1
	var f2 float64 = 0.2
	var f3 float64 = 0.3

	// 직접 비교 (false)
	fmt.Printf("f1 + f2 == f3 ? \t: %t\n", f1+f2 == f3)

	// 차이값을 구해서 매우 작은 값(Epsilon)보다 작은지 확인
	const epsilon = 0.000001
	diff := math.Abs((f1 + f2) - f3)
	
	if diff < epsilon {
		fmt.Printf("차이값(%v)이 매우 작으므로 [같음]으로 판정!\n", diff)
	}

	fmt.Println("\n--- 방법 3: math.Nextafter 활용 (참고) ---")
	// 아주 미세하게 값을 조정하여 비교하는 방법도 있습니다.
	// 하지만 보통 방법 1(정수 변환)이 실무에서 가장 많이 쓰입니다.
}
