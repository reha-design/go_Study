# Go 언어 기초 학습 및 예제 저장소

비전공자 및 입문자를 위한 **Go 언어 7일 완성 마스터 플랜** 저장소입니다.
초심자가 이해하기 쉽도록 핵심 내용을 재구성하고 모든 예제에 상세한 주석을 추가했습니다.

> 💡 **학습 목표**: 7일 안에 Go의 기초 문법부터 고루틴(동시성)까지 마스터하여 실전 활용 역량을 갖춥니다.

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

### [Day 01] 기본 구조와 변수, 상수, 형변환
- **01_hello**: `day01/01_hello/main.go` - Go의 기본 구조 이해 및 출력
- **02_variable**: `day01/02_variable/main.go` - 변수와 기본 자료형 실습
- **03_var_variants**: `day01/03_var_variants/main.go` - 다양한 변수 선언 방식 (타입 추론, 단축 선언)
- **04_constants**: `day01/04_constants/main.go` - 상수(const) 선언 및 `iota` 활용
- **05_type_conversion**: `day01/05_type_conversion/main.go` - 명시적 타입 변환 및 `strconv` 패키지

### [Day 02] 연산자 포맷팅 및 함수 기초
- **06_float_precision**: `day02/06_float_precision/main.go` - 실수의 표현 및 정밀도 오차 극복
- **07_fmt_specifiers**: `day02/07_fmt_specifiers/main.go` - 출력 서식 문자(Printf) 및 정렬
- **08_arithmetic**: `day02/08_arithmetic/main.go` - 산술 연산자 및 우선순위
- **09_bitwise**: `day02/09_bitwise/main.go` - 비트 연산자 기초
- **10_comparison**: `day02/10_comparison/main.go` - 비교 연산자 및 실수 비교 주의사항
- **11_functions**: `day02/11_functions/` - 함수의 기본 선언, 반환값, 유지보수 측면의 함수 활용 장점

### [Day 03] 조건 제어 및 순환 알고리즘
- **12_if_conditional**: `day03/12_if_conditional/main.go` - if-else 분기 및 if 초기화 문법
- **13_logical_operators**: `day03/13_logical_operators/main.go` - 논리 연산자(&&, ||, !) 활용
- **14_switch_conditional**: `day03/14_switch_conditional/main.go` - switch 매칭 및 fallthrough
- **15_for_loop**: `day03/15_for_loop/main.go` - for 기본 루프, break/continue 분기 제어
- **16_nested_loop**: `day03/16_nested_loop/main.go` - 이중 중첩 반복문 (구구단 출력 실습)

### [Day 04] 배열, 슬라이스, 데이터 맵
- **17_array**: `day04/17_array/main.go` - 배열의 선언, 초기화, 인덱스 접근 및 순회
- **18_slice**: `day04/18_slice/main.go` - 슬라이스 선언, append 동작, 슬라이싱, len/cap의 이해
- **19_slice_ops**: `day04/19_slice_ops/main.go` - 얕은 복사와 깊은 복사, 특정 원소 삭제, 2D 슬라이스
- **20_map**: `day04/20_map/main.go` - 맵 선언 및 조회, ok 패턴을 통한 안전한 존재 여부 확인
- **21_map_ops**: `day04/21_map_ops/main.go` - 단어 빈도수 계산, 중첩 맵, 맵과 슬라이스 조합 활용

### [Day 05] 객체지향 설계와 포인터, 다형성
- **22_struct**: `day05/22_struct/main.go` - 구조체 기초 선언, 다양한 초기화, 중첩 구조체
- **23_struct_methods**: `day05/23_struct_methods/main.go` - 구조체 메서드, 값 수신자 vs 포인터 수신자
- **24_pointer**: `day05/24_pointer/main.go` - 포인터 주소 연산(&) 및 역참조(*), 값 전달 vs 포인터 전달
- **25_interface**: `day05/25_interface/main.go` - 인터페이스를 활용한 덕 타이핑, 다형성 구현 및 타입 단언

### [Day 06] 동시성 프로그래밍 (Goroutine & Channel)
- **26_goroutine**: `day06/26_goroutine/main.go` - 고루틴 생성, 비동기 실행 및 `sync.WaitGroup` 동기화
- **27_channel**: `day06/27_channel/main.go` - 버퍼드/언버퍼드 채널 통신, 단방향 채널 제어
- **28_channel_sync**: `day06/28_channel_sync/main.go` - 채널 닫기(close)와 `for range` 채널 순회
- **29_select**: `day06/29_select/main.go` - `select`를 이용한 다중 채널 모니터링 및 타임아웃(Timeout) 처리
- **30_mutex**: `day06/30_mutex/main.go` - 경쟁 상태(Race Condition) 방지를 위한 `sync.Mutex` 임계 구역 보호

### [Day 07] 패키지 모듈 관리 및 실전 미니 프로젝트 (준비 중)
- **31_package_module**: `day07/31_package_module/` - 모듈 생성 및 패키지 분리 가이드
- **32_error_handling**: `day07/32_error_handling/` - 에러 처리 및 defer/panic/recover
- **33_cli_project**: `day07/33_cli_project/` - [미니 프로젝트] CLI 도구 만들기 종합 실습

---

## 🚀 Next Step (기초 이후의 단계)

본 기초 스터디(7일 과정)를 마친 후에는 Go 언어의 실무 활용을 위해 **웹 프레임워크 학습**으로 넘어가는 것을 추천합니다. 

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
