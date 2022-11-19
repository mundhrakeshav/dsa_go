package gcd

// gcd(a, b) = gcd(a*b, b)

func swap(a, b *uint32) {
	*a, *b = *b, *a
}

func GCD_Division_Iterative(a, b uint32) uint32 {
	for b != 0 {
		a = a % b
		swap(&a, &b)
	}
	return a
}

func GCD_Division_Recursive(a, b uint32) uint32 {
	if b == 0 {
		return a
	} else {
		return GCD_Division_Recursive(b, a%b)
	}
}
