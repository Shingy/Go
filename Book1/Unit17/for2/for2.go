// for2
package main

import "fmt"

func main() {
	// break 사용하기
	i1 := 0
	for {
		if i1 > 4 { // i가 4가 되는 순간 반복문 중단
			break
		}

		fmt.Println(i1)
		i1 += 1 // 변화식에서 조건 변경
	}

	fmt.Println()

	// break 키워드에 레이블 지정(레이블 이름은 변수 이름을 짓는 규칙과 동일함)
Loop: // Loop 레이블을 지정
	// fmt.Println("begin for loop")	//==> 들어가면 안되는 코드 : 레이블과 for 키워드 사이에 다른 코드가 있으면 안됨
	for i := 0; i < 3; i++ { // 반복문1
		for j := 0; j < 3; j++ { // 반복문2
			if j == 2 { // j가 2일 때, 중첩된 반복문을 빠져나감
				break Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")

}
