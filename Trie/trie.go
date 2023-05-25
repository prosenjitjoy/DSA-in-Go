package main

// import "fmt"

// // Node represents each node in the trie
// type Node struct {
// 	children [26]*Node
// 	isEnd    bool
// }

// // Trie represent a trie and has pointer to the root node
// type Trie struct {
// 	root *Node
// }

// // NewTrie will create new Trie
// func NewTrie() *Trie {
// 	return &Trie{root: &Node{}}
// }

// // Insert will take in a word and add it to the trie
// func (t *Trie) Insert(word string) {
// 	currentNode := t.root
// 	for _, v := range word {
// 		charIndex := v - 'a'
// 		if currentNode.children[charIndex] == nil {
// 			currentNode.children[charIndex] = &Node{}
// 		}
// 		currentNode = currentNode.children[charIndex]
// 	}
// 	currentNode.isEnd = true
// }

// // Search will take in a word and return true if that word is included in the trie
// func (t *Trie) Search(word string) bool {
// 	currentNode := t.root
// 	for _, v := range word {
// 		charIndex := v - 'a'
// 		if currentNode.children[charIndex] == nil {
// 			return false
// 		}
// 		currentNode = currentNode.children[charIndex]
// 	}

// 	if currentNode.isEnd {
// 		return true
// 	}

// 	return false
// }

// func main() {
// 	myTrie := NewTrie()
// 	toAdd := []string{
// 		"aragorn",
// 		"aragon",
// 		"argon",
// 		"eragon",
// 		"oregon",
// 		"oregano",
// 		"oreo",
// 	}

// 	for _, v := range toAdd {
// 		myTrie.Insert(v)
// 	}
// 	fmt.Println(myTrie.Search("arrgon"))
// }
