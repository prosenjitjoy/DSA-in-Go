package main

import "fmt"

const HashSize = 100

type node struct {
	key  string
	next *node
}

type HashTable struct {
	arr []*node
}

func NewHashTable() *HashTable {
	array := make([]*node, HashSize)
	for i := range array {
		array[i] = &node{key: ""}
	}
	return &HashTable{arr: array}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % HashSize
}

func (ht *HashTable) Insert(key string) {
	currentNode := ht.arr[hash(key)]

	for currentNode.next != nil {
		if currentNode.next.key == key {
			return
		}
		currentNode = currentNode.next
	}

	currentNode.next = &node{key: key}
}

func (ht *HashTable) Search(key string) bool {
	currentNode := ht.arr[hash(key)]

	for currentNode.next != nil {
		if currentNode.next.key == key {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (ht *HashTable) Delete(key string) {
	currentNode := ht.arr[hash(key)]

	for currentNode != nil && currentNode.next != nil {
		if currentNode.next.key == key {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func main() {

	hashTable := NewHashTable()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable.Search("ERIC"))
	hashTable.Delete("ERIC")
	fmt.Println(hashTable.Search("ERIC"))
	fmt.Println(hashTable.Search("TOKEN"))
}
