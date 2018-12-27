package chop

import "testing"

var searches = []struct {
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

func TestMain(t *testing.T) {
	for _, s := range searches {
		result := Chop(s.value, s.numbers)
		if result != s.result {
			t.Errorf("Chop(%d, %v): expected %d, got %d", s.value, s.numbers, s.result, result)
		}
	}
}
