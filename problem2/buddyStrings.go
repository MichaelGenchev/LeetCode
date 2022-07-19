package main

import "fmt"


func main(){

	string := "ab"

	goal := "ab"

	result := buddyStrings(string, goal)

	fmt.Println(result)
}

func buddyStrings(A, B string) bool {
	if len(A) != len(B) {
		return false
	}
	diff := make([]int, 0)
	dic := make(map[rune]int)
	for rk, r := range A {
		dic[r]++
		if A[rk] != B[rk] {
			diff = append(diff, rk)
			if len(diff) > 2 {
				return false
			}
		}
	}
	if len(diff) == 0 {
		for _, num := range dic {
			if num > 1 {
				return true
			}
		}
		return false
	}
	if len(diff) == 1 {
		return false
	}
	tmp := []byte(A)
	tmp[diff[0]], tmp[diff[1]] = tmp[diff[1]], tmp[diff[0]]
	return string(tmp) == B
}