package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Println("--- 1. 정수 비교 (Integer Comparison) ---")
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a == b \t: %t\n", a == b) // 같음
	fmt.Printf("a != b \t: %t\n", a != b) // 다름
	fmt.Printf("a < b  \t: %t\n", a < b)  // 작음
	fmt.Printf("a <= b \t: %t\n", a <= b) // 작거나 같음
	fmt.Printf("a > b  \t: %t\n", a > b)  // 큼
	fmt.Printf("a >= b \t: %t\n", a >= b) // 크거나 같음

	fmt.Println("\n--- 2. 문자열 비교 (String Comparison) ---")
	str1 := "Apple"
	str2 := "Banana"
	str3 := "Apple"

	fmt.Printf("str1 = %q, str2 = %q, str3 = %q\n", str1, str2, str3)
	fmt.Printf("str1 == str3 \t: %t\n", str1 == str3) // 같음
	fmt.Printf("str1 != str2 \t: %t\n", str1 != str2) // 다름
	// 문자열은 사전식(Dictionary) 순서로 비교됨 (A가 B보다 앞에 있음)
	fmt.Printf("str1 < str2  \t: %t\n", str1 < str2) 

	fmt.Println("\n--- 3. 실수 비교 시 주의사항 (Float Comparison) ---")
	var f1 float64 = 0.1
	var f2 float64 = 0.2
	var f3 float64 = 0.3
	
	// 부동소수점 정밀도 오차로 인해 0.1 + 0.2는 정확히 0.3이 되지 않음
	fmt.Printf("f1 + f2 == f3 \t: %t\n", f1+f2 == f3) 
	fmt.Printf("f1 + f2 값 \t: %.18f\n", f1+f2)
	fmt.Printf("f3 값      \t: %.18f\n", f3)
}
