package gcd

// gcd(a, b) = gcd(a-b, b)

func GCD_Subtraction(a, b uint32) uint32 {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a;
}
