package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day13"
)

func main() {
	machines, err := day13.ParseArcades(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("failed to parse arcades: %v", err))
	}

	totalPrice := 0
	for _, machine := range machines {
		totalPrice += machine.OptimalSolution()
	}
	fmt.Println(totalPrice)
}
