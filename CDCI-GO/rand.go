// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	_"strconv"
// // // // )

// // // // type Stack struct {
// // // // 	items []int
// // // // }

// // // // func (s *Stack) isEmpty() bool {
// // // // 	return len(s.items) == 0
// // // // }

// // // // func (s *Stack) push(value int) {
// // // // 	s.items = append(s.items, value)
// // // // }

// // // // func (s *Stack) pop() int {
// // // // 	if s.isEmpty() {
// // // // 		return -1
// // // // 	}
// // // // 	top := s.items[len(s.items)-1]
// // // // 	s.items = s.items[:len(s.items)-1]
// // // // 	return top
// // // // }

// // // // func nextSmallerOfNextGreater(array []int) map[int]int {
// // // // 	nextGreaterStack := Stack{}
// // // // 	// nextSmallerStack := Stack{}

// // // // 	result := make(map[int]int)

// // // // 	for i := len(array) - 1; i >= 0; i-- {
// // // // 		// Find the next greater element
// // // // 		for !nextGreaterStack.isEmpty() && array[i] >= nextGreaterStack.items[len(nextGreaterStack.items)-1] {
// // // // 			nextGreaterStack.pop()
// // // // 		}

// // // // 		// Find the next smaller element
// // // // 		if !nextGreaterStack.isEmpty() {
// // // // 			result[array[i]] = nextGreaterStack.items[len(nextGreaterStack.items)-1]
// // // // 		} else {
// // // // 			result[array[i]] = -1
// // // // 		}

// // // // 		// Push the current element to the nextGreaterStack
// // // // 		nextGreaterStack.push(array[i])
// // // // 	}

// // // // 	return result
// // // // }

// // // // func main() {
// // // // 	array := []int{5, 1, 9, 2, 5, 1, 7}//{4, 5, 2, 10, 8}
// // // // 	result := nextSmallerOfNextGreater(array)
// // // // 	fmt.Println(result)

// // // // 	// Print the result
// // // // 	for key, value := range result {
// // // // 		fmt.Printf("For %d: Next Smaller of Next Greater is %d\n", key, value)
// // // // 	}
// // // // }

// // // // // import (
// // // // // 	"errors"
// // // // // 	"fmt"
// // // // // )

// // // // // type Stack struct {
// // // // // 	items []int
// // // // // }

// // // // // func NewStack() *Stack {
// // // // // 	return &Stack{items: []int{}}
// // // // // }

// // // // // func (s *Stack) Push(value int) {
// // // // // 	s.items = append(s.items, value)
// // // // // }

// // // // // func (s *Stack) Pop() (int, error) {
// // // // // 	if s.IsEmpty() {
// // // // // 		return 0, errors.New("stack is empty")
// // // // // 	}
// // // // // 	value := s.items[len(s.items)-1]
// // // // // 	s.items = s.items[:len(s.items)-1]
// // // // // 	return value, nil
// // // // // }

// // // // // func (s *Stack) IsEmpty() bool {
// // // // // 	return len(s.items) == 0
// // // // // }

// // // // // func (s *Stack) Size() int {
// // // // // 	return len(s.items)
// // // // // }

// // // // // func DeleteMiddleElement(stack *Stack) error {
// // // // // 	auxStack := NewStack()

// // // // // 	// Move half of the elements to the auxiliary stack
// // // // // 	fmt.Println("half")
// // // // // 	fmt.Println(stack.Size()/2)
// // // // // 	for i := 0; i < stack.Size()/2; i++ {
// // // // // 		value, err := stack.Pop()
// // // // // 		if err != nil {
// // // // // 			return err
// // // // // 		}
// // // // // 		auxStack.Push(value)
// // // // // 	}

// // // // // 	// Remove the middle element from the original stack
// // // // // 	stack.Pop()

// // // // // 	// Move the elements back to the original stack
// // // // // 	for !auxStack.IsEmpty() {
// // // // // 		value, _ := auxStack.Pop()
// // // // // 		stack.Push(value)
// // // // // 	}

// // // // // 	return nil
// // // // // }

// // // // // func PrintStack(stack *Stack) {
// // // // // 	for i := len(stack.items) - 1; i >= 0; i-- {
// // // // // 		fmt.Println(stack.items[i])
// // // // // 	}
// // // // // }

// // // // // func main() {
// // // // // 	stack := NewStack()

// // // // // 	// Push elements onto the stack
// // // // // 	for i := 1; i <= 5; i++ {
// // // // // 		stack.Push(i)
// // // // // 	}

// // // // // 	fmt.Println("Original Stack:")
// // // // // 	PrintStack(stack)

// // // // // 	// Delete the middle element
// // // // // 	err := DeleteMiddleElement(stack)
// // // // // 	if err != nil {
// // // // // 		fmt.Println("Error:", err)
// // // // // 	}

// // // // // 	fmt.Println("Stack after deleting the middle element:")
// // // // // 	PrintStack(stack)
// // // // // }


// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strconv"
// // // )

// // // type StackNSE struct {
// // // 	items []string
// // // }

// // // func (s *StackNSE) isEmpty() bool {
// // // 	return len(s.items) == 0
// // // }

