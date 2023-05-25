package main

import "fmt"

type node struct {
	children [26]*node
	isEnd    bool
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{root: &node{}}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, v := range word {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, v := range word {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd {
		return true
	}

	return false
}

func main() {
	myTrie := NewTrie()
	myTrie.Insert("joy")
	fmt.Println(myTrie.Search("prosenjit"))
	fmt.Println(myTrie.Search("joy"))
}
