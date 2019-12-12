package algorithms

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	type testElements struct {
		lst         *LList
		sizeResult  int
		emptyResult bool
	}

	//lst := new(List)

	lst0 := &LList{}

	lst := &LList{}
	lst.LLAddHead(1)
	lst.LLAddHead(2)
	lst.LLAddHead(3)

	lst1 := &LList{}
	lst1.LLAddHead(11)
	lst1.LLAddHead(2)
	lst1.LLAddHead(32)

	lst2 := &LList{}
	lst2.LLAddHead(11)

	lst3 := &LList{}
	lst3.LLAddHead(11)
	lst3.LLAddHead(-2)
	lst3.LLAddHead(32)
	lst3.LLAddHead(35)
	lst3.LLAddHead(-12)

	var tests = []testElements{
		{lst0, 0, true},
		{lst, 3, false},
		{lst1, 3, false},
		{lst2, 1, false},
		{lst3, 5, false},
	}

	t.Run("Size of List", func(t *testing.T) {

		for _, element := range tests {

			got := element.lst.LLSize() //lst.LLSize()
			want := element.sizeResult

			if got != want {
				t.Error(
					"For", lst,
					"expected", want,
					"got", got,
				)
			}
		}
	})

	t.Run("List is empty", func(t *testing.T) {

		for _, element := range tests {

			got := element.lst.LLIsEmpty() //lst.LLSize()
			want := element.emptyResult

			if got != want {
				t.Error(
					"For", lst,
					"expected", want,
					"got", got,
				)
			}
		}

	})

}
