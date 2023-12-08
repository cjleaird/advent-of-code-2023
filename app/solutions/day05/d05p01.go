package day05

import (
	"aoc/app/helper"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type D5P1 struct{}

func (solution *D5P1) Name() string {
	return "Day 5 (Part 1)"
}

func (solution *D5P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	seeds := make([]int, 0)
	maps := [][]transformer{}
	for _, line := range lines {
		regex := regexp.MustCompile(`seeds:`)
		if regex.FindString(line) != "" {
			regex = regexp.MustCompile(`\d+`)
			for _, s := range regex.FindAllString(strings.Split(line, ":")[1], -1) {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
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

	locations := make([]int, 0)
	for _, seed := range seeds {
		result := seed
		for _, transformers := range maps {
			for _, transformer := range transformers {
				if result >= transformer.sourceValue && result < transformer.sourceValue+transformer.rangeValue {
					result = transformer.destValue + result - transformer.sourceValue
					break
				}
			}
		}
		locations = append(locations, result)
	}
	return slices.Min(locations)
}
