package main

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




