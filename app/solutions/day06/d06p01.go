package day06

import (
	"aoc/app/helper"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type D6P1 struct{}

func (solution *D6P1) Name() string {
	return "Day 6 (Part 1)"
}

func (solution *D6P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	regex := regexp.MustCompile(`\d+`)
	values := make([][]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		values = append(values, []int{})
		for _, s := range regex.FindAllString(line, -1) {
			value, _ := strconv.Atoi(s)
			values[len(values)-1] = append(values[len(values)-1], value)
		}
	}
	if len(values) != 2 {
		log.Fatal(err)
	}
	pairs, _ := helper.Zip[int, int](values[0], values[1])
	product := 1
	for _, pair := range pairs {
		recordTime := pair.First
		recordDistance := pair.Second
		count := 0
		for i := 0; i < recordTime; i++ {
			if isNewRecordRace(recordTime, recordDistance, i) {
				count++
			}
		}
		product *= count
	}

	return product
}
