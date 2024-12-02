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
	regexGetDigits := regexp.MustCompile(`^(\d*) *(\d*)$`)
	left, right := []int{}, []int{}
	for _, line := range lines {
		digits := regexGetDigits.FindAllStringSubmatch(line, -1)
		leftValue, _ := strconv.Atoi(digits[0][1])
		left = append(left, leftValue)
		rightValue, _ := strconv.Atoi(digits[0][2])
		right = append(right, rightValue)
	}
	// sort.Ints(left)
	// sort.Ints(right)
	totalSimilarities := 0
	for _, leftValue := range left {
		leftSimilarities := 0
		for _, rightValue := range right {
			if rightValue == leftValue {
				leftSimilarities++
			}
		}
		totalSimilarities += (leftValue * leftSimilarities)
	}
	fmt.Println(totalSimilarities)
}
