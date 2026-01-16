package parity

import "fmt"

func ExampleCheckParityBit() {
	fmt.Println(CheckParityBit(5))   // 5   = 00000101 -> 2 Einsen -> true
	fmt.Println(CheckParityBit(42))  // 42  = 00101010 -> 3 Einsen -> false
	fmt.Println(CheckParityBit(43))  // 43  = 00101011 -> 4 Einsen -> true
	fmt.Println(CheckParityBit(0))   // 0   = 00000000 -> 0 Einsen -> true
	fmt.Println(CheckParityBit(127)) // 127 = 01111111 -> 7 Einsen -> false

	// Output:
	// true
	// false
	// true
	// true
	// false
}
