package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Stack struct {
// 	items []int
// }

// func NewStack() *Stack {
// 	return &Stack{items: []int{}}
// }

// func (s *Stack) Push(value int) {
// 	s.items = append(s.items, value)
// }

// func (s *Stack) Pop() (int, error) {
// 	if s.IsEmpty() {
// 		return 0, errors.New("stack is empty")
// 	}
// 	value := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return value, nil
// }

// func (s *Stack) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// func (s *Stack) Size() int {
// 	return len(s.items)
// }

// func DeleteMiddleElement(stack *Stack) error {
// 	auxStack := NewStack()

// 	// Move half of the elements to the auxiliary stack
// 	fmt.Println("half")
// 	fmt.Println(stack.Size()/2)
// 	for i := 0; i < stack.Size()/2; i++ {
// 		value, err := stack.Pop()
// 		if err != nil {
// 			return err
// 		}
// 		auxStack.Push(value)
// 	}

// 	// Remove the middle element from the original stack
// 	stack.Pop()

// 	// Move the elements back to the original stack
// 	for !auxStack.IsEmpty() {
// 		value, _ := auxStack.Pop()
// 		stack.Push(value)
// 	}

// 	return nil
// }

// func PrintStack(stack *Stack) {
// 	for i := len(stack.items) - 1; i >= 0; i-- {
// 		fmt.Println(stack.items[i])
// 	}
// }

// func main() {
// 	stack := NewStack()

// 	// Push elements onto the stack
// 	for i := 1; i <= 5; i++ {
// 		stack.Push(i)
// 	}

// 	fmt.Println("Original Stack:")
// 	PrintStack(stack)

// 	// Delete the middle element
// 	err := DeleteMiddleElement(stack)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	fmt.Println("Stack after deleting the middle element:")
// 	PrintStack(stack)
// }
