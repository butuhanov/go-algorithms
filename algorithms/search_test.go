package algorithms

import (
	"reflect"
	"testing"
)

func TestLinearSearchSorted(t *testing.T) {

	type testelements struct {
		values []int
		item   int
		result bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, false},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 6, false},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, false},
		{[]int{-1, 1}, 0, false},
		{[]int{0, 0}, 0, true},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -321, true},
	}

	for _, element := range tests {
		v := LinearSearchSorted(element.values, element.item)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
func TestFindDublicatesExhaustive(t *testing.T) {

	type testelements struct {
		values []int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, []int{}},
		{[]int{3, 1, 2, 1, 4, 7, 1}, []int{1}},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{}},
		{[]int{1, 2, 3, 4, 3, 6, 7}, []int{3}},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, []int{3, 1, 6}},
		{[]int{-1, 1}, []int{}},
		{[]int{0, 0}, []int{0}},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, []int{}},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, []int{-321, 700}},
	}

	for _, element := range tests {
		v := FindDublicatesExhaustive(element.values)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindDublicatesSorting(t *testing.T) {

	type testelements struct {
		values []int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, []int{}},
		{[]int{3, 1, 2, 1, 4, 7, 1}, []int{1}},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{}},
		{[]int{1, 2, 3, 4, 3, 6, 7}, []int{3}},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, []int{1, 3, 6}},
		{[]int{-1, 1}, []int{}},
		{[]int{0, 0}, []int{0}},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, []int{}},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, []int{-321, 700}},
	}

	for _, element := range tests {
		v := FindDublicatesSorting(element.values)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindDublicatesCounting(t *testing.T) {

	type testelements struct {
		values   []int
		intrange int
		result   []int
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, []int{}},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 8, []int{1}},
		{[]int{6, 1, 1, 1, 1, 1, 1}, 7, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, []int{}},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 8, []int{3}},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 9, []int{3, 1, 6}},
		{[]int{1, 1}, 2, []int{1}},
		{[]int{0, 0}, 2, []int{0}},
		{[]int{400, 321, 1, 110, 145, 234, 300, 780}, 781, []int{}},
		{[]int{400, 321, 1, 700, 110, 145, 234, 321, 780, 700}, 781, []int{321, 700}},
	}

	for _, element := range tests {
		v := FindDublicatesCounting(element.values, element.intrange)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestGetMaxAppearingExhaustive(t *testing.T) {

	type testelements struct {
		values []int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 1},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 1},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 3},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 3},
		{[]int{-1, 1}, -1},
		{[]int{0, 0}, 0},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -400},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, -321},
	}

	for _, element := range tests {
		v := GetMaxAppearingExhaustive(element.values)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
