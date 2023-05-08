package recursionbacktracking

// Incomplete
func RemoveAdjacentDuplicates(inputString string) string {
	len_input := len(inputString);
	if len_input <= 1 {
		return inputString
	}
	
	if inputString[0] == inputString[1] {
		// lastRemoved := inputString[0]; 
		for len(inputString) > 1 && inputString[0] == inputString[1] {
			return RemoveAdjacentDuplicates(inputString[1:])
		}
	}
	return inputString
}


// c aaa bbb aa c dddd