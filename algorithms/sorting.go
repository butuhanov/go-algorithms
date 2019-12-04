package algorithms

// BubbleSort sorts an array using a bubble sorting algorithm
func BubbleSort(arr []int, comp func(int, int) bool) []int {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

// BubbleSortModified is modified bubble sorting algorithm
func BubbleSortModified(arr []int, comp func(int, int) bool) []int {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
	return arr
}

// resulting output will be in descending order.
func less(value1 int, value2 int) bool {
	return value1 < value2
}

// sorted output will be increasing in order
func more(value1 int, value2 int) bool {
	return value1 > value2
}
