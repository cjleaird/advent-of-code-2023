package day01

import (
	"aoc/app/helper"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type D1P1 struct{}

func (solution *D1P1) Name() string {
	return "Day 1 (Part 1)"
}

func (solution *D1P1) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	regex := regexp.MustCompile(`[0-9]{1}`)
	sum := 0
	for _, line := range lines {
		digits := regex.FindAllString(line, -1)
		if digits == nil {
			continue
		}
		firstDigit, _ := strconv.Atoi(digits[0])
		lastDigit, _ := strconv.Atoi(digits[len(digits)-1])
		sum += firstDigit*10 + lastDigit
	}
	return sum
}
