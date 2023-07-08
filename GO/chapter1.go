package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func IsUnique(s string) bool{

	if len(s) > 128 {
		return false
	}

	ss := SortString(s)
	for i :=0; i< len(ss); i++{
		for j :=i +1; j< len(ss); j++{
			if (ss[i]==ss[j]){
				fmt.Print("false")
				return false
			}
		}
	}
	fmt.Print("true")
	return true

}
func areCharactersUnique(s string) bool{
	checker := 0
	for i :=0; i <len(s); i++ {
		bitAtIndex  := s[i] - 'a';
		if ((checker & (1 << bitAtIndex)) > 0) {
			fmt.Print("falsez")
            return false
        }
		checker = checker | (1 << bitAtIndex);
	}
	fmt.Print("truez")
	return true
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range(nums){

		if v, found := m[target- num]; found { return []int {v, idx}}
		

		m[num] = idx
	}
	return nil
}
func TwoSum(nums []int, target int) bool {
	i := 0
	j := len(nums) - 1
	sort.Ints(nums)

	for(i < j){
		if (nums[i] + nums[j] == target){
			return true // []int {nums[i], nums[j]}
		} else if (nums[i] + nums[j] < target){
			i++;
            
		} else {
            j--;
		}
			
	}
	return false;
}

// func twoSum(nums []int, target int) []int {
// 	i,j := 0, len(nums)
// 	sort.Ints(nums)
// 	for (i<j){
// 		if(nums[i]== nums[j]){
// 			return []int [nums[i], nums[j]]
// 		} else if ((nums[i] + nums[j]) < target){
// 			i++;
            
// 		} else {
//             j--;
// 		}
// 	}
// 	return nil
// }

func permuation(s1,s2 string) bool {
	if len(s1) != len(s2){
		fmt.Print("kole work")
		return false
	}
	s1 = SortString(s1)
	s2 = SortString(s2)

	for i:=0; i<len(s1); i++{
		if s1[i] != s2[i]{
			fmt.Print("ko work")
			return false
		}

	}
	fmt.Print("o work")
	return true
}

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool) // A set to keep track of characters encountered
	maxLength := 0                  // Maximum length of substring
	left := 0                       // Left pointer of the sliding window

	for right := 0; right < len(s); right++ {
		// If the current character is already in the set, move the left pointer to the right
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}
		// Add the current character to the set
		charSet[s[right]] = true
		// Update the maximum length if necessary
		if length := right - left + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func isValid(s string) bool {
	stack := []rune{} // A stack to store the opening brackets encountered

	// Define a mapping of opening brackets to their corresponding closing brackets
	bracketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		// If the current character is an opening bracket, push it onto the stack
		if isOpeningBracket(char) {
			stack = append(stack, char)
		} else {
			// If the current character is a closing bracket
			// and the stack is empty or the top of the stack does not match the current character,
			// the string is not valid
			if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != char {
				fmt.Print("false")
				return false
			}
			// Pop the top element from the stack since it matches the current closing bracket
			stack = stack[:len(stack)-1]
		}
	}

	// If there are remaining opening brackets in the stack, the string is not valid
	fmt.Print("true")
	return len(stack) == 0
}

// Helper function to check if a character is an opening bracket
func isOpeningBracket(char rune) bool {
	return char == '(' || char == '{' || char == '['
}



func LengthOfLongestSubstring(s string) int {
	l, r := 0, 0
    chset := make(map[rune]int)
    result := 0
    for r < len(s){
        ch :=  rune(s[r])
        if _, exist := chset[ch]; !exist {
            chset[ch] = 0
        } else {
            for l < r {
                if rune(s[l]) == ch {
                    l += 1
                    break
                } else {
                    delete(chset,  rune(s[l]))
                    l += 1
                }
            }
        }
        if r - l + 1 > result {
            result = r - l + 1 
        }
        r += 1
    }
	// fmt.Print(result)
    return result

}

func max(numbers map[int]int) (maxNumber int) {
    maxNumber = math.MinInt32
    for n := range numbers {
        if n > maxNumber {
            maxNumber = n
        }
    }
    return maxNumber
}

// func LengthOfLongestSubstring(s string) int{
// 	set := make(map[int]int)
// 	n,p,q,s := 0,0,0,0
// 	for i:=0; i< len(s); i++{
// 		ch := int(s[i])
// 		if _, exist := set[ch]; !exist && len(set) <= 1 {
			
// 			set[n] = p++
// 		} else if  _, exist := set[ch]; exist{
// 			set[n+1] ++
// 		} else {
// 			set[n+1] ++
// 		}
// 		n++
// 	}
// 	result :=max(set)
// 	fmt.Print(result)
// 	return result
// }

func maxProfit(prices []int) int {
	result := 0
	l,r := 0, len(prices)-1
	if r <= 1 {
 		return 0 // If the array has less than 2 elements, no profit can be achieved
	}
	for l<r {
		if prices[l]<prices[r] && prices[r]-prices[l]> result {
			result = prices[r]-prices[l]
		}
		l++
		r--
	}
	return result
    
}

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n <= 1 {
// 		return 0 // If the array has less than 2 elements, no profit can be achieved
// 	}

// 	// Initialize pointers for buying and selling days
// 	buyDay := 0
// 	sellDay := 1
// 	maxProfit := 0

// 	for sellDay < n {
// 		// Calculate the profit if we sell on the current day
// 		profit := prices[sellDay] - prices[buyDay]

// 		// If the profit is greater than the maximum profit so far, update it
// 		if profit > maxProfit {
// 			maxProfit = profit
// 		}

// 		// If the current price is lower than the price on the buy day, update the buy day
// 		if prices[sellDay] < prices[buyDay] {
// 			buyDay = sellDay
// 		}

// 		// Move the sell day pointer to the next day
// 		sellDay++
// 	}

// 	return maxProfit
// }

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0 // If there are 2 or fewer bars, no water can be trapped
	}

	leftMax := make([]int, n)  // Maximum height to the left of each bar
	rightMax := make([]int, n) // Maximum height to the right of each bar

	// Compute the maximum height to the left of each bar
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = MAX(leftMax[i-1], height[i])
	}

	fmt.Print("leftMax", leftMax)

	// Compute the maximum height to the right of each bar
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = MAX(rightMax[i+1], height[i])
	}

	// Compute the total amount of trapped water
	water := 0
	for i := 1; i < n-1; i++ {
		minHeight := MIN(leftMax[i], rightMax[i])
		if minHeight > height[i] {
			water += minHeight - height[i]
		}
	}
	fmt.Print(water)

	return water
}

// Helper function to find the maximum of two integers
func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to find the minimum of two integers
func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	// Use two boolean arrays to keep track of rows and columns that need to be set to 0
	rows := make([]bool, m)
	cols := make([]bool, n)

	// Scan the matrix to mark rows and columns that contain 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	// Set rows to 0
	for i := 0; i < m; i++ {
		if rows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// Set columns to 0
	for j := 0; j < n; j++ {
		if cols[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}


var adg = []int {0,1,0,2,1,0,1,3,2,1,2,1}
// var ze = []int [[1,1,1],[1,0,1],[1,1,1]]
func main(){
	IsUnique("dfata")
	areCharactersUnique("abcdd")
	permuation("dad", "dda")
	// twoSum([2, 7, 11, 15], 9)
	LengthOfLongestSubstring("abcabcbb")
	isValid("()")
	fmt.Print("cook")
	trap(adg)
	// setZeroes
}

