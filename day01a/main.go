package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	lines, err := readLines(true)
	if err != nil {
		panic(err)
	}
	// var answer int
	regexGetDigits := regexp.MustCompile(`^(\d*) *(\d*)$`)
	left, right := []int{}, []int{}
	var totalDistance int
	for _, line := range lines {
		digits := regexGetDigits.FindAllStringSubmatch(line, -1)
		leftValue, _ := strconv.Atoi(digits[0][1])
		left = append(left, leftValue)
		rightValue, _ := strconv.Atoi(digits[0][2])
		right = append(right, rightValue)
	}
	sort.Ints(left)
	sort.Ints(right)
	for i, _ := range left {
		if right[i] < left[i] {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}
	fmt.Println(totalDistance)
}
