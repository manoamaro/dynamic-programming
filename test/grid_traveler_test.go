package test

import (
	"fmt"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestGridTraveler(t *testing.T) {
	cases := []struct {
		m, n int
		want int
	}{
		{1, 1, 1},
		{2, 3, 3},
		{3, 2, 3},
		{3, 3, 6},
		{18, 18, 2333606220},
		{30, 30, 30067266499541040},
	}
	for _, c := range cases {
		got := internal.GridTraveler(c.m, c.n)
		if got != c.want {
			t.Errorf("GridTraveler(%d, %d) == %d, want %d", c.m, c.n, got, c.want)
		}
	}
}

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
