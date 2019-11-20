package algorithms

// SumList - returns the sum of all the elements of the integer list, list is an input argument.
func SumList(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total

}
