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

 // SequentialSearch - look up the value in unknown array
func SequentialSearch(data []int, value int) bool {
size := len(data)
for i := 0; i < size; i++ {
if value == data[i] {
return true
}
}
return false
}
