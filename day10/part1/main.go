package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day10"
)

func main() {
	terrain, err := day10.New(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(day10.TerrainScore(terrain))

}
