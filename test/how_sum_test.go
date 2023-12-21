package test

import (
	"fmt"
	"slices"
	"testing"

	"manoamaro.github.com/dp/internal"
)

var howSumCases = []struct {
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

func TestHowSum(t *testing.T) {
	for _, c := range howSumCases {
		result := internal.HowSum(c.targetSum, c.numbers)
		if slices.Compare(result, c.expected) != 0 {
			t.Errorf("HowSum(%d, %v) == %v, want %v", c.targetSum, c.numbers, result, c.expected)
		}
	}
}

func TestHowSumTabulation(t *testing.T) {
	for _, test := range howSumCases {
		t.Run(fmt.Sprintf("HowSumTabulation(%d,%v)", test.targetSum, test.numbers), func(t *testing.T) {
			output := internal.HowSumTabulation(test.targetSum, test.numbers)
			if slices.Compare(output, test.expected) != 0 {
				t.Errorf("Expected, %d, Received, %d", test.expected, output)
			}
		})
	}
}
