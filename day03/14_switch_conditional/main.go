package main

import "fmt"

func main() {
	// [switch문: 값이 딱딱 맞아떨어지는 분기 처리]
	// if-else if 구조가 너무 길어질 때 사용하면 가독성이 획기적으로 올라갑니다.
	rank := 2

	fmt.Println("--- 1. 기본 switch문 ---")
	switch rank {
	case 1:
		fmt.Println("금메달입니다!")
	case 2:
		fmt.Println("은메달입니다!")
	case 3:
		fmt.Println("동메달입니다!")
	default:
		// 매칭되는 case가 없을 때 최종 실행됩니다.
		fmt.Println("아쉽게도 메달 획득에 실패했습니다.")
	}

	// [2. 다중 조건 case 및 조건식이 생략된 switch문]
	// case 뒤에 쉼표(,)를 사용해 여러 매칭 조건을 나열할 수 있습니다.
	// switch 뒤에 대상을 생략하면 내부 case 조건식을 if문처럼 유연하게 활용할 수 있습니다.
	fmt.Println("\n--- 2. 다중 조건 및 조건식 생략 switch문 ---")
	score := 88
	switch {
	case score >= 90:
		fmt.Println("성적 평가: 우수")
	case score >= 70 && score < 90:
		fmt.Println("성적 평가: 보통")
	default:
		fmt.Println("성적 평가: 노력 요함")
	}

	// [3. fallthrough 키워드]
	// Go 언어의 switch문은 기본적으로 일치하는 case가 실행되면 자동으로 종료됩니다. (break 자동 적용)
	// 만약 다음 case의 동작까지 강제로 이어서 실행하고 싶다면 fallthrough 키워드를 사용합니다.
	fmt.Println("\n--- 3. fallthrough 키워드 활용 ---")
	num := 1
	switch num {
	case 1:
		fmt.Println("case 1 실행")
		fallthrough // 바로 아래 case 2도 조건 검사 없이 강제로 들어갑니다.
	case 2:
		fmt.Println("case 2 실행 (fallthrough 효과)")
	case 3:
		fmt.Println("case 3 실행")
	}
}
