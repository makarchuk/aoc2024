package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day6"
	"github.com/makarchuk/aoc2024/pkg/set"
)

func main() {
	input, err := day6.New(os.Stdin)
	if err != nil {
		panic(err)
	}

	guard := input.GetGuard()
	visitedPoints := set.New[day6.Point]()
	for {
		visitedPoints.Add(guard.Position())
		err := guard.Move(*input)
		if err != nil {
			break
		}

		visitedPoints.Add(guard.Position())
	}

	fmt.Println(visitedPoints.Len())
}
