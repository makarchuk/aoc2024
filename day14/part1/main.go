package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day14"
	"github.com/makarchuk/aoc2024/pkg/field"
)

func main() {
	inp, err := day14.ParseInput(os.Stdin)
	if err != nil {
		panic(err)
	}

	inp.Size = field.Point{X: 101, Y: 103}

	for _ = range 100 {
		inp.Step()
	}

	fmt.Println(inp.SafetyFactor())
}