// // // func NewStackNSE() *StackNSE {
// // // 	return &StackNSE{items: []string{}}
// // // }

// // // func (s *StackNSE) Push(value string) {
// // // 	s.items = append(s.items, value)
// // // }

// // // func NSNG(array []int) *StackNSE {
// // // 	stack1 := NewStackNSE()
// // // 	length := len(array) - 2
// // // 	var seen bool = false

// // // 	for i := 1; i <= length; i++ {
// // // 		n := i + 1
// // // 		for n <= length {
// // // 			if array[n] > array[i] && !seen {
// // // 				seen = true
// // // 				for b := n; b <= length; b++ {
// // // 					if array[b] < array[n] {
// // // 						stack1.Push(strconv.Itoa(array[b]))
// // // 						seen = false
// // // 						break
// // // 					}
// // // 				}
// // // 				stack1.Push(strconv.Itoa(-1))
// // // 				seen = false
// // // 			}
// // // 			n = n + 1
// // // 		}
// // // 	}
// // // 	stack1.Push(strconv.Itoa(-1))
// // // 	return stack1
// // // }

// // // func main() {
// // // 	array := []int{5, 1, 9, 2, 5, 1, 7}
// // // 	result := NSNG(array)
// // // 	fmt.Println(result)

// // // 	// Print the result
// // // 	// for _, value := range result.items {
// // // 	// 	fmt.Printf("Next Smaller of Next Greater is %d\n", value)
// // // 	// }
// // // }

// // package main

// // import (
// // 	"fmt"
// // )

// // type Stack struct {
// // 	items []int
// // }

// // func (s *Stack) isEmpty() bool {
// // 	return len(s.items) == 0
// // }

// // func (s *Stack) push(value int) {
// // 	s.items = append(s.items, value)
// // }

// // func (s *Stack) pop() int {
// // 	if s.isEmpty() {
// // 		return -1
// // 	}
// // 	top := s.items[len(s.items)-1]
// // 	s.items = s.items[:len(s.items)-1]
// // 	return top
// // }

// // func nextSmallerOfNextGreater(array []int) map[int]int {
// // 	nextGreaterStack := Stack{}
// // 	// nextSmallerStack := Stack{}

// // 	result := make(map[int]int)

// // 	for i := len(array) - 1; i >= 0; i-- {
// // 		// Find the next greater element
// // 		for !nextGreaterStack.isEmpty() && array[i] >= nextGreaterStack.items[len(nextGreaterStack.items)-1] {
// // 			nextGreaterStack.pop()
// // 		}

// // 		// Find the next smaller element
// // 		if !nextGreaterStack.isEmpty() {
// // 			result[array[i]] = nextGreaterStack.items[len(nextGreaterStack.items)-1]
// // 		} else {
// // 			result[array[i]] = -1
// // 		}

// // 		// Push the current element to the nextGreaterStack
// // 		nextGreaterStack.push(array[i])
// // 	}

// // 	return result
// // }

// // func main() {
// // 	array := []int{5, 1, 9, 2, 5, 1, 7}//{4, 5, 2, 10, 8}
// // 	result := nextSmallerOfNextGreater(array)
// // 	fmt.Println(result)
// // 	// Print the result
// // 	// for key, value := range result {
// // 	// 	fmt.Printf("For %d: Next Smaller of Next Greater is %d\n", key, value)
// // 	// }
// // }


// package main

// import (
// 	"fmt"
// )

// type StackSortR struct {
// 	items []int
// }

// func (s *StackSortR) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// func NewStackSortR() *StackSortR {
// 	return &StackSortR{items: []int{}}
// }

// func (s *StackSortR) Push(value int) {
// 	s.items = append(s.items, value)
// }

// func (s *StackSortR) Pop() int {
// 	if s.IsEmpty() {
// 		return -1
// 	}
// 	top := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return top
// }

// func (s *StackSortR) Peek() int{
// 	data:= s.items[len(s.items)-1]
// 	return data
// }

// func stockSpan(SP []int) []int {
// 	tempStack := NewStackSortR()
// 	span := make([]int, len(SP))

// 	for i := 0; i < len(SP); i++ {
// 		for !tempStack.IsEmpty() && SP[i] >= SP[tempStack.Pop()] {
// 			tempStack.Pop()
// 		}

// 		if tempStack.IsEmpty() {
// 			span[i] = i + 1
// 		} else {
// 			span[i] = i - tempStack.Peek()
// 		}

// 		tempStack.Push(i)
// 	}

// 	return span
// }

// func main() {
// 	stockPrices := []int{100, 80, 60, 70, 60, 75, 85}
// 	result := stockSpan(stockPrices)

// 	fmt.Println("Stock Span for each day:", result)
// }

package main

import "fmt"

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
		fmt.Println(index, char)
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

	fmt.Println(result)
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
    s := "PAYPALISHIRING"
    numRows := 3
    converted := convert(s, numRows)
    fmt.Println(converted) // Output: "PAHNAPLSIIGYIR"

    d := "Daso"
    fmt.Println(string(d[0]), len(d))
}
