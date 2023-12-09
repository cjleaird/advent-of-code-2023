package day06

import (
	"aoc/app/helper"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type D6P2 struct{}

func (solution *D6P2) Name() string {
	return "Day 6 (Part 2)"
}

func (solution *D6P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	regex := regexp.MustCompile(`\d+`)
	values := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(strings.Join(regex.FindAllString(line, -1), ""))
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	if len(values) != 2 {
		log.Fatal(err)
	}
	recordTime := values[0]
	recordDistance := values[1]
	count := 0
	for i := 0; i < recordTime; i++ {
		if isNewRecordRace(recordTime, recordDistance, i) {
			count++
		} else if count > 0 {
			// if we've reached the point where we've encountered
			// winning races and are no longer winning - we're done
			// checking
			break
		}
	}
	return count
}
