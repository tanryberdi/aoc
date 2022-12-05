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
		opponent := string(scanner.Text()[0])
		me := string(scanner.Text()[2])

		switch opponent {
		case "A":
			switch me {
			case "X":
				result += 3
				result += 1
				break
			case "Y":
				result += 2
				result += 6
				break
			case "Z":
				result += 3
				break
			}
		case "B":
			switch me {
			case "X":
				result += 1
				break
			case "Y":
				result += 2
				result += 3
				break
			case "Z":
				result += 3
				result += 6
				break
			}
		case "C":
			switch me {
			case "X":
				result += 1
				result += 6
				break
			case "Y":
				result += 2
				break
			case "Z":
				result += 3
				result += 3
				break
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
