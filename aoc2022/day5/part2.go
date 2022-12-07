package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	elements []rune
}

func (s *stack) push(r []rune) {
	s.elements = append(s.elements, r...)
}

func (s *stack) pop(n int) (r []rune) {
	r = s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}

func main() {
	//Read input file
	input, _ := os.Open("input1.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	//create slice of stacks
	stacks := make([]stack, 9)

	//Parsing the input
	scanner.Scan()
	for scanner.Text() != " 1   2   3   4   5   6   7   8   9" {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].addToBottom(r)
			}
		}
		scanner.Scan()
	}
	//Read empty line
	scanner.Scan()

	for scanner.Scan() {
		var toMove, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &toMove, &from, &to)
		//Move a block of elements
		stacks[to-1].push(stacks[from-1].pop(toMove))
	}

	for _, s := range stacks {
		fmt.Print(string(s.pop(1)))
	}

}
