package algorithms

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
