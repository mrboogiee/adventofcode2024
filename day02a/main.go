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
	regexGetDigits := regexp.MustCompile(`(\d*)`)
	// left, right := []int{}, []int{}
	invalids := 0
	for _, line := range lines {
		digits := regexGetDigits.FindAllString(line, -1)
		direction := "none"
		lastDigit := 0
		valid := true
		for i, digit := range digits {
			intValue, _ := strconv.Atoi(digit)
			if i == 0 {
				lastDigit = intValue
			} else if i == 1 {
				direction = getDirection(lastDigit, intValue)
			} else if i > 1 {
				newDirection := getDirection(lastDigit, intValue)
				if newDirection != direction {
					valid = false
				}
			}
			if valid && i > 0 {
				distance := getDistance(lastDigit, intValue)
				if distance < 1 || distance > 3 {
					valid = false
				}
			}
			lastDigit = intValue
		}
		if !valid {
			invalids++
		}
	}
	fmt.Println(len(lines))
	fmt.Println(invalids)
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
