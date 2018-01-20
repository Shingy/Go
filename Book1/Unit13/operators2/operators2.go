// operators2
package main

import "fmt"

func main() {
	// 1) ==(같다) 두 값이 같은지 비교
	//	  * 실수형은 값을 연산한 뒤에는 오차가 발생하므로 ==로 비교할 때 주의
	//	  * 문자열은 내용이 같은지 비교
	//	  * 배열은 요소의 내용이 모두 같은지 비교
	//	  * 슬라이스와 맵은 배열과는 달리 내용을 비교할 수 없고, 메모리에 할당되어 있는지 확인
	//	  * 포인터는 주소가 같은지 비교
	fmt.Println(1 == 1)             // true : 두 정수가 같으므로 true
	fmt.Println(3.5 == 3.5)         // true : 두 실수가 같으므로 true
	fmt.Println("Hello" == "Hello") // true : 두 문자열이 같으므로 true

	a1 := [3]int{1, 2, 3}
	b1 := [3]int{1, 2, 3}
	fmt.Println(a1 == b1) // true : 두 배열이 같으므로 true

	c1 := []int{1, 2, 3}
	fmt.Println(c1 == nil) // false : 슬라이스를 nil과 비교하여 메모리가 할당되었는지 확인

	d1 := map[string]int{"Hello": 1}
	fmt.Println(d1 == nil) // false : 맵을 nil과 비교하여 메모리가 할당되었는지 확인

	e1 := 1
	var p11 *int = &e1
	var p12 *int = &e1
	fmt.Println(p11 == p12) // true : 포인터에 저장된 메모리 주소가 같으므로 true

	// 2) !=(같지않다) 두 값이 다른지 비교
	fmt.Println(1 != 2)             // true : 두 정수가 다르므로 true
	fmt.Println(3.5 != 5.5)         // true : 두 실수가 다르므로 true
	fmt.Println("Hello" != "world") // true : 두 문자열이 다르므로 true

	a2 := [3]int{1, 2, 3}
	b2 := [3]int{3, 2, 1}
	fmt.Println(a2 != b2) // true : 두 배열이 다르므로 true

	c2 := []int{1, 2, 3}
	fmt.Println(c2 != nil) // true : 슬라이스를 nil과 비교하여 메모리가 할당되었는지 확인

	d2 := map[string]int{"Hello": 1}
	fmt.Println(d2 != nil) // true : 맵을 nil과 비교하여 메모리가 할당되었는지 확인

	e2 := 1
	f2 := 1
	var p21 *int = &e2
	var p22 *int = &f2
	fmt.Println(p21 != p22) // true : 포인터에 저장된 메모리 주소가 다르므로 true

	// 3) <(작다) 앞의 값이 작은지 비교, 문자열은 ASCII 코드값을 기준으로 판단, 첫 글자가 같다면 그 다음 글자부터 차례대로 비교하여 최종값을 구함
	fmt.Println(1 < 2)             // true : 1이 2보다 작으므로 true
	fmt.Println(3.5 < 5.5)         // true : 3.5가 5.5보다 작으므로 true
	fmt.Println("Hello" < "world") // true : H가 w보다 ASCII 코드값이 작으므로 true

	// 4) <=(작거나 같다) 앞의 값이 작거나 같은지 비교
	fmt.Println(2 <= 2)             // true : 2가 2보다 작거나 같으므로 true
	fmt.Println(3.5 <= 5.5)         // true : 3.5가 5.5보다 작거나 같으므로 true
	fmt.Println("Hello" <= "world") // true : H가 w보다 ASCII 코드값이 작거나 같으므로 true

	// 5) >(크다) 앞의 값이 큰지 비교
	fmt.Println(2 > 1)             // true : 2가 1보다 크므로 true
	fmt.Println(5.5 > 3.5)         // true : 5.5가 3.5보다 크므로 true
	fmt.Println("world" > "Hello") // true : w가 H보다 ASCII 코드값이 크므로 true

	// 6) >(=크거나 같다) 앞의 값이 크거나 같은지 비교
	fmt.Println(2 >= 2)             // true : 2가 2보다 크거나 같으므로 true
	fmt.Println(5.5 >= 3.5)         // true : 5.5가 3.5보다 크거나 같으므로 true
	fmt.Println("world" >= "Hello") // true : w가 H보다 ASCII 코드값이 크거나 같으므로 true

	// 7) &&(AND 논리 연산) 두 BOOL 값이 모두 참인지 확인
	fmt.Println(true && true)   // true : 두 값이 모두 true이므로 true
	fmt.Println(true && false)  // false : 두 값 중 하나가 true이므로 false
	fmt.Println(false && false) // false : 두 값이 모두 false이므로 false

	// 8) ||(OR 논리 연산) 두 BOOL 값 중 하나라도 참인지 확인
	fmt.Println(true || true)   // true : 두 값이 모두 true이므로 true
	fmt.Println(true || false)  // true : 두 값 중 하나가 true이므로 true
	fmt.Println(false || false) // false : 두 값이 모두 false이므로 false

	// 9) !(NOT 논리 연산) BOOL 값을 반대로 연산
	fmt.Println(!true)  // false : true의 반대는 false
	fmt.Println(!false) // true : false의 반대는 true

	// 10) &(참조/레퍼런스 연산) 현재 변수의 메모리 주소를 구함
	a10 := 1
	b10 := &a10      // a10의 메모리 주소를 b10에 대입
	fmt.Println(b10) // 메모리 주소 출력

	// 11) *(역참조 연산) 현재 포인터 변수에 저장된 메모리에 접근하여 값을 가져오거나 저장
	a11 := new(int)
	*a11 = 1          // a11에 저장된 메모리에 접근하여 1을 저장
	fmt.Println(*a11) // 1 : a11에 저장된 메모리에 접근하여 값을 가져옴

	// 12) <-(채널 수신 연산) 채널에 값을 보내거나 값을 가져옴
	c12 := make(chan int)
	go func() {
		c12 <- 1 // 채널 c12에 1을 보냄
	}()

	a12 := <-c12 // 채널 c12에서 값을 가져와서 a12에 대입
	fmt.Println(a12)

	// 13) ++(증가) 변수의 값을 1 증가 시킴, 사용할 수 있는 자료형은 정수, 실수, 복소수
	//	  Go 언어에서는 ++ 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있음
	//	  따라서 ++ 연산자는 단독으로 사용하거나 if 조건문, for 반복문 안에서 사용
	a13 := 1
	a13++ // 2 : 정수 1을 1 증가시켜서 2

	b13 := 1.5
	b13++ // 2.5 : 실수 1.5를 1 증가시켜서 2.5

	c13 := 1 + 2i
	c13++ // (2+2i) : 복소수 1+2i를 1 증가시켜서 2+2i

	fmt.Println(a13) // 2
	fmt.Println(b13) // 2.5
	fmt.Println(c13) // (2+2i)

	// a13 := 1
	// b13 := a13++	// 컴파일에러
	// c13 := ++a13	// 컴파일에러
	// ++a13		// 컴파일에러

	// 14) --(감소) 변수의 값을 1 감소 시킴, 사용할 수 있는 자료형은 정수, 실수, 복소수
	//	  Go 언어에서는 -- 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있음
	//	  따라서 -- 연산자는 단독으로 사용하거나 if 조건문, for 반복문 안에서 사용
	a14 := 1
	a14-- // 0 : 정수 1에서 1 감소시켜서 0

	b14 := 1.5
	b14-- // 0.5 : 실수 1.5를 1 증가시켜서 0.5

	c14 := 1 + 2i
	c14-- // (0+2i) : 복소수 1+2i를 1 증가시켜서 0+2i

	fmt.Println(a14) // 0
	fmt.Println(b14) // 0.5
	fmt.Println(c14) // (0+2i)

	// a14 := 1
	// b14 := a14--	// 컴파일에러
	// c14 := --a14	// 컴파일에러
	// --a14		// 컴파일에러

	/**
	 *	연산자 우선 순위(숫자가 높을 수록 먼저 계산)
	 *	5 : *  /  %  <<  >>  &  &^
	 *	4 : +  -  |  ^
	 *	3 : ==  !=  <  <=  >  >=
	 *	2 : &&
	 *	1 : ||
	 */

}
