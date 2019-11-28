package algorithms

import (
	"sort"
)

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
// result - a list of numbers that appears in the list more than once in the order of their appearance
// The Time Complexity is O(n2) and Space Complexity is O(1)
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

// FindDublicatesSorting - finds dublicates in the list using an sorting first approach
// Sort all the elements in the list and after this in a single scan, we can find the duplicates.
// result - a list of numbers that appears in the list more than once in sorted order
// Sorting algorithms take O(n log n) time and single scan take O(n) time.
// The Time Complexity of an algorithm is O(n log n) and Space Complexity is O(1)
func FindDublicatesSorting(data []int) []int {
	size := len(data)
	result := []int{}
	sort.Ints(data) // Sort(data,size)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			if !SequentialSearch(result, data[i]) { // check if there is already an item as a result to eliminate dublicates
				result = append(result, data[i])
			}
		}
	}
	return result
}

// FindDublicatesHashTable - finds dublicates in the list using hash table
// using Hash-Table, we can keep track of the elements we have already seen and we can find the duplicates in just one scan.
// Hash-Table insert and find take constant time O(1)
// so the total Time Complexity of the algorithm is O(n) time. Space Complexity is also O(n)
// in func below error undefined: Set
// func FindDublicatesHashTable(data []int) {
// 	s := make(Set)
// 	s := make(map[int]int)
// 	size := len(data)
// 	result := []int{}
// 	for i := 0; i < size; i++ {
// 		if s.Find(data[i]) {
// 			fmt.Print(" ", data[i])
// 		} else {
// 			s.Add(data[i])
// 		}
// 	}
// 	return result
// }

// FindDublicatesCounting - finds dublicates in the list using counting
// this approach is only possible if we know the range of the input.
// If we know that, the elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can increase its count. In just one single scan, we know the duplicates. If we know the range of the elements, then this is the fastest way to find the duplicates.
// intrange is the maximum number of the array plus one
// data - is the array of positive integers
func FindDublicatesCounting(data []int, intrange int) []int {
	size := len(data)
	result := []int{}
	count := make([]int, intrange) // create a slice with length intrange

	// for i := 0; i < size; i++ {
	// 	count[i] = 0 // fill count with nulls
	// }

	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			if !SequentialSearch(result, data[i]) { // check if there is already an item as a result to eliminate dublicates
				result = append(result, data[i])
			}
		} else {
			count[data[i]]++
		}
	}
	return result
}

// GetMaxAppearingExhaustive - returns the element that appears maximum number of times in the given list of integers.
func GetMaxAppearingExhaustive(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	return max
}
