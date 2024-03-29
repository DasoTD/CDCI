package main

import (
	"container/list"
	"errors"
	"fmt"
	_ "fmt"
	"strconv"
	"time"
)

type Stack2 struct {
	value []int
}

type StackMin struct {
	value  []int
	min    int
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

func (stack *StackMin) Push(data int) {
	if stack.length == 0 {
		stack.min = data
	} else {
		if stack.min > data {
			stack.min = data
		}

	}
	stack.value = append(stack.value[:len(stack.value)-1], data)
	stack.length = stack.length + 1
	// stack2.value = append(stack2.value, data)

}

func (stack *StackMin) Peek() (int, bool) {
	if stack.length == 0 {
		return -1, false
	}
	data := stack.value[stack.length-1]
	stack.length = stack.length - 1
	return data, true
}

func (stack *StackMin) Pop() (int, bool) {
	if stack.length == 0 {
		return -1, false
	}
	data := stack.value[stack.length-1]
	stack.value = stack.value[:stack.length-1]
	stack.length = stack.length - 1
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
		stacks:   list.New(),
		current:  nil,
	}
}

func (sos *SetOfStacks) Push(value int) {
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

func (s *Stack) Push(item int) {
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

// STACK sort
type StackSort struct {
	items []int
}

func NewStackSort() *StackSort {
	return &StackSort{items: []int{}}
}

func (s *StackSort) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *StackSort) Push(value int) {
	s.items = append(s.items, value)
}

func (s *StackSort) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	length := len(s.items)
	data := s.items[length-1]
	s.items = s.items[:length-1]
	return data, true
}

func (s *StackSort) Peek() (int, bool) {

	if s.isEmpty() {
		return -1, false
	}

	return s.items[len(s.items)-1], true

}
func (s *StackSort) isEmpty() bool {
	return len(s.items) == 0
}

func (s *StackSort) size() int {
	return len(s.items)
}

