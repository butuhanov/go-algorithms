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
	}
	for _, c := range cases {
		got := ChangeStringElement(c.in, c.position, c.symbol)
		if got != c.want {
			t.Errorf("ChangeStringElement(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}
