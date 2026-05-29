package main

import "fmt"

// [구조체 메서드(Method): 구조체에 기능(함수)을 붙이는 방법]
// Go는 클래스가 없지만, 구조체에 메서드를 정의해 비슷한 효과를 냅니다.
// func (수신자변수 구조체타입) 메서드이름(매개변수) 반환타입 { ... }

type Rectangle struct {
	Width  float64
	Height float64
}

// [값 수신자(Value Receiver): 구조체의 복사본을 받아 처리]
// 구조체의 데이터를 읽기만 할 때 사용합니다. 원본이 변경되지 않습니다.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 출력용 메서드: 구조체의 정보를 문자열로 표현
func (r Rectangle) String() string {
	return fmt.Sprintf("사각형(너비: %.1f, 높이: %.1f)", r.Width, r.Height)
}

// [포인터 수신자(Pointer Receiver): 구조체의 원본 주소를 받아 처리]
// 구조체의 필드 값을 직접 변경해야 할 때 사용합니다.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor  // 원본 구조체의 Width를 직접 변경
	r.Height *= factor // 원본 구조체의 Height를 직접 변경
}

// 원 구조체
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	// [1. 값 수신자 메서드 호출]
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("1. 구조체 메서드 호출:")
	fmt.Println("  ", rect.String())
	fmt.Printf("   넓이: %.2f\n", rect.Area())
	fmt.Printf("   둘레: %.2f\n", rect.Perimeter())

	// [2. 포인터 수신자 메서드로 원본 값 변경]
	fmt.Println("\n2. 포인터 수신자로 크기 변경 (2배):")
	fmt.Printf("   변경 전: %s\n", rect.String())
	rect.Scale(2.0) // 내부적으로 &rect를 전달 (Go가 자동 처리)
	fmt.Printf("   변경 후: %s\n", rect.String())

	// [3. 여러 구조체에 같은 이름의 메서드 정의]
	fmt.Println("\n3. 사각형 vs 원 넓이 비교:")
	r := Rectangle{Width: 7, Height: 3}
	c := Circle{Radius: 5}
	fmt.Printf("   사각형 넓이: %.2f\n", r.Area())
	fmt.Printf("   원 넓이: %.2f\n", c.Area())

	// [4. 구조체 메서드 체이닝 개념]
	fmt.Println("\n4. 구조체와 메서드의 활용:")
	shapes := []Rectangle{
		{Width: 4, Height: 6},
		{Width: 10, Height: 2},
		{Width: 3, Height: 3},
	}
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
		fmt.Printf("   %s → 넓이: %.2f\n", s.String(), s.Area())
	}
	fmt.Printf("   총 넓이 합계: %.2f\n", totalArea)
}
