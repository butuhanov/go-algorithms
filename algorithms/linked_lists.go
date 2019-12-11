package algorithms

// List is linked list item
type List struct {
head *Node
count int
}

// Node is an item in the list
type Node struct {
value int
next *Node
}

