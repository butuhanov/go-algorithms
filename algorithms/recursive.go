package algorithms

import (
	"fmt"
	"strconv"
)

// Factorial calculates factorial of given number
// Time Complexity is O(N)
func Factorial(i int) int {
	// Termination Condition
	if i <= 1 {
		return 1
	}
	// Body, Recursive Expansion
	return i * Factorial(i-1)
}

// TowersOfHanoi returns a sequence of actions to solve an Tower of Hanoi puzzle (https://en.wikipedia.org/wiki/Tower_of_Hanoi)
func TowersOfHanoi(num int) string {
	return tohUtil(num, "A", "C", "B")
}

// GCD finds the greatest common divisor (largest common factor) using the Euclid’s (Euclidean) algorithm
func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}

// Fibonacci finds a given number in a Fibonacci sequence (https://en.wikipedia.org/wiki/Fibonacci_number)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)

}

// Permutation generates all permutations of an integer list.
// at each recursive call number at index, “i” is swapped with all the numbers that are right of it. Since the number is swapped with all the numbers in its right one by one it will produce all the permutation possible.
func Permutation(data []int, i int, length int) [][]int {
	var s [][]int
	if length == i {
		fmt.Println("data", data)
		s = append(s, data)
		fmt.Println("s", s)
		return s
	}
	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
	return s
}

// BinarySearchRecursive recusively searches for a value in an increasing order sorted list of integers.
// low - minimum index of array, always 0, high maximum index of array, is amount of elements of array - 1, value - desired value
// returns index of the value we are looking for and if that not found returns -1
// Time Complexity: O(logn). Space Complexity: O(logn) for system stack in recursion
func BinarySearchRecursive(data []int, low int, high int, value int) int {

	if low > high {
		return -1
	}
	mid := low + (high-low)/2 // To afunc the overflow
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}

// LLReverseRecurse - recursively revers linked list
func (list *LList) LLReverseRecurse() {
	list.head = list.reverseLLRecurseUtil(list.head, nil)
}

func (list *LList) reverseLLRecurseUtil(currentNode *NodeLL, nextNode *NodeLL) *NodeLL {
	var ret *NodeLL
	if currentNode == nil {
		return nil
	}
	if currentNode.next == nil {
		currentNode.next = nextNode
		return currentNode
	}

	ret = list.reverseLLRecurseUtil(currentNode.next, currentNode)
	currentNode.next = nextNode
	return ret
}

func tohUtil(num int, from string, to string, temp string) string {
	var result string

	if num < 1 {
		return result
	}

	// To move N disks from source to destination
	// we first move N-1 disks from source to temp then move the lowest Nth disk from source to destination
	result += tohUtil(num-1, from, temp, to)
	result += (strconv.Itoa(num) + from + to)
	// Then will move N-1 disks from temp to destination.
	result += tohUtil(num-1, temp, to, from)

	return result

}

func swap(data []int, x int, y int) []int {
	data[x], data[y] = data[y], data[x]
	return data
}
