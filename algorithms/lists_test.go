package algorithms

import "testing"

type testpair struct {
	values []int
	result int
}

var tests = []testpair{
	{[]int{1, 2}, 3},
	{[]int{1, 1, 1, 1, 1, 1}, 6},
	{[]int{-1, 1}, 0},
	{[]int{0, 0}, 0},
}

func TestSumList(t *testing.T) {

	for _, pair := range tests {
		v := SumList(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
