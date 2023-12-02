package day01

import (
	"aoc/app/helper"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type D1P2 struct{}

func (solution *D1P2) Name() string {
	return "Day 1 (Part 2)"
}

func (solution *D1P2) Evaluate() any {
	content, err := helper.GetContent(contentUrl)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(content, "\n")
	sum := 0
	for _, line := range lines {
		firstDigit, lastDigit := findFirstAndLastDigit(line)
		sum += firstDigit*10 + lastDigit
	}
	return sum
}

func findFirstAndLastDigit(line string) (int, int) {
	var indexToNumber = make(map[int]string)
	for key, value := range wordsToDigits {
		keyIndices := findFirstAndLastIndexOfNumber(line, key)
		if keyIndices != nil {
			indexToNumber[keyIndices[0]] = key
			indexToNumber[keyIndices[1]] = key
		}
		valueIndices := findFirstAndLastIndexOfNumber(line, value)
		if valueIndices != nil {
			indexToNumber[valueIndices[0]] = value
			indexToNumber[valueIndices[1]] = value
		}
	}
	// we didn't find any matches, so let's return zero
	if len(indexToNumber) == 0 {
		return 0, 0
	}
	// we've added all these indices to our map as keys,
	// let's make it easy to grab the smallest and largest
	// by sorting 'em
	keys := make([]int, 0)
	for key := range indexToNumber {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	firstDigit, err := toDigit(indexToNumber[keys[0]])
	if err != nil {
		log.Fatal("uh oh... this seems... impossible...")
	}
	lastDigit, err := toDigit(indexToNumber[keys[len(keys)-1]])
	if err != nil {
		log.Fatal("uh oh... this seems... impossible...")
	}
	return firstDigit, lastDigit
}

func findFirstAndLastIndexOfNumber(line string, number string) []int {
	regex := regexp.MustCompile(number)
	indices := regex.FindAllStringIndex(line, -1)
	if indices != nil {
		return []int{indices[0][0], indices[len(indices)-1][0]}
	}
	return nil
}

func toDigit(match string) (int, error) {
	value, exists := wordsToDigits[match]
	if exists {
		return strconv.Atoi(value)
	}
	// if not exists, then our match is a digit
	return strconv.Atoi(match)
}

var wordsToDigits = map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
