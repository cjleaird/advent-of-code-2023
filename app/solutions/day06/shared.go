package day06

func isNewRecordRace(recordTime int, recordDistance int, pressedTime int) bool {
	timeRemaining := recordTime - pressedTime
	velocity := pressedTime
	distance := (velocity * timeRemaining)
	return distance > recordDistance
}

const contentUrl = "https://adventofcode.com/2023/day/6/input"
