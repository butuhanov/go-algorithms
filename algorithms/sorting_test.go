package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	type testelements struct {
		values []int
		method func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, less, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, less, []int{9, 8, 6, 5, 3, 2, 1, -4, -7}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, less, []int{8, 6, 5, 2, 1, 1, 1, -4, -7}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, less, []int{9, 8, 7, 4, 3, 3, 0, -1, -6}},
	}

	for _, element := range tests {
		v := BubbleSort(element.values, element.method)
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
		method func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, less, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, less, []int{9, 8, 6, 5, 3, 2, 1, -4, -7}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, less, []int{8, 6, 5, 2, 1, 1, 1, -4, -7}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, less, []int{9, 8, 7, 4, 3, 3, 0, -1, -6}},
	}

	for _, element := range tests {
		v := BubbleSortModified(element.values, element.method)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestInsertionSort(t *testing.T) {

	type testelements struct {
		values []int
		method func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, less, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, less, []int{9, 8, 6, 5, 3, 2, 1, -4, -7}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, less, []int{8, 6, 5, 2, 1, 1, 1, -4, -7}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, less, []int{9, 8, 7, 4, 3, 3, 0, -1, -6}},
	}

	for _, element := range tests {
		v := InsertionSort(element.values, element.method)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

// func TestSelectionSort(t *testing.T) {

// 	type testelements struct {
// 		values []int
// 		// method   func(int, int) bool
// 		result []int
// 	}

// 	var tests = []testelements{
// 		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
// 		{[]int{9, 5, 1, 2, 7, 4, 6, 24, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},

// 	}

// 	for _, element := range tests {
// 		v := SelectionSort(element.values)
// 		if !reflect.DeepEqual(v, element.result) {
// 			t.Error(
// 				"For", element.values,
// 				"expected", element.result,
// 				"got", v,
// 			)
// 		}
// 	}
// }

func TestMergeSort(t *testing.T) {

	type testelements struct {
		values []int
		method func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, less, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, less, []int{9, 8, 6, 5, 3, 2, 1, -4, -7}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, less, []int{8, 6, 5, 2, 1, 1, 1, -4, -7}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, less, []int{9, 8, 7, 4, 3, 3, 0, -1, -6}},
	}

	for _, element := range tests {
		v := MergeSort(element.values, element.method)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}

func TestQuickSort(t *testing.T) {

	type testelements struct {
		values []int
		method func(int, int) bool
		result []int
	}

	var tests = []testelements{
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, more, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, more, []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, more, []int{-7, -4, 1, 1, 1, 2, 5, 6, 8}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, more, []int{-6, -1, 0, 3, 3, 4, 7, 8, 9}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, less, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{9, 1, 8, 2, -7, 3, 6, -4, 5}, less, []int{9, 8, 6, 5, 3, 2, 1, -4, -7}},
		{[]int{1, 1, 8, 2, -7, 1, 6, -4, 5}, less, []int{8, 6, 5, 2, 1, 1, 1, -4, -7}},
		{[]int{9, -1, 8, 0, 7, 3, -6, 4, 3}, less, []int{9, 8, 7, 4, 3, 3, 0, -1, -6}},
	}

	for _, element := range tests {
		v := QuickSort(element.values, element.method)
		if !reflect.DeepEqual(v, element.result) {
			t.Error(
				"For", element.values,
				"expected", element.result,
				"got", v,
			)
		}
	}
}
