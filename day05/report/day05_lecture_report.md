# Go 강의 보고서

## 1. 강의 목표 및 성취 수준

본 강의는 Go 언어에서 객체지향 설계를 구현하는 핵심 메커니즘인 구조체(Struct)와 메서드(Method), 메모리를 직접 제어하는 포인터(Pointer), 그리고 Go 고유의 인터페이스(Interface) 기반 다형성 설계 패턴을 학습하고, 각 개념을 직접 코드로 작성·실행해 보는 것을 목표로 진행하였습니다.

수강생들은 강의와 실습을 통해 다음 내용을 수행할 수 있도록 학습하였습니다.

1. 구조체(struct)를 정의하여 연관된 데이터를 하나의 타입으로 묶고, 다양한 초기화 방식 및 필드 접근 방법을 활용할 수 있다.
2. 중첩 구조체와 구조체 슬라이스를 조합하여 계층적 데이터 구조를 표현할 수 있다.
3. 구조체에 메서드를 정의하고, 값 수신자(Value Receiver)와 포인터 수신자(Pointer Receiver)의 동작 차이를 이해하고 선택적으로 적용할 수 있다.
4. 포인터의 주소 연산자(`&`)와 역참조 연산자(`*`)를 이해하고, 값 전달(Call by Value)과 포인터 전달(Call by Pointer)의 차이를 코드로 검증할 수 있다.
5. 인터페이스를 정의하고, 구조체가 암묵적으로 인터페이스를 구현하는 Go의 덕 타이핑(Duck Typing) 방식을 이해할 수 있다.
6. 인터페이스 슬라이스와 타입 단언(Type Assertion)을 활용하여 다형성(Polymorphism) 기반의 코드를 작성할 수 있다.

---

## 2. 실습 및 제출 산출물 현황

강의 과정에서는 각 주제별 실습 코드를 작성하고 실행 결과를 확인하였습니다. 수강생들은 아래 구조의 예제 코드를 직접 작성하며 Go 언어의 타입 시스템과 메모리 모델을 검증하였습니다.

```text
day05/
├── 22_struct/
│   └── main.go           (구조체 선언, 다양한 초기화, 중첩 구조체, 구조체 슬라이스 실습)
├── 23_struct_methods/
│   └── main.go           (값 수신자·포인터 수신자 메서드 정의 및 동작 차이 검증)
├── 24_pointer/
│   └── main.go           (주소·역참조 연산자, 값 전달·포인터 전달 비교, nil 포인터 실습)
└── 25_interface/
    └── main.go           (인터페이스 정의, 다형성, 인터페이스 슬라이스, 타입 단언 실습)
```

### 각 챕터별 상세 실습 소스코드

#### ① `22_struct/main.go`

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	// [1. 필드명 지정 초기화]
	p1 := Person{Name: "김철수", Age: 25, City: "서울"}
	fmt.Println("1. 구조체 초기화 (필드명 지정):", p1)

	// [2. var 선언 후 필드 개별 대입]
	var p2 Person
	p2.Name = "이영희"
	p2.Age = 30
	p2.City = "부산"
	fmt.Println("2. 구조체 초기화 (var 후 대입):", p2)

	// [3. 필드 접근 및 변경]
	fmt.Printf("   이름: %s, 나이: %d세, 도시: %s\n", p1.Name, p1.Age, p1.City)
	p1.Age = 26
	fmt.Println("4. 필드 값 변경 후:", p1)

	// [5. 중첩 구조체]
	emp := Employee{
		Name: "박민준", Age: 32,
		Address: Address{Street: "테헤란로 123", City: "서울", Zip: "06234"},
	}
	fmt.Printf("   이름: %s, 주소: %s %s\n", emp.Name, emp.Address.City, emp.Address.Street)

	// [6. 구조체 슬라이스]
	team := []Person{
		{Name: "홍길동", Age: 28, City: "인천"},
		{Name: "강감찬", Age: 35, City: "대전"},
		{Name: "이순신", Age: 42, City: "광주"},
	}
	for _, member := range team {
		fmt.Printf("   이름: %-6s  나이: %d세  도시: %s\n", member.Name, member.Age, member.City)
	}
}
```

#### ② `23_struct_methods/main.go`

```go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// [값 수신자: 읽기 전용, 원본 불변]
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) String() string {
	return fmt.Sprintf("사각형(너비: %.1f, 높이: %.1f)", r.Width, r.Height)
}

