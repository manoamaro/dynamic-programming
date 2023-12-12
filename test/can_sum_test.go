package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestCanSum(t *testing.T) {
	cases := []struct {
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
	for _, c := range cases {
		got := internal.CanSum(c.targetSum, c.numbers)
		if got != c.want {
			t.Errorf("CanSum(%d, %v) == %t, want %t", c.targetSum, c.numbers, got, c.want)
		}
	}
}
