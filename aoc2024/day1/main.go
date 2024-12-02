package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absByMe(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var nums []string
	nums1 := make([]int, 1000)
	nums2 := make([]int, 1000)
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), " ")

		nums = strings.Split(line, " ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		nums1[index] = num1
		nums2[index] = num2
		index++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)
	result := 0
	for i := 0; i < 1000; i++ {
		result += absByMe(nums1[i], nums2[i])
	}

	fmt.Print(result)
}
