package test

import (
	"fmt"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestFibTabulation(t *testing.T) {
	cases := []struct {
		arg, expected int
	}{
		{6, 8},
		{7, 13},
		{8, 21},
		{50, 12586269025},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("fib of %d", c.arg), func(t *testing.T) {
			got := internal.FibTabulation(c.arg)
			if got != c.expected {
				t.Errorf("Expected=%d, Actual=%d", c.expected, got)
			}
		})
	}
}
