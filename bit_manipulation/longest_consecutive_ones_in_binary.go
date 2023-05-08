package bitmanipulation

func LongestConsecutiveOnesInBinary(num uint) uint {
	var count uint = 0;

	for num != 0 {
		num = num & (num << 1)
		count++;
	}
	return count;
}