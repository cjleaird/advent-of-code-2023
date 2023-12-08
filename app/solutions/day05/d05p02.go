package day05

import (
	"aoc/app/helper"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type D5P2 struct{}

func (solution *D5P2) Name() string {
	return "Day 5 (Part 2)"
}

func (solution *D5P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	categoryToRangeList := [][2]int{}
	maps := [][]transformer{}
	for _, line := range lines {
		regex := regexp.MustCompile(`seeds:`)
		if regex.FindString(line) != "" {
			regex = regexp.MustCompile(`\d+ \d+`)
			seedsAndRanges := regex.FindAllString(strings.Split(line, ":")[1], -1)
			for _, s := range seedsAndRanges {
				result := strings.Split(s, " ")
				seedValue, _ := strconv.Atoi(result[0])
				rangeValue, _ := strconv.Atoi(result[1])
				categoryToRangeList = append(categoryToRangeList, [2]int{seedValue, rangeValue})
			}
		} else {
			regex = regexp.MustCompile(`map`)
			if regex.FindString(line) != "" {
				maps = append(maps, make([]transformer, 0))
			} else {
				regex = regexp.MustCompile(`\d+`)
				values := regex.FindAllString(line, -1)
				if len(values) != 3 {
					continue
				}
				destValue, _ := strconv.Atoi(values[0])
				sourceValue, _ := strconv.Atoi(values[1])
				rangeValue, _ := strconv.Atoi(values[2])
				transformer := transformer{destValue: destValue, sourceValue: sourceValue, rangeValue: rangeValue}
				maps[len(maps)-1] = append(maps[len(maps)-1], transformer)
			}
		}
	}
	// we have to approach this differently than we did for part one - essentially,
	// instead of iterating through each seed one-by-one to obtain its final location,
	// we iterate through each map, obtaining all possible destination category-to-range values
	// for *all* input seeds after applying the associated transformations - we then carry
	// these destination category-to-range values through each map, until we end up with our
	// final location-to-range values
	for _, transformers := range maps {
		categoryToRangeList = transformCategoryToRangeList(transformers, categoryToRangeList)
	}
	// we now have all possible location-to-range values after the final humidity-to-location
	// transformations - all that's left to do is find the minimum location value
	minLocation := math.MaxInt
	for _, categoryToRange := range categoryToRangeList {
		if categoryToRange[0] < minLocation {
			minLocation = categoryToRange[0]
		}
	}
	return minLocation
}

func transformCategoryToRangeList(transformers []transformer, prevCategoryToRangeList [][2]int) [][2]int {
	currCategoryToRangeList := [][2]int{}
	for _, transformer := range transformers {
		sourceValue := transformer.sourceValue
		maxSourceValue := sourceValue + transformer.rangeValue - 1
		destValue := transformer.destValue
		delta := destValue - sourceValue
		transformerCategoryToRangeList := [][2]int{}
		for _, categoryToRange := range prevCategoryToRangeList {
			categoryValue := categoryToRange[0]
			categoryRangeValue := categoryToRange[1]
			maxCategoryValue := categoryValue + categoryRangeValue - 1
			if categoryValue < sourceValue {
				if maxCategoryValue < sourceValue {
					transformerCategoryToRangeList = append(transformerCategoryToRangeList, [2]int{categoryValue, categoryRangeValue})
				} else {
					transformerCategoryToRangeList = append(transformerCategoryToRangeList, [2]int{categoryValue, sourceValue - categoryValue})
					currCategoryToRangeList = append(currCategoryToRangeList, [2]int{sourceValue + delta, maxCategoryValue - sourceValue + 1})
				}
			} else if categoryValue < maxSourceValue {
				if maxCategoryValue < maxSourceValue {
					currCategoryToRangeList = append(currCategoryToRangeList, [2]int{categoryValue + delta, categoryRangeValue})
				} else {
					currCategoryToRangeList = append(currCategoryToRangeList, [2]int{categoryValue + delta, maxSourceValue - categoryValue})
					transformerCategoryToRangeList = append(transformerCategoryToRangeList, [2]int{maxSourceValue, maxCategoryValue - maxSourceValue + 1})
				}
			} else {
				transformerCategoryToRangeList = append(transformerCategoryToRangeList, [2]int{categoryValue, categoryRangeValue})
			}
		}
		prevCategoryToRangeList = transformerCategoryToRangeList
	}
	return append(currCategoryToRangeList, prevCategoryToRangeList...)
}
