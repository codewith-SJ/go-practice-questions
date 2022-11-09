package linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// Len returns the length of the linked list
func (ll *LinkedList) Len() int {
	return ll.length
}

// PushBack adds a node at the end of the linked list
func (ll *LinkedList) PushBack(n *Node) {
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		ll.length++
	} else {
		ll.tail.next = n
		ll.tail = n
		ll.length++
	}
}

// Display prints data of all the nodes in the linked list
func (ll LinkedList) Display() {
	for ll.head != nil {
		fmt.Printf("%v -> ", ll.head.val)
		ll.head = ll.head.next
	}
	fmt.Println() // For next line
}

// Front returns the value of the head or returns error if head does not exist (empty linked list)
func (ll LinkedList) Front() (int, error) {
	if ll.head == nil {
		return 0, fmt.Errorf("Head does not exist. Empty linked list.")
	}
	return ll.head.val, nil
}

// End returns the value of the tail or returns error if tail does not exist (empty linked list)
func (ll LinkedList) End() (int, error) {
	if ll.tail == nil {
		return 0, fmt.Errorf("Tail does not exist. Empty linked list.")
	}
	return ll.tail.val, nil
}

// Delete deletes a node entered
func (ll *LinkedList) Delete(element int) {
	if ll.head.val == element {
		ll.head = ll.head.next
		ll.length--
		return
	}

	var previousNode *Node = nil
	currNode := ll.head

	for currNode != nil && currNode.val != element {
		previousNode = currNode
		currNode = currNode.next
	}
	if currNode == nil {
		fmt.Println("Node does not exist.")
		return
	}
	previousNode.next = currNode.next
	ll.length--
	fmt.Println("Node deleted.")
}

// ReverseLL reverses the Linked List
func (ll *LinkedList) ReverseLL() {
	currentNode := ll.head
	ll.tail = ll.head

	var previousNode *Node
	for currentNode != nil {
		temp := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = temp
	}
	ll.head = previousNode
}

// containsNodeWithValue finds a value if it present in the linked list
func (ll LinkedList) containsNodeWithValue(value int) *Node {
	node := ll.head
	for node != nil && node.val != value {
		node = node.next
	}
	return node
}
