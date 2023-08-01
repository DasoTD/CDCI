package main

import (
	"fmt"
	// "math"
	"errors"
)
// type ElementS struct {
// 	elementValue int
// }

// type StackS struct {
// 	elements []*ElementS
// 	count int
// 	min *ElementS
// }

// func(stack *StackS) New(){
// 	stack.elements = make([]*ElementS, 0)
// }

// func (stack *StackS) Push(element *ElementS) {
// 	if stack.count == 0 {
// 		stack.min = element
// 	}
// 	if stack.min.elementValue > element.elementValue {
// 		stack.min = element
// 	}
// 	stack.elements = append((stack.elements)[:stack.count],element )
// 	stack.count = stack.count +1
// }

// func(stack *StackS) getMin()*ElementS{
// 	if stack.count == 0 {
// 		return nil
// 	} else {
// 		return stack.min
// 	}
// }

// func (stack *StackS) pop()(*ElementS, error) {
// 	if stack.count == 0 {
// 		return 0, nil
// 	}
// 	// var length int = len(stack.elements)
// 	val := stack.elements[len(stack.elements)-1]
// 	stack.elements = stack.elements[:len(stack.elements)-1]
// 	if val == stack.getMin() {
// 		stack.min = stack.min.elementValue[:len(stack.min)-1]
// 	}
// 	return val, nil
	
// }

// func main() {
// 	SST := &StackS{}
// 	SST.New()
// 	element1 := &ElementS{3}
// 	element2 := &ElementS{5}
// 	element3 := &ElementS{8}
// 	element4 := &ElementS{1}

// 	SST.Push(element1)
// 	SST.Push(element2)
// 	SST.Push(element3)
// 	SST.Push(element4)

// 	fmt.Println(SST.getMin())

// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

type Stacks struct {
	values []int
	mins   []int
}

func NewStack() *Stacks {
	return &Stacks{
		values: []int{},
		mins:   []int{},
	}
}

func (s *Stacks) Push(value int){
	s.values = append(s.values, value)

	if len(s.mins) == 0 || s.GetMin() > value {
		s.mins = append(s.mins, value)
	}
}

func (s *Stacks) Pop() (int, error) {
	if len(s.values)==0 {
		return 0, errors.New("Stacks is empty")
	}
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	if val == s.GetMin() {
		s.mins = s.mins[:len(s.mins)-1]
	}

	return val, nil
}

func (s *Stacks) GetMin() int {
	if len(s.mins) == 0 {
		// Return a large value as there's no actual minimum
		return 1<<31 - 1 // Max int32 value
	}
	return s.mins[len(s.mins)-1]
}

func main() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	stack.Push(2)
	stack.Push(7)

	fmt.Println("Minimum element:", stack.GetMin()) // Output: 2

	stack.Pop()
	stack.Pop()

	fmt.Println("Minimum element:", stack.GetMin()) // Output: 3
}