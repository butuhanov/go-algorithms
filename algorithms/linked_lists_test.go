package algorithms

import (
	"reflect"
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
	lst0.LLRemoveHead()

	lst := &LList{}
	lst.LLRemoveHead()
	lst.LLAddHead(1)
	lst.LLAddHead(2)
	lst.LLAddHead(3)
	lst.LLRemoveHead()
	lst.LLAddHead(3)
	lst.LLDeleteNode(2)
	lst.LLAddHead(1)
	lst.LLDeleteNode(4)

	lst1 := &LList{}
	lst1.LLAddHead(11)
	lst1.LLAddHead(2)
	lst1.LLAddHead(32)

	lst2 := &LList{}
	lst2.LLAddHead(11)

	lst3 := &LList{}
	lst3.LLAddHead(11)
	lst3.LLAddHead(-2)
	lst3.LLRemoveHead()
	lst3.LLAddHead(-2)
	lst3.LLAddHead(32)
	lst3.LLAddHead(35)
	lst3.LLAddHead(-12)
	lst3.FreeLList()

	lst4 := &LList{}
	lst4.LLAddTail(1)

	lst5 := &LList{}
	lst5.LLAddTail(-1)
	lst5.LLAddTail(1)

	lst6 := &LList{}
	// lst6.LLPrint()
	lst6.LLAddHead(11)
	// lst6.LLPrint()
	lst6.LLAddHead(-2)
	// lst6.LLPrint()
	lst6.LLSortedInsert(5)
	// lst6.LLPrint()
	lst6.LLRemoveHead()
	// lst6.LLPrint()
	lst6.LLRemoveHead()
	// lst6.LLPrint()
	lst6.LLAddHead(-2)
	lst6.LLAddTail(32)
	lst6.LLSortedInsert(5)
	lst6.LLSortedInsert(-124)
	// lst6.LLPrint()
	lst6.LLAddHead(35)
	// lst6.LLPrint()
	lst6.LLSortedInsert(6)
	// lst6.LLPrint()
	lst6.LLAddHead(-12)
	// lst6.LLPrint()
	lst6.LLAddTail(2)
	lst6.LLRemoveHead()
	// lst6.LLPrint()
	lst6.LLAddHead(32)
	// lst6.LLPrint()
	lst6.LLAddTail(62)
	// lst6.LLPrint()
	lst6.LLDeleteNodes(32)
	// lst6.LLPrint()
	lst6.LLDeleteNode(62)
	// lst6.LLPrint()
	lst6.LLAddHead(-122)
	// lst6.LLPrint()
	lst6.LLAddHead(62)
	// lst6.LLPrint()
	lst6.LLAddHead(12)
	// lst6.LLPrint()
	lst6.LLAddTail(-125)
	// lst6.LLPrint()
	lst6.LLAddTail(162)
	lst6.LLAddTail(42)
	lst6.LLAddHead(42)
	// lst6.LLPrint()

	var tests = []testElements{
		{lst0, 0, true, 1, false},
		{lst, 3, false, 1, true},
		{lst1, 3, false, 32, true},
		{lst1, 3, false, 1, false},
		{lst2, 1, false, 1, false},
		{lst3, 0, true, -2, false},
		{lst4, 1, false, 1, true},
		{lst5, 2, false, 1, true},
		{lst6, 12, false, 12, true},
		{lst6, 12, false, -122, true},
		{lst6, 12, false, 6, true},
		{lst6, 12, false, 1, false},
	}

	t.Run("Size of List", func(t *testing.T) {

		for _, element := range tests {
			// fmt.Printf("%#v\n", element.lst)

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

	t.Run("Reverse elements", func(t *testing.T) {
		want := []int{42, 162, -125, 2, 11, 5, -2, -124, 35, 6, -122, 62, 12, 42}
		// lst6.LLPrint()
		lst6.LLReverse() //lst.LLSize()
		// want := [0, 1, 2, 3]
		get := lst6.LLPrint()

		if !reflect.DeepEqual(want, get) {
			t.Error(
				"For", lst6,
				"expected", want,
				"got", get,
			)
		}

	})

	t.Run("Recursive reverse elements", func(t *testing.T) {
		want := []int{42, 12, 62, -122, 6, 35, -124, -2, 5, 11, 2, -125, 162, 42}
		// lst6.LLPrint()
		lst6.LLReverseRecurse() //lst.LLSize()
		// want := [0, 1, 2, 3]
		get := lst6.LLPrint()

		if !reflect.DeepEqual(want, get) {
			t.Error(
				"For", lst6,
				"expected", want,
				"got", get,
			)
		}

	})

}
