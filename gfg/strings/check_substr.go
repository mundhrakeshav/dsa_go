package strings

func Check_Substr(s1, s2 string) int {
	res := -1

	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[0] {
			j := 0
			for ; j < len(s1); j++ {
				if s2[i + j] != s1[j] {
					break;
				}
			}
			if j == len(s1) {
				return i
			}
		}
	}

	return res
}
