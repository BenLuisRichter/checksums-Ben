package digits

import (
	"fmt"
	"slices"
)

// func BinDigits(n int) byte {
//	return byte(n) //explizite Typumwandlung in byte
//}

// BinDigits erwartet eine Zahl n und liefert eine Liste von Ziffern.
func BinDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 2
		result = append(result, last_digit)
		//result = append({int(last_digit}, result...} //direkt umgekehrt anhängen.
		n /= 2 //n = n / 2
	}

	//Umkehren der Liste
	slices.Reverse(result)

	return result
}

// Digits erwartet eine Zahl n und eine Basis base und liefert eine Liste von Ziffern.
func Digits(n, base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append(result, last_digit)
		n /= base
	}

	slices.Reverse(result)

	return result
}

// Summe erwartet eine Liste von Zahlen und berechnet deren Summe.
func Sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

// ParatiyBit erwartet eine Zahl n und liefert:
// 1: falls die Anzahl der Einsen in der Binärdarstellung von n ungerade ist
// 0: falls die Anzahl der Einsen in der Binärdarstellung von n gerade ist
func ParityBit(n int) int {
	return DigitSum(n, 2) % 2
	//return Sum(Digits(n, 2)) % 2
}

// DigitSum berechnet die Quersumme von n (zur gegebenen Basis).
func DigitSum(n, base int) int {
	return Sum(Digits(n, base))
}

func ExampleDigits() {
	fmt.Println(Digits(42, 2))
	fmt.Println(Digits(42, 16))
	fmt.Println(Digits(42, 10))
	fmt.Println(Digits(42, 8))

	fmt.Println(ParityBit(42))
	fmt.Println(ParityBit(43))

	//Output:
	//[1 0 1 0 1 0]
	//[2 10]
	//[4 2]
	//[5 2]
	//1
	//0
}
