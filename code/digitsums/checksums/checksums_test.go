package checksums

import "fmt"

//9780 345 39 180-3
//1313 131 31 313 1
//9*1+7*3+8*1+0*3+3*1+4*3+5*1+3*3+9*1+1*3+8*1+0*3+3*1=

func Sum(numbers []int) int {
	result := 0
	for i, n := range numbers {
		//Als Einzeiler: result += (n + 2*n*(i+1)%2)
		if i%2 == 0 {
			result += n
		} else {
			result += 3 * n
		}

	}
	return result
}

//9780 345 39 180-3
//EAN (european Article Number) Prüfziffernberechnung
//EAN erwartet eine Liste von Ziffern und liefert die EAN Prüfziffer dazu.

func Ean(digits []int) int {
	return (10 - (Sum(digits)%10)%10)

}

func ExampleEan() {
	fmt.Println(Ean([]int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0}))

	//Output:
	//3
}
