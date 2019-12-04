package algorithms

import (
	"fmt"
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

// GetMaxAppearingExhaustive - returns the element that appears maximum number of times in the given list of integers using bruteforce.
// if more than one element is found, the first is returned
// if no items are found, the first is returned
func GetMaxAppearingExhaustive(data []int) (int, int) {
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
	return max, maxCount
}

// GetMaxAppearingSorting - returns the element that appears maximum number of times in the given list of integers using sorting.
// if more than one element is found, the first minimal is returned
// if no items are found, the first minimal is returned
func GetMaxAppearingSorting(data []int) int {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

// GetMaxAppearingCounting - returns the element that appears maximum number of times in the given list of integers using counting.
// If we know that, the elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can increase its count. In just one single scan, we know the duplicates. If we know the range of the elements, then this is the fastest way to find the max count.
func GetMaxAppearingCounting(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}

// GetMajorityExhaustive - finds the majority element, which appears more than n/2 times in given list of elements.
// Return max appearing number and false in case there is no majority element.
// Similarly other approaches can be used
func GetMajorityExhaustive(data []int) (int, bool) {

	max, maxCount := GetMaxAppearingExhaustive(data)

	if maxCount > len(data)/2 {
		return max, true
	}
	return max, false
}

// MooresVotingAlgorithm - implements Moore’s Voting Algorithm, is an algorithm for finding the majority of a sequence of elements using linear time and constant space (see also https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm)  This is a cancelation approach. The elements stand against the majority and each element is cancelled with one element of majority if there is majority then majority prevails.
func MooresVotingAlgorithm(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)

	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	//  Now scan through the list once again to see if that candidate we found above have appeared more than n/2 times.
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	return 0, false
}

// FindMissingNumberExhaustive finds missing number in array using brute force
// Given a list of n-1 elements, which are in the range of 1 to n. There are no duplicates in the list. One of the integer is missing. Find the missing element.
// First loop to select value in the range 1 to n and the second loop to find if this element is in the list or not.
// The Time Complexity is O(n2) and Space Complexity is O(1)
func FindMissingNumberExhaustive(data []int) (int, bool) {
	var found int
	size := len(data)
	for i := 1; i <= size; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if data[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i, true
		}
	}
	return 0, false // fmt.Println("NoNumberMissing")
}

// FindMissingNumberSorting finds missing number in array using sorting
// Sorting algorithms take O(n.logn) time and single scan take O(n) time.
// The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
func FindMissingNumberSorting(data []int) (int, bool) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	for i := 1; i <= (size - 1); i++ {
		if data[i-1] != data[i]-1 {
			return data[i] - 1, true
		}

	}
	return 0, false // fmt.Println("NoNumberMissing")
}

// FindMissingNumberHashTable finds missing number in array using hash table
// Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space Complexity is also O(n)
// TODO

// FindMissingNumberCounting finds missing number in array using counting approach
// we know the range of the input so counting will work. As we know that, the elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can increase its count. In just one single scan, we know the missing element.
// Counting approach just uses a list so insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space Complexity for creating count list is also O(n)
func FindMissingNumberCounting(data []int) (int, bool) {
	size := len(data)
	count := make([]int, size+2) // the range of 1 to n
	// fmt.Println(count)
	for i := 0; i < size; i++ {
		count[data[i]]++
	}
	// fmt.Println(size)
	// fmt.Println(count)
	for i := 1; i < size+1; i++ { // ignore the first and last elements
		// fmt.Println(i, data[i], count[data[i]])
		if count[i] == 0 {
			return i, true
		}
	}
	return 0, false
}

// FindMissingNumberSummation  finds missing number in array using summation approach
// Arithmetic progression . The sum is (a1+an)*n/2
func FindMissingNumberSummation(data []int) (int, bool) {
	var count int
	size := len(data)
	sum := (1 + size + 1) * (size + 1) / 2 // the sum of arithmetic progression
	for i := 0; i < size; i++ {
		count += data[i]
	}
	// fmt.Println(size, sum, count, sum-size-1)
	if sum-size-1 == count {
		return 0, false
	}
	return sum - count, true

}

