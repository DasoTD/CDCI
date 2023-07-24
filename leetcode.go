package main

import "fmt"


func setDifference(a,b [] int) []int{
	d := [2001]bool{}
	ans := []int{}

	for _, x := range a{
		d[x + 1000] = true
	}

	for _,x := range b {
		if !d[x + 1000]{
			ans = append(ans, x)
			d[x + 1000] = true
		}
	}

	return ans
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := [][]int{ setDifference(nums1, nums2), setDifference(nums2, nums1) }
	fmt.Print(res)
    return [][]int{ setDifference(nums1, nums2), setDifference(nums2, nums1) }
}

func twoSums(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found { return []int{v,idx} } 

		m[num]=idx
	}

	return nil

}

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	maxLength:=0
	left:=0
	for right :=0; right<len(s); right++{
		for charSet[s[right]]{//////abbcdcdb
			delete(charSet, s[left])
			left++
		}

		charSet[s[right]]= true

		if length := right - left + 1; length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}


func isValid(s string) bool{
	stack := []rune{}
	bracketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if isOpeningBracket(char){
			stack = append(stack, char)
		} else {
			if len(stack)==0 || bracketMap[stack[len(stack)-1]] != char{
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Helper function to check if a character is an opening bracket
func isOpeningBracket(char rune) bool {
	return char == '(' || char == '{' || char == '['
}



var num2 = []int {2,4,6}
var num1 = []int {1,2,3}
var nums = [] int {2, 7, 11, 15}
func main (){
	findDifference(num1, num2)
	fmt.Print(twoSums(nums, 9))
	fmt.Print(lengthOfLongestSubstring("aabcdbcd"))
}