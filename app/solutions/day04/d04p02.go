package day04

import (
	"aoc/app/helper"
	"log"
	"strings"
)

type D4P2 struct{}

func (solution *D4P2) Name() string {
	return "Day 4 (Part 2)"
}

func (solution *D4P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	sum := 0
	cardsToCount := map[int]int{}
	for index, line := range lines {
		winningNumbers := map[int]struct{}{}
		pickedNumbers := map[int]struct{}{}
		err := populateCardNumbers(line, winningNumbers, pickedNumbers)
		if err != nil {
			continue
		}
		// add one for visiting the original card
		cardsToCount[index] = cardsToCount[index] + 1
		matchCount := findMatchCount(winningNumbers, pickedNumbers)
		// for all succeeding cards, up to the match count, add one for all
		// instances of this card (original + copies)
		for i := 1; i <= matchCount; i++ {
			cardsToCount[index+i] = cardsToCount[index+i] + cardsToCount[index]
		}
	}
	for _, count := range cardsToCount {
		sum += count
	}
	return sum
}
