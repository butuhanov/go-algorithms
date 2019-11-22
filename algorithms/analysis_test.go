package algorithms

import (
	"testing"
)

func TestAnalyse(t *testing.T) {

	type testdata struct {
		function func(int) int
		value    int
		result   int
	}

	var tests = []testdata{
		{fun1, 100, 100},
		{fun2, 100, 10000},
		{fun3, 100, 4950},
		{fun4, 100, 7},
		{fun5, 100, 7},
		{fun6, 100, 1000000},
		{fun7, 100, 20000},
		{fun8, 100, 1000},
		{fun9, 100, 197},
		{fun10, 100, 4950},
		{fun11, 100, 166650},
		{fun12, 100, 100},
		{fun13, 100, 134},
	}

	for _, data := range tests {
		v := Analyse(data.function, data.value)
		if v != data.result {
			t.Error(
				"For", data.value,
				"expected", data.result,
				"got", v,
			)
		}
	}
}
