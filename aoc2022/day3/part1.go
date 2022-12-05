package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		rucksack1 := make(map[uint8]int)
		rucksack2 := make(map[uint8]int)

		for i := 0; i < len(rucksack)/2; i++ {
			rucksack1[rucksack[i]]++
		}

		for i := len(rucksack) / 2; i < len(rucksack); i++ {
			rucksack2[rucksack[i]]++
		}

		for key := range rucksack1 {
			if _, ok := rucksack2[key]; ok {
				//fmt.Println(key)
				if key > 96 {
					result += int(key - 96)
				} else {
					result += int(key - 38)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
