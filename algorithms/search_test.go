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
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 7, true},
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
func TestFindPairSorting(t *testing.T) {

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
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 7, true},
		{[]int{1, 3}, 2, false},
		{[]int{1, 4, 3}, 2, false},
		{[]int{1, 2, 4}, 3, true},
	}

	for _, element := range tests {
		v := FindPairSorting(element.values, element.value)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindMinAbsSumPairExhaustive(t *testing.T) {

	type testelements struct {
		values  []int
		result1 int
		result2 int
	}

	var tests = []testelements{
		{[]int{1, 2}, 1, 2},
		{[]int{3, 1, 2, 6, -4, 7, 5}, 3, -4},
		{[]int{6, 4, -5, -2, 1, 7}, 6, -5},
		{[]int{1, 2, -3, 5, 6}, 2, -3},
		{[]int{1, 2, 3, 4, 6, 7}, 1, 2},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 2, 1},
		{[]int{223, 22, 112, 224, 225, 72, 92, 26, 102, 121}, 22, 26},
		{[]int{1, 3}, 1, 3},
		{[]int{1, 4, 3}, 1, 3},
		{[]int{1, 2, 4}, 1, 2},
	}

	for _, element := range tests {
		res1, res2 := FindMinAbsSumPairExhaustive(element.values)
		if res1 != element.result1 || res2 != element.result2 {
			t.Error(
				"For", element.values,
				"expected", element.result1, element.result2,
				"got", res1, res2,
			)
		}
	}
}

func TestFindMinAbsSumPairSorting(t *testing.T) {

	type testelements struct {
		values  []int
		result1 int
		result2 int
	}

	var tests = []testelements{
		{[]int{1, 2}, 1, 2},
		{[]int{3, 1, 2, 6, -4, 7, 5}, -4, 5},
		{[]int{6, 4, -5, -2, 1, 7}, -5, 6},
		{[]int{1, 2, -3, 5, 6}, -3, 2},
		{[]int{1, 2, 3, 4, 6, 7}, 1, 2},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 1, 2},
		{[]int{223, 22, 112, 224, 225, 72, 92, 26, 102, 121}, 22, 26},
		{[]int{1, 3}, 1, 3},
		{[]int{1, 4, 3}, 1, 3},
		{[]int{1, 2, 4}, 1, 2},
	}

	for _, element := range tests {
		res1, res2 := FindMinAbsSumPairSorting(element.values)
		if res1 != element.result1 || res2 != element.result2 {
			t.Error(
				"For", element.values,
				"expected", element.result1, element.result2,
				"got", res1, res2,
			)
		}
	}
}
func TestSearchBitonicArrayMax(t *testing.T) {

	type testelements struct {
		values     []int
		resultInt  int
		resultBool bool
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, false},
		{[]int{3, 5, 10, 4, 1}, 2, true},
		{[]int{-3, -1, 2, 4, 6, 12}, 5, true},
		{[]int{13, 12, 10, 4, 1, -2, -4}, 0, true},
		{[]int{10, 14, 5, 1, 0}, 1, true},
		{[]int{1, 2, 3, 0, -1}, 2, true},
		{[]int{1, 2, 3, 5, 6, 10, 11, 4, 2, -1}, 6, true},
		{[]int{1, 2, 3, 4, 6, 7}, 5, true},
		{[]int{3, 2, 1, 4, 5, 7, 9}, 6, true},
		{[]int{223, 22, 11, 224, 225}, 4, true},
		{[]int{1, 3}, 0, false},
		{[]int{1, 4, 3}, 1, true},
		{[]int{1, 2, 4}, 2, true},
	}

	for _, element := range tests {
		v, ok := SearchBitonicArrayMax(element.values)
		if v != element.resultInt || ok != element.resultBool {
			t.Error(
				"For", element.values,
				"expected", element.resultInt, element.resultBool,
				"got", v, ok,
			)
		}
	}
}

