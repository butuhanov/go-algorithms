package algorithms

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	//lst := new(List)
	lst := &LList{}
	lst.LLAddHead(1)
	lst.LLAddHead(2)
	lst.LLAddHead(3)

	t.Run("Size of List", func(t *testing.T) {

		got := lst.LLSize()
		want := 3

		if got != want {
			t.Error(
				"For", lst,
				"expected", want,
				"got", got,
			)
		}
	})

	t.Run("List is empty", func(t *testing.T) {

		got := lst.LLIsEmpty()
		want := false

		if got != want {
			t.Error(
				"For", lst,
				"expected", want,
				"got", got,
			)
		}
	})

}
