package digitsums

// DigitSum berechnet die Quersumme einer nicht-negativen Ganzzahl n.
func DigitSum(n int) int {
	// Hinweis:
	// Verwenden Sie die Funktion `Digits`, um die Ziffern der Zahl zu erhalten,
	// und summieren Sie diese Ziffern dann in einer Schleife auf.
	digits := Digits(n)
	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum
}
