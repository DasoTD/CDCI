package main

import (
	"fmt"
	"sort"
)


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


func maxProfit(prices []int) int {
	n := len(prices)
	if n <1{
		return 0
	}
	buyday, sellday, maxProfit :=0, 1,0

	for sellday <n{
		profit := prices[sellday] - prices[buyday]
		if profit > maxProfit{
			maxProfit = profit
		}
		if prices[sellday]< prices[buyday]{
			buyday = sellday
		}
		sellday++
	}
	return maxProfit
}


func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}
		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}

var num2 = []int {2,4,6}
var num1 = []int {1,2,3}
var nums = [] int {2, 7, 11, 15}
var dd = [] int {7,1,5,3,6,4}
func main (){
	findDifference(num1, num2)
	fmt.Print(twoSums(nums, 9))
	fmt.Println(lengthOfLongestSubstring("aabcdbcd"))
	fmt.Println(maxProfit(dd))
}