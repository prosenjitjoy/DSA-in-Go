package main

// import "fmt"

// const ArraySize = 100

// // HashTable structure
// type HashTable struct {
// 	array [ArraySize]*bucket
// }

// // bucket structure
// type bucket struct {
// 	head *bucketNode
// }

// // bucketNode structure
// type bucketNode struct {
// 	key  string
// 	next *bucketNode
// }

// // Insert will take in a key and add it to the hash table array
// func (h *HashTable) Insert(key string) {
// 	index := hash(key)
// 	h.array[index].insert(key)
// }

// // // Search will take in a key and return true if that key is stored in the hash table
// func (h *HashTable) Search(key string) bool {
// 	index := hash(key)
// 	return h.array[index].search(key)
// }

// // // Delete will take in a key and delete it from the hash table
// func (h *HashTable) Delete(key string) {
// 	index := hash(key)
// 	h.array[index].delete(key)
// }

// // insert will take in a key, create a node with the key and insert the node in the bucket
// func (b *bucket) insert(k string) {
// 	if !b.search(k) {
// 		newNode := &bucketNode{key: k}
// 		newNode.next = b.head
// 		b.head = newNode
// 	} else {
// 		fmt.Println(k, "already exists")
// 	}

// }

// // search
// func (b *bucket) search(k string) bool {
// 	currentNode := b.head
// 	for currentNode != nil {
// 		if currentNode.key == k {
// 			return true
// 		}
// 		currentNode = currentNode.next
// 	}
// 	return false
// }

// // delete
// func (b *bucket) delete(k string) {
// 	if b.head.key == k {
// 		b.head = b.head.next
// 		return
// 	}
// 	previousNode := b.head
// 	for previousNode.next != nil {
// 		if previousNode.next.key == k {
// 			previousNode.next = previousNode.next.next
// 		}
// 		previousNode = previousNode.next
// 	}
// }

// // Hash
// func hash(key string) int {
// 	sum := 0
// 	for _, v := range key {
// 		sum += int(v)
// 	}
// 	return sum % ArraySize
// }

// // Init will create a bucket in each slot of the hash table
// func Init() *HashTable {
// 	result := &HashTable{}
// 	for i := range result.array {
// 		result.array[i] = &bucket{}
// 	}
// 	return result
// }

// func main() {
// 	// testHashTable := &HashTable{}
// 	// testHashTable := Init()
// 	// fmt.Println(testHashTable)
// 	// fmt.Println(hash("RANDY"))

// 	// testBucket := &bucket{}
// 	// testBucket.insert("RANDY")
// 	// testBucket.delete("RANDY")
// 	// fmt.Println(testBucket.search("RANDY"))
// 	// fmt.Println(testBucket.search("ERIC"))

// 	hashTable := Init()
// 	list := []string{
// 		"ERIC",
// 		"KENNY",
// 		"KYLE",
// 		"STAN",
// 		"RANDY",
// 		"BUTTERS",
// 		"TOKEN",
// 	}
// 	for _, v := range list {
// 		hashTable.Insert(v)
// 	}

// 	hashTable.Delete("ERIC")
// }
