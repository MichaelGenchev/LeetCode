package main

import "fmt"



func canBeIncreasing(nums []int) bool {

	res := checkIfIncreasing(nums)
	if res == true {
		return true
	}

	for i := range nums {
		res = deleteItemAndCheckIfIncreasing(i, nums)
	}

	return res

}

func main() {

	nums := []int{2,3,1,2}
	result := canBeIncreasing(nums)
	fmt.Println(result)

}
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func checkIfIncreasing(nums []int) bool {

	var prevNum int

	checked := true

	for i, num := range nums {
		if i == 0 {
			prevNum = num
			continue
		}
		
		
		if prevNum >= num {
			checked = false
		}

		prevNum = num
		
	}

	return checked

}

func deleteItemAndCheckIfIncreasing(index int, nums []int) bool {

	copyNums := nums

	newNums := RemoveIndex(copyNums, index)

	res := checkIfIncreasing(newNums)

	if res == true {
		return true
	}

	return false

}
