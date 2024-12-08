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

	cleanGuard := input.OriginalGuard

	guard := cleanGuard.Clone()

	//clean walk, to discover all visited points
	possibleLoops := set.New[field.Point]()
	for {
		possibleLoops.Add(guard.Position())
		err := guard.Move(input.Field)
		if err != nil {
			break
		}

		possibleLoops.Add(guard.Position())
	}

	possibleLoops.Remove(cleanGuard.Position())

	detectedLoops := set.New[field.Point]()
	for _, possibleStone := range possibleLoops.List() {
		guardCopy := cleanGuard.Clone()
		if possibleStone == guardCopy.Position() {
			continue
		}
		newField := input.Field.Replace(possibleStone, '#')
		if checkLoop(guardCopy, newField) {
			detectedLoops.Add(possibleStone)
		}
	}

	fmt.Println(detectedLoops.Len())
}

func checkLoop(guard day6.Guard, field field.Field) bool {
	visited := set.New[day6.Guard]()

	localGuard := guard.Clone()
	for {
		if visited.Contains(localGuard) {
			return true
		}
		visited.Add(localGuard.Clone())
		err := localGuard.Move(field)
		if err != nil {
			return false
		}
	}
}
