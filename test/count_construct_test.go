package test

import (
	"testing"

	"manoamaro.github.com/dp/internal"
)

type countConstructArgs struct {
	target   string
	wordBank []string
}

var countConstructCases = []struct {
	name string
	args countConstructArgs
	want int
}{
	{"", countConstructArgs{"purple", []string{"purp", "p", "ur", "le", "purpl"}}, 2},
	{"", countConstructArgs{"abcdef", []string{"ab", "abc", "cd", "def", "abcd"}}, 1},
	{"", countConstructArgs{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}}, 0},
	{"", countConstructArgs{"enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}}, 4},
	{"", countConstructArgs{"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}}, 0},
}

func TestCountConstruct(t *testing.T) {
	for _, tt := range countConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CountConstruct(tt.args.target, tt.args.wordBank); got != tt.want {
				t.Errorf("CountConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountConstructTabulation(t *testing.T) {
	for _, tt := range countConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.CountConstructTabulation(tt.args.target, tt.args.wordBank); got != tt.want {
				t.Errorf("CountConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
