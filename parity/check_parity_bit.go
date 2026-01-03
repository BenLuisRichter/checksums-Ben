package parity

// CheckParityBit erwartet eine Zahl als `byte` und prüft, ob die Gesamtzahl
// der Einsen in der Binärdarstellung gerade ist.
func CheckParityBit(b byte) bool {
	// Hinweis:
	// Verwenden Sie die Funktion `OneCount`, um die Anzahl der Einsen zu ermitteln.
	// Wenn die Anzahl gerade ist, liefern Sie `true`, andernfalls `false`.
	return OneCount(b)%2 == 0
}
