package main

import "fmt"



func canBeIncreasing(nums []int) bool {

	res := checkIfIncreasing(nums)
	if res == true {
		return true
	}

	for i := range nums {
		res = deleteItemAndCheckIfIncreasing(i, nums)
		if res == true {
			return true
		}
	}
	
	return res

}

func main() {

	nums := []int{1,2, 10, 5,7}
	result := canBeIncreasing(nums)
	fmt.Println(result)


	arr := []int{1,2,5,7}
	r := canBeIncreasing(arr)
	fmt.Println(r)

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

	fmt.Println(newNums)

	res := checkIfIncreasing(newNums)

	fmt.Println(res)

	if res == true {
		return true
	}

	return false

}
