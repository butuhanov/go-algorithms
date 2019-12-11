package algorithms

// List is linked list item
type List struct {
	head  *NodeLL
	count int
}

// Node is an item in the list
type NodeLL struct {
	value int
	next  *NodeLL
}
