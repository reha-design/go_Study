// package main: 이 파일이 실행 가능한 프로그램의 시작점임을 선언
// Go 언어에서 실행 파일을 만들려면 반드시 package main이라는 선언이 필요
package main

// import "fmt": 표준 라이브러리인 fmt(format의 약자) 패키지를 가져온다.
// fmt 패키지는 콘솔에 텍스트를 출력하거나 입력을 받는 등의 포맷팅 기능을 제공
import "fmt"

// main 함수: 프로그램이 실행될 때 가장 먼저 호출되는 함수
// 프로그램의 '입구(Entry Point)' 역할을 수행
func main() {
	// fmt.Println: 콘솔(터미널)에 괄호 안의 내용을 출력하고 줄바꿈을 한다.
	// 여기서는 "Hello World"라는 문자열을 화면에 표시한다.
	fmt.Println("Hello World")
}
