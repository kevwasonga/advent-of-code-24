package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var count int // Counter for safe lines
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var nums []int
		line := scanner.Text()

		// Split the line into parts and convert to integers
		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Skipping invalid input:", part)
				continue
			}
			nums = append(nums, num)
		}

		// Check safety for the current line
		if isAdjacent(nums) && isCrease(nums) {
			count++
		}
	}

	// Handle scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the count of safe lines
	fmt.Println("Count of safe lines:", count)
}

// Checks if the sequence is strictly increasing or decreasing
func isCrease(crease []int) bool {
	isIncrease := true
	isDecrease := true

	for i := 0; i < len(crease)-1; i++ {
		if crease[i] < crease[i+1] {
			isDecrease = false
		} else if crease[i] > crease[i+1] {
			isIncrease = false
		}

		if !isIncrease && !isDecrease {
			return false
		}
	}
	return isDecrease || isIncrease
}

// Checks if the differences between adjacent numbers are within 1, 2, or 3
func isAdjacent(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		diff := abs(input[i] - input[i+1])
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

// Computes the absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
