package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day2"
	"github.com/makarchuk/aoc2024/pkg/helpers"
)

type Input struct {
	Reports [][]int
}

func main() {
	var input Input

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		levels, err := helpers.ParseIntsArray(scanner.Text(), " ")
		if err != nil {
			panic(err)
		}
		input.Reports = append(input.Reports, levels)
	}

	safeReports := 0

	for _, report := range input.Reports {
		if day2.CheckReport(report) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
