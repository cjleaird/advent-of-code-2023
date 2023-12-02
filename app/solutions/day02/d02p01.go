package day02

import (
	"aoc/app/helper"
	"log"
	"strings"
)

type D2P1 struct{}

func (solution *D2P1) Name() string {
	return "Day 2 (Part 1)"
}

func (solution *D2P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		id, sets, err := findIdAndCubeSets(line)
		// we'll just ignore the error and continue, we assume valid input
		// so this almost assuredly would be caused by an EOF
		if err != nil {
			continue
		}
		isGamePossible := true
		for _, set := range sets {
			red, green, blue := findRgbCubeCounts(set)
			if red > RedLimit || green > GreenLimit || blue > BlueLimit {
				isGamePossible = false
				break
			}
		}
		if isGamePossible {
			sum += id
		}
	}
	return sum
}

const RedLimit = 12
const GreenLimit = 13
const BlueLimit = 14
