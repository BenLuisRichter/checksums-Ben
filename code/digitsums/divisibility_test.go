package digitsums

import "fmt"

func ExampleIsDivisibleBy3() {
	fmt.Println(IsDivisibleBy3(9))   // true
	fmt.Println(IsDivisibleBy3(10))  // false
	fmt.Println(IsDivisibleBy3(123)) // true
	fmt.Println(IsDivisibleBy3(124)) // false
	fmt.Println(IsDivisibleBy3(0))   // true

	// Output:
	// true
	// false
	// true
	// false
	// true
}
