// adding the higher int in right and smaller in left
// Also it will print the higher "righ side" first
package main

import "fmt"

type Tree struct {
	node *Node
}

type Node struct {
	value       int
	right, left *Node
}

func (n *Node) addValue(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.addValue(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.addValue(value)
		}
	}

}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.addValue(value)
	}
	return t
}

func printNodes(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	printNodes(n.right)
	printNodes(n.left)
}
func main() {
	// create the tree
	t := &Tree{}
	t.insert(10).
		insert(9).
		insert(22).
		insert(7).
		insert(30).
		insert(25).
		insert(3).
		insert(14)
	printNodes(t.node)
}
