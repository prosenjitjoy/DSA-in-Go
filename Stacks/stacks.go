package main

import "fmt"

type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop will remove a value at the end
// and return the removed value
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}
