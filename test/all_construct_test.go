package test

import (
	"slices"
	"testing"

	"manoamaro.github.com/dp/internal"
)

type allConstructArgs struct {
	target   string
	wordBank []string
}

var allConstructCases = []struct {
	name string
	args allConstructArgs
	want [][]string
}{
	{"", allConstructArgs{"abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}}, [][]string{{"ab", "cd", "ef"}, {"ab", "c", "def"}, {"abc", "def"}, {"abcd", "ef"}}},
	{"", allConstructArgs{"purple", []string{"purp", "p", "ur", "le", "purpl"}}, [][]string{{"purp", "le"}, {"p", "ur", "p", "le"}}},
	{"", allConstructArgs{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}}, [][]string{}},
	{"", allConstructArgs{"enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}}, [][]string{
		{"enter", "a", "p", "ot", "ent", "p", "ot"},
		{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
		{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
		{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
	}},
	{"", allConstructArgs{"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}}, [][]string{}},
}

func TestAllConstruct(t *testing.T) {

	for _, tt := range allConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.AllConstruct(tt.args.target, tt.args.wordBank)
			if len(got) != len(tt.want) {
				t.Errorf("AllConstruct() = %v, want %v", got, tt.want)
			}
			for i, way := range got {
				if slices.Compare(way, tt.want[i]) != 0 {
					t.Errorf("AllConstruct() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAllConstructTabulation(t *testing.T) {

	for _, tt := range allConstructCases {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.AllConstruct(tt.args.target, tt.args.wordBank)
			if len(got) != len(tt.want) {
				t.Errorf("AllConstruct() = %v, want %v", got, tt.want)
			}

			for _, want := range tt.want {
				contained := slices.ContainsFunc(got, func(s []string) bool {
					return slices.Compare(want, s) == 0
				})
				if !contained {
					t.Errorf("Want=%v, Got=%v", tt.want, got)
				}
			}
		})
	}
}
