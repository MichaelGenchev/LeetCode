package main

import (
	"fmt"
	"sort"
)


func thirdMax(nums []int) int {

	if len(nums) < 3{

		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		return nums[len(nums)-1]
	}
    
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nums = unique(nums)

	return nums[len(nums)-3]
}

func main() {
	nums := []int{1,2, 1, 1, 2 , 3}
	// nums := []int{2,2,3,1}
	result := thirdMax(nums)

	fmt.Println(result)
}


func unique(s []int) []int {
    inResult := make(map[int]bool)
    var result []int
    for _, int := range s {
        if _, ok := inResult[int]; !ok {
            inResult[int] = true
            result = append(result, int)
        }
    }
    return result
}