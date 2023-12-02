package solutions

import (
	"aoc/app/solutions/day01"
	"aoc/app/solutions/day02"
	"aoc/app/solutions/day03"
	"aoc/app/solutions/day04"
	"aoc/app/solutions/day05"
	"aoc/app/solutions/day06"
	"aoc/app/solutions/day07"
	"aoc/app/solutions/day08"
	"aoc/app/solutions/day09"
	"aoc/app/solutions/day10"
	"aoc/app/solutions/day11"
	"aoc/app/solutions/day12"
	"aoc/app/solutions/day13"
	"aoc/app/solutions/day14"
	"aoc/app/solutions/day15"
	"aoc/app/solutions/day16"
	"aoc/app/solutions/day17"
	"aoc/app/solutions/day18"
	"aoc/app/solutions/day19"
	"aoc/app/solutions/day20"
	"aoc/app/solutions/day21"
	"aoc/app/solutions/day22"
	"aoc/app/solutions/day23"
	"aoc/app/solutions/day24"
	"aoc/app/solutions/day25"
)

type Solution[T any] interface {
	Name() string
	Evaluate() T
}

func GetAll() []Solution[any] {
	return []Solution[any]{
		&day01.D1P1{},
		&day01.D1P2{},
		&day02.D2P1{},
		&day02.D2P2{},
		&day03.D3P1{},
		&day03.D3P2{},
		&day04.D4P1{},
		&day04.D4P2{},
		&day05.D5P1{},
		&day05.D5P2{},
		&day06.D6P1{},
		&day06.D6P2{},
		&day07.D7P1{},
		&day07.D7P2{},
		&day08.D8P1{},
		&day08.D8P2{},
		&day09.D9P1{},
		&day09.D9P2{},
		&day10.D10P1{},
		&day10.D10P2{},
		&day11.D11P1{},
		&day11.D11P2{},
		&day12.D12P1{},
		&day12.D12P2{},
		&day13.D13P1{},
		&day13.D13P2{},
		&day14.D14P1{},
		&day14.D14P2{},
		&day15.D15P1{},
		&day15.D15P2{},
		&day16.D16P1{},
		&day16.D16P2{},
		&day17.D17P1{},
		&day17.D17P2{},
		&day18.D18P1{},
		&day18.D18P2{},
		&day19.D19P1{},
		&day19.D19P2{},
		&day20.D20P1{},
		&day20.D20P2{},
		&day21.D21P1{},
		&day21.D21P2{},
		&day22.D22P1{},
		&day22.D22P2{},
		&day23.D23P1{},
		&day23.D23P2{},
		&day24.D24P1{},
		&day24.D24P2{},
		&day25.D25P1{},
		&day25.D25P2{},
	}
}
