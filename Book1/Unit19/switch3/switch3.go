// switch3
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fallthrough : 다음 case의 문장을 실행하고 싶을 때 사용
	i1 := 3

	switch i1 {

	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3: // 3과 변수의 값이 일치하므로 이 부분 실행
		fmt.Println("3 이상")
		fallthrough // fallthrough를 사용했으므로 아래 case를 모두 실행
	case 2:
		fmt.Println("2 이상") // 실행
		fallthrough
	case 1:
		fmt.Println("1 이상") // 실행
		fallthrough
	case 0:
		fmt.Println("0 이상") // 실행, 마지막 case에는 fallthrough를 사용할 수 없음
	}

	fmt.Println()

	// 여러 조건 함께 처리
	i2 := 3

	switch i2 {
	case 2, 4, 6: // i2가 2,4,6일 때
		fmt.Println("짝수")
	case 1, 3, 5: // i2가 1,3,5일 때
		fmt.Println("홀수")
	}

	fmt.Println()

	// 조건식으로 분기
	i3 := 7

	switch { // case에 조건식을 지정했으므로 판단할 변수는 생략
	case i3 >= 5 && i3 < 10: // i3이 5보다 크거나 같으면서 10보다 작을 때
		fmt.Println("5 이상, 10 미만")
	case i3 >= 0 && i3 < 5: // i3이 0보다 크거나 같으면서 5보다 작을 때
		fmt.Println("0 이상, 5 미만")
	}

	fmt.Println()

	// switch 분기문 안에서 함수를 실행하고 결과값으로 분기 가능
	// 이 때, 함수를 호출하고 뒤에 ;(세미콜론)을 붙여줌
	// 또한 case에서는 값으로 분기할 수 없고 조건식만 사용 가능
	// * math/rand : 무작위(랜덤) 패키지
	//	 - Seed: 시드값 설정
	//	 - Intn: 랜덤값 생성, 랜덤 값의 범위는 0부터 매개변수로 설정한 값까지
	// * time : 시간 패키지
	//	 - Now: 현재 시간을 구하는 함수
	//	 - UnixNano: 유닉스 시간을 나노초 단위로 리턴
	randSeed := time.Now().UnixNano()
	fmt.Printf("Current Time - UnixNano : %d\n", randSeed)
	rand.Seed(randSeed)
	switch i4 := rand.Intn(10); {
	case i4 >= 3 && i4 < 6:
		fmt.Printf("3이상, 6미만\ni4 = %d", i4)
	case i4 == 9:
		fmt.Println("9")
	default:
		fmt.Printf("i4 = %d\n", i4)
	}
}
