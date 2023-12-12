package test

import (
	"fmt"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestGridTravelerTabulation(t *testing.T) {
	cases := []struct {
		m, n     int
		expected int
	}{
		{3, 3, 6},
		{2, 3, 3},
		{3, 2, 3},
		{3, 0, 0},
		{0, 3, 0},
		{18, 18, 2333606220},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d,%d", c.m, c.n), func(t *testing.T) {
			got := internal.GridTravelerTabulation(c.m, c.n)
			if got != c.expected {
				t.Errorf("Expected=%d, Got=%d", c.expected, got)
			}
		})
	}
}
