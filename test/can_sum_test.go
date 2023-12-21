package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

var canSumCases = []struct {
	targetSum int
	numbers   []int
	want      bool
}{
	{7, []int{2, 3}, true},
	{7, []int{5, 3, 4, 7}, true},
	{7, []int{2, 4}, false},
	{8, []int{2, 3, 5}, true},
	{300, []int{7, 14}, false},
}

func TestCanSum(t *testing.T) {
	for _, c := range canSumCases {
		got := internal.CanSum(c.targetSum, c.numbers)
		if got != c.want {
			t.Errorf("CanSum(%d, %v) == %t, want %t", c.targetSum, c.numbers, got, c.want)
		}
	}
}

func TestCanSumTabulation(t *testing.T) {
	for _, c := range canSumCases {
		got := internal.CanSumTabulation(c.targetSum, c.numbers)
		if got != c.want {
			t.Errorf("CanSumTabulation(%d, %v) == %v, want %v", c.targetSum, c.numbers, got, c.want)
		}
	}
}
