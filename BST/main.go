package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(500)
	tree.Insert(400)
	fmt.Println(tree.Search(500))
	fmt.Println(tree.Search(40))
	fmt.Println(tree.Search(400))
	fmt.Println(tree.Search(401))
}
