# Go 강의  보고서
## 1. 강의 목표 및 성취 수준

본 강의는 Go 언어의 기본 실행 구조와 자료형, 변수·상수 선언 방식, 명시적 형변환 개념을 학습하고, 각 개념을 직접 코드로 작성·실행해 보는 것을 목표로 진행하였습니다.

수강생들은 강의와 실습을 통해 다음 내용을 수행할 수 있도록 학습하였습니다.

1. Go 프로그램의 기본 구조인 `package main`, `import`, `func main()`의 역할을 이해하고 실행 가능한 프로그램을 작성할 수 있다.
2. Go의 기본 자료형인 `int`, `float64`, `string`, `bool`의 특징을 이해하고 목적에 맞게 변수를 선언할 수 있다.
3. 변수와 상수의 차이를 이해하고, `const` 및 `iota`를 활용하여 고정값과 열거형 형태의 데이터를 정의할 수 있다.
4. Go의 강타입 특성을 이해하고, 명시적 형변환과 `strconv` 패키지를 활용하여 문자열과 숫자 간 변환을 처리할 수 있다.
5. `strconv.Atoi` 사용 시 발생할 수 있는 변환 오류를 확인하고, Go의 기본적인 에러 처리 흐름을 이해할 수 있다.

---

## 2. 실습 및 제출 산출물 현황

강의 과정에서는 각 주제별 실습 코드를 작성하고 실행 결과를 확인하였습니다. 수강생들은 아래 구조의 예제 코드를 직접 작성하며 Go 프로그램의 기본 문법과 실행 흐름을 검증하였습니다.

```text
day01/
├── 01_hello/
│   └── main.go           (패키지 선언, 외부 패키지 임포트, 표준 출력 실습)
├── 02_variable/
│   └── main.go           (기본 타입 변수 선언 및 값 재대입 실습)
├── 03_var_variants/
│   └── main.go           (타입 추론, 단축 선언, 그룹 선언, 제로 값 검증)
├── 04_constants/
│   └── main.go           (상수 선언, iota 열거형 구현, 타입 없는 상수 유연성 검증)
└── 05_type_conversion/
    └── main.go           (명시적 형변환, strconv 패키지 사용, Atoi 에러 핸들링)
```

### 각 챕터별 상세 실습 소스코드

#### ① `01_hello/main.go`

```go
// package main: 이 파일이 실행 가능한 프로그램의 시작점임을 선언합니다.
// Go 언어에서 실행 파일을 만들려면 반드시 package main이라는 선언이 필요합니다.
package main

// import "fmt": 표준 라이브러리인 fmt(format의 약자) 패키지를 가져옵니다.
// fmt 패키지는 콘솔에 텍스트를 출력하거나 입력을 받는 등의 포맷팅 기능을 제공합니다.
import "fmt"

// main 함수: 프로그램이 실행될 때 가장 먼저 호출되는 함수입니다.
// 프로그램의 '입구(Entry Point)' 역할을 수행합니다.
func main() {
	// fmt.Println: 콘솔(터미널)에 괄호 안의 내용을 출력하고 줄바꿈을 합니다.
	// 여기서는 "Hello World"라는 문자열을 화면에 표시합니다.
	fmt.Println("Hello World")
}
```

#### ② `02_variable/main.go`

```go
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
```

#### ③ `03_var_variants/main.go`

```go
package main

import "fmt"

func main() {
	// 1. 타입 추론 (Type Inference)
	// 타입을 명시하지 않아도 대입되는 값을 보고 Go가 자동으로 타입을 결정합니다.
	var name = "이순신" // string으로 추론
	var age = 30      // int로 추론
	var score = 95.5  // float64로 추론

	// %v는 값(value), %T는 타입(Type)을 의미합니다.
	fmt.Println("1. 타입 추론:", name, age, score)
	// Printf는 서식(%T 등)을 미리 정해놓고 출력할 때 사용합니다.
	// 문자열 안의 %T 자리에 뒤에 나열한 변수들이 순서대로 하나씩 대입됩니다.
	fmt.Printf("   (각 변수의 타입: %T, %T, %T)\n", name, age, score)

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
```

