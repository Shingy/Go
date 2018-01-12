// complexNum
package main

import "fmt"

func main() {
	var num1 complex64 = 1 + 2i           // 실수부 1, 허수부 2i
	var num2 complex64 = 4.2342 + 2.3527i // 실수부 고정소수점 4.2342, 허수부 고정소수점 2.3527i

	var num3 complex64 = 1.e+3i       // 실수부 없음, 허수부 부동소수점 1.e+3i
	var num4 complex64 = 7.27151e-10i // 실수부 없음, 허수부 부동소수점 7.27151e-10i

	var num5 complex128 = 1 + 2i                   // 실수부, 허수부 2i
	var num6 complex128 = 5.32521e-10 + .12345E+2i // 실수부 부동소수점 5.32521e-10, 허수부 부동소수점 .12345E+2i

	fmt.Print("num1 : ")
	fmt.Println(num1)
	fmt.Print("num2 : ")
	fmt.Println(num2)

	fmt.Print("\nnum3 : ")
	fmt.Println(num3)
	fmt.Print("num4 : ")
	fmt.Println(num4)

	fmt.Print("\nnum5 : ")
	fmt.Println(num5)
	fmt.Print("num6 : ")
	fmt.Println(num6)

	var r1 float32 = real(num1) // num1의 실수부 1
	var i1 float32 = imag(num1) // num1의 허수부 2

	var r2 float64 = real(num5) // num5의 실수부 1
	var i2 float64 = imag(num5) // num5의 허수부 2

	fmt.Print("\nr1 : ")
	fmt.Println(r1)
	fmt.Print("i1 : ")
	fmt.Println(i1)
	fmt.Print("\nr2 : ")
	fmt.Println(r2)
	fmt.Print("i2 : ")
	fmt.Println(i2)

	var comnum1 complex64 = complex(1, 2)                    // 실수부 1, 허수부 2i
	var comnum2 complex128 = complex(5.32521e-10, .12345E+2) // 실수부 부동소수점 5.32521e-10, 허수부 부동소수점 .12345E+2i

	fmt.Print("\ncomnum1 : ")
	fmt.Println(comnum1)
	fmt.Print("comnum2 : ")
	fmt.Println(comnum2)

}
