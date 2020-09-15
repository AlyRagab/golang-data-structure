// It is not only pointing to the next node as implemented in the single linked list , but also pointing to the previous one as well.
/* The output should be like :
The next node is : 1
The next node is : 2
The next node is : 3
The next node is : 4
The next node is : 5
>> The Previous node is :  5
>> The Previous node is :  4
>> The Previous node is :  3
>> The Previous node is :  2
>> The Previous node is :  1
*/
package main

import "fmt"

// List struct which has the information of the node
type List struct {
	tail, head *Node
}

// Node struct which has the information to the next one
type Node struct {
	value        int
	second, prev *Node
}

// Geeting into the first node
func (l *List) first() *Node {
	return l.head
}

// Getting into the next node
func (n *Node) next() *Node {
	return n.second
}

// Getting into the last node
func (l *List) last() *Node {
	return l.tail
}

// Getting into the previous node
func (n *Node) previous() *Node {
	return n.prev
}

func (l *List) move(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.second = node
		node.prev = l.tail
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

	n = list.last()
	for {
		fmt.Println(">> The Previous node is : ", n.value)
		n = n.previous()
		if n == nil {
			break
		}
	}

}
