// strings
package main

import "fmt"
import "unicode/utf8"

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Printf("%s 길이 : %d\n", s1, len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
	fmt.Printf("%s 길이 : %d\n", s2, len(s2)) // 5: 알파벳 5글자이므로 5

	fmt.Printf("\n문자열의 실제 길이 : %d\n\n", utf8.RuneCountInString(s1)) // 2: 문자열의 실제 길이를 구함

	// 문자열 연산
	var ss1 string = "한글"
	var ss2 string = "한글"
	var ss3 string = "Go"

	fmt.Println(ss1 == ss2)      // true: 두 문자열이 같으므로 true
	fmt.Println(ss1 + ss3 + ss2) // 한글Go한글
	fmt.Println("안녕하세요 " + ss1)  // 안녕하세요 한글

	// 문자열 & 배열
	var s12 = "Hello"
	fmt.Printf("\n%c\n", s12[1]) // e: 문자열에서 두 번째 문자 e
	/**
	 *	제어문자
	 *	\a : 경고음, 벨(u0007)
	 *	\b : 백스페이스(u0008)
	 *	\f : 폼 피드(u000c)
	 *	\n : 라인 피드, 새 줄(u000a)
	 *	\r : 캐리지 리턴(u000d)
	 *	\t : 수평 탭(u0009)
	 *	\v : 수직 탭(u000b)
	 *	\\ : 백슬래시(u005c)
	 *	\' : 작은따옴표(u0027), rune 변수에 저장할 때 사용할 수 있습니다.
	 *	\" : 큰따옴표(u0022), string 변수에 저장할 때 사용할 수 있습니다.
	 */
}
