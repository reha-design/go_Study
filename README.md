# Go 언어 기초 학습 및 예제 저장소

비전공자 및 입문자를 위한 **Go 언어 5일 완성 마스터 플랜** 저장소입니다.
기존 **Tucker의 Go 언어 프로그래밍** 강의를 기반으로 하되, 초심자가 이해하기 쉽도록 핵심 내용을 재구성하고 모든 예제에 상세한 주석을 추가했습니다.

> 💡 **학습 목표**: 5일 안에 Go의 기초 문법부터 고루틴(동시성)까지 마스터하여 실전 CLI 도구를 제작합니다.

## 🚀 실행 방법

각 예제는 Go CLI 도구를 사용하여 실행하거나 빌드할 수 있습니다.

### 1. 즉시 실행 (테스트용)
컴파일과 실행을 한 번에 수행합니다.
```bash
go run <폴더명>/<하위폴더명>/main.go
```
*예시:* `go run day01/01_hello/main.go`

### 2. 실행 파일 빌드
실행 가능한 바이너리 파일(`.exe` 등)을 생성합니다.
```bash
go build <폴더명>/<하위폴더명>/main.go
```

---

## 📂 학습 기록

### [Day 01] Hello World & Variables
- **01_hello**: `day01/01_hello/main.go` - Go의 기본 구조 이해
- **02_variable**: `day01/02_variable/main.go` - 변수와 자료형 실습
- **03_var_variants**: `day01/03_var_variants/main.go` - 다양한 변수 선언 방식 (타입 추론, 단축 선언 등)
- **04_constants**: `day01/04_constants/main.go` - 상수(const) 선언 및 iota 활용
- **05_type_conversion**: `day01/05_type_conversion/main.go` - 명시적 타입 변환 및 strconv 패키지

### [Day 02] Precision, Operators & Functions
- **06_float_precision**: `day02/06_float_precision/main.go` - 실수의 표현 및 정밀도 오차
- **07_fmt_specifiers**: `day02/07_fmt_specifiers/main.go` - 출력 서식 문자(Printf) 활용
- **08_arithmetic**: `day02/08_arithmetic/main.go` - 산술 연산자 및 우선순위
- **09_bitwise**: `day02/09_bitwise/main.go` - 비트 연산자 기초 (가볍게 읽고 넘어가기)
- **10_comparison**: `day02/10_comparison/main.go` - 비교 연산자 주의사항
- **11_functions**: `day02/11_functions/`
    - `01_basic/main.go`: 함수의 기본 선언 및 반환값
    - `02_without_func/main.go`: 함수가 필요한 이유
    - `03_with_func/main.go`: 함수의 장점

### [Day 03] Control Flow & Data Structures (준비 중)
- **12_control_flow**: `day03/12_control_flow/` - 조건문(if, switch) 및 반복문(for) 기초
- **13_array_slice**: `day03/13_array_slice/` - 배열과 슬라이스 활용 및 **range를 이용한 순회**
- **14_map**: `day03/14_map/` - 키-값 쌍을 저장하는 맵(Map) 자료구조와 **데이터 존재 여부 확인(ok 필터)**

### [Day 04] Object-Oriented (준비 중)
- **15_pointers**: `day04/15_pointers/` - 메모리 주소와 포인터 기초 및 **문자열과 룬(rune)의 차이**
- **16_structs_methods**: `day04/16_structs_methods/` - 구조체와 메서드, **대문자/소문자 노출 규칙(Visibility)**
- **17_interfaces**: `day04/17_interfaces/` - 인터페이스(Interface)와 다형성

### [Day 05] Concurrency & Practical App (준비 중)
- **18_error_handling**: `day05/18_error_handling/` - 에러 처리 및 defer/panic/recover
- **19_goroutines**: `day05/19_goroutines/` - 고루틴(Goroutine)을 이용한 경량 스레드 기초
- **20_channels**: `day05/20_channels/` - 채널(Channel)을 통한 고루틴 간 데이터 통신
- **21_cli_project**: `day05/21_cli_project/` - [미니 프로젝트] 할 일 관리(Todo) CLI 도구 만들기

---

## 🚀 Next Step (기초 이후의 단계)

본 기초 스터디(5일 과정)를 마친 후에는 Go 언어의 실무 활용을 위해 **웹 프레임워크 학습**으로 넘어가는 것을 추천합니다. 

- **대상 프레임워크**: [Gin](https://github.com/gin-gonic/gin) 또는 [Fiber](https://github.com/gofiber/fiber)
- **학습 방법**: 별도의 레포지토리(예: `go_Web_Project`)를 생성하여 실습 진행
- **핵심 학습 내용**: REST API 설계, 미들웨어 활용, 데이터베이스(GORM 등) 연동

---

## 📚 참고 자료

| 자료 | 링크 |
|------|------|
| Go 공식 문서 (Overview) | [go.dev/doc](https://go.dev/doc) |
| A Tour of Go (한글/영문) | [tour.go.dev](https://tour.go.dev/) |
| Effective Go (Go다운 코드 작성법) | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) |
| Go 언어 명세 (Language Spec) | [go.dev/ref/spec](https://go.dev/ref/spec) |
| Go 표준 라이브러리 (Package) | [pkg.go.dev/std](https://pkg.go.dev/std) |