#### ④ `04_constants/main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("=== 1. 변수(var)와 상수(const)의 차이점 ===")
	// [변수 vs 상수]
	// 변수(var): 값을 언제든지 바꿀 수 있는 그릇 (예: 나이, 오늘의 기온)
	// 상수(const): 한 번 정하면 절대 바꿀 수 없는 고정된 값 (예: 파이(3.14), 생년월일)
	var changeableVar = "나는 변수야 (변경 가능)"
	const unchangeableConst = "나는 상수야 (변경 불가)"

	fmt.Println("초기 상태 ->", changeableVar, "|", unchangeableConst)

	// 변수는 값을 덮어쓸 수 있습니다.
	changeableVar = "변수 값 변경 완료!"
	fmt.Println("변경 후   ->", changeableVar, "|", unchangeableConst)

	// 상수는 값 변경을 시도하면 컴파일 에러가 발생합니다.
	// 실수로 값이 바뀌면 안 되는 중요한 데이터를 지키기 위해 사용합니다.
	// unchangeableConst = "상수 값 변경 시도!"

	fmt.Println("\n=== 2. 그룹 상수 선언 ===")
	// 괄호를 사용해 연관된 상수를 한 번에 묶어서 선언할 수 있어 코드가 깔끔해집니다.
	const (
		Sunday  = 0
		Monday  = 1
		Tuesday = 2
	)
	fmt.Println("요일별 번호:", Sunday, Monday, Tuesday)

	fmt.Println("\n=== 3. iota 키워드 (자동 번호표) ===")
	// iota는 '아이오타'라고 읽으며, 0부터 시작해서 아래로 갈수록 1씩 자동으로 늘어납니다.
	// 일일이 0, 1, 2... 적기 번거로울 때 자동으로 번호를 부여할 수 있습니다.
	const (
		Apple  = iota // 0이 자동으로 들어갑니다.
		Banana        // 1이 자동으로 들어갑니다.
		Grape         // 2가 자동으로 들어갑니다.
	)
	fmt.Printf("자동 번호 - 사과: %d, 바나나: %d, 포도: %d\n", Apple, Banana, Grape)

	fmt.Println("\n=== 4. 타입이 없는 상수의 유연함 ===")
	// 타입을 정하지 않은 상수는 대입되는 변수의 타입에 맞춰 유연하게 사용될 수 있습니다.
	const maxCount = 100

	var currentCount int = maxCount     // 정수(int) 변수에 들어가면 정수가 됨
	var currentPrice float64 = maxCount // 실수(float64) 변수에 들어가면 실수가 됨

	fmt.Printf("정수 변수에 담긴 100: %v (타입: %T)\n", currentCount, currentCount)
	fmt.Printf("실수 변수에 담긴 100: %v (타입: %T)\n", currentPrice, currentPrice)
}
```

#### ⑤ `05_type_conversion/main.go`

```go
package main

import (
	"fmt"
	"strconv" // string conversion의 약자입니다. 문자열 변환을 도와주는 도구상자입니다.
)

