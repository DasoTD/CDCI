package main

import (
	"fmt"
)

func maximumOperations(s string) int {
	uniqueSubstrings := make(map[string]bool)
	n := len(s)

	for i := 0; i < n; i++ {
		substr := ""
		for j := i+1; j < n; j++ {
			substr += string(s[j])

			if s[j] == s[i] && (j-i+1)%2 == 0 {
				uniqueSubstrings[substr] = true
			}
		}
	}

	maxDeletions := 0
	for range uniqueSubstrings {
		maxDeletions++
	}

	return maxDeletions
}



// 
func ddeleteString(s string) int {
	left, right := 0, len(s)/2+1

	for left < right {
		mid := left + (right-left)/2

		if canDelete(s, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}

func deleteString(s string) int {
	n := len(s)
    // fmt.Println(n)
	dp := make([]int, n)

	// for i := n - 2; i >= 0; i--
	for i:=0; i<n/2; i++ {
        // fmt.Println(i)
		dp[i] = 1 // at least one operation needed (deleting the entire string)

        for j := i + 1; j < n; j++ {
            // fmt.Println(j)
            // fmt.Println(string(s[j]))
            // break
            if (2*j - i) < n && s[i] == s[j] && s[j:2*j-i] == s[i:j] {
                dp[i] = max(dp[i], 2+dp[j])
            }
        }
	}
	result := 0
	for _, val := range dp {
		if val > result {
			result = val
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canDelete(s string, operations int) bool {
	// n := len(s)

	for i := 0; i < operations; i++ {
		if s[i] != s[i+operations] {
			return false
		}
	}

	return true
}



// package main

// import "fmt"

func longestCommonPrefix(s string) int {
	n := len(s)
	result := n
	for i := 0; i < n-1; i++ {
		j := 0
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		result += j-1
	}
	return result -n
	// return prefix
}

// func main() {
// 	strs := []string{"apple", "appetizer", "apparatus"}
// 	result := longestCommonPrefix(strs)
// 	fmt.Println(result) // Output: "app"
// }



func main() {
	s := "abcabcdabc"
	result := deleteString(s)
	fmt.Println(result) // Output: 1

	strs := "azbazbzaz"
	results := longestCommonPrefix(strs)
	fmt.Println(results) // Output: "app"

}
