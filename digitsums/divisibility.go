package digitsums

// IsDivisibleBy3 prüft, ob die gegebene Zahl n durch 3 teilbar ist.
//
// Randbedingung: Verwenden Sie NICHT den Modulo-Operator (%).
func IsDivisibleBy3(n int) bool {
	// Hinweis:
	// Eine Zahl ist durch 3 teilbar, wenn ihre Quersumme durch 3 teilbar ist.
	for n >= 10 {
		n = DigitSum(n)
	}
	return n == 0 || n == 3 || n == 6 || n == 9
}
