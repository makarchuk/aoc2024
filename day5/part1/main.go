package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day5"
)

func main() {
	input, err := day5.NewInput(os.Stdin)
	if err != nil {
		panic(err)
	}

	rules := day5.NewRules()
	for _, rule := range input.Rules {
		rules.AddRule(rule[0], rule[1])
	}

	checkSum := 0
	for _, update := range input.Updates {
		if rules.Part1Check(update) {
			checkSum += update[len(update)/2]
		}
	}

	fmt.Println(checkSum)
}
