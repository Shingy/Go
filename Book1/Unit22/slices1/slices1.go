// slices1
package main

import "fmt"

func main() {
	a1 := []int{32, 29, 78, 16, 81} // 슬라이스를 생성하면서 값을 초기화
	b1 := []int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	fmt.Println(a1, b1)

	fmt.Println()

	var s2 = make([]int, 5, 10) // 슬라이스 길이가 5이고 용량이 10인 슬라이스 생성

	fmt.Printf("size of s2 : %d\n", len(s2))

	fmt.Println()

	a3 := make([]int, 5, 10)
	fmt.Println(len(a3)) // 길이는 5
	fmt.Println(cap(a3)) // 용량은 10

	fmt.Println()

	/*
		a4 := make([]int, 5, 10) // 길이가 5이면 a[0],a[1],a[2],a[3],a[4]가 생성
		fmt.Println(a4[4])       // 0 : make 함수를 사용하면 슬라이스의 요소는 모두 0으로 초기화
		fmt.Println(a4[5])       // Runtime Error : 길이를 벗어난 인덱스 접근
		fmt.Println(a4[8])       // Runtime Error : 길이를 벗어난 인덱스 접근

		fmt.Println()
	*/

	a5 := []int{1, 2, 3}
	a5 = append(a5, 4, 5, 6)
	fmt.Println(a5) // [1 2 3 4 5 6]

	fmt.Println()

	a6 := []int{1, 2, 3}
	b6 := []int{4, 5, 6}
	a6 = append(a6, b6...) // 슬라이스 a6에 슬라이스 b6를 붙일 때는 b6...을 씀
	fmt.Println(a6)        // [1 2 3 4 5 6]

	fmt.Println()

	a7 := [3]int{1, 2, 3}
	var b7 [3]int

	b7 = a7   // 배열의 요소가 모두 복사됨
	b7[0] = 9 // b7[0]에 9를 대입하면 b7의 첫 번째 요소만 바뀜

	fmt.Println(a7) // [1 2 3]
	fmt.Println(b7) // [9 2 3]

	fmt.Println()

	// Reference 타입
	a8 := []int{1, 2, 3}
	var b8 []int // 슬라이스로 선언

	b8 = a8   // a8을 b8에 대입해도 요소가 모두 복사되지 않고 참조만 함
	b8[0] = 9 // 슬라이스는 참조이므로, a8[0], b8[0]의 값이 모두 바뀜

	fmt.Println(a8) // [9 2 3]
	fmt.Println(b8) // [9 2 3]

	fmt.Println()

	// 슬라이스 복사
	a9 := []int{1, 2, 3, 4, 5}
	b9 := make([]int, 3) // make 함수로 공간을 할당

	copy(b9, a9) // 슬라이스 a9의 요소를 슬라이스 b9에 복사

	fmt.Println(a9) // [1 2 3 4 5]
	fmt.Println(b9) // [1 2 3] : 슬라이스 b9의 길이는 3이므로 a9의 요소 3개만 복사됨

	fmt.Println()

	a10 := []int{1, 2, 3}
	b10 := make([]int, 3)

	copy(b10, a10) // 슬라이스를 복사했으므로
	b10[0] = 9     // b10[0]만 바뀌고, a10[0]은 바뀌지 않음

	fmt.Println(a10) // [1 2 3]
	fmt.Println(b10) // [9 2 3]

	fmt.Println()

}
