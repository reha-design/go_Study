package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("--- [Jeff Dean Style] 데이터 압축 사고방식 ---")

	// 상황: 게임 서버에서 수천 명의 유저 정보를 전송해야 함.
	// 1. 유저 ID (0~1023) -> 10비트 필요
	// 2. 레벨 (0~255) -> 8비트 필요
	// 3. 상태 (살았나/죽었나/무적 등 4가지) -> 2비트 필요
	// 총 20비트면 충분한데, 일반적인 구조체로 만들면 최소 8~16바이트(64~128비트)를 씁니다.

	userID := uint64(1023) // 11 1111 1111
	level := uint64(255)   // 1111 1111
	status := uint64(3)    // 11 (무적 상태라고 가정)

	// [압축하기] 비트 시프트를 이용해 64비트 공간 하나에 다 밀어넣기
	// ID는 그대로, 레벨은 10칸 밀고, 상태는 18칸 밀어서 합침
	var packedData uint64
	packedData |= userID
	packedData |= (level << 10)
	packedData |= (status << 18)

	fmt.Printf("압축된 데이터 (2진수): %064b\n", packedData)
	fmt.Printf("압축된 데이터 크기: %d 바이트\n", unsafe.Sizeof(packedData))

	// [해독하기] 다시 비트 연산으로 꺼내기
	extractedID := packedData & 0x3FF           // 하위 10비트만 가져오기 (0x3FF = 1111111111)
	extractedLevel := (packedData >> 10) & 0xFF // 10비트 밀고 8비트만 가져오기
	extractedStatus := (packedData >> 18) & 0x3 // 18비트 밀고 2비트만 가져오기

	fmt.Println("\n--- 압축 해제 결과 ---")
	fmt.Printf("유저 ID: %d\n", extractedID)
	fmt.Printf("레벨: %d\n", extractedLevel)
	fmt.Printf("상태: %d\n", extractedStatus)

	fmt.Println("\n[제프 딘의 한마디]")
	fmt.Println("\"단순한 변수 선언보다 데이터가 메모리에서 차지하는 공간을 1비트라도 아끼는 것이,")
	fmt.Println(" 수백만 대의 서버가 돌아가는 구글 환경에서는 수십억 원의 비용 절감이 됩니다.\"")
}
