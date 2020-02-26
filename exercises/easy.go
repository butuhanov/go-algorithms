package main

import (
	"fmt"
	algorithm "github.com/butuhanov/go-algorithms/algorithms"
)

func main() {
	fmt.Println(algorithm.Hello())

	llist1 := algorithm.LList{}
	llist1.LLAddHead(1)
	llist1.LLAddTail(5)
	llist1.LLAddTail(4)
	llist1.LLPrint()
	llist1.LLSortedInsert(6)
	llist1.LLSortedInsert(2)
	llist1.LLSortedInsert(5)
	llist1.LLPrint()
	llist1.RemoveLLDuplicate()
	llist1.LLPrint()
	llist1.LLReverse()
	llist1.LLPrint()
}
