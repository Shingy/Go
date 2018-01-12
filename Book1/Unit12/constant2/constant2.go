// constant2
package main

import "fmt"

func main() {
	/**
	 *	Case 1
	 */
	/*
		const Sunday = 0
		const Monday = 1
		const Tuesday = 2
		const Wednesday = 3
		const Thursday = 4
		const Friday = 5
		const Saturday = 6
		const numberOfDays = 7
	*/

	/**
	 *	Case 2
	 */
	/*
		const (
			Sunday       = 0
			Monday       = 1
			Tuesday      = 2
			Wednesday    = 3
			Thursday     = 4
			Friday       = 5
			Saturday     = 6
			numberOfDays = 7
		)
	*/

	/**
	 *	Case 3 : use iota
	 */
	const (
		Sunday       = iota // 0
		Monday              // 1
		Tuesday             // 2
		Wednesday           // 3
		Thursday            // 4
		Friday              // 5
		Saturday            // 6
		numberOfDays        // 7
	)
	// Sunday를 1부터 시작하고 싶다면 iota + 1

	fmt.Println(Thursday)     // 4
	fmt.Println(numberOfDays) // 7

}
