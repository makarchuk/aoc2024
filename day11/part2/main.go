package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day11"
)

func main() {
	stones, err := day11.NewStones(os.Stdin)
	if err != nil {
		panic(err)
	}
	counter := day11.NewCounter(stones.Stones())
	for _ = range 75 {
		counter = counter.NextGeneration()
	}
	fmt.Println(counter.Len())
}
