package main

import "strings"

type StringParser struct {
	vowel string
}

func (p *StringParser) IsNice(input string) bool {
	if !p.enoughVowels(input) {
		return false
	}
	if !p.doubleLetter(input) {
		return false
	}
	if !p.containsGoodCharacters(input) {
		return false
	}
	return true
}

func (p *StringParser) containsGoodCharacters(input string) bool {
	if strings.Contains(input, "ab") {
		return false
	}
	if strings.Contains(input, "cd") {
		return false
	}
	if strings.Contains(input, "pq") {
		return false
	}
	if strings.Contains(input, "xy") {
		return false
	}
	return true
}

func (p *StringParser) enoughVowels(input string) bool {
	vowels := 0
	for _, s := range input {
		for _, vowel := range p.vowel {
			if s == vowel {
				vowels++
			}
		}
	}
	if vowels < 3 {
		return false
	}
	return true
}

func (p *StringParser) doubleLetter(input string) bool {
	var lastChar = ""
	for _, c := range input {
		if lastChar == string(c) {
			return true
		} else {
			lastChar = string(c)
		}
	}
	return false
}

func NewStringParser() *StringParser {
	return &StringParser{
		vowel: "aeiou",
	}
}
