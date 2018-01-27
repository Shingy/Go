// for3
package main

import "fmt"

func main() {
	// continue 사용하기
	for i1 := 0; i1 < 5; i1++ {
		if i1 == 2 { // i1이 2일 때, 아래 부분 코드를 실행하지 않고 넘어감
			continue
		}

		fmt.Println(i1)
	}

	fmt.Println()

	// continue 키워드에 레이블 지정(레이블 이름은 변수 이름을 짓는 규칙과 동일함)

Loop: // Loop 레이블을 지정
	// fmt.Println("begin for loop")	//==> 들어가면 안되는 코드 : 레이블과 for 키워드 사이에 다른 코드가 있으면 안됨
	for i := 0; i < 3; i++ { // 반복문1
		for j := 0; j < 3; j++ { // 반복문2
			if j == 2 { // j가 2일 때, 아래 부분 코드를 실행하지 않고 반복문 1부터 이어서 실행
				continue Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")

	fmt.Println()

	// 반복문에서 변수 여러개 사용하기
	for i2, j2 := 0, 0; i2 < 10; i2, j2 = i2+1, j2+2 {
		fmt.Println(i2, j2)
	}
}
