// A single pointing linked list which points only to the next node
/* The ourput should be like :
The next node is : 1
The next node is : 2
The next node is : 3
The next node is : 4
The next node is : 5
*/
package main

import "fmt"

// List struct which has the information of the node
type List struct {
	tail, head *Node
}

// Node struct which has the information to the next one
type Node struct {
	value  int
	second *Node
}

func (l *List) first() *Node {
	return l.head
}

func (n *Node) next() *Node {
	return n.second
}

func (l *List) move(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.second = node
	}
	l.tail = node
}
func main() {
	// Create a list of 5 nodes
	list := &List{}
	for i := 1; i <= 5; i++ {
		list.move(i)
	}

	// Move to the first node
	n := list.first()

	// Move to the rest of nodes
	for {
		fmt.Println("The next node is :", n.value)
		n = n.next()
		if n == nil {
			break
		}
	}
}