// FindMissingNumberXOR  finds missing number in array using XOR approach
//  XOR approach to find the sum of n numbers from 1 to n.
// the values stored in the list and you will have your missing number. The Time Complexity of this algorithm is O(n). Space Complexity is O(1)
func FindMissingNumberXOR(data []int) (int, bool) {
	size := len(data)
	xorSumRange := 0
	xorSumNumbers := 0

	// get the XOR of all the numbers from 1 to dataRange
	for i := 1; i <= size; i++ {
		xorSumRange ^= i
	}

	// loop through the array and get the XOR of elements
	for i := 0; i < size; i++ {
		xorSumNumbers ^= data[i]

	}

	if xorSumRange == xorSumNumbers {
		return 0, false
	}
	// fmt.Println(size, xorSumRange, xorSumNumbers)
	return xorSumNumbers, true
}

// TODO
// 1. There are numbers in the range of 1-n out of which all appears single time but one that appear two times.
// 2. All the elements in the range 1-n are appearing 16 times and one element appear 17 times. Find the element that appears 17 times.

//
// FindPairExhaustive - finds a pair using exhaustive search, returns true if the pair exists. In a given list of n numbers, find two elements such that their sum is equal to “value”
// The Time Complexity is O(n2) and Space Complexity is O(1)
func FindPairExhaustive(data []int, value int) bool {
	size := len(data)
	ret := false
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				// fmt.Println("The pair is : ", data[i], ",", data[j])
				ret = true
			}
		}
	}
	return ret
}

// FindPairSorting - finds a pair using sorting approach
func FindPairSorting(data []int, value int) bool {
	size := len(data)
	first := 0
	second := size - 1
	ret := false
	sort.Ints(data) // Sort(data, size)
	for first < second {
		curr := data[first] + data[second]
		if curr == value {
			// fmt.Println("The pair is ", data[first], ",", data[second])
			ret = true
		}
		if curr < value {
			first++
		} else {
			second--
		}
	}
	return ret
}

// FindPairHashTable - finds a pair using hash tables
// using Hash-Table, we can keep track of the elements we have already seen and we can find the pair in just one scan.
// Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space Complexity is also O(n)
// func FindPairHashTable(data []int, value int) bool {
// 	s := make(Set)
// 	size := len(data)
// 	ret := false
// 	for i := 0; i < size; i++ {
// 		if s.Find(value - data[i]) {
// 			fmt.Println(i, "The pair is:", data[i], ",", (value - data[i]))
// 			ret = true
// 		} else {
// 			s.Add(data[i])
// 		}
// 	}
// 	return ret
// }

// FindMinAbsSumPairExhaustive finds two elements whose sum is closest to zero
// Given a List of integers, both +ve and -ve. You need to find the two elements such that their sum is closest to zero.
func FindMinAbsSumPairExhaustive(data []int) (value1, value2 int) {
	var sum int
	size := len(data)
	// Array should have at least two elements
	if size < 2 {
		fmt.Println("InvalidInput")
	}
	// Initialization of values
	minFirst := 0
	minSecond := 1
	minSum := data[0] + data[1]
	if minSum < 0 {
		minSum *= -1
	}
	for l := 0; l < size-1; l++ {
		for r := l + 1; r < size; r++ {
			sum = data[l] + data[r]
			if sum < 0 {
				sum *= -1
			}
			if sum < minSum {
				minSum = sum
				minFirst = l
				minSecond = r
			}
		}
	}
	// fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ",
	// 	data[minSecond])
	return data[minFirst], data[minSecond]
}

