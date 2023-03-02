package strings

func Reverse_Words(str string) (res []string) {
	lastBreak := len(str) - 1
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == string(" ") {
			res = append(res, str[i:lastBreak])
			lastBreak = i;
		}
	}
	res = append(res, str[:lastBreak])
	return res;
}