package main

import "fmt"

// [구조체(Struct): 여러 타입의 필드를 하나로 묶는 사용자 정의 타입]
// Go에는 클래스(Class)가 없습니다. 대신 구조체(struct)로 데이터를 묶습니다.

// 구조체 정의는 함수 밖(패키지 레벨)에 작성합니다.
// type 이름 struct { 필드명 타입 }
type Person struct {
	Name string
	Age  int
	City string
}

// [중첩 구조체: 구조체 안에 구조체를 포함]
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // 구조체 안에 다른 구조체를 필드로 포함
}

func main() {
	// [1. 구조체 선언 및 초기화 방법 1: 필드명 지정]
	p1 := Person{
		Name: "김철수",
		Age:  25,
		City: "서울",
	}
	fmt.Println("1. 구조체 초기화 (필드명 지정):", p1)

	// [2. 구조체 초기화 방법 2: var로 선언 후 필드에 개별 대입]
	var p2 Person
	p2.Name = "이영희"
	p2.Age = 30
	p2.City = "부산"
	fmt.Println("2. 구조체 초기화 (var 후 대입):", p2)

	// [3. 구조체 필드 접근]
	// 구조체변수.필드명 으로 접근합니다.
	fmt.Println("\n3. 필드 개별 접근:")
	fmt.Printf("   이름: %s, 나이: %d세, 도시: %s\n", p1.Name, p1.Age, p1.City)

	// [4. 필드 값 변경]
	p1.Age = 26
	fmt.Println("4. 필드 값 변경 후:", p1)

	// [5. 중첩 구조체 접근]
	fmt.Println("\n5. 중첩 구조체:")
	emp := Employee{
		Name: "박민준",
		Age:  32,
		Address: Address{
			Street: "테헤란로 123",
			City:   "서울",
			Zip:    "06234",
		},
	}
	// 중첩 구조체의 필드는 .을 연속으로 사용합니다.
	fmt.Printf("   이름: %s, 주소: %s %s\n", emp.Name, emp.Address.City, emp.Address.Street)

	// [6. 구조체 슬라이스: 구조체를 여러 개 관리하기]
	fmt.Println("\n6. 구조체 슬라이스 (팀 명단):")
	team := []Person{
		{Name: "홍길동", Age: 28, City: "인천"},
		{Name: "강감찬", Age: 35, City: "대전"},
		{Name: "이순신", Age: 42, City: "광주"},
	}
	for _, member := range team {
		fmt.Printf("   이름: %-6s  나이: %d세  도시: %s\n", member.Name, member.Age, member.City)
	}
}
