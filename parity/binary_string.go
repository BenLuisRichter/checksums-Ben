package parity

import "fmt"

// BinaryString erwartet eine Zahl als `byte und liefert einen
// String mit der Binärdarstellung dieser Zahl.
func BinaryString(b byte) string {
	// Hinweis:
	// Verwenden Sie eine Hilfsvariable (z.B. `result`), um die Binärdarstellung
	// aufzubauen. Nutzen Sie eine Schleife, um die Binärziffern zu ermitteln.
	// Sie können die letzte Ziffer bestimmen, indem Sie `b % 2` berechnen,
	// und Sie können die letzte Ziffer entfernen, indem Sie `b` durch 2 teilen.`
	if b == 0 {
		return "0"
	}
	result := ""
	for b != 0 {
		lastbit := b % 2
		b = b / 2
		result = fmt.Sprintf("%d", lastbit) + result
	}
	return result
}
