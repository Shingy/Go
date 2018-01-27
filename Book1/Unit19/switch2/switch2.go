// switch2
package main

import "fmt"

func main() {
	// break 사용
	s := "Hello"
	i := 2

	switch i { // 값을 판단할 변수 설정
	case 1:
		fmt.Println(1)
	case 2: // i가 2이고, s가 "Hello" 이면 Hello 2를 출력하고 switch 분기문 실행 중단
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}

		fmt.Println(2)
	}
}
