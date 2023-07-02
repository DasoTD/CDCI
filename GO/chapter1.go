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


func main(){
	IsUnique("dfata")
	areCharactersUnique("abcdd")
	permuation("dad", "dda")
	// twoSum([2, 7, 11, 15], 9)
	LengthOfLongestSubstring("abcabcbb")
}

