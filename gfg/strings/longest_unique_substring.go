package strings

func Longest_Unique_Substring(str string) int {
	i := 0
	max := 0
	hm := make(map[byte]int, 0)
	j := i
	for j < len(str) {
		if _, exists := hm[str[j]]; exists {
			if (j - i > max ) {max = j - i};
			delete(hm, str[j])
			i++;
		} else {
			hm[str[j]] = 1;
			j++
		}
	}
	return max
}
