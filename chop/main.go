package chop

import (
	"math"
)

func main() {

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
