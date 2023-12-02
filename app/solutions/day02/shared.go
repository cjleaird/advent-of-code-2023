package day02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func findIdAndCubeSets(str string) (int, []string, error) {
	regex := regexp.MustCompile(`Game (?P<id>[0-9]+): (?P<sets>.*)`)
	matches := findAllNamedMatches(regex, str)
	_, exists := matches["id"]
	if !exists {
		return 0, nil, errors.New("invalid line: missing <id>")
	}
	_, exists = matches["sets"]
	if !exists {
		return 0, nil, errors.New("invalid line: missing <sets>")
	}
	id, _ := strconv.Atoi(matches["id"])
	sets := strings.Split(matches["sets"], "; ")
	return id, sets, nil
}

func findRgbCubeCounts(str string) (int, int, int) {
	regex := regexp.MustCompile(`(?:(?:(?P<blue>\d+) blue|(?P<red>\d+) red)|(?P<green>\d+) green)`)
	matches := findAllNamedMatches(regex, str)
	red := findCubeCount(matches, "red")
	green := findCubeCount(matches, "green")
	blue := findCubeCount(matches, "blue")
	return red, green, blue
}

func findCubeCount(matches map[string]string, color string) int {
	_, exists := matches[color]
	if !exists {
		return 0
	}
	count, err := strconv.Atoi(matches[color])
	if err != nil {
		count = 0
	}
	return count
}

func findAllNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	matches := regex.FindAllStringSubmatch(str, -1)

	results := map[string]string{}
	for i := range matches {
		for j, value := range matches[i] {
			name := regex.SubexpNames()[j]
			if name == "" || value == "" {
				continue
			}
			results[name] = value
		}
	}
	return results
}

const contentUrl = "https://adventofcode.com/2023/day/2/input"
