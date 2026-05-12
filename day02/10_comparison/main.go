package main

import "fmt"

func main() {
	// [비교 연산자: 참(true)과 거짓(false) 판별기]
	// 비교 연산의 결과는 항상 불리언(bool) 타입인 true 또는 false로 나옵니다.
	
	var a int = 10
	var b int = 20

	fmt.Println("--- 1. 숫자 비교 ---")
	fmt.Println("a =", a, ", b =", b)
	fmt.Println("a == b (둘이 같아?)      :", a == b) // 결과: false
	fmt.Println("a != b (둘이 달라?)      :", a != b) // 결과: true (다르니까!)
	fmt.Println("a <  b (a가 더 작아?)     :", a < b)  // 결과: true
	fmt.Println("a >  b (a가 더 커?)      :", a > b)  // 결과: false

	fmt.Println("\n--- 2. 문자열 비교 (사전 찾기 방식) ---")
	str1 := "Apple"
	str2 := "Banana"
	str3 := "Apple"

	fmt.Printf("str1 = %q, str2 = %q, str3 = %q\n", str1, str2, str3)
	// 글자도 비교가 가능합니다! "Apple"과 "Apple"은 완벽히 같습니다.
	fmt.Println("str1 == str3 (글자가 똑같아?) :", str1 == str3) 
	
	// 글자의 크기 비교는 '사전 순서'입니다. 
	// A가 B보다 앞에 나오므로 "Apple"이 "Banana"보다 작다고 판단합니다.
	fmt.Println("str1 < str2  (A가 B보다 앞에 있어?) :", str1 < str2) 

	fmt.Println("\n--- 3. 실수 비교의 함정 (복습) ---")
	var f1 float64 = 0.1
	var f2 float64 = 0.2
	var f3 float64 = 0.3
	
	// [중요!] 컴퓨터의 미세한 오차 때문에 0.1 + 0.2는 0.3과 같지 않다고 나옵니다.
	// 그래서 실수끼리는 '==' 로 직접 비교하면 위험합니다.
	fmt.Printf("0.1 + 0.2 == 0.3 이 과연 맞을까? : %t\n", f1+f2 == f3) 
	fmt.Printf("실제 0.1+0.2 계산값: %.18f\n", f1+f2)
	fmt.Printf("실제 0.3 값:         %.18f\n", f3)
}
