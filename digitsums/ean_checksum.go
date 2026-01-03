package digitsums

// EANChecksum berechnet die EAN-Prüfziffer für eine Liste von Ziffern.
// Annahme: Die Eingabeliste `digits` enthält genau 12 Ziffern.
// Die Prüfziffer nach dem EAN-13-Standard wird wie folgt berechnet:
//  1. Summe der Ziffern an den ungeraden Positionen (1., 3., 5., 7., 9., 11.)
//  2. Summe der Ziffern an den geraden Positionen (2., 4., 6., 8., 10., 12.) multipliziert mit 3
//  3. Addition der beiden Summen
//  4. Die Prüfziffer ist die Zahl, die zu dieser Summe addiert werden muss,
//     um die nächste durch 10 teilbare Zahl zu erreichen.
func EANChecksum(digits []int) int {
	// Hinweis:
	// Die benötigte Ziffernsumme können Sie ähnlich wie in der Funktion `DigitSum`
	// berechnen. Verwenden Sie entweder eine Schleife mit Schrittweite 1 und
	// prüfen Sie die Position (ungerade/gerade) mit dem Modulo-Operator,
	// oder verwenden Sie zwei Schleifen mit Schrittweite 2 und starten Sie
	// einmal bei Index 0 (gerade Positionen) und einmal bei Index 1 (ungerade Positionen).
	// Um die Prüfziffer zu berechnen, muss i.W. diese Summe von 10 subtrahiert werden
	// und dabei der Modulo-Operator genutzt werden, das Ergebnis in den Bereich 0-9 zu bringen.
	digitsum := 0
	for i, digit := range digits {
		digitsum += digit * (1 + 2*(i%2)) // ungerade Positionen: *1, gerade Positionen: *3
	}
	checksum := (10 - (digitsum % 10)) % 10
	return checksum
}

// / EANisValid prüft, ob eine Liste von Ziffern eine gültige EAN-13-Zahl darstellt.
// Annahme: Die Eingabeliste `digits` enthält genau 13 Ziffern.
func EANisValid(digits []int) bool {
	// Hinweis:
	// Verwenden Sie die Funktion `EANChecksum`, um die Prüfziffer zu berechnen.
	// Sie können die Prüfung auf zwei Arten durchführen:
	// 1. Berechnen Sie die Prüfziffer für die ersten 12 Ziffern und vergleichen Sie
	//    sie mit der 13. Ziffer.
	// 2. Berechnen Sie die Prüfziffer für alle 13 Ziffern und prüfen Sie, ob das Ergebnis 0 ist.
	return EANChecksum(digits) == 0
}
