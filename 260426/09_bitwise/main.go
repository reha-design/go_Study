package main

import "fmt"

func main() {
	var a int = 12 // 1100 (2진수)
	var b int = 10 // 1010 (2진수)

	fmt.Println("--- 비트 연산자 (Bitwise Operators) ---")
	fmt.Printf("a = %2d (%04b)\n", a, a)
	fmt.Printf("b = %2d (%04b)\n", b, b)

	fmt.Println("\n--- 1. 비트 AND (&) ---")
	// 둘 다 1일 때만 1
	// 1100 & 1010 = 1000 (8)
	fmt.Printf("a & b \t= %2d (%04b)\n", a&b, a&b)

	fmt.Println("\n--- 2. 비트 OR (|) ---")
	// 둘 중 하나라도 1이면 1
	// 1100 | 1010 = 1110 (14)
	fmt.Printf("a | b \t= %2d (%04b)\n", a|b, a|b)

	fmt.Println("\n--- 3. 비트 XOR (^) ---")
	// 서로 다를 때만 1
	// 1100 ^ 1010 = 0110 (6)
	fmt.Printf("a ^ b \t= %2d (%04b)\n", a^b, a^b)

	fmt.Println("\n--- 4. 비트 클리어 (&^) ---")
	// a의 비트 중 b의 비트가 1인 위치를 0으로 만듦 (AND NOT)
	// 1100 &^ 1010 = 0100 (4)
	fmt.Printf("a &^ b \t= %2d (%04b)\n", a&^b, a&^b)

	fmt.Println("\n--- 5. 비트 시프트 (<<, >>) ---")
	var c int = 1 // 0001
	fmt.Printf("c \t= %2d (%04b)\n", c, c)
	
	// 왼쪽으로 2칸 이동 (c * 2^2)
	// 0001 << 2 = 0100 (4)
	fmt.Printf("c << 2 \t= %2d (%04b)\n", c<<2, c<<2)
	
	var d int = 8 // 1000
	fmt.Printf("d \t= %2d (%04b)\n", d, d)

	// 오른쪽으로 2칸 이동 (d / 2^2)
	// 1000 >> 2 = 0010 (2)
	fmt.Printf("d >> 2 \t= %2d (%04b)\n", d>>2, d>>2)
}
