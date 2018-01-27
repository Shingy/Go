// goto1
package main

import "fmt"

func main() {
	// Code 1
	/*
		var a int = 1

		if a == 1 {
			fmt.Println("Error 1") // 에러 1 상황
			return
		}
		if a == 2 {
			fmt.Println("Error 2") // 에러 2 상황
			return
		}
		if a == 3 {
			fmt.Println("Error 1") // 에러 1 상황, 중복 코드
			return
		}
		fmt.Println(a)
		fmt.Println("Success") // 정상 실행
		return
	*/

	// Code 2
	// goto 키워드와 레이블을 사용하면 중복 코드 없이 간단하게 작성 가능
	var a2 int = 1
	if a2 == 1 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}
	if a2 == 2 {
		goto ERROR2 // 에러 2 상황이면 ERROR2 레이블로 이동
	}
	if a2 == 3 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	fmt.Println(a2)
	fmt.Println("Success") // 정상실행
	return

ERROR1: // 에러 처리 1
	fmt.Println("Error 1")
	return

ERROR2: // 에러 처리 2
	fmt.Println("Error 2")
	return
}
