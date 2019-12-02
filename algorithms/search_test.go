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
		{[]int{1, 0, 0}, 0},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -400},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, -321},
	}

	for _, element := range tests {
		v, _ := GetMaxAppearingExhaustive(element.values)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestGetMaxAppearingSorting(t *testing.T) {

	type testelements struct {
		values []int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{1, 1, 2, 2}, 1},
		{[]int{2, 2, 1, 1}, 1},
		{[]int{2, 2, 2, 1}, 2},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 1},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 3},
		{[]int{1, 2, 3, 4, 3, 6, 7, 7, 4, 6, 4}, 4},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 3},
		{[]int{-1, 1}, -1},
		{[]int{0, 0}, 0},
		{[]int{1, 0, 0}, 0},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -400},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, -321},
	}

	for _, element := range tests {
		v := GetMaxAppearingSorting(element.values)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestGetMaxAppearingCounting(t *testing.T) {

	type testelements struct {
		values    []int
		dataRange int
		result    int
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 1}, 4, 3},
		{[]int{1, 2, 3}, 4, 1},
		{[]int{3, 1, 2}, 4, 3},
		{[]int{1, 1, 2, 2}, 3, 1},
		{[]int{2, 2, 1, 1}, 3, 2},
		{[]int{2, 2, 2, 1}, 3, 2},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 8, 1},
		{[]int{6, 1, 1, 1, 1, 1, 1}, 7, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, 1},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 8, 3},
		{[]int{1, 2, 3, 4, 3, 6, 7, 7, 4, 6, 4}, 8, 4},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 9, 3},
		{[]int{1, 1}, 2, 1},
		{[]int{0, 0}, 1, 0},
		{[]int{1, 0, 0}, 2, 0},
		{[]int{400, 321, 1, 110, 145, 234, 300, 780}, 781, 400},
		{[]int{400, 321, 1, 700, 110, 145, 234, 321, 780, 700}, 781, 321},
	}

	for _, element := range tests {
		v := GetMaxAppearingCounting(element.values, element.dataRange)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestGetMajorityExhaustive(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 1, false},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 1, false},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, false},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 3, false},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 3, false},
		{[]int{-1, 1}, -1, false},
		{[]int{0, 0}, 0, true},
		{[]int{1, 0, 0}, 0, true},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, -400, false},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, -321, false},
	}

	for _, element := range tests {
		v, ok := GetMajorityExhaustive(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestMooresVotingAlgorithm(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 0, false},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 0, false},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 0, false},
		{[]int{-1, 1}, 0, false},
		{[]int{0, 0}, 0, true},
		{[]int{1, 0, 0}, 0, true},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, 0, false},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, 0, false},
	}

	for _, element := range tests {
		v, ok := MooresVotingAlgorithm(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindMissingNumberExhaustive(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 0}, 3, true},
		{[]int{6, 8, 5, 9, 4, 7, 11}, 1, true},
		{[]int{1, 2, 3, 5, 6, 7}, 4, true},
		{[]int{1, 2, 3, 4, 8, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11, 13}, 8, true},
		{[]int{1, 2}, 0, false},
		{[]int{1, 4, 3}, 2, true},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v, ok := FindMissingNumberExhaustive(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindMissingNumberSorting(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 0}, 3, true},
		{[]int{6, 8, 5, 9, 4, 7, 11}, 10, true},
		{[]int{1, 2, 3, 5, 6, 7}, 4, true},
		{[]int{1, 2, 3, 4, 8, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11, 13}, 8, true},
		{[]int{1, 2}, 0, false},
		{[]int{1, 4, 3}, 2, true},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v, ok := FindMissingNumberSorting(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindMissingNumberCounting(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 7}, 3, true},
		{[]int{1, 2, 3, 5, 6}, 4, true},
		{[]int{1, 2, 3, 4, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 8, true},
		{[]int{1, 3}, 2, true},
		{[]int{1, 4, 3}, 2, true},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v, ok := FindMissingNumberCounting(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindMissingNumberSummation(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 7}, 3, true},
		{[]int{1, 2, 3, 5, 6}, 4, true},
		{[]int{1, 2, 3, 4, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 8, true},
		{[]int{1, 3}, 2, true},
		{[]int{1, 4, 3}, 2, true},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v, ok := FindMissingNumberSummation(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindMissingNumberXOR(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{2, 1}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 7}, 3, true},
		// {[]int{1, 2, 3, 5, 6}, 4, true},
		{[]int{1, 2, 3, 4, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 8, true},
		{[]int{1, 3}, 2, true},
		// {[]int{1, 4, 3}, 2, true},
		// {[]int{1, 2, 4}, 3, true},
		{[]int{1, 2, 3}, 0, false},
		{[]int{2, 3, 1}, 0, false},
		{[]int{3, 2, 1}, 0, false},
		{[]int{1, 2, 4, 3}, 0, false},
		{[]int{1, 3, 4, 2}, 0, false},
		{[]int{4, 3, 2, 1}, 0, false},
		// {[]int{5, 3, 2, 1}, 4, true},
		// {[]int{1, 5, 3, 4}, 2, true},
	}

	for _, element := range tests {
		v, ok := FindMissingNumberXOR(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestFindPairExhaustive(t *testing.T) {

	type testelements struct {
		values []int
		value  int
		result bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 1, 2, 6, 4, 7, 5}, 0, false},
		{[]int{6, 4, 5, 2, 1, 7}, 3, true},
		{[]int{1, 2, 3, 5, 6}, 4, true},
		{[]int{1, 2, 3, 4, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 7, false},
		{[]int{1, 3}, 2, false},
		{[]int{1, 4, 3}, 2, false},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v := FindPairExhaustive(element.values, element.value)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
