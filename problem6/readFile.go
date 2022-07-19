package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("/Users/genchevm/leetcode/LeetCode/problem6/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


	counter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        text := scanner.Text()
		counter ++
		if counter == 10 {
			fmt.Println(text)
		}
    }

	if counter < 10 {
		fmt.Println("Not enough lines")
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}


