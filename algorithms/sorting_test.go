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
