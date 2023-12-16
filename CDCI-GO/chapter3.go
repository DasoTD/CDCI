package main

import (
	"container/list"
	"errors"
	"fmt"
	_ "fmt"
	"time"
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


//STACK sort
type StackSort struct {
	items []int
}

func NewStackSort() *StackSort {
	return &StackSort{items: []int{}}
}

func(s *StackSort) IsEmpty() bool{
	if len(s.items)==0 {
		return false
	}
	return true
}

func(s *StackSort) Push(value int) {
	s.items = append(s.items, value)
}

func(s *StackSort) Pop() (int, bool){
	if s.IsEmpty(){
		return -1, false
	}
	length := len(s.items)
	data := s.items[length -1]
	s.items = s.items[:length-1]
	return data, true
}

func(s *StackSort) Peek() (int, bool){

	if s.isEmpty() {
		return -1, false
	}
	
	return s.items[len(s.items)-1], true 
	
}
func(s *StackSort) isEmpty() bool{
	if len(s.items) == 0 {
		return false
	}
	return true
}
func(s *StackSort) size() int{
	return len(s.items)
}


func SortStack(stack *StackSort) *StackSort {
	tempStack := NewStack()

	for !stack.isEmpty() {
		// Pop the top element from the original stack
		current, _ := stack.Pop()

		// Move elements from the temporary stack to the original stack
		for !tempStack.IsEmpty() && current < tempStack.items[len(tempStack.items)-1] {
			top, _ := tempStack.Pop()
			stack.Push(top)
		}

		// Push the current element onto the temporary stack
		tempStack.Push(current)
	}
	// Move sorted elements from the temporary stack to the original stack
	for !tempStack.IsEmpty() {
		top, _ := tempStack.Pop()
		stack.Push(top)
	}

	return stack

}


// ANIMAL SHELTER

type Animal struct {
	Name string
	Timestamp time.Time
}

type AnimalShelter  struct  {
	dogs *list.List
	cats *list.List
}

func NewAnimalShelter()  *AnimalShelter {
	return &AnimalShelter{
		dogs: list.New(),
		cats: list.New(),
	}
}


func(shelter *AnimalShelter) Enqueue(Name, AnimalType string) {
	animal := &Animal{
		Name: Name,
		Timestamp: time.Now(),
	}
	if AnimalType == "dog"{
		shelter.dogs.PushBack(animal)
	} else if AnimalType == "cat" {
		shelter.cats.PushBack(animal)
	} else {
		fmt.Println("Invalid animal type.")

	}
}

func(shelter *AnimalShelter) DequeueAny()(Animal, error){
	if shelter.cats.Len() == 0 && shelter.dogs.Len() == 0 {
		return Animal{}, errors.New("shelter is empty")
	}

	var oldestAnimal Animal

	if shelter.dogs.Len() == 0 {
		oldestAnimal = shelter.dequeueCat()
	} else if shelter.cats.Len() == 0 {
		oldestAnimal = shelter.dequeueDog()
	} else {
		oldestDog := shelter.dogs.Front().Value.(Animal)
		oldestCat := shelter.cats.Front().Value.(Animal)

		if oldestDog.Timestamp.Before(oldestCat.Timestamp) {
			oldestAnimal = shelter.dequeueDog()
		} else {
			oldestAnimal = shelter.dequeueCat()
		}
	}
	return oldestAnimal, nil

}


func(shelter *AnimalShelter) dequeueCat() Animal {
	if shelter.cats.Len() ==0 {
		return Animal{}
	}
	oldestCat := shelter.cats.Front().Value.(Animal)
	shelter.cats.Remove(shelter.cats.Front())
	return oldestCat
}

func (shelter *AnimalShelter) dequeueDog() Animal {
	if shelter.dogs.Len() == 0 {
		return Animal{}
	}

	oldestDog := shelter.dogs.Front().Value.(Animal)
	shelter.dogs.Remove(shelter.dogs.Front())

	return oldestDog
}



