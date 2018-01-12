// constant
package main

//import "fmt"

func main() {
	const age int = 10
	const name string = "Maria"
	const score int		// 컴파일 에러

	age = 20		// 컴파일 에러
	name = "Grace"	// 컴파일 에러

	const age = 10			// int
	const name = "Maria"	// string
	const address			// 컴파일 에러

}
