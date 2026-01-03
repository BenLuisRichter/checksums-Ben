package parity

// OneCount erwartet eine Zahl als `byte` und liefert die Anzahl der
// Einsen in der Binärdarstellung dieser Zahl.
func OneCount(b byte) int {
	// Hinweis:
	// Gehen Sie analog zur Funktion `BinaryString` vor, aber zählen Sie
	// dabei die Einsen, statt die Binärdarstellung zu erzeugen.
	count := 0
	for b != 0 {
		lastbit := b % 2
		if lastbit == 1 {
			count++
		}
		b = b / 2
	}
	return count
}
