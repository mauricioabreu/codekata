package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You must pass value and numbers")
		os.Exit(1)
	}
	stringNumbers := strings.Split(os.Args[2], ",")
	numbers := make([]int, len(stringNumbers))

	for i, n := range stringNumbers {
		nInt, _ := strconv.Atoi(n)
		numbers[i] = nInt
	}

	value, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Binary chop: value %d in %v numbers: %d\n", value, numbers, Chop(value, numbers))
}

// Chop : finds the position of value in a sorted array of values
func Chop(value int, numbers []int) int {
	leftSide := 0
	rightSide := len(numbers) - 1

	for leftSide <= rightSide {
		middle := int(math.Floor(float64((leftSide + rightSide) / 2)))
		if numbers[middle] > value {
			rightSide = middle - 1
		} else if numbers[middle] < value {
			leftSide = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
