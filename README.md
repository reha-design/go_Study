# Go 언어 기초 학습 및 예제 저장소

Go 언어를 학습하며 작성한 예제 코드와 학습 내용을 정리하는 저장소입니다.

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
