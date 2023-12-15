package main

import (
	"container/list"
	"errors"
	_ "fmt"
)



type Stack2 struct {
	value     []int

}

type StackMin struct {
	value  []int
	min int
	length int
}

type Element struct {
	elementValue int
}

// NewStack returns a new stack.
// func (stack *StackMin) New() {
// 	stack.value = make([]*Element, 0)
// }

// var stack2  Stack2

func(stack *StackMin) Push(data int) {
	if stack.length == 0 {
		stack.min = data
	} else {
		if stack.min > data{
			stack.min = data
		}
		
	}
	stack.value = append(stack.value[:len(stack.value)-1], data)
	stack.length = stack.length + 1
	// stack2.value = append(stack2.value, data)
	
}

func(stack *StackMin) Peek() (int, bool){
	if stack.length == 0 {
		return -1, false
	}
	data := stack.value[stack.length-1]
	stack.length = stack.length -1
	return data, true
}

func (stack *StackMin) Pop() (int, bool){
	if stack.length == 0 {
		return -1, false
	}
	data := stack.value[stack.length-1]
	stack.value = stack.value[:stack.length-1]
	stack.length = stack.length -1
	return data, true

}


////STACK OF PLATES


type SetOfStacks struct {
	capacity int           // Maximum capacity of each individual stack
	stacks   *list.List    // List of stacks
	current  *list.Element // Pointer to the current stack
}


type Stack struct {
	items []int
}


func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{
		capacity: capacity,
		stacks: list.New(),
		current: nil,
	}
}


func(sos *SetOfStacks) Push(value int) {
	if sos.current == nil || sos.current.Value.(*Stack).Size() == sos.capacity {
		newStack := NewStack()
		sos.current = sos.stacks.PushBack(newStack)
	}
	sos.current.Value.(*Stack).Push(value)

}

func (s *SetOfStacks) pop() (int, error) {
	if s.stacks.Len() == 0 {
		return 0, errors.New("stack is empty")
	}

	value, err := s.current.Value.(*Stack).Pop()
	if err != nil {
		// Current stack is empty, remove it from the set of stacks
		s.stacks.Remove(s.current)
		if s.stacks.Len() > 0 {
			// Update current stack to the last stack in the list
			s.current = s.stacks.Back()
		} else {
			s.current = nil
		}
	}
	return value, err
}

func (s *SetOfStacks) popAt(index int) (int, error) {
	if index < 0 || index >= s.stacks.Len() {
		return 0, errors.New("invalid stack index")
	}

	stackElement := s.stacks.Front()
	for i := 0; i < index; i++ {
		stackElement = stackElement.Next()
	}

	value, err := stackElement.Value.(*Stack).Pop()
	if stackElement.Value.(*Stack).IsEmpty() {
		// If the stack is empty, remove it from the set of stacks
		s.stacks.Remove(stackElement)
		if s.stacks.Len() > 0 {
			// Update current stack to the last stack in the list
			s.current = s.stacks.Back()
		} else {
			s.current = nil
		}
	}
	return value, err
}



func NewStack() *Stack {
	return &Stack{items: []int{}}
}


func(s *Stack) Push(item int){
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}