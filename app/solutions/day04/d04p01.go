package day04

import (
	"aoc/app/helper"
	"log"
	"math"
	"strings"
)

type D4P1 struct{}

func (solution *D4P1) Name() string {
	return "Day 4 (Part 1)"
}

func (solution *D4P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		winningNumbers := map[int]struct{}{}
		pickedNumbers := map[int]struct{}{}
		err := populateCardNumbers(line, winningNumbers, pickedNumbers)
		if err != nil {
			continue
		}
		matchCount := findMatchCount(winningNumbers, pickedNumbers)
		if matchCount == 0 {
			continue
		}
		// one point for first matching number, and then doubled for every additional
		// matching number
		sum += int(math.Pow(2, float64(matchCount)-1))

	}
	return sum
}
