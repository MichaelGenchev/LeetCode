package main

import "fmt"


func canPlaceFlowers(flowerbed []int, n int) bool {

	mapFlowerbed := makeMap(flowerbed)

	counter := 0
	

	for k, _ := range flowerbed {

		if mapFlowerbed[k] == false {
			continue
		}
		res := checkIfPossible(k, mapFlowerbed)
		if res == true {
			counter++
		}

		
	}

	fmt.Println(mapFlowerbed)
	if counter >= n {
		return true
	}
	
	return false

}



func main(){
	// flowerbed := []int{0,0,0,0,0,0,0,0,1}
	flowerbed := []int{1,0,0,0,1}
	n := 2


	result := canPlaceFlowers(flowerbed, n)

	fmt.Println(result)
}

func makeMap(flowerbed []int) map[int]bool {

	mapFlowerbed := make(map[int]bool)

	for i, value := range flowerbed {

		
		if value == 1 {
			mapFlowerbed[i] = false
			continue
		}
		mapFlowerbed[i] = true
	}

	return mapFlowerbed
}

func checkIfPossible(index int, mapFlowerbed map[int]bool) bool {
	
	before := mapFlowerbed[index - 1]
	after := mapFlowerbed[index + 1]
	

	if index == 0 && after == true {
		mapFlowerbed[index] = false
		return true
	}

	if before == true && after == true {
		mapFlowerbed[index] = false
		return true
	}
	
	return false

}