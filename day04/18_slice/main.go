package main

import "fmt"

func main() {
	// [슬라이스(Slice): 크기가 유동적인 배열]
	// 배열은 크기가 고정이지만, 슬라이스는 필요에 따라 원소를 추가하고 제거할 수 있습니다.
	// Go 실무 개발에서는 배열보다 슬라이스를 훨씬 더 많이 사용합니다.

	// [1. 슬라이스 선언 - make 함수 사용]
	// make([]타입, 길이) 또는 make([]타입, 길이, 용량)
	s1 := make([]int, 5)
	fmt.Println("1. make로 만든 슬라이스:", s1)

	// [2. 슬라이스 리터럴 (선언과 동시에 초기화)]
	// 배열 리터럴과 다르게 대괄호 안에 크기를 적지 않습니다.
	s2 := []string{"Go", "Python", "Java"}
	fmt.Println("2. 슬라이스 리터럴:", s2)
	fmt.Printf("   len(길이)=%d\n", len(s2))

	// [3. append: 슬라이스에 원소 추가하기]
	// 슬라이스 = append(슬라이스, 추가할값)
	// 한 번에 여러 개를 추가하는 것도 가능합니다.
	s2 = append(s2, "Kotlin")
	s2 = append(s2, "Swift", "Rust")
	fmt.Println("3. append로 원소 추가 후:", s2)

	// [4. 슬라이싱: 슬라이스의 일부를 잘라 새 슬라이스 만들기]
	// 슬라이스[시작인덱스:끝인덱스]
	// 끝인덱스는 결과에 포함되지 않습니다. (반개방 구간)
	fmt.Println("\n4. 슬라이싱:")
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("   원본:", nums)
	fmt.Println("   nums[1:3] =", nums[1:3]) // 20, 30
	fmt.Println("   nums[:3]  =", nums[:3])  // 처음부터 인덱스 2까지
	fmt.Println("   nums[2:]  =", nums[2:])  // 인덱스 2부터 끝까지

	// [5. len(길이)과 cap(용량)]
	// len: 현재 슬라이스에 담긴 원소의 수
	// cap: 재할당 없이 담을 수 있는 최대 원소 수 (내부 배열 크기)
	fmt.Println("\n5. len과 cap:")
	s3 := make([]int, 3, 10)
	fmt.Printf("   len=%d, cap=%d, 내용=%v\n", len(s3), cap(s3), s3)

	// append 시 cap을 초과하면 Go가 자동으로 더 큰 내부 배열을 새로 할당합니다.
	var autoGrow []int
	for i := 0; i < 5; i++ {
		autoGrow = append(autoGrow, i)
		fmt.Printf("   append 후: len=%d, cap=%d\n", len(autoGrow), cap(autoGrow))
	}
}
