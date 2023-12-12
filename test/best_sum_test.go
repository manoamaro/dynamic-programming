package test

import (
	"slices"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestBestSum(t *testing.T) {
	cases := []struct {
		targetSum int
		numbers   []int
		expected  []int
	}{
		{7, []int{5, 3, 4, 7}, []int{7}},
		{8, []int{2, 3, 5}, []int{5, 3}},
		{8, []int{1, 4, 5}, []int{4, 4}},
		{100, []int{1, 2, 5, 25}, []int{25, 25, 25, 25}},
		{300, []int{7, 14}, nil},
	}

	for _, c := range cases {
		got := internal.BestSum(c.targetSum, c.numbers)
		if slices.Compare(got, c.expected) != 0 {
			t.Errorf("BestSum(%d, %v) == %v, expected %v", c.targetSum, c.numbers, got, c.expected)
		}
	}
}
