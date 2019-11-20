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
// we do not have information about the structure of the array, so we sequentially look for the element one by one
// This is suitable if the data is not sorted. If the data is sorted, a binary search can be used.
func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

// BinarySearch looking up the value in the sorted array
// We examine the middle position at each step. Depending upon the data that we are searching is greater or smaller than the middle value. We will search either the left or the right portion of the array. At each step, we are eliminating half of the search space there by making this algorithm very efficient.
func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2
		switch {
		case data[mid] == value:
			return true
		case data[mid] < value:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return false
}
