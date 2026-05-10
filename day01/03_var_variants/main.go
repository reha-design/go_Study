package main

import "fmt"

func main() {
	// 1. 타입 추론 (Type Inference)
	// 타입을 명시하지 않아도 대입되는 값을 보고 Go가 자동으로 타입을 결정합니다.
	var name = "이순신" // string으로 추론
	var age = 30      // int로 추론
	var score = 95.5  // float64로 추론

	// %v는 값(value), %T는 타입(Type)을 의미합니다.
	fmt.Printf("1. 타입 추론: %v(%T), %v(%T), %v(%T)\n", name, name, age, age, score, score)

	// 2. 단축 변수 선언 (Short Variable Declaration)
	// var 키워드와 타입 없이 ':='를 사용하여 선언과 동시에 할당합니다.
	// ※ 주의: 함수 안에서만 사용 가능합니다.
	city := "Seoul"
	population := 1000

	fmt.Println("2. 단축 선언:", city, population)

	// 3. 여러 변수 동시 선언
	var a, b, c int = 1, 2, 3
	x, y, z := 10, "Hello", true // 서로 다른 타입도 동시에 선언 가능

	fmt.Println("3. 동시 선언:", a, b, c, x, y, z)

	// 4. 그룹 선언
	// 연관된 변수들을 괄호로 묶어서 선언할 수 있습니다.
	var (
		company = "Google"
		rank    = 1
		isOpen  = false
	)

	fmt.Println("4. 그룹 선언:", company, rank, isOpen)

	// 5. 초깃값 없이 선언 (Zero Value)
	// 초깃값을 주지 않으면 각 타입의 기본값으로 초기화됩니다.
	var defaultInt int       // 0
	var defaultFloat float64 // 0
	var defaultString string // ""
	var defaultBool bool     // false

	fmt.Println("5. 기본값(Zero Value):", defaultInt, defaultFloat, defaultString, defaultBool)
}
