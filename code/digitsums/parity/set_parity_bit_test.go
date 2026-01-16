package parity

import "fmt"

func ExampleSetParityBit() {
	fmt.Println(SetParityBit(5))   // 5   = 00000101 -> 00000101 = 5
	fmt.Println(SetParityBit(42))  // 42  = 00101010 -> 10101010 = 170
	fmt.Println(SetParityBit(43))  // 43  = 00101011 -> 00101011 = 43
	fmt.Println(SetParityBit(0))   // 0   = 00000000 -> 00000000 = 0
	fmt.Println(SetParityBit(127)) // 127 = 01111111 -> 11111111 = 255

	// Output:
	// 5
	// 170
	// 43
	// 0
	// 255
}
