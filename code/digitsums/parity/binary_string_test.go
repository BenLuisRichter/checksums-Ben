package parity

import "fmt"

func ExampleBinaryString() {
	fmt.Println(BinaryString(5))
	fmt.Println(BinaryString(42))
	fmt.Println(BinaryString(43))
	fmt.Println(BinaryString(0))
	fmt.Println(BinaryString(127))

	// Output:
	// 101
	// 101010
	// 101011
	// 0
	// 1111111
}
