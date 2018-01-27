// switch1
package main

import "fmt"

func main() {
	i := 3

	switch i { // 값을 판단할 변수 설정 : 각 조건에 일치하는 코드 실행
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3: // 3과 변수의 값이 일치하므로 이 부분을 실행하고 이후 실행을 중단
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(-1)
	}

	fmt.Println()

	// 문자열도 사용 가능
	s := "world"

	switch s { // 값을 판단할 변수 설정 : 각 조건에 일치하는 코드 실행
	case "hello":
		fmt.Println("hello")
	case "world": // 문자열 "world"와 변수의 값이 일치하므로 이 부분을 실행하고 이후 실행을 중단
		fmt.Println("world")
	default:
		fmt.Println("일치하는 문자열이 없습니다.")
	}
}
