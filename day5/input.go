package day5

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/helpers"
)

type Input struct {
	Rules   [][2]int
	Updates [][]int
}

func NewInput(in io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(in)
	input := &Input{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var rule [2]int
		if _, err := fmt.Sscanf(line, "%d|%d", &rule[0], &rule[1]); err != nil {
			return nil, err
		}
		input.Rules = append(input.Rules, rule)
	}

	for scanner.Scan() {
		update, err := helpers.ParseIntsArray(scanner.Text(), ",")
		if err != nil {
			return nil, err
		}
		input.Updates = append(input.Updates, update)
	}
	return input, nil
}
