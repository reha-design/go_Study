package main

import (
	"fmt"
	"strconv" // 문자열 변환을 위한 패키지
)

func main() {
	// 1. 숫자 간의 명시적 형변환
	// Go에서는 큰 타입에서 작은 타입으로든, 작은 타입에서 큰 타입으로든 반드시 명시해야 합니다.
	var a int = 10
	var b float64 = float64(a) // int -> float64
	var c int64 = int64(a)    // int -> int64
	
	fmt.Printf("1. 숫자 변환: %v(%T), %v(%T)\n", b, b, c, c)

	// 2. 실수 -> 정수 변환 (소수점 버림)
	var f float64 = 10.7
	var i int = int(f) // 10.7 -> 10 (소수점 이하 소실)
	
	fmt.Printf("2. 실수->정수: %v(%T) -> %v(%T)\n", f, f, i, i)

	// 3. 문자열과 숫자 간의 변환 (strconv 패키지 사용)
	// 단순히 string(65)라고 하면 문자 'A'가 나오므로, 숫자 자체를 문자로 바꾸려면 strconv를 써야 합니다.
	
	// 3-1. 숫자를 문자열로 (Itoa: Integer to ASCII)
	var num int = 123
	var str string = strconv.Itoa(num)
	fmt.Printf("3-1. 숫자->문자열: %v(%T)\n", str, str)

	// 3-2. 문자열을 숫자로 (Atoi: ASCII to Integer)
	var str2 string = "456"
	num2, err := strconv.Atoi(str2) // 에러 처리가 함께 발생함
	if err == nil {
		fmt.Printf("3-2. 문자열->숫자: %v(%T)\n", num2, num2)
	}

	// 4. 상수와 타입 변환
	// 앞서 배운 것처럼 '타입이 없는 상수'는 변수에 대입될 때 자동으로 변환되지만,
	// '타입이 있는 변수'끼리는 무조건 명시적 변환이 필요합니다.
	var x int = 10
	var y float64 = 20.5
	// res := x + y // 에러! (int + float64 연산 불가)
	res := float64(x) + y // OK! (타입을 맞춰줘야 연산 가능)
	
	fmt.Println("4. 연산 시 타입 맞춤:", res)
}
