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
	// TODO
	return 0
}

// / EANisValid prüft, ob eine Liste von Ziffern eine gültige EAN-13-Zahl darstellt.
// Annahme: Die Eingabeliste `digits` enthält genau 13 Ziffern.
func EANisValid(digits []int) bool {
	// TODO
	return false
}