// [포인터 수신자: 원본 직접 변경]
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println(rect.String())
	fmt.Printf("넓이: %.2f, 둘레: %.2f\n", rect.Area(), rect.Perimeter())

	// 포인터 수신자로 크기 2배 확장
	rect.Scale(2.0)
	fmt.Printf("2배 확장 후: %s\n", rect.String())

	// 슬라이스와 메서드 조합
	shapes := []Rectangle{{4, 6}, {10, 2}, {3, 3}}
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	fmt.Printf("총 넓이 합계: %.2f\n", totalArea)
}
```

#### ③ `24_pointer/main.go`

```go
package main

import "fmt"

// [값 전달: 원본 불변]
func doubleByValue(n int) {
	n = n * 2
	fmt.Printf("함수 내부: n = %d\n", n)
}

// [포인터 전달: 원본 직접 변경]
func doubleByPointer(n *int) {
	*n = *n * 2
	fmt.Printf("함수 내부: *n = %d\n", *n)
}

func main() {
	// 포인터 기본
	x := 42
	p := &x
	fmt.Printf("x의 값: %d, 주소: %p\n", x, &x)
	fmt.Printf("*p(역참조): %d\n", *p)

	*p = 100
	fmt.Printf("*p = 100 후: x = %d\n", x)

	// 값 전달 vs 포인터 전달
	a := 10
	doubleByValue(a)
	fmt.Printf("값 전달 후: a = %d (불변)\n", a)

	doubleByPointer(&a)
	fmt.Printf("포인터 전달 후: a = %d (변경됨)\n", a)

	// nil 포인터
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nilPtr은 nil — 역참조 시 패닉 발생 주의!")
	}
}
```

#### ④ `25_interface/main.go`

```go
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
	Describe() string
}

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) Describe() string {
	return fmt.Sprintf("사각형(%.1f x %.1f)", r.Width, r.Height)
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }
func (c Circle) Describe() string {
	return fmt.Sprintf("원(반지름: %.1f)", c.Radius)
}

// 인터페이스를 매개변수로 받는 함수 — 다형성의 핵심
func printShapeInfo(s Shape) {
	fmt.Printf("도형: %-25s 넓이: %8.2f  둘레: %8.2f\n",
		s.Describe(), s.Area(), s.Perimeter())
}

