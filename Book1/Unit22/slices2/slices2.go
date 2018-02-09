// slices2
package main

import "fmt"

func main() {
	// 슬라이스와 용량
	a1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a1), cap(a1)) // 5 5 : 길이가 5이며 용량이 5인 슬라이스

	a1 = append(a1, 6, 7)         // 슬라이스 a1에 값 6,7을 추가
	fmt.Println(len(a1), cap(a1)) // 7 10 : 길이가 7이며 용량이 10인 슬라이스, 용량이 늘어남!

	fmt.Println()

	// 부분 슬라이스 만들기
	a2 := []int{1, 2, 3, 4, 5}
	b2 := a2[0:5]   // a2의 인덱스 0부터 5까지 참조
	fmt.Println(a2) // [1 2 3 4 5]
	fmt.Println(b2) // [1 2 3 4 5]

	fmt.Println()

	a3 := []int{1, 2, 3, 4, 5}
	fmt.Println(a3[0:3]) // [1 2 3]
	fmt.Println(a3[1:3]) // [2 3]
	fmt.Println(a3[2:5]) // [3 4 5]

	fmt.Println()

	a4 := []int{1, 2, 3, 4, 5}
	fmt.Println(a4[:])         // [1 2 3 4 5]
	fmt.Println(a4[0:])        // [1 2 3 4 5]
	fmt.Println(a4[:5])        // [1 2 3 4 5]
	fmt.Println(a4[0:len(a4)]) // [1 2 3 4 5]

	fmt.Println(a4[3:])  // [4 5]
	fmt.Println(a4[:3])  // [1 2 3]
	fmt.Println(a4[1:3]) // [2 3]

	fmt.Println()

	a5 := [5]int{1, 2, 3, 4, 5} // 배열 선언
	b5 := a5[:2]                // 배열 a5의 일부를 부분 슬라이스로 참조
	b5[0] = 9                   // 부분 슬라이스는 참조이므로 a[0], b[0]의 값이 모두 바뀜

	fmt.Println(a5) // [9 2 3 4 5]
	fmt.Println(b5) // [9 2]

	fmt.Println()

	a6 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b6 := a6[0:6:8]               // 인덱스 0부터 6까지 가져와서 부분 슬라이스로 만들고 용량을 8로 설정
	fmt.Println(len(b6), cap(b6)) // 6 8 : 길이가 6이며 용량이 8인 슬라이스
	fmt.Println(b6)               // [1 2 3 4 5 6]
}
