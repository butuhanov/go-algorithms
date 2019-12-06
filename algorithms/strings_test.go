package algorithms

import "testing"

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
