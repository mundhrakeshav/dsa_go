package recursionbacktracking

// Incomplete
func RemoveAdjacentDuplicates(inputString string) string {
	if len(inputString) <= 1 {
		return inputString
	}
	
	if inputString[0] == inputString[1] {
		// lastRemoved := inputString[0]; 
		for len(inputString) > 1 && inputString[0] == inputString[1] {
			return RemoveAdjacentDuplicates(inputString[1:])
		}
	}

	// rem_str := RemoveAdjacentDuplicates(inputString[1:])

	return inputString
}


// c aaa bbb aa c dddd