package main

import "testing"

var simpleSearches = []struct {
	value   int   // value we are looking for
	numbers []int // array of numbers to look up
	result  int
}{
	{3, []int{}, -1},
	{3, []int{1}, -1},
	{1, []int{1}, 0},

	{1, []int{1, 3, 5}, 0},
	{3, []int{1, 3, 5}, 1},
	{5, []int{1, 3, 5}, 2},
	{0, []int{1, 3, 5}, -1},
	{2, []int{1, 3, 5}, -1},
	{4, []int{1, 3, 5}, -1},
	{6, []int{1, 3, 5}, -1},

	{1, []int{1, 3, 5, 7}, 0},
	{3, []int{1, 3, 5, 7}, 1},
	{5, []int{1, 3, 5, 7}, 2},
	{7, []int{1, 3, 5, 7}, 3},
	{0, []int{1, 3, 5, 7}, -1},
	{2, []int{1, 3, 5, 7}, -1},
	{4, []int{1, 3, 5, 7}, -1},
	{6, []int{1, 3, 5, 7}, -1},
	{8, []int{1, 3, 5, 7}, -1},
}

var rotatedSearches = []struct {
	value   int   // value we are looking for
	numbers []int // array of numbers to look up
	result  int
}{
	{1, []int{3, 4, 5, 1, 2}, 3},
	{3, []int{5, 6, 7, 8, 9, 10, 1, 2, 3}, 8},
	{30, []int{5, 6, 7, 8, 9, 10, 1, 2, 3}, -1},
	{10, []int{30, 40, 50, 10, 20}, 3},
	// fail when using a sorted array
	{1, []int{1, 3, 5}, -1},
}

func TesChop(t *testing.T) {
	for _, s := range simpleSearches {
		result := Chop(s.value, s.numbers)
		if result != s.result {
			t.Errorf("Chop(%d, %v): expected %d, got %d", s.value, s.numbers, s.result, result)
		}
	}
}

func TestRecursiveChop(t *testing.T) {
	for _, s := range simpleSearches {
		result := RecursiveChop(s.value, s.numbers)
		if result != s.result {
			t.Errorf("RecursiveChop(%d, %v): expected %d, got %d", s.value, s.numbers, s.result, result)
		}
	}
}

func TestChopInRotatedArray(t *testing.T) {
	for _, s := range rotatedSearches {
		result := ChopInRotatedArray(s.value, s.numbers)
		if result != s.result {
			t.Errorf("ChopInRotatedArray(%d, %v): expected %d, got %d", s.value, s.numbers, s.result, result)
		}
	}
}

func TestFindPivot(t *testing.T) {
	values := []int{3, 4, 5, 1, 2}
	result := FindPivot(values)
	if result != 2 {
		t.Errorf("FindPivot(%v): expected %d, got %d", values, 2, result)
	}
}
