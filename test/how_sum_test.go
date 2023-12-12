package test

import (
	"slices"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestHowSum(t *testing.T) {
	cases := []struct {
		targetSum int
		numbers   []int
		expected  []int
	}{
		{7, []int{2, 3}, []int{3, 2, 2}},
		{7, []int{5, 3, 4, 7}, []int{4, 3}},
		{7, []int{2, 4}, nil},
		{8, []int{2, 3, 5}, []int{2, 2, 2, 2}},
		{300, []int{7, 14}, nil},
	}

	for _, c := range cases {
		result := internal.HowSum(c.targetSum, c.numbers)
		if slices.Compare(result, c.expected) != 0 {
			t.Errorf("HowSum(%d, %v) == %v, want %v", c.targetSum, c.numbers, result, c.expected)
		}
	}
}
