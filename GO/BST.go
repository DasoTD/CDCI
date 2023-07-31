///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"sync"
)

// TreeNodes class
type TreeNodes struct {
	value     int
	leftNode  *TreeNodes
	rightNode *TreeNodes
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNodes
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement( value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNodes
	treeNode = &TreeNodes{ value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

func insertTreeNode(rootNode *TreeNodes, newTreeNode *TreeNodes ){
	if newTreeNode.value < rootNode.value {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else if newTreeNode.value > rootNode.value {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	} else {
		return 
	}
}

func (tree *BinarySearchTree) Print() {
	if tree != nil {

		fmt.Println(" Value", tree.rootNode.value)
		fmt.Println("Root Tree Node")
		printTreeNode(tree.rootNode)
	} else {
		fmt.Println("Nil\n")
	}
}

func printTreeNode(treeNode *TreeNodes) {
	if treeNode != nil {
		fmt.Println(" Valuez", treeNode.value)
		fmt.Println("TreeNodes Left")
		printTreeNode(treeNode.leftNode)
		fmt.Println("TreeNodes Right")
		printTreeNode(treeNode.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}

}
// SearchNode method
func (tree *BinarySearchTree) SearchNode(value int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, value)
}

var count int

func searchNode(tree *TreeNodes, value int) bool {
	count++
	if tree == nil {
		return false
	}
	if value < tree.value{
		return searchNode(tree.leftNode, value)
	}
	if value > tree.value {
		return searchNode(tree.rightNode, value)
	}

	return true
}

func (tree *BinarySearchTree) RemoveNode(value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, value)
}


func removeNode(treeNode *TreeNodes, value int) *TreeNodes {
	if treeNode == nil {
		return nil
	}
	if value < treeNode.value {
		treeNode.leftNode = removeNode(treeNode.leftNode, value)
		return treeNode
	}
	if value > treeNode.value {
		treeNode.rightNode = removeNode(treeNode.rightNode, value)
		return treeNode
	}

	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	var leftmostrightNode *TreeNodes
	leftmostrightNode = treeNode.rightNode
	for {

		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.value, treeNode.value = leftmostrightNode.value, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.value)
	return treeNode
}


func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNodes, function func(int)) {
	if treeNode != nil {
	inOrderTraverseTree(treeNode.leftNode, function)
	function(treeNode.value)
	inOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNodes, function func(int)) {
	if treeNode != nil {
	function(treeNode.value)
	preOrderTraverseTree(treeNode.leftNode, function)
	preOrderTraverseTree(treeNode.rightNode, function)
}
}

func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNodes, function func(int)) {
	if treeNode != nil {
	postOrderTraverseTree(treeNode.leftNode, function)
	postOrderTraverseTree(treeNode.rightNode, function)
	function(treeNode.value)
	}
}

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	treeNode := &TreeNodes{}
	// var treeNode *TreeNodes
	// treeNode = tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

func (tree *BinarySearchTree) MaxNode() *int{
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	treeNode := &TreeNodes{}

	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode ==nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}
func maingg(){
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(100)
	tree.InsertElement(200)
	tree.InsertElement(300)
	tree.InsertElement(32)
	tree.InsertElement(16)
	tree.InsertElement(65)
	tree.InsertElement(43)
	tree.InsertElement(25)
	// fmt.Println(tree)
	fmt.Println(tree.SearchNode(25))
	fmt.Println(count)
	// tree.Print()
	// tree.RemoveNode(25)
	// fmt.Println("remove")
	// tree.Print()
	fmt.Println(tree.MinNode())
	fmt.Println(tree.MaxNode())
}