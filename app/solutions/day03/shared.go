package day03

import (
	"regexp"
	"strconv"
)

// we have to find the coordinates of the symbols and numbers in the grid - we use
// a map structure for our symbol coordinates to optimize look-ups when we are
// determining whether a number is adjacent to a symbol
func populateCoordinates(row int, line string, numberCoordinates map[string][][2]int, symbolCoordinates map[[2]int]struct{}) {
	regex := regexp.MustCompile(`(\d+|[^\d.])`)
	indices := regex.FindAllStringIndex(line, -1)
	for _, index := range indices {
		col := index[0]
		value := line[col:index[1]]
		_, err := strconv.Atoi(value)
		// if true, we've found ourselves a symbol
		if err != nil {
			symbolCoordinates[[2]int{row, col}] = struct{}{}
		} else {
			_, exists := numberCoordinates[value]
			if !exists {
				numberCoordinates[value] = make([][2]int, 0)
			}
			numberCoordinates[value] =
				append(numberCoordinates[value], [2]int{row, col})
		}
	}
}

func findAdjacentSymbolCoordinates(number string, numberCoordinates [2]int, symbolCoordinates map[[2]int]struct{}) ([2]int, bool) {
	row := numberCoordinates[0]
	col := numberCoordinates[1]
	adjacentCoordinates := make([][2]int, 0)
	// check for directly to the left
	adjacentCoordinates = append(adjacentCoordinates, [2]int{row, col - 1})
	// check for directly to the right
	adjacentCoordinates = append(adjacentCoordinates, [2]int{row, col + len(number)})
	for i := col - 1; i <= col+len(number); i++ {
		// check for above
		adjacentCoordinates = append(adjacentCoordinates, [2]int{row - 1, i})
		// check for below
		adjacentCoordinates = append(adjacentCoordinates, [2]int{row + 1, i})
	}
	for _, coordinates := range adjacentCoordinates {
		_, exists := symbolCoordinates[coordinates]
		if exists {
			return coordinates, true
		}
	}
	return [2]int{}, false
}

const contentUrl = "https://adventofcode.com/2023/day/3/input"
