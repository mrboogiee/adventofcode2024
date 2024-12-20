package main

import (
	"bufio"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(production bool) ([]string, error) {
	path := "input.txt"
	if production == false {
		path = "example.txt"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
