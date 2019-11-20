package algorithms

import (
	"reflect"
	"testing"
)

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

func TestBinarySearch(t *testing.T) {

	type testelements struct {
		values []int
		item   int
		result bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, false},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 6, false},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, false},
		{[]int{-1, 1}, 0, false},
		{[]int{0, 0}, 0, true},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -321, true},
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

func TestReverseArray(t *testing.T) {

	type testelements struct {
		values []int
		start  int
		end    int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, 1, []int{2, 1}},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 0, 6, []int{1, 1, 1, 1, 1, 1, -6}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, 6, []int{7, 6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 5, []int{1, 6, 5, 4, 3, 2, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, 5, []int{1, 2, 3, 6, 5, 4, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, 4, []int{1, 2, 3, 5, 4, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 4, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{-1, 1}, 0, 1, []int{1, -1}},
		{[]int{-1, 1}, 1, 1, []int{-1, 1}},
		{[]int{0, 0}, 0, 1, []int{0, 0}},
	}

	for _, element := range tests {
		v := ReverseArray(element.values, element.start, element.end)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestRotateArrayLeft(t *testing.T) {

	type testelements struct {
		values []int
		steps  int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2}, 4, []int{1, 2}},
		{[]int{1, 2}, 5, []int{2, 1}},

		{[]int{1}, 2, []int{1}},
		{[]int{-1}, 2, []int{-1}},
		{[]int{0}, 2, []int{0}},

		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{5, 1, 2, 3, 4}},

		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{10, 20, 30, 40, 50, 60}, 2, []int{30, 40, 50, 60, 10, 20}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, []int{3, 4, 5, 6, 7, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3}},
	}

	for _, element := range tests {
		v := RotateArrayLeft(element.values, element.steps)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestRotateArrayRight(t *testing.T) {

	type testelements struct {
		values []int
		steps  int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2}, 4, []int{1, 2}},
		{[]int{1, 2}, 5, []int{2, 1}},

		{[]int{1}, 2, []int{1}},
		{[]int{-1}, 2, []int{-1}},
		{[]int{0}, 2, []int{0}},

		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{2, 3, 4, 5, 1}},

		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{10, 20, 30, 40, 50, 60}, 2, []int{50, 60, 10, 20, 30, 40}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, []int{10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, element := range tests {
		v := RotateArrayRight(element.values, element.steps)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestMaxSubArraySum(t *testing.T) {

	type testelements struct {
		values []int
		result int
	}

	var tests = []testelements{
		{[]int{0, 1, 1, 1, 0, 1, 2, 1, 2}, 9},
		{[]int{0, 1, -1, -1, 0, 1, -2, 1, -2}, 1},
		{[]int{5, -1, -1, -1, 0, 1, 2, 1, -2}, 6},
		{[]int{-1, -1, -1, -1, 0, -2, -2, -4, -2}, 0},
		{[]int{-3, -1, -1, -1, -5, -2, -2, -4, -2}, 0},
		{[]int{1, -2, 3, 4, -4, 6, -14, 8, 2}, 10},
		{[]int{1, 2, -1, -1, 0, -1, 4, -1, -2}, 4},
		{[]int{1, 2, -1, -1, 0, -1, 2, -1, -2}, 3},
	}

	for _, element := range tests {
		v := MaxSubArraySum(element.values)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
