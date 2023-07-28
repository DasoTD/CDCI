package main

import "fmt"

type DLLNode struct {
	prev *DLLNode
	next *DLLNode
	data int
}

type DLL struct{
	head *DLLNode

}

func (dll *DLL) NodeBTWValues( first, second int) *DLLNode{
	// var node *DLLNode
	// var nodeWith *DLLNode

	// for node= dll.head; node != nil; node = node.next {
	// 	if node.prev != nil && node.next != nil {
	// 		if node.prev.data == first && node.next.data == second{
	// 			nodeWith = node
	// 			break;
	// 		}
	// 	}
	// }
	// return *nodeWith
	curr := dll.head
	value := &DLLNode{}
	for curr != nil {
			if curr.prev.data == first && curr.next.data == second {
				value = curr
				break;
			}
		
	}
	return value
}

func (dll *DLL) insert(data int ) {
	node := &DLLNode{data: data}
	head := dll.head

	if head == nil {
		head = node
		head.prev = nil
		head.next = nil
	} else {
	
			for head.next != nil {
				head = head.next
			}

			head.next = node
			node.prev = head
		}
	// if dll.length == 0 {
	// 	dll.head = node
	// 	dll.tail = node
	// 	dll.head.next = dll.tail
	// } else {
	// 	dll.tail.next = node
	// 	node.prev = dll.tail
	// 	dll.tail = node
	// }
	// dll.length ++
}

func (dll *DLL) printDLL(){
	if dll.head == nil {
		return
	}
	curr := dll.head
	for curr != nil {
		fmt.Println(curr.data)
		curr =curr.next
	}
	// var node *DLLNode
	// for node = dll.head; node != nil; node = node.next {

	// 	fmt.Println(node.data)
	// }
}

func maind(){
	dll := &DLL{}
	dll.insert(4)
	dll.printDLL()
	// fmt.Println(dll.NodeBTWValues(2,3) )
}