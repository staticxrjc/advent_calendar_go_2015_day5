package main

import (
	"testing"
)

func TestNiceStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "ugknbfddgicrmopn is nice",
			input: "ugknbfddgicrmopn",
			want:  true,
		},
		{
			name:  "aaa is nice",
			input: "aaa",
			want:  true,
		},
		{
			name:  "jchzalrnumimnmhp is not nice",
			input: "jchzalrnumimnmhp",
			want:  false,
		},
		{
			name:  "haegwjzuvuyypxyu is not nice",
			input: "haegwjzuvuyypxyu",
			want:  false,
		},
		{
			name:  "dvszwmarrgswjxmb is not nice",
			input: "dvszwmarrgswjxmb",
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewStringParser()
			got := s.IsNice(test.input)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
