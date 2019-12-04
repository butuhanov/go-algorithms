package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	type testelements struct {
		values []int
		item   func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
	}

	for _, element := range tests {
		v := BubbleSort(element.values, element.item)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestBubbleSortModified(t *testing.T) {

	type testelements struct {
		values []int
		item   func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
	}

	for _, element := range tests {
		v := BubbleSortModified(element.values, element.item)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