func TestSearchBitonicArray(t *testing.T) {

	type testelements struct {
		values []int
		key    int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, -1},
		{[]int{3, 5, 10, 4, 1}, 2, -1},
		{[]int{-3, -1, 2, 4, 6, 12}, 5, -1},
		{[]int{-3, -1, 2, 4, 6, 12}, 12, 5},
		{[]int{-3, -1, 2, 4, 6, 12}, -3, 0},
		{[]int{13, 12, 10, 4, 1, -2, -4}, 10, 2},
		{[]int{10, 14, 5, 1, 0}, 1, 3},
		{[]int{1, 2, 3, 0, -1}, 2, 1},
		{[]int{1, 2, 3, 5, 6, 10, 11, 4, 2, -1}, 6, 4},
		{[]int{1, 2, 3, 4, 6, 7}, 5, -1},
		{[]int{3, 2, 1, 4, 5, 7, 9}, 6, -1},
		{[]int{223, 22, 11, 224, 225}, 4, -1},
		{[]int{223, 22, 11, 224, 225}, 223, 0},
		{[]int{4, 3, 2, 1, 2, 3, 4, 5}, 4, 6},
		{[]int{223, 22, 11, 224, 225}, 225, 4},
		{[]int{1, 3}, 0, -1},
		{[]int{1, 4, 3}, 1, 0},
		{[]int{1, 2, 4}, 2, 1},
	}

	for _, element := range tests {
		v := SearchBitonicArray(element.values, element.key)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
func TestFindFirstSortedIndex(t *testing.T) {

	type testelements struct {
		values []int
		start  int
		end    int
		key    int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, 1, 3, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 9, 9, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 7, 9, -1},
		{[]int{1, 1, 1, 4, 5, 6, 7, 8, 9, 10}, 0, 7, 1, 0},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 0, 7, 4, 3},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 5, 7, 4, 5},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 72, 6},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 85, 7},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 97, -1},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 5, 7, 43, -1},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 3, 7, 43, 3},
	}

	for _, element := range tests {
		v := FindFirstSortedIndex(element.values, element.start, element.end, element.key)
		if v != element.result {
			t.Error(
				"For", element.values, "find", element.key,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindLastSortedIndex(t *testing.T) {

	type testelements struct {
		values []int
		start  int
		end    int
		key    int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, 1, 3, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 9, 9, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 7, 9, -1},
		{[]int{1, 1, 1, 4, 5, 6, 7, 8, 9, 10}, 0, 7, 1, 2},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 0, 7, 4, 7},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 5, 6, 4, 6},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 72, 6},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 85, 7},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 0, 7, 97, -1},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 5, 7, 43, -1},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 3, 7, 43, 3},
	}

	for _, element := range tests {
		v := FindLastSortedIndex(element.values, element.start, element.end, element.key)
		if v != element.result {
			t.Error(
				"For", element.values, "find", element.key,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestOccurrenceCount(t *testing.T) {

	type testelements struct {
		values []int
		key    int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 3, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1},
		{[]int{1, 1, 1, 4, 5, 6, 7, 8, 9, 10}, 1, 3},
		{[]int{10, 1, 3, 4, 15, 64, 1, 28, 19, 1}, 1, 3},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 4, 5},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, 5, 0},
		{[]int{1, 1, 1, 4, 4, 2, 4, 4, 9, 10, 1, 12, 33, -3, 112, 44, -323, -22, 12, -3, 4}, 4, 5},
		{[]int{1, 1, 1, 4, 4, 2, 4, 4, 9, 10, 1, 12, 33, -3, 112, 44, -323, -22, 12, -3, 4}, 1, 4},
		{[]int{1, 1, 1, 4, 4, 2, 4, 4, 9, 10, 1, 12, 33, -3, 112, 44, -323, -22, 12, -3, 4}, -3, 2},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 72, 1},
		{[]int{1, 22, 132, 433, 54, 65, 172, 815, 197, 132}, 132, 2},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, 85, 1},
	}

	for _, element := range tests {
		v := OccurrenceCount(element.values, element.key)
		if v != element.result {
			t.Error(
				"For", element.values, "find", element.key,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestSeperateEvenAndOdd(t *testing.T) {

	type testelements struct {
		values []int
		result []int
	}

	var tests = []testelements{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, -3, 4, 5, -6}, []int{-6, 2, 4, -3, 5, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 2, 8, 4, 6, 5, 7, 3, 9, 1}},
		{[]int{10, 1, 3, 4, 15, 64, 1, 28, 19, 1}, []int{10, 28, 64, 4, 15, 3, 1, 1, 19, 1}},
		{[]int{1, 2, 3, 4, 4, 4, 4, 4, 9, 10}, []int{10, 2, 4, 4, 4, 4, 4, 3, 9, 1}},
		{[]int{1, 1, 1, 4, 4, 2, 4, 10, 1, 12, 33, -3, 112, 44, -323, -22, 12, -3, 4}, []int{4, 12, -22, 4, 4, 2, 4, 10, 44, 12, 112, -3, 33, 1, -323, 1, 1, -3, 1}},
		{[]int{1, 2, 31, 43, 54, 65, 72, 85, 97, 132}, []int{132, 2, 72, 54, 43, 65, 31, 85, 97, 1}},
		{[]int{1, 22, 132, 433, 54, 65, 172, 815, 197, 132}, []int{132, 22, 132, 172, 54, 65, 433, 815, 197, 1}},
	}

	for _, element := range tests {
		v := SeperateEvenAndOdd(element.values)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindMaxProfit(t *testing.T) {

	type testelements struct {
		stocks []int
		buy    int
		ack    int
		profit int
	}

	var tests = []testelements{
		{[]int{1, 2}, 0, 1, 1},
		{[]int{3, 1, 2, 6, -4, 7, 5}, 4, 5, 11},
		{[]int{6, 4, -5, -2, 1, 7}, 2, 5, 12},
		{[]int{1, 2, -3, 5, 6}, 2, 4, 9},
		{[]int{1, 2, 3, 4, 6, 7}, 0, 5, 6},
		{[]int{3, 2, 1, 4, 5, 7, 9, 6, 10, 11}, 2, 9, 10},
		{[]int{223, 22, 112, 224, 225, 72, 92, 26, 102, 121}, 1, 4, 203},
		{[]int{1, 3}, 0, 1, 2},
		{[]int{1, 4, 3}, 0, 1, 3},
		{[]int{1, 2, 4}, 0, 2, 3},
		{[]int{23, 22, 11, 224, 225, 72, 1, 26, 102, 121}, 2, 4, 214},
	}

	for _, element := range tests {
		res1, res2, res3 := FindMaxProfit(element.stocks)
		if res1 != element.buy || res2 != element.ack || res3 != element.profit {
			t.Error(
				"For", element.stocks,
				"expected", element.buy, element.ack, element.profit,
				"got", res1, res2, res3,
			)
		}
	}
}

func TestFindMedian(t *testing.T) {

	type testelements struct {
		values []int
		result int
	}

	var tests = []testelements{
		{[]int{1, 2}, 2},
		{[]int{3, 2, 1}, 2},
		{[]int{1, 2, 3}, 2},
		{[]int{3, 1, 2}, 2},
		{[]int{1, 1, 2, 2}, 2},
		{[]int{2, 2, 1, 1}, 2},
		{[]int{2, 2, 2, 1}, 2},
		{[]int{3, 1, 2, 1, 4, 7, 1}, 2},
		{[]int{-6, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
		{[]int{1, 2, 3, 4, 3, 6, 7}, 3},
		{[]int{1, 2, 3, 4, 3, 6, 7, 7, 4, 6, 4}, 4},
		{[]int{3, 2, 1, 4, 3, 6, 7, 8, 1, 6, 3, 0}, 3},
		{[]int{-1, 1}, 1},
		{[]int{0, 0}, 0},
		{[]int{1, 0, 0}, 0},
		{[]int{-400, -321, 1, 110, 145, 234, 300, 780}, 145},
		{[]int{-400, -321, 1, 700, -110, 145, -234, -321, 780, 700}, 1},
		{[]int{11, 9, 3, 5, 5}, 5},
		{[]int{1, 3, 5, 7}, 5},
		{[]int{3, 5}, 5},
		{[]int{1, 3, 3, 6, 7, 8, 9}, 6},
	}

	for _, element := range tests {
		v := FindMedian(element.values)
		if v != element.result {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestFindMedianTwoLists(t *testing.T) {

	type testelements struct {
		values1 []int
		values2 []int
		result  int
	}

	var tests = []testelements{
		{[]int{1, 2}, []int{1, 2}, 1},
		{[]int{1, 2, 4, 5, 2, 5, 4, 8, 9, 3, 4}, []int{1, 2, 3, 4, 5}, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 1, 1, 1, 3}, 3},
	}

	for _, element := range tests {
		v := FindMedianTwoLists(element.values1, element.values2)
		if v != element.result {
			t.Error(
				"For", element.values1, element.values2,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
