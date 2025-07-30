package main

import "testing"

func TestNiceStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "ugknbfddgicrmopn is nice",
			input: "ugknbfddgicrmopn ",
			want:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := StringParser{}
			got := s.IsNice(test.input)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
