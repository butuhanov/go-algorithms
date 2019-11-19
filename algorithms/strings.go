package algorithms

// ChangeStringElement change the character by its position in the stirng
func ChangeStringElement(in string, position int, symbol rune) string {

r := []rune(in)
r[position] = symbol
return string(r)

}