// FindMinAbsSumPairSorting finds two elements whose sum is closest to zero using sorting approach
func FindMinAbsSumPairSorting(data []int) (value1, value2 int) {
	var sum int
	size := len(data)
	// Array should have at least two elements
	if size < 2 {
		fmt.Println("InvalidInput")
	}
	sort.Ints(data) // Sort(data, size)
	// Initialization of values
	minFirst := 0
	minSecond := size - 1
	minSum := data[minFirst] + data[minSecond]
	if minSum < 0 {
		minSum *= -1
	}
	for l, r := 0, (size - 1); l < r; {
		sum = (data[l] + data[r])
		if sum < 0 {
			sum *= -1
		}
		if sum < minSum {
			minSum = sum
			minFirst = l
			minSecond = r
		}
		if sum < 0 {
			l++
		} else if sum > 0 {

			r--
		} else {
			break
		}
	}
	// fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", 		data[minSecond])
	return data[minFirst], data[minSecond]
}

// SearchBitonicArrayMax finds maxima in a bitonic list
// A bitonic list comprises of an increasing sequence of integers immediately followed by a decreasing sequence of integers.
func SearchBitonicArrayMax(data []int) (int, bool) {
	size := len(data)
	start := 0
	end := size - 1
	mid := (start + end) / 2
	maximaFound := 0
	if size < 3 {
		// fmt.Println("InvalidInput")
		return 0, false
	}

	for start <= end {
		mid = (start + end) / 2

		// fmt.Println(start, end, mid)

		if mid == 0 { // start of the list
			if data[mid] < data[mid+1] {
				maximaFound = 1
				mid = mid + 1
				break
			} else {
				maximaFound = 1
				break
			}
		}
		if mid == size-1 { // end of the list
			if data[mid] < data[mid-1] {
				maximaFound = 1
				break
			} else {
				maximaFound = 1
				break
			}
		}

		if data[mid-1] < data[mid] && data[mid+1] < data[mid] { //maxima
			maximaFound = 1
			break
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] { // increasing
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] { // decreasing
			end = mid - 1
		} else if data[mid-1] > data[mid] && data[mid] < data[mid+1] { // local minimum
			maximaFound = 1
			if data[0] > data[size-1] {
				mid = 0
				break
			} else {
				mid = size - 1
				break
			}
		} else {
			break
		}
	}

	if maximaFound == 0 {
		// fmt.Println("NoMaximaFound")
		return 0, false
	}
	return mid, true
}

// SearchBitonicArray searches for an item in bitonic list
// Search element in a bitonic list
func SearchBitonicArray(data []int, key int) int {
	size := len(data)
	maxIndex, _ := SearchBitonicArrayMax(data)
	k := BinarySearchBitonic(data, 0, maxIndex, key, true)
	if k != -1 {
		return k
	}
	return BinarySearchBitonic(data, maxIndex+1, size-1, key, false)
}

// BinarySearchBitonic - modifyed recursive binary search algorithm (BinarySearch in list.go)
func BinarySearchBitonic(data []int, start int, end int, key int, isInc bool) int {
	// fmt.Println(start, end)
	if end < start {
		// fmt.Println("end < start")
		if key == data[0] {
			return 0
		}
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] {
		return mid
	}
	if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
		return BinarySearchBitonic(data, start, mid-1, key, isInc)
	}
	return BinarySearchBitonic(data, mid+1, end, key, isInc)

}

// FindFirstSortedIndex - recursivly finds first index of the key in the sorted range of data
func FindFirstSortedIndex(data []int, start int, end int, key int) int {
	// sort.Ints(data) // Sort(data,size)
	// fmt.Println(start, end)
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	// fmt.Println(key, data[mid], mid)
	if key == data[mid] && (mid == start || data[mid-1] != key) {
		return mid
	}

	if key <= data[mid] {
		return FindFirstSortedIndex(data, start, mid-1, key)
	}
	return FindFirstSortedIndex(data, mid+1, end, key)
}

// FindLastSortedIndex - recursivly finds last index of the key in the sorted range of data
func FindLastSortedIndex(data []int, start int, end int, key int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] && (mid == end || data[mid+1] != key) {
		return mid
	}

	if key < data[mid] {
		return FindLastSortedIndex(data, start, mid-1, key)
	}
	return FindLastSortedIndex(data, mid+1, end, key)
}

