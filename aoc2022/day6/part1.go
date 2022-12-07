package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStream := scanner.Text()
		fmt.Println(dataStream)
		for i := 0; i < len(dataStream)-3; i++ {
			chars := make(map[string]int)
			chars[string(dataStream[i])]++
			for j := i + 1; j < i+4; j++ {
				chars[string(dataStream[j])]++
			}

			if len(chars) == 4 {
				fmt.Println(i + 4)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
