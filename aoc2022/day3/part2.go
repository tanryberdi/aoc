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

	var rucksacks []string
	counter := 0
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
		rucksack := scanner.Text()
		rucksacks = append(rucksacks, rucksack)

		if counter%3 == 0 {
			rucksack1 := make(map[uint8]int)
			rucksack2 := make(map[uint8]int)
			rucksack3 := make(map[uint8]int)

			for i := 0; i < len(rucksacks[0]); i++ {
				rucksack1[rucksacks[0][i]]++
			}

			for i := 0; i < len(rucksacks[1]); i++ {
				rucksack2[rucksacks[1][i]]++
			}

			for i := 0; i < len(rucksacks[2]); i++ {
				rucksack3[rucksacks[2][i]]++
			}

			for key := range rucksack1 {
				_, ok2 := rucksack2[key]
				_, ok3 := rucksack3[key]

				if ok2 && ok3 {
					//fmt.Println(key)
					if key > 96 {
						result += int(key - 96)
					} else {
						result += int(key - 38)
					}
				}
			}
			rucksacks = nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
