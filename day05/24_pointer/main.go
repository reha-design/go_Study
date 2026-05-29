package main

import "fmt"

// [포인터(Pointer): 변수의 메모리 주소를 저장하는 변수]
// 포인터를 사용하면 함수 밖에 있는 변수의 값을 함수 안에서 직접 변경할 수 있습니다.
// &변수명: 변수의 메모리 주소를 가져옵니다. (주소 연산자)
// *포인터변수: 포인터가 가리키는 주소의 값을 읽거나 씁니다. (역참조 연산자)

// [값으로 전달 (Call by Value): 원본이 변경되지 않음]
func doubleByValue(n int) {
	n = n * 2
	fmt.Printf("   함수 내부: n = %d\n", n)
}

// [포인터로 전달 (Call by Pointer): 원본을 직접 변경]
func doubleByPointer(n *int) {
	*n = *n * 2 // *n: 포인터가 가리키는 주소의 값을 변경
	fmt.Printf("   함수 내부: *n = %d\n", *n)
}

func main() {
	// [1. 포인터 기본 개념]
	fmt.Println("--- 1. 포인터 기본 개념 ---")
	x := 42
	p := &x // p는 x의 메모리 주소를 저장하는 포인터

	fmt.Printf("   x의 값: %d\n", x)
	fmt.Printf("   x의 주소(&x): %p\n", &x)
	fmt.Printf("   p의 값(주소): %p\n", p)
	fmt.Printf("   p가 가리키는 값(*p): %d\n", *p)

	// [2. 포인터로 값 변경]
	fmt.Println("\n--- 2. 포인터로 원본 값 변경 ---")
	fmt.Printf("   변경 전: x = %d\n", x)
	*p = 100 // p가 가리키는 주소(= x)의 값을 100으로 변경
	fmt.Printf("   *p = 100 대입 후: x = %d\n", x) // x도 바뀜!

	// [3. 값 전달 vs 포인터 전달 비교]
	fmt.Println("\n--- 3. 값 전달 vs 포인터 전달 ---")
	a := 10

	fmt.Println("   [값 전달] doubleByValue:")
	fmt.Printf("   함수 호출 전: a = %d\n", a)
	doubleByValue(a)
	fmt.Printf("   함수 호출 후: a = %d (원본 변경 안 됨!)\n", a)

	fmt.Println("\n   [포인터 전달] doubleByPointer:")
	fmt.Printf("   함수 호출 전: a = %d\n", a)
	doubleByPointer(&a) // &a: a의 주소를 전달
	fmt.Printf("   함수 호출 후: a = %d (원본 변경됨!)\n", a)

	// [4. new 함수로 포인터 생성]
	// new(타입): 해당 타입의 제로값으로 초기화된 메모리를 할당하고 포인터를 반환
	fmt.Println("\n--- 4. new 함수 ---")
	np := new(int) // *int 타입의 포인터, 초기값 0
	fmt.Printf("   new(int) 초기값: %d\n", *np)
	*np = 77
	fmt.Printf("   *np = 77 대입 후: %d\n", *np)

	// [5. nil 포인터: 아무 주소도 가리키지 않는 포인터]
	fmt.Println("\n--- 5. nil 포인터 확인 ---")
	var nilPtr *int // 선언만 하면 nil (아무것도 가리키지 않음)
	fmt.Printf("   nilPtr 값: %v\n", nilPtr)
	if nilPtr == nil {
		fmt.Println("   nilPtr은 nil입니다. 역참조(*nilPtr)하면 프로그램이 패닉(Panic)합니다!")
	}
}
