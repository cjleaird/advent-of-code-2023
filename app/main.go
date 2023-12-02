package main

import (
	"aoc/app/config"
	"aoc/app/solutions"
	"fmt"
	"log"
)

func init() {
	var err error
	config.AppConfig, err = config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config - required COOKIE for authorized access to AOC", err)
	}
}

func main() {
	for _, solution := range solutions.GetAll() {
		name := solution.Name()
		answer := solution.Evaluate()
		switch answer.(type) {
		case int:
			fmt.Printf("Name: %s, Answer: %d\n", name, answer)
		default:
			fmt.Printf("Name: %s, Answer: %s\n", name, answer)
		}
	}
}
