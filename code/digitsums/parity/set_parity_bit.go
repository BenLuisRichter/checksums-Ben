package parity

// SetParityBit erwartet eine Zahl als `byte` und liefert ein neues `byte`.
// Das neue `byte` enthält die ursprünglichen Bits, wobei das höchstwertige
// Bit so gesetzt ist, dass die Gesamtzahl der Einsen in der Binärdarstellung gerade
// ist. Dabei wird angenommen, dass die ursprüngliche Zahl nur die unteren 7 Bits
// nutzt, d.h. dass das höchste Bit (Bit 7) immer 0 ist.
//
// Beispiele siehe Tests.
func SetParityBit(b byte) byte {
	// TODO
	return 0
}