func SortStack(stack *StackSort) *StackSort {
	tempStack := NewStack()

	for !stack.isEmpty() {
		// Pop the top element from the original stack
		current, _ := stack.Pop()

		// Move elements from the temporary stack to the original stack
		for !tempStack.IsEmpty() && current > tempStack.items[len(tempStack.items)-1] {
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
	Name      string
	Timestamp time.Time
}

type AnimalShelter struct {
	dogs *list.List
	cats *list.List
}

func NewAnimalShelter() *AnimalShelter {
	return &AnimalShelter{
		dogs: list.New(),
		cats: list.New(),
	}
}

func (shelter *AnimalShelter) Enqueue(Name, AnimalType string) {
	animal := &Animal{
		Name:      Name,
		Timestamp: time.Now(),
	}
	if AnimalType == "dog" {
		shelter.dogs.PushBack(animal)
	} else if AnimalType == "cat" {
		shelter.cats.PushBack(animal)
	} else {
		fmt.Println("Invalid animal type.")

	}
}

func (shelter *AnimalShelter) DequeueAny() (Animal, error) {
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

func (shelter *AnimalShelter) dequeueCat() Animal {
	if shelter.cats.Len() == 0 {
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

// OTHERS

func isBalancedParentheses(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}

type StackR struct {
	items []rune
}

func NewStackR() *StackR {
	return &StackR{items: []rune{}}
}

func (s *StackR) Push(value rune) {
	s.items = append(s.items, value)
}

func (s *StackR) Pop() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	data := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return data, nil
}

func (s *StackR) isEmpty() bool {
	return len(s.items) == 0
}

func reverseStringWithStack(s string) string {
	stack := NewStackR()

	for _, char := range s {
		stack.Push(char)
	}

	reversed := ""
	for !stack.isEmpty() {
		char, _ := stack.Pop()
		reversed += string(char)
	}

	return reversed
}

// Reverse a Stack using Queue

type Queue struct {
	items []rune
}

func NewQueue() *Queue {
	return &Queue{
		items: []rune{},
	}
	// return &Stack{items: []int{}}
}

func (q *Queue) Push(value rune) {
	q.items = append(q.items, value)
}

func (q *Queue) Pop() (rune, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	data := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return data, nil
}

func (q *Queue) Remove() (rune, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	data := q.items[0]
	q.items = q.items[1:]
	return data, nil
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0

}

func ReverseStackWithQueue(s StackR) {
	queue := NewQueue()
	NS := NewStackR()

	for !s.isEmpty() {
		top, _ := s.Pop()
		queue.Push(top)
	}

	for !queue.isEmpty() {
		first, _ := queue.Remove()
		NS.Push(first)
	}
}

func (s *StackR) Size() int {
	return len(s.items)
}

func DeleteMiddleElement(stack *Stack) error {
	auxStack := NewStack()

	// Move half of the elements to the auxiliary stack
	for i := 0; i < stack.Size()/2; i++ {
		value, err := stack.Pop()
		if err != nil {
			return err
		}
		auxStack.Push(value)
	}

	// Remove the middle element from the original stack
	stack.Pop()

	// Move the elements back to the original stack
	for !auxStack.IsEmpty() {
		value, _ := auxStack.Pop()
		stack.Push(value)
	}

	return nil
}

func (s *Stack) ReverseWordWithStack(words string) string {
	stack := NewStack()

	for _, char := range words {
		stack.Push(int(char))
	}

	reversed := ""
	if !stack.IsEmpty() {
		char, _ := stack.Pop()
		reversed += string(rune(char))

	}

	return reversed
}

func (q *Queue) QueueUsing2Stack() *Queue {
	stack1 := NewStackR()
	stack2 := NewStackR()

	resultQueue := NewQueue()

	if !stack1.isEmpty() {
		top, _ := stack1.Pop()
		resultQueue.Push(top)
	}
	if !stack2.isEmpty() {
		top, _ := stack2.Pop()
		resultQueue.Push(top)
	}

	return resultQueue

}

func StackUsingSingleQueue(q *Queue) *StackR {
	stack := NewStackR()
	queue := NewQueue()

	for !queue.isEmpty() {
		first, _ := queue.Remove()
		stack.Push(first)
	}

	return stack

}

func PostfixExpression(v string) (rune, error) {
	stack := NewStackR()
	for _, char := range v {
		if isOperand(char) {
			operand, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			stack.items = append(stack.items, rune(operand))

		} else if isOperator(char) {
			if len(stack.items) < 2 {
				return 0, fmt.Errorf("insufficient operands for operator %c", char)

			}
			operand2 := stack.items[len(stack.items)-1]
			stack.items = stack.items[:len(stack.items)-1]

			operand1 := stack.items[len(stack.items)-1]
			stack.items = stack.items[:len(stack.items)-1]

			result := performOperation(int(operand1), int(operand2), char)
			stack.items = append(stack.items, rune(result))

		}
	}
	if len(stack.items) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack.items[0], nil

}

func performOperation(operand1, operand2 int, operator rune) int {
	switch operator {
	case '+':
		return operand1 + operand2
	case '-':
		return operand1 - operand2
	case '*':
		return operand1 * operand2
	case '/':
		return operand1 / operand2
	default:
		return 0
	}
}

func isOperand(char rune) bool {
	return char >= '0' && char <= '9'
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}



// func (s *StackR) Pop() (rune, error) {
// 	if s.isEmpty() {
// 		return 0, errors.New("stack is empty")
// 	}
// 	data := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return data, nil
// }



type StackNGE struct {
	items map[string]string
}

func (s *StackNGE) isEmpty() bool {
	return len(s.items) == 0
}

func NewStackNGE() *StackNGE {
	return &StackNGE{items: make(map[string]string)}
}

func (s *StackNGE) Push( key,value string) {
	s.items[key] ="=>" + value
}

func NGE(array []int) *StackNGE{ 
	stack := NewStackNGE()
	lengthMinus2 := len(array)-2
	var i int
	for i =1; i< lengthMinus2; i++ {
		n := i + 1
		for i <= lengthMinus2{ //74,5,2,25
			if array[n] > array[i] {
				stack.Push(strconv.Itoa(i), strconv.Itoa(n))
				break
			}
			n = n+1
		}
		stack.Push(strconv.Itoa(i), strconv.Itoa(-1))

		// for n <= lengthMinus2 { //this will also work
		// 	if array[n] > array[i] {
		// 		stack.Push(strconv.Itoa(array[i]), strconv.Itoa(array[n]))
		// 		break
		// 	}
		// 	n = n + 1
		// }
		// if n > lengthMinus2 {
		// 	stack.Push(strconv.Itoa(array[i]), strconv.Itoa(-1))
		// }
	}
	stack.Push(strconv.Itoa(len(array)-1), strconv.Itoa(-1))
	return stack
}

type StackNSE struct {
	items []string //map[string]string
}

func (s *StackNSE) isEmpty() bool {
	return len(s.items) == 0
}

func NewStackNSE() *StackNSE {
	return &StackNSE{items: []string{}}
}

func (s *StackNSE) Push(value string) {
	s.items = append(s.items, value)
}



func NSE(array [] int) *StackNSE{
	stack := NewStackNSE()
	var i int
	for i=0; i<len(array); i++ {
		// fmt.Println("solving for: ", array[i])
		if stack.isEmpty(){
			stack.Push("_")
		}
		n := i-1
		for n >=0 {
			if array[n] < array[i]{
				stack.Push(strconv.Itoa(array[n]))
				break
			}
			n--
			if n == 0{
				stack.Push("_")
			}
		}

	}
	return stack
}

func convert(s string, numRows int) string {
    // Handle special cases where numRows is 1 or numRows is greater than or equal to the length of the string
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    // Create an array of strings to represent each row
    result := make([]string, numRows)
    index, step := 0, 1

    // Iterate through the characters in the input string
    for _, char := range s {
        // Append the character to the current row
        result[index] += string(char)

        // Update the direction of movement based on the current row
        if index == 0 {
            step = 1
        } else if index == numRows-1 {
            step = -1
        }

        // Move to the next row
        index += step
    }

    // Concatenate the rows to form the converted string
    return join(result)
}

func join(strs []string) string {
    result := ""
    for _, str := range strs {
        result += str
    }
    return result
}

func main() {
	array := []int{1, 3, 0, 2, 5}//{4, 5, 2, 10, 8}
	result := NSE(array)
	fmt.Println(result)

	array2 := []int{5, 1, 9, 2, 5, 1, 7}
	result2 := NSNG(array2)
	fmt.Println(result2)

	// array3 := []int {34, 3, 31, 98, 92, 23}
	stack := NewStackSort()
	stack.Push(34)
	stack.Push(3)
	stack.Push(31)
	stack.Push(98)
	stack.Push(92)
	stack.Push(23)
	result3 := StackSorting(stack)
	fmt.Println(result3)

	arraySP := []int{100, 80, 60, 70, 60, 75, 85}
	resultSP := stockSpan(arraySP)
	fmt.Println(resultSP)


	s := "PAYPALISHIRING"
    numRows := 4
    converted := convert(s, numRows)
    fmt.Println(converted) // Output: "PAHNAPLSIIGYIR"

	// Print the result
	// for _, value := range result.items {
	// 	fmt.Printf("Smallest Element to the Left: %s\n", value)
	// }
}

func NSNG(array []int) *StackNSE{
	stack1 := NewStackNSE()
	length := len(array)-2
	var i int //{5, 1, 9, 2, 5, 1, 7}
	var seen bool = false
	for i =1; i<length; i++ {
		n := i+1
		for n <= length {
			if array[n] > array[i] && !seen{
				seen = true
				for b:= n; b<length; b++{
					if array[b] <array[n]{
						stack1.Push(strconv.Itoa(array[b]))
						seen = false
						break
					}
				}
				stack1.Push(strconv.Itoa(-1))
				seen = false
			}
			n = n+1 
		}
	}
	stack1.Push(strconv.Itoa(-1))
	return stack1
}



func StackSorting(stack *StackSort) *StackSort{
	tempStack := NewStack()
	for !stack.isEmpty(){
		current, _ := stack.Pop()
// func SortStack(stack *StackSort) *StackSort {
		for !tempStack.IsEmpty() && current > tempStack.items[len(tempStack.items)-1]  {
			top, _ := tempStack.Pop()
			stack.Push(top)

		}
		tempStack.Push(current) // = append(tempStack.items, current)
	}
	for !tempStack.IsEmpty() {
		top, _ := tempStack.Pop()
		stack.Push(top)
	}
	return stack
}

func stockSpan(SP []int) *StackSort{
	tempStack := NewStackSort()

	for i :=0; i<len(SP); i++{
		// fmt.Println("for", SP[i])
		if tempStack.IsEmpty(){
			tempStack.Push(1)
			continue //[100 80 60 70 60 75 85]
		}
		if SP[i] > SP[i-1]{
			tempStack.Push(2)
		} else {
			tempStack.Push(1)
		}
	}
	return tempStack //(*StackSort)(tempStack)
}