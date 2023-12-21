package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestMinimizeSteps(t *testing.T) {
	cases := []struct {
		k        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 3},
		{5, 4},
		{6, 4},
		{7, 5},
		{8, 4},
		{9, 5},
		{10, 5},
		{14, 6},
	}

	for _, c := range cases {
		actual := internal.MinimizeSteps(c.k)
		if actual != c.expected {
			t.Errorf("MinimizeSteps(%d) == %d, expected %d", c.k, actual, c.expected)
		}
	}
}

func TestMinimizeStepsTabulation(t *testing.T) {
	cases := []struct {
		k        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 3},
		{5, 4},
		{6, 4},
		{7, 5},
		{8, 4},
		{9, 5},
		{10, 5},
		{14, 6},
	}

	for _, c := range cases {
		actual := internal.MinimizeStepsTabulation(c.k)
		if actual != c.expected {
			t.Errorf("MinimizeSteps(%d) == %d, expected %d", c.k, actual, c.expected)
		}
	}
}
