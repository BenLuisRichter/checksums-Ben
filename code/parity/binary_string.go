package parity

// BinaryString erwartet eine Zahl als `byte und liefert einen
// String mit der Binärdarstellung dieser Zahl.
func BinaryString(b byte) string {
	result := ""
	if b == 0 {
		return "0"
	}
	for b > 0 {
		if b%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		b /= 2
	}
	return result
}

// 0x42 = Hexadezimalzahl für 42
// 42 = 01010002 = 2^1+2^3+2^5
// 42%2 = 0
// 42/2 = 21
