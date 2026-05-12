package main

import (
	"fmt"
	"strconv" // string conversion의 약자입니다. 문자열 변환을 도와주는 도구상자에요!
)

func main() {
	// [1. 숫자 간의 형변환]
	// Go는 타입에 대해 매우 깐깐합니다. '정수'와 '실수'는 엄연히 다른 존재라고 생각해요.
	// 그래서 작은 그릇(int)에서 큰 그릇(float64)으로 옮길 때도 "내가 책임지고 바꾸겠다!"라고 명시해야 합니다.
	var a int = 10
	var b float64 = float64(a) // a를 실수로 변신시켜서 b에 넣습니다.
	var c int64 = int64(a)     // a를 더 큰 정수(int64)로 변신시켜서 c에 넣습니다.
	
	fmt.Println("1. 숫자 변환 결과:", b, c)
	fmt.Printf("   (각 변수의 타입 확인: %T, %T)\n", b, c)

	// [2. 실수 -> 정수 변환 시 주의사항]
	// 실수를 정수로 바꾸면 소수점 뒷부분은 '반올림'이 아니라 그냥 '싹둑' 잘려 나갑니다.
	var f float64 = 10.7
	var i int = int(f) // 10.7에서 .7이 사라지고 10만 남습니다.
	
	fmt.Println("2. 실수 -> 정수 변환 (소수점 버림):", f, "->", i)

	// [3. 문자열과 숫자 간의 변환]
	// 문자열 "123"과 숫자 123은 다릅니다. "123"은 그냥 글자 그림일 뿐이에요.
	// 그래서 단순히 float64(str) 처럼은 바꿀 수 없고, strconv라는 전문 도구를 써야 합니다.
	
	// 3-1. 숫자를 문자열로 (Itoa: Integer to ASCII)
	var num int = 123
	var str string = strconv.Itoa(num) // 숫자 123을 글자 "123"으로 만듭니다.
	fmt.Println("3-1. 숫자 -> 문자열 변환:", str)
	fmt.Printf("     (타입 확인: %T)\n", str)

	// 3-2. 문자열을 숫자로 (Atoi: ASCII to Integer)
	var str2 string = "456"
	// Atoi는 결과물(num2)과 함께 '에러(err)'라는 것도 같이 줍니다.
	// 글자 "abc"를 숫자로 바꾸라고 하면 에러가 나기 때문이죠!
	num2, err := strconv.Atoi(str2) 
	
	if err == nil { // err가 'nil(비어있음)'이라면 에러 없이 성공했다는 뜻입니다.
		fmt.Println("3-2. 문자열 -> 숫자 변환 성공:", num2)
		fmt.Printf("     (타입 확인: %T)\n", num2)
	}

	// [4. 타입이 다른 변수끼리의 연산]
	// Go에서는 타입이 다르면 아예 더하기(+)조차 안 됩니다.
	// 반드시 한쪽 타입을 다른 쪽에 맞춰줘야 합니다.
	var x int = 10
	var y float64 = 20.5
	
	// 정수인 x를 실수인 float64(x)로 잠시 변신시켜야 y와 계산할 수 있습니다.
	res := float64(x) + y 
	
	fmt.Println("4. 타입이 다른 연산 결과(int+float):", res)
}
