// Implementing Stacks as its main options "Push" and "Pull"
package main

import "fmt"

type Stack struct {
	items []int
}

// Append a Data
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	item, items := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = items
	return item

}

func main() {

	s := Stack{}
	for i := 1; i <= 5; i++ {
		s.push(i)
		fmt.Println("Pushed in  : ", i)
		fmt.Println("Pulled out : ", s.pop())
	}

}
