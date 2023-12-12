package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CanConstruct(tt.target, tt.wordBank); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
