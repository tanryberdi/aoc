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
		//fmt.Println(dataStream)
		for i := 0; i < len(dataStream)-13; i++ {
			chars := make(map[string]int)
			chars[string(dataStream[i])]++
			for j := i + 1; j < i+14; j++ {
				chars[string(dataStream[j])]++
			}

			if len(chars) == 14 {
				fmt.Println(i + 14)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
