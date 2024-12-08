package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day6"
	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

func main() {
	input, err := day6.New(os.Stdin)
	if err != nil {
		panic(err)
	}

	guard := input.OriginalGuard
	visitedPoints := set.New[field.Point]()
	for {
		visitedPoints.Add(guard.Position())
		err := guard.Move(input.Field)
		if err != nil {
			break
		}

		visitedPoints.Add(guard.Position())
	}

	fmt.Println(visitedPoints.Len())
}
