package main

import "fmt"

func main() {
	// [변수 선언의 정석] var 변수이름 타입 = 값
	// 1. var: "컴퓨터야, 나 값을 저장할 바구니(변수) 하나 만들게!"라는 선언입니다.
	// 2. age: 바구니에 붙인 이름입니다. 나중에 이 이름으로 값을 꺼내 써요.
	// 3. int: 바구니의 종류입니다. 'integer'의 약자로 정수(숫자)만 담겠다는 뜻입니다.
	var age int = 20

	// float64: 소수점이 있는 숫자(실수)를 담는 바구니입니다.
	var height float64 = 175.5

	// string: 글자(문자열)를 담는 바구니입니다. 반드시 큰따옴표(" ")로 감싸야 합니다.
	var name string = "홍길동"

	// bool: 'boolean'의 약자로 참(true) 또는 거짓(false) 딱 두 가지만 담습니다.
	var isStudent bool = true

	// fmt.Println: 쉼표(,)를 사용해서 여러 개를 동시에 출력할 수 있습니다.
	fmt.Println("이름:", name)
	fmt.Println("나이:", age)
	fmt.Println("키:", height, "cm")
	fmt.Println("학생 여부:", isStudent)

	fmt.Println("-------------------")

	// [값의 변경]
	// 한 번 만든 바구니는 이름과 종류를 다시 말할 필요가 없습니다. (var, int 생략)
	// '=' 기호는 "같다"가 아니라 "오른쪽에 있는 값을 왼쪽 바구니에 넣어라"는 뜻입니다.
	age = 21
	isStudent = false

	fmt.Println("변경된 나이:", age)
	fmt.Println("변경된 학생 여부:", isStudent)
}
