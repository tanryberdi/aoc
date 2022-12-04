package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	var allSums []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			allSums = append(allSums, sum)
			sum = 0
			continue
		}

		item := scanner.Text()
		itemToInt, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		sum += itemToInt
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	allSums = append(allSums, sum)
	sort.Ints(allSums)

	itemsSum := allSums[len(allSums)-1] + allSums[len(allSums)-2] + allSums[len(allSums)-3]
	fmt.Println(itemsSum)

}
