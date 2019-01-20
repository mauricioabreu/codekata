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

// Chop finds the position of value in a sorted array of values
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

// RecursiveChop finds the position of value in a sorted array of values (recursively)
func RecursiveChop(value int, numbers []int) int {
	return chop(value, numbers, 0, len(numbers)-1)
}

func chop(value int, numbers []int, leftSide int, rightSide int) int {
	middle := int(math.Floor(float64((leftSide + rightSide) / 2)))
	if leftSide <= rightSide {
		if numbers[middle] > value {
			return chop(value, numbers, leftSide, middle-1)
		} else if numbers[middle] < value {
			return chop(value, numbers, middle+1, rightSide)
		} else {
			return middle
		}
	}
	return -1
}

// ChopInRotatedArray finds the position of a value in a rotate array with
// a pivot. This pivot element is the only element for which next element to it is smaller than it.
// For example: 3 4 5 1 2 -> pivot is 5 (index 2)
func ChopInRotatedArray(value int, numbers []int) int {
	var index int
	pivot := FindPivot(numbers)

	if pivot == -1 {
		return Chop(value, numbers)
	}

	if numbers[pivot] < value {
		index = Chop(value, numbers[0:pivot-1])
	} else {
		index = Chop(value, numbers[pivot+1:len(numbers)])
	}
	return index + len(numbers[0:pivot+1])
}

func FindPivot(numbers []int) int {
	return findPivot(numbers, 0, len(numbers)-1)
}

func findPivot(numbers []int, leftSide int, rightSide int) int {
	var compareTo int

	// We don't have a pivot
	if rightSide < leftSide {
		return -1
	}

	middle := int(math.Floor(float64((leftSide + rightSide) / 2)))
	if numbers[middle] > rightSide {
		leftSide = middle + 1
		compareTo = leftSide
	} else if numbers[middle] < leftSide {
		rightSide = middle - 1
		compareTo = rightSide
	}

	if numbers[middle] > numbers[compareTo] {
		return middle
	}

	return findPivot(numbers, leftSide, rightSide)
}
