package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/makarchuk/aoc2024/day5"
	"github.com/makarchuk/aoc2024/pkg/dag"
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
			continue
		}

		relevantRules := [][2]int{}
		for _, rule := range input.Rules {
			if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
				relevantRules = append(relevantRules, rule)
			}
		}

		affectedNodes := map[int]bool{}
		for _, rule := range relevantRules {
			affectedNodes[rule[0]] = true
			affectedNodes[rule[1]] = true
		}

		if len(affectedNodes) != len(update) {
			panic("Some nodea are not part of the graph. My solution does not work here")
		}

		dag := dag.NewDag[int]()

		for _, rule := range relevantRules {
			dag.AddEdge(rule[1], rule[0])
			dag.AddNode(rule[0])
			dag.AddNode(rule[1])
		}

		linearized := dag.Linearize()
		checkSum += linearized[len(linearized)/2]
	}

	fmt.Println(checkSum)
}
