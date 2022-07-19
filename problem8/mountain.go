package main

import "fmt"

func validMountainArray(arr []int) bool {

	if len(arr) < 3{
		return false
	}

	index := 0

	
	index = up(index, arr)


	if index ==0 || index == len(arr) - 1{
		return false
	}

	index = down(index, arr)
	

	return index == len(arr)-1
    
}

func main() {
	// mountain := []int{0,3,2,1}
	mountain := []int{4,3,2,1}

	result := validMountainArray(mountain)

	fmt.Println(result)
}


func up(index int, arr []int) int {
	for index+1 < len(arr) && arr[index] < arr[index+1]{
		index++
	}
	return index
}


func down(index int , arr []int) int {
	for index + 1 < len(arr) && arr[index] > arr[index+1]{
		index++
	}
	return index
}