package day25_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/makarchuk/aoc2024/day25"
	"github.com/makarchuk/aoc2024/pkg/set"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	type testCase struct {
		input string
		keys  []day25.Key
		locks []day25.Lock
	}

	for _, tc := range []testCase{
		{
			input: `
#####
.####
.####
.####
.#.#.
.#...
.....
`,
			keys:  nil,
			locks: []day25.Lock{{0, 5, 3, 4, 3}},
		},
		{
			input: `
.....
#....
#....
#...#
#.#.#
#.###
#####
`,
			keys:  []day25.Key{{5, 0, 2, 1, 3}},
			locks: nil,
		},
	} {
		input := strings.Trim(tc.input, "\n")
		in := strings.NewReader(input)
		got, err := day25.ParseInput(in)
		require.NoError(t, err)
		require.True(
			t,
			set.From(tc.keys).Equal(set.From(got.Keys)),
			fmt.Sprintf("expected: %v, got: %v", tc.keys, got.Keys),
		)
		require.True(
			t,
			set.From(tc.locks).Equal(set.From(got.Locks)),
			fmt.Sprintf("expected: %v, got: %v", tc.locks, got.Locks),
		)
	}

}
