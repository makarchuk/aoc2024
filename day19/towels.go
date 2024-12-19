package day19

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/priorityqueue"
	"github.com/makarchuk/aoc2024/pkg/set"
)

func Part1(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	constructible := 0
	for _, pattern := range input.Patterns {
		if input.ConstructPattern(pattern) > 0 {
			constructible++
		}
	}

	return fmt.Sprintf("%d", constructible), nil
}

type Input struct {
	Towels   []string
	Patterns []string
}

func ParseInput(in io.Reader) (Input, error) {
	input := Input{}

	scanner := bufio.NewScanner(in)

	if !scanner.Scan() {
		return Input{}, fmt.Errorf("expected towels")
	}
	input.Towels = strings.Split(scanner.Text(), string(", "))

	if !scanner.Scan() {
		return Input{}, fmt.Errorf("expected empty line")
	}

	if scanner.Text() != "" {
		return Input{}, fmt.Errorf("expected empty line, found %v", scanner.Text())
	}

	for scanner.Scan() {
		input.Patterns = append(input.Patterns, scanner.Text())
	}

	return input, nil
}

type patternConstructorState struct {
	// Towels    []string
	Towels    int
	Remaining string
}

func (in *Input) ConstructPattern(pattern string) int {
	state := patternConstructorState{
		// Towels:    []string{},
		Towels:    0,
		Remaining: pattern,
	}

	cameFrom := map[string]set.Set[string]{}

	states := priorityqueue.PriorityQueue[patternConstructorState]{}
	states.Push(-len(state.Remaining), state)
	for {
		state, ok := states.Pop()
		if !ok {
			break
		}

		for _, towel := range in.Towels {
			if !strings.HasPrefix(state.Remaining, towel) {
				continue
			}

			// towels := make([]string, len(state.Towels))
			// copy(towels, state.Towels)
			// towels = append(towels, towel)

			remaining := strings.TrimPrefix(state.Remaining, towel)
			remainingPaths, ok := cameFrom[remaining]
			if !ok {
				remainingPaths = set.New[string]()
			}
			//we're storing where did we came from to reach this state
			remainingPaths.Add(state.Remaining)

			// fmt.Printf("remaining before:%v, after: %v, towel: `%v`\n", state.Remaining, remaining, towel)
			if !ok {
				states.Push(-len(remaining), patternConstructorState{
					// Towels:    towels,
					Towels:    state.Towels + 1,
					Remaining: remaining,
				})
			}

		}
	}

	countWays := 0

	surface := []string{pattern}

	for {
		newSurface := []string{}

		for _, path := range surface {
			cameFromPaths, ok := cameFrom[path]

			if !ok {
				panic(fmt.Sprintf("don't know how we reached %v", path))
			}

		}
	}

}
