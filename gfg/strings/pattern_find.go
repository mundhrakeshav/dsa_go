package strings

func Pattern_find(txt, pat string) []int {
	res := make([]int, 0)
	for i := 0; i < len(txt) - len(pat) + 1; i++ {
		j := 0;
		for ; j < len(pat) && txt[i + j] == pat[j]; j++ {
			
		}
		if j == len(pat) {
			res = append(res, i)
		}
	}
	return res
}
