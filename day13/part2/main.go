package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day13"
	"github.com/makarchuk/aoc2024/pkg/field"
)

var targetOffset = field.Point{X: 10_000_000_000_000, Y: 10_000_000_000_000}

func main() {
	machines, err := day13.ParseArcades(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("failed to parse arcades: %v", err))
	}

	totalPrice := 0
	for _, machine := range machines {
		machine.Target = machine.Target.Add(targetOffset)
		totalPrice += machine.OptimalSolution()
	}
	fmt.Println(totalPrice)

}
