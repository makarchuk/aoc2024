package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day7"
)

func main() {
	input := []day7.Expression{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		expr := day7.Expression{}
		err := expr.Parse(line)
		if err != nil {
			panic(err)
		}
		input = append(input, expr)
	}

	var correctSum int64
	for _, expr := range input {
		if expr.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}) {
			correctSum += int64(expr.Result)
		}
	}

	fmt.Println(correctSum)
}
