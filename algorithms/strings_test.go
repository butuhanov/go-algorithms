package algorithms

import (
	"testing"
)

func TestChangeStringElement(t *testing.T) {
	cases := []struct {
		in       string
		position int
		symbol   rune
		want     string
	}{
		{"hello, World!", 0, 'H', "Hello, World!"},
		{"hello", 3, 'H', "helHo"},
		{"привет!", 0, 'П', "Привет!"},
		{"♥♥♥", 1, 'П', "♥П♥"},
		{"hello", 5, 'П', "hello"},
		{"hello", 4, '_', "hell_"},
		{"♥♥♥", 3, 'П', "♥♥♥"},
	}
	for _, c := range cases {
		got := ChangeStringElement(c.in, c.position, c.symbol)
		if got != c.want {
			t.Errorf("ChangeStringElement(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

func TestTransformStringAB1(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"12341234", "11223344"},
		{"12344321", "14233241"},
		{"111222", "121212"},
		{"11122", "11122"},
		{"aaabb", "aaabb"},
		{"aaabbb", "ababab"},
		{"aaabbc", "ababac"},
		{"11112222", "12121212"},
		{"hello, World!", "h eWlolrol,d!"},
		{"hello", "hlelo"},
		{"привет!", "пвреит!"},
		{"♥♥♥", "♥♥♥"},
		{"hello", "hlelo"},
		{"hello", "hlelo"},
		{"♥♥♥", "♥♥♥"},
	}
	for _, c := range cases {
		got := TransformStringAB1(c.in)
		if got != c.want {
			t.Errorf("ChangeStringElement(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

func TestRobinKarp(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		result  int
	}{
		{"hello, World!", "ll", 2},
		{"hello, World!", "!", 12},
		{"hello, World!", " ", 6},
		{"hello, World!", ",", 5},
		// {"hello, World!", "?", -1},
	}
	for _, element := range cases {
		v := RobinKarp(element.text, element.pattern)
		if v != element.result {
			t.Error(
				"For", element.text, "find", element.pattern,
				"expected", element.result,
				"got", v,
			)
		}
	}

}

func TestKMP(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		result  int
	}{
		{"hello, World!", "ll", 2},
		{"hello, World!", "!", 12},
		{"hello, World!", " ", 6},
		{"hello, World!", ",", 5},
		{"hello, World!", "?", -1},
	}
	for _, element := range cases {
		v := KMP(element.text, element.pattern)
		if v != element.result {
			t.Error(
				"For", element.text, "find", element.pattern,
				"expected", element.result,
				"got", v,
			)
		}
	}

}

func TestKMPFindCount(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		result  int
	}{
		{"hello, World!", "ll", 1},
		{"hello, World!", "l", 3},
		{"hello, World!", "!", 1},
		{"heLLo, World!", "L", 2},
		{"hello, World!", ",", 1},
		{"hello, World!", "?", 0},
		{" hello,  World! ", " ", 4},
	}
	for _, element := range cases {
		v := KMPFindCount(element.text, element.pattern)
		if v != element.result {
			t.Error(
				"For", element.text, "find", element.pattern,
				"expected", element.result,
				"got", v,
			)
		}
	}

}

func TestStringTree(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		result  bool
	}{
		{"Test 1", "Aynsh", true},
		{"Test 3", "3", false},
		{"Test 4", "Test 4", true},
		{"Parth", "Test 4", true},
		{"Aynsh", "Test 4", true},
		{"Rayan", "Test 4", true},
		{"Aditya", "Test 4", true},
		{"Pranav", "Test 4", true},
		{"Samar", "Test 4", true},
	}
	tt := new(StringTree)
	for _, element := range cases {
		tt.InsertBSTree(element.text)
	}
	for _, element := range cases {
		v := tt.FindBSTree(element.pattern)
		if v != element.result {
			t.Error(
				"For", element.text, "find", element.pattern,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestOrderMatching(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		result  int
	}{
		{"hello, World!", "ll", 1},
		{"hello, World!", "lord", 1},
		{"hello, World!", "leh", 0},
		{"heLLo, World!", "Ll!", 1},
		{"hello, World!", "W,", 0},
		{"hello, World!", "?", 0},
		{" hello,  World! ", " !", 1},
	}
	for _, element := range cases {
		v := OrderMatching(element.text, element.pattern)
		if v != element.result {
			t.Error(
				"For", element.text, "find", element.pattern,
				"expected", element.result,
				"got", v,
			)
		}
	}

}

func TestCheckUniqueChar(t *testing.T) {
	cases := []struct {
		text   string
		result bool
	}{
		{"ghjdthrf", false},
		{"ghjdtrf", true},
		{"test", false},
		{"tes", true},
	}
	for _, element := range cases {
		v := CheckUniqueChar(element.text)
		if v != element.result {
			t.Error(
				"For", element.text, "find", v,
				"expected", element.result,
				"got", v,
			)
		}
	}

}

func TestCheckPermutation(t *testing.T) {
	cases := []struct {
		string1 string
		string2 string
		result  bool
	}{
		{"test", "test", true},
		{"test", "estt", true},
		{"test", "stte", true},
		{"test", "tets", true},
		{"test", "ttee", false},
	}
	for _, element := range cases {
		v := CheckPermutation(element.string1, element.string2)
		if v != element.result {
			t.Error(
				"For", element.string1, element.string2, "find", v,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
func TestCheckPalindrome(t *testing.T) {
	cases := []struct {
		str    string
		result bool
	}{
		{"test", false},
		{"tset", false},
		{"Sator Arepo Tenet Opera Rotas", false},
		{"sator arepo tenet opera rotas", true},
		{"madam im adam", false},
		{"313", true},
		{"tattarrattat", true},
	}
	for _, element := range cases {
		v := CheckPalindrome(element.str)
		if v != element.result {
			t.Error(
				"For", element.str, "find", v,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

// func TestAnagram(t *testing.T) {
// 	cases := []struct {
// 		str    string
// 		result []string
// 	}{
// 		{"Abc", []string{}},
// 	}

// 	for _, element := range cases {
// 		v := Anagram(element.str)
// 		if !reflect.DeepEqual(v, element.result) {
// 			t.Error(
// 				"For", element.str,
// 				"expected", element.result,
// 				"got", v,
// 			)
// 		}
// 	}
// }

func TestReverseString(t *testing.T) {
	cases := []struct {
		text   string
		result string
	}{
		{"hello", "olleh"},
		{"hello, World", "dlroW ,olleh"},
		{"test", "tset"},
		{"тест", "тсет"},
		{"фы♥♥♥ва", "ав♥♥♥ыф"},
		{"ф♥♥ы♥ва", "ав♥ы♥♥ф"},
	}
	for _, element := range cases {
		v := ReverseString(element.text)
		if v != element.result {
			t.Error(
				"For", element.text,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
func TestReverseWords(t *testing.T) {
	cases := []struct {
		text   string
		result string
	}{
		{"hello", "hello"},
		{"hello, World", "World hello,"},
		// {"проверка тест", "тест проверка"},
	}
	for _, element := range cases {
		v := ReverseWords(element.text)
		if v != element.result {
			t.Error(
				"For", element.text,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
