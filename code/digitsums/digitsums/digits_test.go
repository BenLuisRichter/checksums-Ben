package digitsums

import (
	"fmt"
)

func ExampleDigits() {
	fmt.Println(Digits(5))
	fmt.Println(Digits(42))
	fmt.Println(Digits(43))
	fmt.Println(Digits(0))
	fmt.Println(Digits(127))

	// Output:
	// [5]
	// [4 2]
	// [4 3]
	// [0]
	// [1 2 7]
}
