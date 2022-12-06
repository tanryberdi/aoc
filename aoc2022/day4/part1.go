package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		frontBack := strings.Split(line, ",")
		//fmt.Println(frontBack)
		frontNumbers := strings.Split(frontBack[0], "-")
		backNumbers := strings.Split(frontBack[1], "-")

		//fmt.Println(frontNumbers, backNumbers)
		front1, _ := strconv.Atoi(frontNumbers[0])
		front2, _ := strconv.Atoi(frontNumbers[1])

		back1, _ := strconv.Atoi(backNumbers[0])
		back2, _ := strconv.Atoi(backNumbers[1])
		//fmt.Println(front1, front2, back1, back2)
		if (front1 <= back1 && front2 >= back2) || (front1 >= back1 && front2 <= back2) {
			result++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
