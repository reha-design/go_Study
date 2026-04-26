package main

import "fmt"

// 1. 비트 마스킹 (권한 관리)에 사용될 상수 (iota 활용 가능)
const (
	Read    = 1 << 0 // 0001 (1)
	Write   = 1 << 1 // 0010 (2)
	Execute = 1 << 2 // 0100 (4)
)

func main() {
	fmt.Println("--- 실무 활용 1: 비트 마스킹 (권한 제어) ---")
	var myPermission int = 0 // 초기 권한: 없음(0000)

	// 권한 부여 (OR 연산)
	myPermission = myPermission | Read | Write // 0011 (읽기, 쓰기 권한)
	fmt.Printf("현재 권한 설정 (Read|Write): %04b\n", myPermission)

	// 권한 확인 (AND 연산)
	if myPermission&Read != 0 {
		fmt.Println(" -> 읽기 권한이 있습니다.")
	}
	if myPermission&Execute == 0 {
		fmt.Println(" -> 실행 권한이 없습니다.")
	}

	// 특정 권한 회수 (비트 클리어 연산)
	myPermission = myPermission &^ Write // 0001 (쓰기 권한 박탈)
	fmt.Printf("쓰기 권한 회수 후: %04b\n", myPermission)


	fmt.Println("\n--- 실무 활용 2: 색상 데이터 (압축 및 추출) ---")
	// 0x (16진수) AARRGGBB 포맷: 투명도(AA), 빨강(FF), 초록(00), 파랑(55)
	var colorCode uint32 = 0xAAFF0055

	// 시프트 연산과 비트 AND 연산을 이용한 개별 색상 데이터 추출
	alpha := (colorCode >> 24) & 0xFF
	red := (colorCode >> 16) & 0xFF
	green := (colorCode >> 8) & 0xFF
	blue := colorCode & 0xFF

	fmt.Printf("원본 색상 코드: 0x%X\n", colorCode)
	fmt.Printf("Alpha (투명도): 0x%02X (%3d)\n", alpha, alpha)
	fmt.Printf("Red   (빨강): 0x%02X (%3d)\n", red, red)
	fmt.Printf("Green (초록): 0x%02X (%3d)\n", green, green)
	fmt.Printf("Blue  (파랑): 0x%02X (%3d)\n", blue, blue)


	fmt.Println("\n--- 실무 활용 3: 암호화 및 복호화 (XOR 연산) ---")
	var data int = 12345
	var secretKey int = 888 // 대칭키 역할

	// XOR를 이용한 암호화
	encrypted := data ^ secretKey
	fmt.Printf("원본 데이터: %d\n", data)
	fmt.Printf("비밀 키로 암호화한 결과: %d\n", encrypted)

	// XOR를 이용한 복호화 (동일한 키로 다시 XOR)
	decrypted := encrypted ^ secretKey
	fmt.Printf("다시 복호화한 결과: %d\n", decrypted)
}
