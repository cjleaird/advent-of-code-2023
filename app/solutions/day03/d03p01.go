package day03

import (
	"aoc/app/helper"
	"log"
	"strconv"
	"strings"
)

type D3P1 struct{}

func (solution *D3P1) Name() string {
	return "Day 3 (Part 1)"
}

func (solution *D3P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	symbolCoordinates := map[[2]int]struct{}{}
	numberCoordinates := map[string][][2]int{}
	// visit each line row and populate our number and symbol coordinates
	for row, line := range lines {
		populateCoordinates(row, line, numberCoordinates, symbolCoordinates)
	}
	// find which numbers are adjacent to a symbol and add 'em up
	sum := 0
	for number, allCoordinates := range numberCoordinates {
		for _, coordinates := range allCoordinates {
			_, found := findAdjacentSymbolCoordinates(number, coordinates, symbolCoordinates)
			if found {
				result, _ := strconv.Atoi(number)
				sum += result
			}
		}
	}
	return sum
}
