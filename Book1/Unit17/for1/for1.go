// for1
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Println()

	i3 := 0
	for i3 < 5 {
		fmt.Println(i3)
		i3 = i3 + 1 // i3++
	}

	fmt.Println()

	// for 키워드에 조건식을 설정하지 않으면 무한 루프
	//for {
	//	fmt.Println("Hello, world!")
	//}
}
