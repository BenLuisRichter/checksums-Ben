package digitsums

import "slices"

// Digits erwartet eine nicht-negative ganze Zahl n als `int`
// und liefert eine Liste der Ziffern dieser Zahl.
func Digits(n int) []int {
	// Hinweis:
	// Verwenden Sie eine Schleife, um die Ziffern zu extrahieren.
	// Sie können die letzte Ziffer mit `n % 10` ermitteln und
	// die letzte Ziffer entfernen, indem Sie `n` durch 10 teilen.
	// Am einfachsten ist es, erstmal die Ziffern in umgekehrter Reihenfolge zu sammeln
	// und die Liste am Ende umzukehren (z.B. mit `slices.Reverse`).
	if n == 0 {
		return []int{0}
	}
	result := []int{}
	for n != 0 {
		lastdigit := n % 10
		n = n / 10
		result = append(result, lastdigit)
	}
	slices.Reverse(result)
	return result
}
