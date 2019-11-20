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
		mid = low + (high-low)/2 //  At each step, we reduce our search space by half
		switch {
		case data[mid] == value:
			return true
		case data[mid] < value: //  search the left half of the list
			low = mid + 1
		default:
			high = mid - 1 //  search the right half of the list
		}
	}
	return false
}

// ReverseArray swaps list items from start to end values
func ReverseArray(data []int, start int, end int) []int {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return data
}
