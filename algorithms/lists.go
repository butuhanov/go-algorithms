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

// RotateArrayLeft is rotates the array to the left by k steps, where k is non-negative
// Using Reverse approach that based on the fact that when we rotate the array k times, k elements from the back end of the array come to the front and the rest of the elements from the front shift backwards.
// In this approach, we firstly reverse all the elements of the array. Then, reversing the first k elements followed by reversing the rest n−k elements gives us the required result.
// Time complexity : O(n). n elements are reversed a total of three times.
// Space complexity : O(1). No extra space is used.
// see also:
// https://www.geeksforgeeks.org/array-rotation/
// https://hackernoon.com/fun-with-array-rotations-add4a335d79a
// https://leetcode.com/problems/rotate-array/solution/ (right rotation)
func RotateArrayLeft(data []int, k int) []int {
	n := len(data)
	if k > n {
		k = k % n
	}
	// k = n - k // for right rotation
	data = ReverseArray(data, 0, k-1)
	data = ReverseArray(data, k, n-1)
	data = ReverseArray(data, 0, n-1)
	return data
}

// RotateArrayRight is rotates the array to the right by k steps, where k is non-negative
// https://leetcode.com/problems/rotate-array/
// Example:
// Original List                   : 1 2 3 4 5 6 7
// After reversing all numbers     : 7 6 5 4 3 2 1
// After reversing first k numbers : 5 6 7 4 3 2 1
// After revering last n-k numbers : 5 6 7 1 2 3 4 --> Result
func RotateArrayRight(data []int, k int) []int {
	n := len(data)
	if k > n {
		k = k % n
	}

	// code below works too
	// 	k = n - k
	// data = RotateArrayLeft(data, k)

	data = ReverseArray(data, 0, n-1) // reversing all numbers
	data = ReverseArray(data, k, n-1) // reversing first k numbers
	data = ReverseArray(data, 0, k-1) // reversing last n-k numbers

	return data
}
