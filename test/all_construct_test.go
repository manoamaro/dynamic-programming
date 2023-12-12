package test

import (
	"slices"
	"testing"

	"manoamaro.github.com/dp/internal"
)

func TestAllConstruct(t *testing.T) {
	type args struct {
		target   string
		wordBank []string
	}
	cases := []struct {
		name string
		args args
		want [][]string
	}{
		{"", args{"purple", []string{"purp", "p", "ur", "le", "purpl"}}, [][]string{{"purp", "le"}, {"p", "ur", "p", "le"}}},
		{"", args{"abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}}, [][]string{{"ab", "cd", "ef"}, {"ab", "c", "def"}, {"abc", "def"}, {"abcd", "ef"}}},
		{"", args{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}}, [][]string{}},
		{"", args{"enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}}, [][]string{
			{"enter", "a", "p", "ot", "ent", "p", "ot"},
			{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
			{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
			{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
		}},
		{"", args{"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}}, [][]string{}},
	}

	for _, tt := range cases {
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
