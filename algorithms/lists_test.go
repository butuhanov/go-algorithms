package algorithms

import "testing"

func TestSumList(t *testing.T) {

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

func TestSequentialSearch(t *testing.T) {

	type testelements struct {
		values []int
		item   int
		result bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, false},
		{[]int{1, 1, 1, 1, 1, 1, -6}, 6, false},
		{[]int{-1, 1}, 0, false},
		{[]int{0, 0}, 0, true},
		{[]int{1, -400, 110, 234, 145, 780, -321, 300}, -321, true},
	}

	for _, element := range tests {
		v := SequentialSearch(element.values, element.item)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
