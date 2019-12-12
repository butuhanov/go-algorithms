package algorithms

import "fmt"

// LList is linked list item
type LList struct {
	head  *NodeLL
	count int
}

// NodeLL is an item in the list
type NodeLL struct {
	value int
	next  *NodeLL
}

// LLSize returns size of a given linked list
func (list *LList) LLSize() int {
	return list.count
}

// LLIsEmpty returns true if the linked list is empty
func (list *LList) LLIsEmpty() bool {
	return (list.count == 0)
}

// LLAddHead adds a head to the linked list
func (list *LList) LLAddHead(value int) {
	list.head = &NodeLL{value, list.head} //  We need to create a new node with the value passed to the function as argument.
	list.count++                          //  Size of the list is increased by one.
}

// LLAddTail inserts an element at the end of linked list
// This operation is un-efficient as each time you want to insert an element you have to traverse to the end of the list. Therefore, the complexity of creation of the list is n2. So how to make it efficient we have to keep track of the last element by keeping a tail reference. Therefore, if it is required to insert element at the end of linked list, then we will keep track of the tail reference also.
func (list *LList) LLAddTail(value int) {
	curr := list.head
	list.count++
	newNode := &NodeLL{value, nil}
	if curr == nil {
		list.head = newNode
		return
	}
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

// LLPrint prints elements of a linked list
func (list *LList) LLPrint() {
	temp := list.head
	for temp != nil {

		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}
