package main

import "fmt"

func main() {
	// 정수형 변수 선언 및 값 할당
	var age int = 20

	// 실수형 변수 선언 및 값 할당
	var height float64 = 175.5

	// 문자열 변수 선언 및 값 할당
	var name string = "홍길동"

	// 불리언 변수 선언 및 값 할당
	var isStudent bool = true

	// 선언된 변수들의 값을 화면에 출력
	fmt.Println("이름:", name)
	fmt.Println("나이:", age)
	fmt.Println("키:", height, "cm")
	fmt.Println("학생 여부:", isStudent)

	// 변수 값 변경

	age = 21
	isStudent = false

	// 변경된 값 출력
	fmt.Println("변경된 나이:", age)
	fmt.Println("변경된 학생 여부:", isStudent)
}
