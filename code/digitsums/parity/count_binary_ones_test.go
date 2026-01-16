package parity

import "fmt"

func ExampleOneCount() {
	fmt.Println(OneCount(5))
	fmt.Println(OneCount(42))
	fmt.Println(OneCount(43))
	fmt.Println(OneCount(0))
	fmt.Println(OneCount(127))

	// Output:
	// 2
	// 3
	// 4
	// 0
	// 7
}
