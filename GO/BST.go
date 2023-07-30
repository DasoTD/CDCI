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



func searchNode(tree *TreeNodes, value int) bool {
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



func main(){
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8)
	tree.InsertElement(3)
	tree.InsertElement(10)
	tree.InsertElement(12)
	// fmt.Println(tree)
	fmt.Println(tree.SearchNode(10))
	// tree.Print()
}