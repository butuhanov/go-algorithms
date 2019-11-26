package algorithms

// see also: SequentialSearch, BinarySearch in list.go, BinarySearchRecursive in recursive.go

// LinearSearchSorted - search in increasing sorted list
// when you encounter a greater value element from the increasing sorted list, you stop searching further.
// Time Complexity: O(n). As we need to traverse the complete list in worst case.
// Worst case is when your desired element is at the last position of the sorted list. However, in the average case this algorithm is more efficient even though the growth rate is same as unsorted.
// Space Complexity: O(1). No extra memory is used to allocate the list.
func LinearSearchSorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}

// FindDublicatesExhaustive - finds dublicates in the list using an exhaustive search approach
// result - a list of numbers that appears in the list more than once
func FindDublicatesExhaustive(data []int) []int {
	size := len(data)
	result := []int{}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				if !SequentialSearch(result, data[i]) { // check if there is already an item as a result to eliminate dublicates
					result = append(result, data[i])
				}
				break
			}
		}
	}
	return result
}
