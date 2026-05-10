package main

import (
	"fmt"
	"log"

	// ONNX Runtime Go 바인딩 라이브러리
	ort "github.com/yalue/onnxruntime_go"
)

func main() {
	fmt.Println("=== Go 언어에서 ONNX 모델 실행하기 예제 ===")

	// 1. ONNX Runtime 라이브러리 경로 설정 (OS에 맞게 수정 필요)
	// Go에서 ONNX를 사용하려면 시스템에 ONNX Runtime C 라이브러리가 있어야 합니다.
	// Windows의 경우 onnxruntime.dll 파일 경로를 지정해줍니다.
	// ort.SetSharedLibraryPath("./onnxruntime.dll") // 주석을 해제하고 dll 경로를 맞춰주세요.

	// 2. ONNX 환경 초기화
	err := ort.InitializeEnvironment()
	if err != nil {
		log.Fatalf("ONNX Runtime 환경 초기화 실패: %v\n(힌트: onnxruntime.dll 라이브러리가 없거나 경로가 틀렸을 수 있습니다.)\n", err)
	}
	defer ort.DestroyEnvironment()

	fmt.Println("ONNX Runtime 환경 초기화 완료!")

	// 3. ONNX 모델 경로 설정
	modelPath := "example_model.onnx" // 파이썬에서 추출한 실제 onnx 파일 경로

	// 4. 입력(Input) / 출력(Output) 텐서 구조 정의
	// 모델 구조에 맞춰서 수정해야 합니다. (예: input 이름이 'x' 이고 output 이름이 'y'인 경우)
	inputNames := []string{"input"}
	outputNames := []string{"output"}

	// 예시: Batch Size 1, Feature 개수 4 (1x4 배열)
	inputShape := ort.NewShape(1, 4)
	outputShape := ort.NewShape(1, 3) // 결과값 분류가 3개 (1x3 배열)

	// 5. 입력 데이터 준비 (예시 더미 데이터)
	inputData := []float32{1.2, 3.4, 5.6, 7.8}

	// 입력 텐서(Tensor) 객체 생성
	inputTensor, err := ort.NewTensor(inputShape, inputData)
	if err != nil {
		log.Fatalf("입력 텐서 생성 오류: %v\n", err)
	}
	defer inputTensor.Destroy()

	// 출력 데이터를 담을 빈 배열 및 텐서 생성
	outputData := make([]float32, 3) 
	outputTensor, err := ort.NewTensor(outputShape, outputData)
	if err != nil {
		log.Fatalf("출력 텐서 생성 오류: %v\n", err)
	}
	defer outputTensor.Destroy()

	// 6. ONNX 세션(Session) 생성 및 모델 로드
	// *주의: 실제 파일(example_model.onnx)이 없으면 여기서 에러가 발생합니다.
	session, err := ort.NewAdvancedSession(
		modelPath,
		inputNames, outputNames,
		[]ort.ArbitraryTensor{inputTensor},
		[]ort.ArbitraryTensor{outputTensor},
		nil, // 추가 옵션 (없으면 nil)
	)
	
	if err != nil {
		log.Fatalf("ONNX 모델 로드 및 세션 생성 실패: %v\n", err)
	}
	defer session.Destroy()

	fmt.Println("ONNX 모델 로드 성공! 추론을 시작합니다...")

	// 7. 모델 추론(Inference) 실행
	err = session.Run()
	if err != nil {
		log.Fatalf("추론 실행 실패: %v\n", err)
	}

	// 8. 최종 결과 출력 (outputData 배열에 결과가 덮어씌워짐)
	fmt.Printf("추론 결과 (Output Data): %v\n", outputData)
	fmt.Println("=== 추론 완료 ===")
}
