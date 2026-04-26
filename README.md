# Go 언어 기초 학습 및 예제 저장소

**Tucker의 Go 언어 프로그래밍** 강의를 바탕으로 실습한 예제 코드와 학습 내용을 정리하는 저장소입니다.

> 📖 **참고 강의**: [Tucker의 Go 언어 프로그래밍](https://www.youtube.com/c/TuckerProgramming) — Tucker(박성재)님이 운영하는 Go 언어 입문 강의 시리즈
> ⚠️ **참고**: 본 저장소의 코드는 개인 학습 과정에서 작성된 것으로, 도서의 원본 예제 코드와 다를 수 있습니다.

## 🚀 실행 방법

각 예제는 Go CLI 도구를 사용하여 실행하거나 빌드할 수 있습니다.

### 1. 즉시 실행 (테스트용)
컴파일과 실행을 한 번에 수행합니다.
```bash
go run <폴더명>/<하위폴더명>/main.go
```
*예시:* `go run 260425/01_hello/main.go`

### 2. 실행 파일 빌드
실행 가능한 바이너리 파일(`.exe` 등)을 생성합니다.
```bash
go build <폴더명>/<하위폴더명>/main.go
```

---

## 📂 학습 기록

### [2026-04-25] Hello World & Variables
- **01_hello**: `260425/01_hello/main.go` - Go의 기본 구조 이해
- **02_variable**: `260425/02_variable/main.go` - 변수와 자료형 실습
- **03_var_variants**: `260425/03_var_variants/main.go` - 다양한 변수 선언 방식 (타입 추론, 단축 선언 등)
- **04_constants**: `260425/04_constants/main.go` - 상수(const) 선언 및 iota 활용
- **05_type_conversion**: `260425/05_type_conversion/main.go` - 명시적 타입 변환 및 strconv 패키지

### [2026-04-26] Floating Point, Precision & Operators
- **06_float_precision**: `260426/06_float_precision/main.go` - 실수의 표현, 지수 표기법, 정밀도 오차
- **07_fmt_specifiers**: `260426/07_fmt_specifiers/main.go` - 다양한 출력 서식 문자(%%v, %%d, %%f 등)
- **08_arithmetic**: `260426/08_arithmetic/main.go` - 산술 연산자 (+, -, *, /, %), 정수/실수 나눗셈, 증감 연산자, 연산자 우선순위
- **09_bitwise**: `260426/09_bitwise/main.go` - 비트 연산자 (&, |, ^, &^, <<, >>) 실습
- **10_bitwise_practical**: `260426/10_bitwise_practical/main.go` - 비트 연산의 실무 활용 (비트 마스킹, 색상 데이터 추출, XOR 암호화)
- **11_comparison**: `260426/11_comparison/main.go` - 비교 연산자 (==, !=, <, >, <=, >=) 및 문자열/실수 비교 주의사항 실습
- **12_float_solution**: `260426/12_float_solution/main.go` - 실수 정밀도 문제 해결법 (정수 변환 및 Epsilon 비교) 실습
- **13_jeff_dean_style**: `260426/13_jeff_dean_style/main.go` - [Jeff Dean Style] 비트 패킹을 활용한 데이터 압축 및 성능 최적화 사고방식 실습
- **14_functions**: `260426/14_functions/`
    - `01_basic/main.go`: 함수의 기본 선언 및 반환값 실습
    - `02_without_func/main.go`: 함수를 안 썼을 때의 불편함 (코드 중복)
    - `03_with_func/main.go`: 함수를 썼을 때의 장점 (재사용성, 유지보수)

---

## 📚 참고 자료

| 자료 | 링크 |
|------|------|
| Tucker의 Go 언어 프로그래밍 (YouTube) | [youtube.com/c/TuckerProgramming](https://www.youtube.com/c/TuckerProgramming) |
| Tucker의 Go 언어 프로그래밍 (도서) | [교보문고 링크](https://product.kyobobook.co.kr/detail/S000001810497) |
| Go 공식 문서 | [go.dev/doc](https://go.dev/doc) |
