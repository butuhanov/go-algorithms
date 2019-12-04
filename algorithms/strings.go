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
// · First swap elements in the middle pair
// · Next swap elements in the middle two pairs
// · Next swap elements in the middle three pairs
// · Iterate n-1 steps.
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
