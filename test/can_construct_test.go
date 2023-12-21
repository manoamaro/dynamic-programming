package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

var canConstructCases = []struct {
	name     string
	target   string
	wordBank []string
	want     bool
}{
	{
		name:     "empty target",
		target:   "",
		wordBank: []string{"cat", "dog", "fish"},
		want:     true,
	},
	{
		name:     "constructible target",
		target:   "fishdog",
		wordBank: []string{"cat", "dog", "fish"},
		want:     true,
	},
	{
		name:     "unconstructible target",
		target:   "fishcat",
		wordBank: []string{"dog", "fish"},
		want:     false,
	},
} 

func TestCanConstruct(t *testing.T) {
	for _, tt := range canConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CanConstruct(tt.target, tt.wordBank); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanConstructTabulation(t *testing.T) {
	for _, tt := range canConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CanConstructTabulation(tt.target, tt.wordBank); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
