package main

import "fmt"

func main() {
	// [맵(Map): 키(Key)로 값(Value)을 즉시 찾는 사전(Dictionary) 구조]
	// 배열/슬라이스는 숫자 인덱스로 접근하지만, 맵은 의미 있는 문자열 키로 접근합니다.
	// 현실의 사전처럼 단어(키)를 주면 뜻(값)을 즉시 꺼내줍니다.

	// [1. 맵 선언 - make 함수 사용]
	// make(map[키타입]값타입)
	studentAge := make(map[string]int)
	studentAge["김철수"] = 20
	studentAge["이영희"] = 22
	studentAge["박민준"] = 21
	fmt.Println("1. 학생 나이 맵:", studentAge)

	// [2. 맵 리터럴 (선언과 동시에 초기화)]
	capitals := map[string]string{
		"한국": "서울",
		"일본": "도쿄",
		"미국": "워싱턴 D.C.",
	}
	fmt.Println("2. 수도 맵:", capitals)

	// [3. 키로 값 조회]
	fmt.Println("\n3. 값 조회:")
	fmt.Println("   한국의 수도:", capitals["한국"])

	// [4. ok 패턴으로 키 존재 여부 확인]
	// 맵에서 없는 키를 조회하면 에러 대신 기본값(int: 0, string: "")이 반환됩니다.
	// 이는 "값이 없음"과 "값이 0(또는 빈 문자열)"을 구분하지 못하는 문제를 일으킵니다.
	// 두 번째 반환값 ok(bool)로 실제로 키가 존재하는지 확인해야 합니다.
	fmt.Println("\n4. ok 패턴으로 존재 여부 확인:")
	if age, ok := studentAge["김철수"]; ok {
		fmt.Printf("   김철수의 나이: %d세\n", age)
	}
	if _, ok := studentAge["홍길동"]; !ok {
		fmt.Println("   홍길동은 맵에 존재하지 않습니다.")
	}

	// [5. delete로 요소 삭제]
	// delete(맵, 키)
	fmt.Println("\n5. 요소 삭제:")
	fmt.Println("   삭제 전:", studentAge)
	delete(studentAge, "박민준")
	fmt.Println("   박민준 삭제 후:", studentAge)

	// [6. range를 활용한 맵 순회]
	// 주의: 맵의 순회 순서는 매 실행마다 무작위로 달라집니다. (순서 보장 없음)
	fmt.Println("\n6. range로 맵 순회:")
	for key, value := range capitals {
		fmt.Printf("   %s의 수도 → %s\n", key, value)
	}
}
