package algorithms

import "strconv"

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

// secondary functions
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
