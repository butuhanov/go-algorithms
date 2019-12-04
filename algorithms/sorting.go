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

// InsertionSort - implements insertion sort algorithm
func InsertionSort(arr []int, comp func(int, int) bool) []int {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
	return arr
}

// SelectionSort - implements selection sort algorithm
// func SelectionSort(arr []int) []int {
// 	size := len(arr)
// 	var i, j, max int
// 	for i = 0; i < size-1; i++ {
// 		max = 0
// 		for j = 1; j < size-1-i; j++ {
// 			if arr[j] > arr[max] {
// 				max = j
// 			}
// 		}
// 		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
// 	}
// 	return arr
// }

// MergeSort -	Merge sort divide the input into half recursive in each step. It sort the two parts separately recursively and finally combine the result into final sorted output.
func MergeSort(arr []int, comp func(int, int) bool) []int {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1, comp)
	return arr
}

// QuickSort - implements QuickSort algorithm
func QuickSort(arr []int, comp func(int, int) bool) []int {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
	return arr
}

func quickSortUtil(arr []int, lower int, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for comp(arr[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(arr[upper], pivot) && lower <= upper {
			upper--
		}
		if lower < upper {
			swapArray(arr, upper, lower)
		}
	}
	swapArray(arr, upper, start)
	// upper is the pivot position
	quickSortUtil(arr, start, upper-1, comp) // pivot -1 is the upper for left sub array.
	quickSortUtil(arr, upper+1, stop, comp)  // pivot + 1 is the lower for right sub array.
}

func swapArray(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int, comp func(int, int) bool) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex
	count := lowerIndex
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if comp(arr[lowerStart], arr[upperStart]) == false {
			tempArray[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[count] = arr[upperStart]
			upperStart++
		}

		count++
	}
	for lowerStart <= lowerStop {
		tempArray[count] = arr[lowerStart]
		count++
		lowerStart++
	}
	for upperStart <= upperStop {
		tempArray[count] = arr[upperStart]
		count++
		upperStart++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArray[i]
	}
}

func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int, comp func(int, int) bool) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex, comp)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex, comp)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex, comp)
}

// resulting output will be in descending order.
func less(value1 int, value2 int) bool {
	return value1 < value2
}

// sorted output will be increasing in order
func more(value1 int, value2 int) bool {
	return value1 > value2
}
