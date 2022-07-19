package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {

	mapDeck := makeMapFromDeck(deck)

	result := checkGroupSize(mapDeck)

	return result

}

func main() {

	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}

	// deck := []int{1,1,1,2,2,2,3,3}

	result := hasGroupsSizeX(deck)

	fmt.Println(result)

}


func makeMapFromDeck(deck []int) map[int][]int{

	mapDeck := make(map[int][]int)

	for _, value := range deck {
		mapDeck[value] = append(mapDeck[value], value)
	}

	return mapDeck
}


func checkGroupSize(mapDeck map[int][]int) bool {

	prevLength := 0
	counter := 0
	for _, v := range mapDeck {

		if counter == 0 {
			prevLength = len(v)
			counter++
			continue
		}
		if prevLength != len(v){
			return false
		}
	
	}

	return true
}