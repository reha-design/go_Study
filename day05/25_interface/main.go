package main

import "fmt"

// [인터페이스(Interface): 메서드의 이름과 서명(시그니처)만 정의하는 약속]
// 인터페이스는 "이 메서드들을 가지고 있으면 이 타입으로 취급한다"는 규약입니다.
// Go의 인터페이스는 명시적 선언 없이 자동으로 구현됩니다. (Duck Typing)

// [인터페이스 정의]
// type 이름 interface { 메서드이름() 반환타입 }
type Shape interface {
	Area() float64
	Perimeter() float64
	Describe() string
}

// [Rectangle: Shape 인터페이스를 구현하는 구조체]
// "implements Shape" 같은 키워드를 쓰지 않아도,
// Area(), Perimeter(), Describe() 메서드가 있으면 자동으로 Shape 구현체가 됩니다.
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) Describe() string {
	return fmt.Sprintf("사각형(%.1f x %.1f)", r.Width, r.Height)
}

// [Circle: Shape 인터페이스를 구현하는 또 다른 구조체]
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }
func (c Circle) Describe() string {
	return fmt.Sprintf("원(반지름: %.1f)", c.Radius)
}

// [Triangle: Shape 인터페이스를 구현하는 세 번째 구조체]
type Triangle struct {
	Base, Height, Side1, Side2, Side3 float64
}

func (t Triangle) Area() float64      { return 0.5 * t.Base * t.Height }
func (t Triangle) Perimeter() float64 { return t.Side1 + t.Side2 + t.Side3 }
func (t Triangle) Describe() string {
	return fmt.Sprintf("삼각형(밑변: %.1f, 높이: %.1f)", t.Base, t.Height)
}

// [인터페이스를 매개변수로 사용하는 함수]
// Shape 인터페이스를 받으므로 Rectangle, Circle, Triangle 모두 전달 가능합니다.
// 이것이 다형성(Polymorphism)입니다.
func printShapeInfo(s Shape) {
	fmt.Printf("   도형: %-25s 넓이: %8.2f  둘레: %8.2f\n",
		s.Describe(), s.Area(), s.Perimeter())
}

func main() {
	// [1. 인터페이스를 통한 다형성]
	fmt.Println("--- 1. 인터페이스 다형성 ---")
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}
	tri := Triangle{Base: 6, Height: 4, Side1: 5, Side2: 5, Side3: 6}

	printShapeInfo(rect) // Rectangle이지만 Shape로 취급
	printShapeInfo(circ) // Circle이지만 Shape로 취급
	printShapeInfo(tri)  // Triangle이지만 Shape로 취급

	// [2. 인터페이스 슬라이스: 여러 타입을 한 묶음으로 관리]
	fmt.Println("\n--- 2. 인터페이스 슬라이스로 일괄 처리 ---")
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 5},
		Triangle{Base: 8, Height: 6, Side1: 7, Side2: 7, Side3: 8},
		Rectangle{Width: 2, Height: 9},
	}

	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
		printShapeInfo(s)
	}
	fmt.Printf("   %-25s 합계: %8.2f\n", "전체 도형 넓이", totalArea)

	// [3. 타입 단언 (Type Assertion): 인터페이스에서 원래 타입 꺼내기]
	fmt.Println("\n--- 3. 타입 단언 (Type Assertion) ---")
	var s Shape = Circle{Radius: 3}

	// s.(Circle): s가 실제로 Circle 타입인지 확인하고 꺼냅니다.
	if c, ok := s.(Circle); ok {
		fmt.Printf("   Circle 타입 확인! 반지름: %.1f\n", c.Radius)
	}
	if _, ok := s.(Rectangle); !ok {
		fmt.Println("   Rectangle 타입이 아닙니다.")
	}

	// [4. 빈 인터페이스 (Empty Interface): 모든 타입을 담을 수 있는 컨테이너]
	// interface{} 또는 any (Go 1.18 이상)
	fmt.Println("\n--- 4. 빈 인터페이스 (any) ---")
	var anything interface{}
	anything = 42
	fmt.Printf("   int 저장: %v (타입: %T)\n", anything, anything)
	anything = "hello"
	fmt.Printf("   string 저장: %v (타입: %T)\n", anything, anything)
	anything = Circle{Radius: 2}
	fmt.Printf("   Circle 저장: %v (타입: %T)\n", anything, anything)
}
