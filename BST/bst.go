package main

// import "fmt"

// var count int

// // Node represents the componets of a binary search tree
// type Node struct {
// 	Key   int
// 	Left  *Node
// 	Right *Node
// }

// // Insert will add a node to the tree
// // the key to add should not be already in the tree
// func (n *Node) Insert(k int) {
// 	if k < n.Key {
// 		if n.Left == nil {
// 			n.Left = &Node{Key: k}
// 		} else {
// 			n.Left.Insert(k)
// 		}
// 	} else if k > n.Key {
// 		if n.Right == nil {
// 			n.Right = &Node{Key: k}
// 		} else {
// 			n.Right.Insert(k)
// 		}
// 	}
// }

// // Search will take in a key value
// // and return true if there is a node with the value
// func (n *Node) Search(k int) bool {
// 	count++

// 	if n == nil {
// 		return false
// 	}

// 	if k < n.Key {
// 		return n.Left.Search(k)
// 	} else if k > n.Key {
// 		return n.Right.Search(k)
// 	}

// 	return true
// }
// func main() {
// 	tree := &Node{Key: 100}
// 	tree.Insert(52)
// 	tree.Insert(203)
// 	tree.Insert(19)
// 	tree.Insert(310)
// 	tree.Insert(150)
// 	tree.Insert(7)
// 	tree.Insert(24)
// 	tree.Insert(88)
// 	tree.Insert(276)
// 	fmt.Println(tree.Search(52))
// 	fmt.Println(count)
// }
