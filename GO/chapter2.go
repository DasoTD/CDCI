package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct {
    head *Node
	tail *Node
	length int
}

func(l *LinkedList) prepend(value int) {
	newNode := &Node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = newNode
		l.length ++
	} else {
		l.head.next = l.head
		l.head = newNode
		// l.tail = newNode
		l.length ++
	}
	
}
func (l *LinkedList) append(value int) bool{
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return true
	}

	curr := l.head
    for curr.next != nil {
        curr = curr.next
    }
    curr.next = newNode
	// l.tail.next = newNode
	// newNode = l.tail
	l.length++
	return true

}

func (l *LinkedList) printList(){
	if l.head == nil {
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func (l *LinkedList) deleteWithValue(value int){
	if l.length ==0 || value < 0 || value > l.length{
		return
	}

	if value == l.head.data && l.length==1{
		l.head = nil
		l.tail = nil
		l.length--
	}

	if value == l.head.data{
		l.head = l.head.next
		l.length--
	}
	
	prev := l.head
	for prev.next.data != value {
		if prev.next.next == nil{
			return
		}
		prev = prev.next
	}
	prev.next = prev.next.next
	l.length --
}

func (l *LinkedList) pop() (int, bool){
	curr := l.head
	if l.length == 0 {
		return 0, false
	}
	if l.length ==1{
		l.head = nil
		l.tail = nil
		l.length--
		fmt.Println(curr.data)
		return curr.data, true
	}

	//issue dey here
	prev := l.head
	for curr.next.data !=l.tail.data {
		prev = curr
		curr = curr.next
	}
	l.tail = prev
	l.tail.next = nil
	l.length--
	return curr.data, true
}

func (l *LinkedList) Get(index int) int{
	if index <0 || index > l.length {
		// return 
	}
	if index == l.length {
		return l.tail.data
	}
	if index == 0 {
		return l.head.data
	}
	curr := l.head
	for i:=0; i<(index -1); i++ {
		curr = curr.next
	}
	return curr.data
}

func nestLoop(data []int )  {
	for a:=0; a<len(data); a++ {
		for j := len(data) - 1; a <= j; j--{
			fmt.Println(a,j)
		}
	}
}

var val = []int {1,2,3,4,5}
func main(){
	list := &LinkedList{}
	list.append(1)
    list.append(2)
    list.append(3)
    list.append(4)

    fmt.Println("Initial List: ")
    list.printList()
	list.deleteWithValue(3)
	fmt.Println("deleted List: ")
	list.prepend(5)
	// list.pop()
	list.printList()
	// nestLoop(val)
}
