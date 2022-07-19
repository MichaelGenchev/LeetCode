package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)


func main() {

	// sentence := "cat and  dog"

	sentence := "!this  1-s b8d!"

	result := countValidWords(sentence)

	fmt.Println(result)

}

func countValidWords(sentence string) int {

	arrayWords := strings.Fields(sentence)

	validWords := filterWords(arrayWords)

	return len(validWords)
}


func filterWords(words []string) []string {
	validwords := []string{}
	for _, word := range words{

		matched, err := regexp.MatchString(`/^([a-z]+(-?[a-z]+)?)?(!|\\.|,)?$/`, word)
		if err != nil {
			log.Fatal(err)
		}

		if matched == true{
			validwords = append(validwords, word)
		}

	}

	return validwords
}