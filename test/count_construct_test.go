package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestCountConstruct(t *testing.T) {
	type args struct {
		target   string
		wordBank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"purple", []string{"purp", "p", "ur", "le", "purpl"}}, 2},
		{"", args{"abcdef", []string{"ab", "abc", "cd", "def", "abcd"}}, 1},
		{"", args{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}}, 0},
		{"", args{"enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}}, 4},
		{"", args{"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CountConstruct(tt.args.target, tt.args.wordBank); got != tt.want {
				t.Errorf("CountConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
