package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Fields(line)

		leftInt, _ := strconv.Atoi(lineSplit[0])
		rightInt, _ := strconv.Atoi(lineSplit[1])

		left = append(left, leftInt)
		right = append(right, rightInt)

	}
	// sort both arrrays
	sort.Ints(left)
	sort.Ints(right)

	var distance int
	for i := 0; i < len(left); i++ {
		distance += abs(left[i] - right[i])
	}

	fmt.Println(distance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
