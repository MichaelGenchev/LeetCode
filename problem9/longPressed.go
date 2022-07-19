package main

import "fmt"


func isLongPressedName(name string, typed string) bool {

	arrayTyped := makeArray(typed)

	arrayName := makeArray(name)

	mapTyped := makeMap(arrayTyped)

	mapName := makeMap(arrayName)
	
	result := true

	for k, _ := range mapTyped {
		
		if mapName[k] > mapTyped[k]{
			result = false
		}
	}

	return result
    
}

func main(){

	// name := "saeed"
	// typed := "ssaaedd"

	name := "alex"
	typed := "aalleexx"


	result := isLongPressedName(name, typed)


	fmt.Println(result)

	

}



func makeArray(str string) []string {

	result := []string{}

	chars := []rune(str)
    for i := 0; i < len(chars); i++ {
        char := string(chars[i])
		result = append(result, char)
    }

	return result
}


func makeMap(nameArr []string) map[string]int {

	mapName := make(map[string]int)

	for _, value := range nameArr{

		mapName[value] += 1
	}
	return mapName
}