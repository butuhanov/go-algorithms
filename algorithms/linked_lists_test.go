package algorithms

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {

	type testElements struct {
		lst          *LList
		sizeResult   int
		emptyResult  bool
		searchData   int
		searchResult bool
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

	lst4 := &LList{}
	lst4.LLAddTail(1)

	lst5 := &LList{}
	lst5.LLAddTail(-1)
	lst5.LLAddTail(1)

	lst6 := &LList{}
	lst6.LLPrint()
	lst6.LLAddHead(11)
	lst6.LLPrint()
	lst6.LLAddHead(-2)
	lst6.LLPrint()
	lst6.LLSortedInsert(5)
	lst6.LLPrint()
	lst6.LLAddTail(32)
	lst6.LLPrint()
	lst6.LLAddHead(35)
	lst6.LLPrint()
	lst6.LLSortedInsert(6)
	lst6.LLPrint()
	lst6.LLAddHead(-12)
	lst6.LLPrint()
	lst6.LLAddTail(2)
	lst6.LLPrint()
	lst6.LLAddHead(32)
	lst6.LLPrint()
	lst6.LLAddTail(62)
	lst6.LLPrint()
	lst6.LLAddHead(-122)
	lst6.LLPrint()
	lst6.LLAddHead(62)
	lst6.LLPrint()
	lst6.LLAddHead(12)
	lst6.LLPrint()
	lst6.LLAddTail(-125)
	lst6.LLPrint()

	var tests = []testElements{
		{lst0, 0, true, 1, false},
		{lst, 3, false, 1, true},
		{lst1, 3, false, 32, true},
		{lst1, 3, false, 1, false},
		{lst2, 1, false, 1, false},
		{lst3, 5, false, -2, true},
		{lst4, 1, false, 1, true},
		{lst5, 2, false, 1, true},
		{lst6, 12, false, -12, true},
		{lst6, 12, false, -122, true},
		{lst6, 12, false, 6, true},
		{lst6, 12, false, 1, false},
	}

	t.Run("Size of List", func(t *testing.T) {

		for _, element := range tests {
			fmt.Printf("%#v\n", element.lst)

			got := element.lst.LLSize() //lst.LLSize()
			want := element.sizeResult

			if got != want {
				t.Error(
					"For", element.lst,
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

	t.Run("Search element", func(t *testing.T) {

		for _, element := range tests {

			got := element.lst.LLElementIsPresent(element.searchData) //lst.LLSize()
			want := element.searchResult

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
