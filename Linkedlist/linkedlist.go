package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := LinkedList{}
	node1 := &node{data: 64}
	node2 := &node{data: 32}
	node3 := &node{data: 16}
	node4 := &node{data: 8}
	node5 := &node{data: 4}
	node6 := &node{data: 2}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)

	mylist.printListData()
	mylist.deleteWithValue(106)
	mylist.printListData()
}
