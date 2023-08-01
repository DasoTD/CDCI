package main

import (
	"errors"
	"fmt"
)

type ThreeStacks struct {
	data     []int
	stackPtr [3]int
	stackSize int // The size of each stack
}

func NewThreeStacks(size int) *ThreeStacks {
	data := make([]int, size*3)
	return &ThreeStacks{
		data:      data,
		stackPtr:  [3]int{-1, -1, -1},
		stackSize: size,
	}
}

func (ts *ThreeStacks) push(stackNum int, value int) error {
	if stackNum < 0 || stackNum > 2 {
		return errors.New("Invalid stack number")
	}

	if ts.stackPtr[stackNum]+1 >= ts.stackSize {
		return errors.New("Stack overflow")
	}

	ts.stackPtr[stackNum]++
	ts.data[ts.index(stackNum)] = value
	return nil
}

func (ts *ThreeStacks) pop(stackNum int) (int, error) {
	if stackNum < 0 || stackNum > 2 {
		return 0, errors.New("Invalid stack number")
	}

	if ts.stackPtr[stackNum] == -1 {
		return 0, errors.New("Stack underflow")
	}

	value := ts.data[ts.index(stackNum)]
	ts.stackPtr[stackNum]--
	return value, nil
}

func (ts *ThreeStacks) index(stackNum int) int {
	return ts.stackSize*stackNum + ts.stackPtr[stackNum]
}

func main() {
	threeStacks := NewThreeStacks(5)

	threeStacks.push(0, 10)
	threeStacks.push(1, 20)
	threeStacks.push(2, 30)
	threeStacks.push(0, 15)
	threeStacks.push(1, 25)

	value, _ := threeStacks.pop(0)
	fmt.Println("Popped from Stack 0:", value)

	value, _ = threeStacks.pop(1)
	fmt.Println("Popped from Stack 1:", value)

	value, _ = threeStacks.pop(2)
	fmt.Println("Popped from Stack 2:", value)
}
