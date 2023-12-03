package day03

import (
	"aoc/app/helper"
	"log"
	"strconv"
	"strings"
)

type D3P2 struct{}

func (solution *D3P2) Name() string {
	return "Day 3 (Part 2)"
}

func (solution *D3P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	// visit each line row and populate our number and symbol coordinates
	symbolCoordinates := map[[2]int]struct{}{}
	numberCoordinates := map[string][][2]int{}
	for row, line := range lines {
		populateCoordinates(row, line, numberCoordinates, symbolCoordinates)
	}
	// let's assemble all part numbers that are adjacent to a symbol, keying on the symbol's coordinates
	symbolCoordinatesToPartNumbers := map[[2]int][]int{}
	for number, coordinatesList := range numberCoordinates {
		for _, coordinates := range coordinatesList {
			symbolCoordinates, found := findAdjacentSymbolCoordinates(number, coordinates, symbolCoordinates)
			if found {
				result, _ := strconv.Atoi(number)
				_, exists := symbolCoordinatesToPartNumbers[symbolCoordinates]
				if !exists {
					symbolCoordinatesToPartNumbers[symbolCoordinates] = make([]int, 0)
				}
				symbolCoordinatesToPartNumbers[symbolCoordinates] =
					append(symbolCoordinatesToPartNumbers[symbolCoordinates], result)
			}
		}
	}
	// find symbol coordinates with *exactly* two adjacent part numbers and add 'em up
	sum := 0
	for _, numbers := range symbolCoordinatesToPartNumbers {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}
	return sum
}
