package digitsums

import "fmt"

func ExampleDigitSum() {
	fmt.Println(DigitSum(5))
	fmt.Println(DigitSum(42))
	fmt.Println(DigitSum(43))
	fmt.Println(DigitSum(0))
	fmt.Println(DigitSum(127))

	// Output:
	// 5
	// 6
	// 7
	// 0
	// 10
}
