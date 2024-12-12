package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day12"
)

func main() {
	finder, err := day12.NewRegionsFinder(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(finder.PlanFences().Price)
}
