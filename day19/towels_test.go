package day19_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day19"
	"github.com/stretchr/testify/require"
)

func TestWaysToreac(t *testing.T) {
	type TestCase struct {
		Towels  []string
		Pattern string
		Result  int
	}

	exampleTowels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}

	for _, tc := range []TestCase{
		{
			Towels:  exampleTowels,
			Pattern: "brwrr",
			Result:  2,
		},
		{
			Towels:  exampleTowels,
			Pattern: "rrbgbr",
			Result:  6,
		},
		{
			Towels:  exampleTowels,
			Pattern: "bwurrg",
			Result:  1,
		},
		{
			Towels:  exampleTowels,
			Pattern: "ubwu",
			Result:  0,
		},
		{
			Towels:  exampleTowels,
			Pattern: "gbbr",
			Result:  4,
		},
		{
			Towels:  exampleTowels,
			Pattern: "bggr",
			Result:  1,
		},
		{
			Towels:  exampleTowels,
			Pattern: "brgr",
			Result:  2,
		},
		{
			Towels:  exampleTowels,
			Pattern: "bbrgwb",
			Result:  0,
		},
	} {
		in := day19.Input{
			Towels:   tc.Towels,
			Patterns: []string{tc.Pattern},
		}

		waysToReach := in.ConstructPattern(tc.Pattern)
		require.Equal(t, tc.Result, waysToReach)
	}
}