func main() {
	// [1. 숫자 간의 형변환]
	// Go는 타입에 대해 매우 엄격합니다. 정수와 실수는 서로 다른 타입으로 취급됩니다.
	// 따라서 타입이 다른 값을 대입하거나 연산하려면 명시적으로 변환해야 합니다.
	var a int = 10
	var b float64 = float64(a) // a를 실수로 변환해서 b에 넣습니다.
	var c int64 = int64(a)     // a를 int64로 변환해서 c에 넣습니다.

	fmt.Println("1. 숫자 변환 결과:", b, c)
	fmt.Printf("   (각 변수의 타입 확인: %T, %T)\n", b, c)

	// [2. 실수 -> 정수 변환 시 주의사항]
	// 실수를 정수로 바꾸면 소수점 뒷부분은 반올림되지 않고 버려집니다.
	var f float64 = 10.7
	var i int = int(f) // 10.7에서 .7이 사라지고 10만 남습니다.

	fmt.Println("2. 실수 -> 정수 변환 (소수점 버림):", f, "->", i)

	// [3. 문자열과 숫자 간의 변환]
	// 문자열 "123"과 숫자 123은 서로 다른 값입니다.
	// 문자열과 숫자 사이의 변환은 strconv 패키지를 사용해야 합니다.

	// 3-1. 숫자를 문자열로 (Itoa: Integer to ASCII)
	var num int = 123
	var str string = strconv.Itoa(num) // 숫자 123을 글자 "123"으로 만듭니다.
	fmt.Println("3-1. 숫자 -> 문자열 변환:", str)
	fmt.Printf("     (타입 확인: %T)\n", str)

	// 3-2. 문자열을 숫자로 (Atoi: ASCII to Integer)
	var str2 string = "456"
	// Atoi는 변환 결과와 함께 에러 값을 반환합니다.
	// 숫자로 바꿀 수 없는 문자열이 들어오면 err에 에러 정보가 담깁니다.
	num2, err := strconv.Atoi(str2)

	if err != nil {
		fmt.Println("3-2. 문자열 -> 숫자 변환 실패:", err)
		return
	}

	fmt.Println("3-2. 문자열 -> 숫자 변환 성공:", num2)
	fmt.Printf("     (타입 확인: %T)\n", num2)

	// [4. 타입이 다른 변수끼리의 연산]
	// Go에서는 타입이 다르면 바로 연산할 수 없습니다.
	// 반드시 한쪽 타입을 다른 쪽에 맞춰줘야 합니다.
	var x int = 10
	var y float64 = 20.5

	// 정수인 x를 float64로 변환해야 y와 계산할 수 있습니다.
	res := float64(x) + y

	fmt.Println("4. 타입이 다른 연산 결과(int+float):", res)
}
```

---

## 3. 핵심 교육 내용 요약

### Go 프로그램의 기본 구조

Go 프로그램이 실행되기 위해 필요한 기본 구성 요소를 학습하였습니다. 실행 가능한 Go 파일은 `package main`에 속해야 하며, 프로그램의 시작점으로 `func main()` 함수가 필요하다는 점을 확인하였습니다.

또한 `fmt` 패키지를 활용하여 `Println`, `Printf` 등의 표준 출력 함수를 사용해 보고, 콘솔에 값을 출력하는 기본 방식을 익혔습니다.

### 변수와 기본 자료형

Go의 기본 자료형인 `int`, `float64`, `string`, `bool`을 중심으로 변수 선언 방식을 실습하였습니다. `var 변수명 타입 = 값` 형태의 기본 선언 방식부터 타입 추론, 단축 변수 선언(`:=`), 복수 변수 선언, 그룹 선언까지 다양한 선언 방식을 확인하였습니다.

초기값을 지정하지 않은 변수에는 타입별 기본값이 자동으로 할당된다는 점도 실습을 통해 확인하였습니다.

- `int`: `0`
- `float64`: `0`
- `string`: `""`
- `bool`: `false`

### 상수와 `iota`

변수와 상수의 차이를 비교하며, 변경 가능한 값과 변경 불가능한 값을 구분하는 방법을 학습하였습니다. `const`를 사용하여 고정값을 선언하고, 상수 블록 안에서 `iota`를 활용해 자동으로 증가하는 값을 정의하는 방식을 실습하였습니다.

이를 통해 요일, 상태값, 등급과 같이 순차적인 값을 다룰 때 `iota`가 코드의 반복을 줄이고 실수를 예방하는 데 유용하다는 점을 확인하였습니다.

### 명시적 형변환과 문자열 변환

Go는 타입이 다른 값 사이의 자동 변환을 허용하지 않는 강타입 언어임을 학습하였습니다. 따라서 정수와 실수 간 연산 또는 대입이 필요한 경우 `float64(value)`, `int(value)`와 같은 명시적 형변환이 필요함을 실습하였습니다.

또한 문자열과 숫자 간 변환은 단순 형변환으로 처리할 수 없으며, `strconv` 패키지를 사용해야 함을 확인하였습니다.

- `strconv.Itoa`: 정수를 문자열로 변환
- `strconv.Atoi`: 문자열을 정수로 변환
- `err` 확인: 변환 실패 가능성에 대한 기본 에러 처리

---

## 4. 강의 평가 및 교육적 성과

이번 1일차 강의에서는 Go 언어의 가장 기본이 되는 실행 구조, 변수, 자료형, 상수, 형변환 개념을 중심으로 학습을 진행하였습니다.

수강생들은 단순히 문법을 암기하는 방식이 아니라, 각 개념을 직접 코드로 작성하고 실행 결과를 확인하면서 Go 언어의 동작 방식을 이해하였습니다. 특히 정적 타입 언어의 엄격한 규칙, 단축 선언의 사용 범위, 타입이 다른 값 사이의 연산 제한 등 초반에 혼동하기 쉬운 부분을 실습을 통해 명확히 정리하였습니다.

또한 `strconv.Atoi`의 반환값과 에러 처리 구조를 살펴보며, Go에서 함수가 여러 값을 반환하는 방식과 기본적인 에러 처리 흐름을 경험하였습니다.

종합적으로 볼 때, 수강생들은 Go 프로그래밍을 시작하기 위한 핵심 기초 문법과 실행 흐름을 성공적으로 학습하였으며, 이후 조건문, 반복문, 함수, 배열 및 슬라이스 등 후속 학습을 진행하기 위한 기반을 충분히 마련하였습니다.
