// constant3
package main

import "fmt"

func main() {
	const (
		a = 1 << iota // a == 1 (1 << 0)
		b = 1 << iota // b == 2 (1 << 1)
		c = 1 << iota // c == 4 (1 << 2)
		d = 1 << iota // d == 8 (1 << 3)
	)

	fmt.Println(b) // 2
	fmt.Println(d) // 8

	const (
		aa = iota * 30 // a == 0 (0 * 30)
		bb = iota * 30 // b == 30 (1 * 30)
		cc = iota * 30 // c == 60 (2 * 30)
		dd = iota * 30 // d == 90 (3 * 30)
	)

	fmt.Println(aa) // 0
	fmt.Println(cc) // 60

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // iota == 2를 생략하여 bit2와 mask2 생략
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)

	fmt.Println(bit1)  // 2
	fmt.Println(mask3) // 7

}
