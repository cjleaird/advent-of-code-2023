package day04

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func populateCardNumbers(line string, winningNumbers map[int]struct{}, pickedNumbers map[int]struct{}) error {
	regex := regexp.MustCompile(`(\d+)`)
	entries := strings.Split(line, ":")
	// indicates an invalid line, or more probably EOF
	if len(entries) != 2 {
		return errors.New("malformatted line input")
	}
	numbers := strings.Split(entries[1], "|")
	// indicates an invalid line, or more probably EOF
	if len(numbers) != 2 {
		return errors.New("malformatted line input")
	}
	for _, match := range regex.FindAllString(numbers[0], -1) {
		number, _ := strconv.Atoi(match)
		winningNumbers[number] = struct{}{}
	}
	for _, match := range regex.FindAllString(numbers[1], -1) {
		number, _ := strconv.Atoi(match)
		pickedNumbers[number] = struct{}{}
	}
	return nil
}

func findMatchCount(winningNumbers map[int]struct{}, pickedNumbers map[int]struct{}) int {
	count := 0
	for number := range pickedNumbers {
		_, exists := winningNumbers[number]
		if exists {
			count++
		}
	}
	return count
}

const contentUrl = "https://adventofcode.com/2023/day/4/input"