// OccurrenceCount searches the number of items in a sorted list
func OccurrenceCount(data []int, key int) int {
	size := len(data)
	sort.Ints(data) // Sort(data, size)
	firstIndex := FindFirstSortedIndex(data, 0, size-1, key)
	lastIndex := FindLastSortedIndex(data, 0, size-1, key)
	// fmt.Println(firstIndex, lastIndex)
	if lastIndex == -1 && firstIndex == -1 {
		return 0
	}
	return (lastIndex - firstIndex + 1)
}

// SeperateEvenAndOdd separates even numbers from odd numbers
func SeperateEvenAndOdd(data []int) []int {
	size := len(data)
	left := 0
	right := size - 1
	for left < right {
		// fmt.Println(data[left], data[right], data[left]%2, data[right]%2)
		if data[left]%2 == 0 {
			left++
		} else if data[right]%2 == 1 || data[right]%2 == -1 {
			right--
		} else {
			data[left], data[right] = data[right], data[left] // swap
			left++
			right--
		}
	}
	return data
}

// FindMaxProfit is a solution to the Stock purchase-sell problem
// Bruteforce approach have The Time Complexity is O(n2) and Space Complexity is O(1)
// Second approach is to keep track of the smallest value seen so far from the start. At each point, we can find the difference and keep track of the maximum profit. This is a linear solution.
// The Time Complexity of the algorithm is O(n) time. Space Complexity for creating count list is also O(1)
func FindMaxProfit(stocks []int) (int, int, int) {
	size := len(stocks)
	buy := 0
	sell := 0
	curMin := 0
	currProfit := 0
	maxProfit := 0
	for i := 0; i < size; i++ {
		if stocks[i] < stocks[curMin] {
			curMin = i
		}
		currProfit = stocks[i] - stocks[curMin]
		if currProfit > maxProfit {
			buy = curMin
			sell = i
			maxProfit = currProfit
		}
	}
	// fmt.Println("Purchase day is- ", buy, " at price ", stocks[buy])
	// fmt.Println("Sell day is- ", sell, " at price ", stocks[sell])
	// fmt.Println("Max Profit :: ", maxProfit)

	return buy, sell, maxProfit

}

// FindMedian get median using sorting approach
// See also  https://en.wikipedia.org/wiki/Median
func FindMedian(data []int) int {
	size := len(data)
	sort.Ints(data)
	return data[size/2]
}

// FindMedianTwoLists - finds median of two lists
func FindMedianTwoLists(dataFirst []int, dataSecond []int) int {
	sizeFirst := len(dataFirst)
	sizeSecond := len(dataSecond)
	sort.Ints(dataFirst)
	sort.Ints(dataSecond)
	// cealing function.
	medianIndex := ((sizeFirst + sizeSecond) + (sizeFirst+sizeSecond)%2) / 2
	i := 0
	j := 0
	count := 0
	for count < medianIndex-1 {
		if i < sizeFirst-1 && dataFirst[i] < dataSecond[j] {
			i++
		} else {
			j++
		}
		count++
	}
	if dataFirst[i] < dataSecond[j] {
		return dataFirst[i]
	}
	return dataSecond[j]
}

// FindElementInSorted2DArray searches for given item in 2D-array (two dimenional list)
func FindElementInSorted2DArray(data [][]int, r int, c int, value int) bool {
	//  Start with element at last column and first row
	row := 0
	column := c - 1
	for row < r && column >= 0 {
		if data[row][column] == value { //  the element is the value we are looking for.
			return true
		} else if data[row][column] > value { //  element is greater than the value we are looking for, go to the element at previous column but same row.
			column--
		} else { // element is less than the value we are looking for, go to the element at next row but same column.
			row++
		}
	}
	return false //  the element is not found after reaching the element of the last row  Return false, if the element is not found after reaching the element ofof the first column. Condition (row < r && column >= 0) is false.

}
