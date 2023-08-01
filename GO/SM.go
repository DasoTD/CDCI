// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Stack struct {
// 	values []int
// 	mins   []int
// }

// func NewStack() *Stack {
// 	return &Stack{
// 		values: []int{},
// 		mins:   []int{},
// 	}
// }

// func (s *Stack) Push(val int) {
// 	s.values = append(s.values, val)

// 	// Update the auxiliary stack with the minimum element
// 	if len(s.mins) == 0 || val <= s.GetMin() {
// 		s.mins = append(s.mins, val)
// 	}
// }

// func (s *Stack) Pop() (int, error) {
// 	if s.IsEmpty() {
// 		return 0, errors.New("Stack is empty")
// 	}

// 	val := s.values[len(s.values)-1]
// 	s.values = s.values[:len(s.values)-1]

// 	// If the popped value is the current minimum, remove it from the auxiliary stack
// 	if val == s.GetMin() {
// 		s.mins = s.mins[:len(s.mins)-1]
// 	}

// 	return val, nil
// }

// func (s *Stack) GetMin() int {
// 	if len(s.mins) == 0 {
// 		// Return a large value as there's no actual minimum
// 		return 1<<31 - 1 // Max int32 value
// 	}
// 	return s.mins[len(s.mins)-1]
// }

// func (s *Stack) IsEmpty() bool {
// 	return len(s.values) == 0
// }

// func main() {
// 	stack := NewStack()

// 	stack.Push(1)
// 	stack.Push(3)
// 	stack.Push(5)
// 	stack.Push(2)
// 	stack.Push(7)

// 	fmt.Println("Minimum element:", stack.GetMin()) // Output: 2

// 	stack.Pop()
// 	stack.Pop()

// 	fmt.Println("Minimum element:", stack.GetMin()) // Output: 3
// }
