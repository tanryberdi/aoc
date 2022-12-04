package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxSum := 0
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if maxSum < sum {
				maxSum = sum
			}
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

	if maxSum < sum {
		maxSum = sum
	}
	fmt.Println(maxSum)
}
