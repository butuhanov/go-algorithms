package algorithms

import (
	"strings"
)

// ChangeStringElement change the character by its position in the stirng
func ChangeStringElement(in string, position int, symbol rune) string {
	r := []rune(in)
	if len(r) <= position {
		return in
	}
	r[position] = symbol
	return string(r)

}

// TransformStringAB1 -transforms string
// Example: swap elements of a list like [a1 a2 a3 a4 b1 b2 b3 b4] to convert it into [a1 b1 a2 b2 a3 b3 a4 b4]
// Approach:
// First swap elements in the middle pair
// Next swap elements in the middle two pairs
// Next swap elements in the middle three pairs
// Iterate n-1 steps.
func TransformStringAB1(str string) string {
	data := []rune(str)
	size := len(data)
	N := size / 2
	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			data[N-i+2*j], data[N-i+2*j+1] = data[N-i+2*j+1], data[N-i+2*j]
		}
	}
	return string(data)
}

// RobinKarp implements Robin-Karp algorithm
// English words only
// Pattern must exists
func RobinKarp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	prime := 101
	powm := 1
	TextHash := 0
	PatternHash := 0
	i, j := 0, 0
	if m == 0 || m > n {
		return -1
	}
	for i = 0; i < m-1; i++ {
		powm = (powm << 1) % prime
	}
	for i = 0; i < m; i++ {
		PatternHash = ((PatternHash << 1) + (int)(pattern[i])) % prime
		TextHash = ((TextHash << 1) + (int)(text[i])) % prime
	}
	for i = 0; i <= (n - m); i++ {
		if TextHash == PatternHash {
			for j = 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					break
				}
			}
			if j == m {
				return i
			}
		}
		TextHash = (((TextHash - (int)(text[i])*powm) << 1) + (int)(text[i+m])) % prime
		if TextHash < 0 {
			TextHash = (TextHash + prime)
		}
	}
	return -1
}

// KMP implements Knuth-Morris-Pratt algorithm
// English words only
func KMP(text string, pattern string) int {
	i, j := 0, 0
	n := len(text)
	m := len(pattern)
	ShiftArr := make([]int, m+1)
	kmpPreprocess(pattern, ShiftArr)
	for i < n {
		for j >= 0 && text[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		if j == m {
			return (i - m)
		}
	}
	return -1
}

// KMPFindCount finds  the number of occurrences of the pattern in the text using Knuth-Morris-Pratt algorithm
func KMPFindCount(text string, pattern string) int {
	i, j := 0, 0
	count := 0
	n := len(text)
	m := len(pattern)
	ShiftArr := make([]int, m+1)
	kmpPreprocess(pattern, ShiftArr)
	for i < n {
		for j >= 0 && text[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		if j == m {
			count++
			j = ShiftArr[j]
		}
	}
	return count
}

// Node is struct for a tree
type Node struct {
	value  string
	count  int
	lChild *Node
	rChild *Node
}

// StringTree is implementation of binary search tree to store string as key.
type StringTree struct {
	root *Node
}

// InsertBSTree inserts an item into a tree
func (t *StringTree) InsertBSTree(value string) {
	t.root = t.insertUtil(value, t.root)
}

// FindBSTree looking for value in a tree
func (t *StringTree) FindBSTree(value string) bool {
	ret := t.findUtil(t.root, value)
	return ret
}

// OrderMatching finds if the characters of pattern string are in the same order in text string.
func OrderMatching(source string, pattern string) int {
	iSource := 0
	iPattern := 0
	sourceLen := len(source)
	patternLen := len(pattern)
	for iSource = 0; iSource < sourceLen; iSource++ {
		if source[iSource] == pattern[iPattern] {
			iPattern++
		}
		if iPattern == patternLen {
			return 1
		}
	}
	return 0
}

// CheckUniqueChar checks if string contains all unique characters
// for english symbols only
func CheckUniqueChar(str string) bool {
	mp := make(map[byte]int)
	size := len(str)
	for i := 0; i < size; i++ {
		c := str[i]
		if mp[c] != 0 {
			return false
		}
		mp[c] = 1
	}
	return true
}

// CheckPermutation checks if one string is a permutation of the second
func CheckPermutation(s1 string, s2 string) bool {
	count := make(map[byte]int)
	length := len(s1)
	if len(s2) != length {
		return false
	}

	for i := 0; i < length; i++ {
		ch := s1[i]
		count[ch]++
		ch = s2[i]
		count[ch]--
	}
	for i := 0; i < length; i++ {
		ch := s1[i]
		if count[ch] != 0 {
			return false
		}
	}
	return true
}

// CheckPalindrome finds if the string is a palindrome
// see https://en.wikipedia.org/wiki/Palindrome
func CheckPalindrome(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j && str[i] == str[j] {
		i++
		j--
	}
	if i < j {
		return false
	}
	return true
}

func kmpPreprocess(pattern string, ShiftArr []int) {
	m := len(pattern)
	i := 0
	j := -1
	ShiftArr[i] = -1
	for i < m {
		for j >= 0 && pattern[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		ShiftArr[i] = j
	}
}

func (t *StringTree) findUtil(curr *Node, value string) bool {
	var compare int
	if curr == nil {
		return false
	}
	compare = strings.Compare(curr.value, value)
	if compare == 0 {
		return true
	}

	if compare == 1 {
		return t.findUtil(curr.lChild, value)
	}
	return t.findUtil(curr.rChild, value)
}
func (t *StringTree) insertUtil(value string, curr *Node) *Node {
	var compare int
	if curr == nil {
		curr = new(Node)
		curr.value = value
	} else {
		compare = strings.Compare(curr.value, value)
		if compare == 0 {
			curr.count++
		} else if compare == 1 {
			curr.lChild = t.insertUtil(value, curr.lChild)
		} else {
			curr.rChild = t.insertUtil(value, curr.rChild)
		}
	}
	return curr
}
