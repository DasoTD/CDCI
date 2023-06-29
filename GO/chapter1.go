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


func main(){
	IsUnique("dfata")
}