func main() {
	// 다형성: 같은 함수에 다른 타입 전달
	printShapeInfo(Rectangle{Width: 10, Height: 5})
	printShapeInfo(Circle{Radius: 7})

	// 인터페이스 슬라이스로 일괄 처리
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 5},
		Rectangle{Width: 2, Height: 9},
	}
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
		printShapeInfo(s)
	}
	fmt.Printf("전체 넓이 합계: %.2f\n", totalArea)

	// 타입 단언
	var s Shape = Circle{Radius: 3}
	if c, ok := s.(Circle); ok {
		fmt.Printf("Circle 확인 — 반지름: %.1f\n", c.Radius)
	}

	// 빈 인터페이스
	var anything interface{}
	anything = 42
	fmt.Printf("int 저장: %v (%T)\n", anything, anything)
	anything = "hello"
	fmt.Printf("string 저장: %v (%T)\n", anything, anything)
}
```

---

## 3. 핵심 교육 내용 요약

### 구조체의 정의와 초기화

Go는 클래스(Class)를 제공하지 않는 대신, `struct` 키워드를 통해 여러 타입의 필드를 하나의 사용자 정의 타입으로 묶는 구조체를 제공합니다. 구조체는 패키지 레벨에서 정의하며, 초기화 방법은 다음과 같이 다양합니다.

- 필드명 지정 초기화: `Person{Name: "김철수", Age: 25}` — 가장 명확하고 권장되는 방식
- `var` 선언 후 개별 대입: 각 필드를 필요에 따라 차례로 채우는 방식
- 구조체 안에 다른 구조체를 필드로 포함하는 중첩 구조체를 통해 계층적 데이터 모델을 표현하였습니다.

### 메서드와 수신자 타입

구조체에 메서드를 부착하는 방식으로 Go에서 객체지향적 설계를 구현합니다. 메서드 정의 시 수신자 타입 선택이 핵심입니다.

- 값 수신자 `(r Rectangle)`: 구조체의 복사본을 받아 처리합니다. 내부에서 필드를 수정해도 원본에 영향이 없으며, 주로 조회(Read) 목적의 메서드에 사용합니다.
- 포인터 수신자 `(r *Rectangle)`: 구조체의 메모리 주소를 직접 받아 처리합니다. 원본 필드를 변경해야 할 때 반드시 사용합니다.

Go는 메서드 호출 시 값 변수에서 포인터 수신자 메서드를 호출하더라도 `&` 연산자를 자동으로 적용해주어 편의성을 제공합니다.

### 포인터와 메모리 제어

포인터는 변수가 저장된 메모리 주소를 값으로 갖는 특수한 타입입니다.

- `&변수`: 해당 변수의 메모리 주소를 반환합니다.
- `*포인터`: 포인터가 가리키는 주소에 있는 값을 읽거나 씁니다.

함수에 값을 전달할 때(Call by Value) Go는 기본적으로 복사본을 생성하므로 원본이 보호됩니다. 그러나 함수 내부에서 원본 변수를 직접 수정해야 하는 경우에는 포인터를 전달(Call by Pointer)하여 `*포인터`를 통해 원본을 조작해야 합니다. 이 패턴은 구조체의 포인터 수신자 메서드에도 동일하게 적용됩니다.

선언만 하고 초기화하지 않은 포인터는 `nil` 상태이며, `nil` 포인터를 역참조(`*`)하면 프로그램이 패닉(Panic) 상태로 종료되므로 반드시 사전 검사가 필요합니다.

### 인터페이스와 덕 타이핑

Go의 인터페이스는 메서드 서명의 집합으로 정의되며, 구조체가 해당 메서드를 모두 갖추면 별도의 선언 없이 자동으로 그 인터페이스의 구현체가 됩니다. 이를 암묵적 구현(Implicit Implementation) 또는 덕 타이핑(Duck Typing)이라 합니다.

이번 실습에서는 `Shape` 인터페이스를 정의하고 `Rectangle`, `Circle`, `Triangle` 세 구조체가 각각 구현하도록 설계하였습니다. 인터페이스를 매개변수 타입으로 사용하면 어떤 구현체든 동일한 함수로 처리할 수 있는 다형성(Polymorphism)이 구현됩니다.

- 인터페이스 슬라이스 `[]Shape`: 서로 다른 구조체 타입을 하나의 컬렉션으로 묶어 일괄 처리
- 타입 단언 `s.(Circle)`: 인터페이스 변수에서 실제 구체 타입을 안전하게 추출
- 빈 인터페이스 `interface{}` / `any`: 모든 타입을 담을 수 있는 범용 컨테이너

---

## 4. 강의 평가 및 교육적 성과

이번 5일차 강의에서는 Go 언어의 타입 시스템 중 가장 핵심적인 구조체, 메서드, 포인터, 인터페이스를 집중적으로 학습하였습니다.

수강생들은 구조체와 메서드를 결합하여 데이터와 행동을 하나의 단위로 묶는 객체지향적 설계를 Go 방식으로 구현하는 경험을 쌓았습니다. 특히 값 수신자와 포인터 수신자의 동작 차이를 코드 실행을 통해 눈으로 확인함으로써, 언제 어떤 수신자를 선택해야 하는지에 대한 실질적인 판단 기준을 습득하였습니다.

포인터 학습에서는 "값 복사"와 "주소 공유"라는 두 개념을 함수 호출 전후의 변수 변화로 직접 비교하여, 메모리 동작 원리를 직관적으로 이해하였습니다. 이는 이후 대용량 구조체 처리나 고루틴 간 데이터 공유 시 올바른 포인터 활용의 토대가 됩니다.

인터페이스 학습에서 가장 두드러진 교육 성과는 "명시적 선언 없는 구현"이라는 Go의 철학을 체득한 것입니다. 수강생들은 `implements` 키워드 없이도 메서드 집합만으로 인터페이스가 자동 충족됨을 실감하고, 이를 통해 확장에 열린 유연한 코드 설계가 가능함을 인식하였습니다.

종합적으로 볼 때, 수강생들은 Go 언어의 타입 설계 철학과 메모리 모델을 실무 수준에서 이해하는 역량을 완성하였으며, 6일차에 다룰 고루틴(Goroutine)과 채널(Channel)을 통한 Go 언어의 동시성(Concurrency) 프로그래밍을 학습하기 위한 준비를 충실히 마쳤습니다.
