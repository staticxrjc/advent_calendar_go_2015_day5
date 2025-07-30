package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	parser := NewStringParser()
	count := 0

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			if parser.IsNice(string(line)) {
				count++
			}
		}
		if err != nil {
			break
		}
	}

	fmt.Println("Contains nice strings", count)
}
