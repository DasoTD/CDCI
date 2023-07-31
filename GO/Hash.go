package main

import (
	"fmt"
	// "hash"
)

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key string
	next *BucketNode
}

func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (b *bucket) insert(value string){
	if !b.search(value){
		newNode := &BucketNode{key:value}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(value, "already exist")
	}
	
}

func (b *bucket) search(value string) bool{
	currentNode := b.head

	for currentNode !=  nil {
		if currentNode.key == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(value string) {
	// if !b.search(value){
	// 	fmt.Println(value, "does not exist")
	// }

	if b.head.key == value {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	

	for prevNode.next.key != value{
		if prevNode.next.next == nil{
			return
		}
		prevNode = prevNode.next
	}

	// for prevNode.next != nil {
	// 	if prevNode.next.key ==value{
	// 		prevNode.next = prevNode.next.next
	// 	}
	// 	prevNode = prevNode.next
	// }
	prevNode.next = prevNode.next.next

	
}

func (h *HashTable) Insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}
func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].search(key)
}
func (h *HashTable) Delete(key string){
	index := hash(key)
	h.array[index].delete(key)
}

func hash(key string) int{
	sum := 0
	for _, v := range key {
		// fmt.Println("v", int(v))
		sum+=  int(v)
	}
	return sum % ArraySize
}

func main (){
	testHT := Init()
	testHT.Insert("DADA")
	testHT.Insert("DAVID")
	testHT.Delete("DAVID")
	testHT.Insert("T>>>D")
	fmt.Println(testHT.Search("DADA"))
	fmt.Println(testHT.Search("DAVID"))
	fmt.Println(testHT.Search("DADAd"))

	fmt.Println(testHT)

	fmt.Println(hash("RANDY"))
	testBucket := &bucket{}
	testBucket.insert("Randy")
	testBucket.delete("Randy")
	fmt.Println(testBucket.search("Randy"))
	fmt.Println(testBucket.search("DADA"))
}