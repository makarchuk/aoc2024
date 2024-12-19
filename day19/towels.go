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

func Part2(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	patternsTotal := 0
	for _, pattern := range input.Patterns {
		waysToConstruct := input.ConstructPattern(pattern)

		patternsTotal += waysToConstruct
	}

	return fmt.Sprintf("%d", patternsTotal), nil
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
	Remaining string
}

func (in *Input) ConstructPattern(pattern string) int {
	state := patternConstructorState{
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
			if !strings.HasSuffix(state.Remaining, towel) {
				continue
			}

			remaining := strings.TrimSuffix(state.Remaining, towel)
			waysToReach, ok := cameFrom[remaining]
			if !ok {
				waysToReach = set.New[string]()
			}

			//we're storing where did we came from to reach this state
			waysToReach.Add(state.Remaining)

			cameFrom[remaining] = waysToReach
			if len(remaining) == 0 {
				break
			}

			// fmt.Printf("remaining before:%v, after: %v, towel: `%v`\n", state.Remaining, remaining, towel)
			if !ok {
				states.Push(-len(remaining), patternConstructorState{
					Remaining: remaining,
				})
			}

		}
	}

	// fmt.Printf("found no path for %v. Bailing\n", pattern)

	if _, ok := cameFrom[""]; !ok {
		return 0
	}
	// fmt.Printf("found path to %v. Scoring\n", pattern)

	cache := make(map[string]int)

	return countWays(pattern, cameFrom, "", cache)
}

func countWays(pattern string, visitsMap map[string]set.Set[string], start string, cache map[string]int) int {
	score := 0

	if pattern == start {
		return 1
	}

	parents := visitsMap[start]
	if parents.Len() == 0 {
		panic("should not be happening")
	}

	for _, parent := range parents.List() {
		if parentScore, ok := cache[parent]; ok {
			score += parentScore
			continue
		}

		parentScore := countWays(pattern, visitsMap, parent, cache)
		cache[parent] = parentScore
		score += parentScore
	}

	return score
}
