package main

import (
	"fmt"
	"strings"
)

func main() {
  str1 := "educative.io"
	fmt.Println(str1, "io", strings.Contains(str1, "io"))
	fmt.Println(str1, "shot", strings.Contains(str1, "shot"))
	fmt.Println(str1, "", strings.Contains(str1, ""))

	var s string = "Hello, World"
  
	for index, character := range(s){
	  fmt.Printf("The character %c is in position %d \n", character, index)
	}
}


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range(nums){

		if v, found := m[target- num]; found { return []int {v, idx}}
		

		m[num] = idx
	}
	return nil
}

// two_sum([2, 7, 11, 15], 9)


// In Golang, make() is used for slices, maps, or channels. make() allocates memory on the heap and initializes and puts zero or empty strings into values. Unlike new(), make() returns the same type as its argument.

// Slice: The size determines the length. The capacity of the slice is equal to or greater than its length. For instance, make([]int, 0, 10) allots an array of size 10 and returns a slice of length 0 and capacity 10.

// Map: A map is allocated with a specified amount of space by default, hence the size argument can be left out.

// Channel: The buffer for the channel is initialized with the given buffer capacity. If the capacity is zero, the channel is unbufferred.