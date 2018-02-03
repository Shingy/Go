// arrays
package main

import "fmt"

func main() {
	var a1 [5]int // int 형이며 길이가 5인 배열 선언

	a1[2] = 7 // 배열의 세 번째 요소에 7 대입

	fmt.Println(a1) // [0 0 7 0 0]

	fmt.Println()

	var a2 [5]int = [5]int{32, 29, 78, 16, 81} // int 형이며 길이가 5인 배열을 선언하고 초기화
	var b2 = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이 생략
	c2 := [5]int{32, 29, 78, 16, 81}           // 배열을 선언할 때 var 키워드, 자료형과 길이 생략

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	fmt.Println()

	a3 := [5]int{32, 29, 78, 16, 81} // 배열을 생성하면서 값을 초기화
	b3 := [5]int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}
	c3 := [...]int{32, 29, 78, 16, 81}           // 초기화할 요소가 5개이며 ...을 사용했으므로 배열 크기는 5로 설정됨
	d3 := [...]string{"Maria", "Andrew", "John"} // 초기화할 요소가 3개이며 ...을 사용했으므로 배열 크기는 3으로 설정됨

	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(c3)
	fmt.Println(d3)

	fmt.Println()

	// 다차원 배열 선언
	a4 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(a4)

	fmt.Println()

	a5 := [5]int{32, 29, 78, 16, 81}
	for i5 := 0; i5 < len(a5); i5++ { // len 함수로 배열의 길이를 구한 뒤 배열의 길이만큼 반복
		fmt.Println(a5[i5])
	}

	fmt.Println()

	a6 := [5]int{32, 29, 78, 16, 81}
	for i6, value6 := range a6 { // i6에는 인덱스, value6에는 배열 요소의 값이 들어감
		fmt.Println(i6, value6)
	}

	fmt.Println()

	a7 := [5]int{32, 29, 78, 16, 81}
	for value7 := range a7 { // value7에는 값 대신 인덱스가 들어감
		fmt.Println(value7)
	}

	fmt.Println()

	a8 := [5]int{32, 29, 78, 16, 81}
	for _, value8 := range a8 { // 인덱스는 생략, value8에 배열 요소의 값이 들어감
		fmt.Println(value8)
	}

	fmt.Println()

	// 배열 복사
	a9 := [5]int{1, 2, 3, 4, 5}
	b9 := a9        // 배열을 대입하면 배열 전체가 복사됨
	fmt.Println(a9) // [1, 2, 3, 4, 5]
	fmt.Println(b9) // [1, 2, 3, 4, 5]
}
