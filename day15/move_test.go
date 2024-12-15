package day15

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"

	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	type testCase struct {
		submarine field.Point
		boxes     []field.Point

		commands []field.Direction

		newSubmarine field.Point
		newBoxes     []field.Point
	}

	for _, tc := range []testCase{
		{
			submarine:    field.Point{X: 1, Y: 1},
			boxes:        []field.Point{},
			commands:     []field.Direction{field.RIGHT, field.RIGHT, field.DOWN, field.LEFT, field.UP, field.UP},
			newSubmarine: field.Point{X: 2, Y: 1},
			newBoxes:     []field.Point{},
		},
		{
			submarine: field.Point{X: 1, Y: 1},
			boxes: []field.Point{
				{X: 2, Y: 1},
				{X: 3, Y: 1},
			},
			commands:     []field.Direction{field.RIGHT, field.RIGHT, field.DOWN, field.LEFT, field.UP, field.UP},
			newSubmarine: field.Point{X: 2, Y: 1},
			newBoxes: []field.Point{
				{X: 4, Y: 1},
				{X: 5, Y: 1},
			},
		},
		{
			submarine: field.Point{X: 1, Y: 1},
			boxes: []field.Point{
				{X: 2, Y: 1},
				{X: 3, Y: 1},
			},
			commands: []field.Direction{
				field.RIGHT,
				field.RIGHT,
				field.RIGHT, //blocked by the wall
				field.DOWN,
				field.LEFT,
				field.UP,
				field.UP,
			},
			newSubmarine: field.Point{X: 2, Y: 1},
			newBoxes: []field.Point{
				{X: 4, Y: 1},
				{X: 5, Y: 1},
			},
		},
	} {
		f := newField(t)
		state := State{
			sub:   tc.submarine,
			boxes: set.From(tc.boxes),
		}

		for _, cmd := range tc.commands {
			state = state.Execute(cmd, f)
		}

		require.Equal(t, tc.newSubmarine, state.sub)
		require.True(t, set.From(tc.newBoxes).Equal(state.boxes))
	}

}

func TestDoubleBoxesMoves(t *testing.T) {
	type testCase struct {
		field          string
		newSubmarine   field.Point
		newDoubleBoxes []field.Point
		commands       []field.Direction
	}

	for i, tc := range []testCase{
		{
			field: `
#############
#....@.[][].#
#############`,
			newSubmarine:   field.Point{X: 7, Y: 1},
			newDoubleBoxes: []field.Point{{X: 8, Y: 1}, {X: 10, Y: 1}},
			commands: []field.Direction{
				field.RIGHT,
				field.RIGHT,
				field.RIGHT,
				field.RIGHT,
			},
		},
		{
			field: `
#####
#...#
#.[]#
#[].#
#.@.#
#...#
#####`,
			newSubmarine:   field.Point{X: 2, Y: 3},
			newDoubleBoxes: []field.Point{{X: 2, Y: 1}, {X: 1, Y: 2}},
			commands: []field.Direction{
				field.UP,
				field.UP,
				field.UP,
			},
		},
		{
			field: `
######
#....#
#[][]#
#.[].#
#.@..#
#....#
######`,
			newSubmarine:   field.Point{X: 2, Y: 3},
			newDoubleBoxes: []field.Point{{X: 1, Y: 1}, {X: 3, Y: 1}, {X: 2, Y: 2}},
			commands: []field.Direction{
				field.UP,
				field.UP,
				field.UP,
			},
		},
		{
			field: `
##########
#.[][]@..#
##########
`,
			newSubmarine:   field.Point{X: 5, Y: 1},
			newDoubleBoxes: []field.Point{{X: 1, Y: 1}, {X: 3, Y: 1}},
			commands: []field.Direction{
				field.LEFT,
				field.LEFT,
				field.LEFT,
			},
		},
		{
			field: `
######
#.@..#
#[][]#
#.[].#
#....#
#....#
######`,
			newSubmarine:   field.Point{X: 2, Y: 3},
			newDoubleBoxes: []field.Point{{X: 3, Y: 2}, {X: 1, Y: 4}, {X: 2, Y: 5}},
			commands: []field.Direction{
				field.DOWN,
				field.DOWN,
				field.DOWN,
			},
		},
		{
			field: `
######
#.@..#
#[]..#
#.[].#
#..#.#
#....#
######`,
			newSubmarine:   field.Point{X: 2, Y: 1},
			newDoubleBoxes: []field.Point{{X: 1, Y: 2}, {X: 2, Y: 3}},
			commands: []field.Direction{
				field.DOWN,
				field.DOWN,
				field.DOWN,
			},
		},
		{
			field: `
######
#....#
#....#
#.[].#
#.[].#
#.@..#
#....#
######`,
			newSubmarine:   field.Point{X: 2, Y: 4},
			newDoubleBoxes: []field.Point{{X: 2, Y: 3}, {X: 2, Y: 2}},
			commands: []field.Direction{
				field.UP,
			},
		},
	} {
		t.Run(fmt.Sprintf("test-case-%v", i), func(t *testing.T) {
			in := parseInput(t, tc.field)
			state := State{
				sub:         in.submarine,
				boxes:       in.boxes,
				doubleBoxes: in.doubleBoxes,
			}
			require.Zero(t, state.boxes.Len())

			for _, cmd := range tc.commands {
				state = state.ExecuteWithDoubleBoxes(cmd, in.warehouse)
			}

			require.Equal(t, tc.newSubmarine, state.sub)
			require.True(
				t,
				set.From(tc.newDoubleBoxes).Equal(state.doubleBoxes),
				fmt.Sprintf("expected: %+v, got: %+v", tc.newDoubleBoxes, state.doubleBoxes.List()),
			)
		})
	}
}

func parseInput(t *testing.T, s string) Input {
	s = strings.TrimPrefix(s, "\n")
	buffer := strings.NewReader(s)
	in, err := ParseInput(buffer)
	require.NoError(t, err)
	return *in
}

func newField(t *testing.T) field.Field {
	f, err := field.New(bytes.NewBuffer([]byte{
		'#', '#', '#', '#', '#', '#', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '.', '.', '.', '.', '.', '#', '\n',
		'#', '#', '#', '#', '#', '#', '#', '\n',
	}))
	require.NoError(t, err)
	return *f
}
