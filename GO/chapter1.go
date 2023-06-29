package main

import (
	"fmt"
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




func main(){
	IsUnique("dfata")
	areCharactersUnique("abcdd")
	permuation("dad", "dda")
	// twoSum([2, 7, 11, 15], 9)
}

