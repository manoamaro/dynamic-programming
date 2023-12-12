package test

import (
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
