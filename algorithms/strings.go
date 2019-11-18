package algorithms

func ChangeStringElement(in string, position int, symbol rune) string {

r := []rune(in)
r[position] = symbol
return string(r)

}
