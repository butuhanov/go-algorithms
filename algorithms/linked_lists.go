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
func (list *LList) LLPrint() []int {
	temp := list.head
	result := []int{}
	for temp != nil {
		result = append(result, temp.value)
		// fmt.Print(temp.value, " ")
		temp = temp.next
	}
	// fmt.Println("")
	return result
}

// LLSortedInsert - inserts an element in sorted order in linked list given Head reference
func (list *LList) LLSortedInsert(value int) {
	newNode := &NodeLL{value, nil}
	curr := list.head

	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
}

// LLElementIsPresent - search element in a linked list
func (list *LList) LLElementIsPresent(data int) bool {
	temp := list.head // a temp variable, which will point to head of the list
	for temp != nil { //  iterate through the list.
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

// LLRemoveHead - Delete First element in a linked list.
func (list *LList) LLRemoveHead() (int, bool) {
	if list.LLIsEmpty() {
		// fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	list.count--
	return value, true
}

// LLDeleteNode - Delete node from the linked list given its value.
func (list *LList) LLDeleteNode(delValue int) bool {
	temp := list.head
	if list.LLIsEmpty() {
		// fmt.Println("EmptyListError")
		return false
	}
	if delValue == list.head.value {
		list.head = list.head.next
		list.count--
		return true
	}
	for temp.next != nil {
		if temp.next.value == delValue {
			temp.next = temp.next.next
			list.count--
			return true
		}
		temp = temp.next
	}
	return false
}

// LLDeleteNodes - deletes all occurences of particular value from linked list
func (list *LList) LLDeleteNodes(delValue int) {
	currNode := list.head
	for currNode != nil && currNode.value == delValue {
		list.head = currNode.next
		currNode = list.head
	}
	for currNode != nil {
		nextNode := currNode.next
		if nextNode != nil && nextNode.value == delValue {
			currNode.next = nextNode.next
		} else {
			currNode = nextNode
		}
	}
}

// FreeLList - deletes all the elements of a linked list
func (list *LList) FreeLList() {
	list.head = nil
	list.count = 0
}

// LLReverse reversed linked list iteratively
func (list *LList) LLReverse() {
	curr := list.head
	var prev, next *NodeLL
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}
