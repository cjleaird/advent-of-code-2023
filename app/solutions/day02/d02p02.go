package day02

import (
	"aoc/app/helper"
	"log"
	"strings"
)

type D2P2 struct{}

func (solution *D2P2) Name() string {
	return "Day 2 (Part 2)"
}

func (solution *D2P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		_, sets, err := findIdAndCubeSets(line)
		// we'll just ignore the error and continue, we assume valid input
		// so this almost assuredly would be caused by an EOF
		if err != nil {
			continue
		}
		// find the min rgb pull counts in our available cube sets
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, set := range sets {
			red, green, blue := findRgbCubeCounts(set)
			if red > minRed {
				minRed = red
			}
			if green > minGreen {
				minGreen = green
			}
			if blue > minBlue {
				minBlue = blue
			}
		}
		sum += minRed * minGreen * minBlue

	}
	return sum
}
