package main

import "fmt"


func canPlaceFlowers(flowerbed []int, n int) bool {

	mapFlowerbed := makeMap(flowerbed)

	counter := 0
	

	for k, v := range mapFlowerbed {
		if v == false {
			continue
		}
		res := checkIfPossible(k, v, mapFlowerbed, flowerbed)
		if res == true {
			counter++
		}

		
	}

	if counter >= n {
		return true
	}
	
	return false

}



func main(){
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

func checkIfPossible(index int, value bool, mapFlowerbed map[int]bool, flowerbed []int) bool {
	
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