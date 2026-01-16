package digitsums

import "fmt"

func ExampleEANChecksum() {
	isbn_1 := []int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0}
	isbn_2 := []int{9, 7, 8, 0, 1, 3, 7, 9, 3, 5, 1, 0}
	isbn_3 := []int{9, 7, 8, 0, 6, 1, 5, 3, 1, 4, 4, 6}
	isbn_4 := []int{9, 7, 8, 3, 7, 7, 0, 4, 3, 6, 1, 0}

	fmt.Println(EANChecksum(isbn_1)) // 3
	fmt.Println(EANChecksum(isbn_2)) // 9
	fmt.Println(EANChecksum(isbn_3)) // 4
	fmt.Println(EANChecksum(isbn_4)) // 1

	// Output:
	// 3
	// 9
	// 4
	// 1
}

func ExampleEANisValid() {
	isbn_1 := []int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0, 3}
	isbn_2 := []int{9, 7, 8, 0, 1, 3, 7, 9, 3, 5, 1, 0, 9}
	isbn_3 := []int{9, 7, 8, 0, 6, 1, 5, 3, 1, 4, 4, 6, 4}
	isbn_4 := []int{9, 7, 8, 3, 7, 7, 0, 4, 3, 6, 1, 0, 1}

	fmt.Println(EANisValid(isbn_1))
	fmt.Println(EANisValid(isbn_2))
	fmt.Println(EANisValid(isbn_3))
	fmt.Println(EANisValid(isbn_4))

	isbn_invalid := []int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0, 4}
	fmt.Println(EANisValid(isbn_invalid))

	// Output:
	// true
	// true
	// true
	// true
	// false
}
