package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, err := readLines(true)
	if err != nil {
		panic(err)
	}
	// left, right := []int{}, []int{}
	invalids := 0
	for _, line := range lines {
		invalids = processLine(line, invalids)
	}
	fmt.Println(len(lines) - invalids)
}

func getDirection(first, second int) string {
	if first > second {
		return "down"
	} else if second > first {
		return "up"
	} else {
		return "none"
	}
}

func getDistance(first, second int) int {
	if first > second {
		return first - second
	} else {
		return second - first
	}
}

func processLine(line string, invalids int) int {
	regexGetDigits := regexp.MustCompile(`(\d*)`)
	digits := regexGetDigits.FindAllString(line, -1)
	valid := processDigits(digits)
	if !valid {
		invalids++
	}
	return invalids
}

func processDigits(digits []string) bool {
	valid := true
	lastDigit := 0
	direction := "none"
	for i, digit := range digits {
		valid, lastDigit, direction = processDigit(i, digit, lastDigit, direction, valid)
	}
	return valid
}
func processDigit(index int, digit string, lastDigit int, direction string, valid bool) (bool, int, string) {
	intValue, _ := strconv.Atoi(digit)
	if index == 0 {
		lastDigit = intValue
	} else if index == 1 {
		direction = getDirection(lastDigit, intValue)
	} else if index > 1 {
		newDirection := getDirection(lastDigit, intValue)
		if newDirection != direction {
			valid = false
		}
	}
	if valid && index > 0 {
		distance := getDistance(lastDigit, intValue)
		if distance < 1 || distance > 3 {
			valid = false
		}
	}
	lastDigit = intValue
	return valid, lastDigit, direction
}
