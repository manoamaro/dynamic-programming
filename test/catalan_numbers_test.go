package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestCatalanNumber(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{3, 5},
		{4, 14},
		{5, 42},
		{6, 132},
		{7, 429},
		{8, 1430},
		{9, 4862},
		{10, 16796},
		{11, 58786},
		{12, 208012},
		{13, 742900},
		{14, 2674440},
		{15, 9694845},
		{16, 35357670},
		{17, 129644790},
		{18, 477638700},
		{19, 1767263190},
		{20, 6564120420},
		{21, 24466267020},
		{22, 91482563640},
		{23, 343059613650},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			got := internal.CatalanNumbers(c.in)
			if got != c.want {
				t.Errorf("CatalanNumbers(%d) == %d, want %d", c.in, got, c.want)
			}
		})
	}
}
