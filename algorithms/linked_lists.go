package algorithms

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
	list.count++ //  Size of the list is increased by one.
}
