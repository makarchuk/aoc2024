package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day8"
)

func main() {
	input, err := day8.ReadInput(os.Stdin)
	if err != nil {
		panic(err)
	}

	antinodes := day8.FindHarmonicAntinodes(input.Space, input.Antennas)
	fmt.Println(antinodes.Len())
}
