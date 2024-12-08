package day8_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/makarchuk/aoc2024/day8"
	"github.com/makarchuk/aoc2024/pkg/field"
)

func TestAntinodes(t *testing.T) {
	require.Equal(t,
		[]field.Point{{X: 1, Y: 1}, {X: 4, Y: 4}},
		day8.AntiNodes(field.Point{X: 2, Y: 2}, field.Point{X: 3, Y: 3}),
	)

	require.Equal(t,
		[]field.Point{{X: 4, Y: -4}, {X: -2, Y: 8}},
		day8.AntiNodes(field.Point{X: 2, Y: 0}, field.Point{X: 0, Y: 4}),
	)

	require.Equal(t,
		[]field.Point{{X: 4, Y: 2}, {X: 1, Y: 5}},
		day8.AntiNodes(field.Point{X: 3, Y: 3}, field.Point{X: 2, Y: 4}),
	)

	require.Equal(t,
		[]field.Point{{X: 0, Y: 5}, {X: 3, Y: 2}},
		day8.AntiNodes(field.Point{X: 1, Y: 4}, field.Point{X: 2, Y: 3}),
	)
}

func TestE2E(t *testing.T) {
	input := bytes.NewBuffer([]byte(strings.Trim(`
...A...
..AaA..
...a...
`, "\n")))

	in, err := day8.ReadInput(input)
	require.NoError(t, err)
	antinodes := day8.FindAntinodes(in.Space, in.Antennas)

	require.Len(t, antinodes.List(), 5)
	fmt.Printf("Antinodes: %+v\n", antinodes.List())
	require.True(t, antinodes.Contains(field.Point{X: 3, Y: 0}))
	require.True(t, antinodes.Contains(field.Point{X: 5, Y: 2}))
	require.True(t, antinodes.Contains(field.Point{X: 1, Y: 2}))
	require.True(t, antinodes.Contains(field.Point{X: 0, Y: 1}))
	require.True(t, antinodes.Contains(field.Point{X: 6, Y: 1}))
}

func TestE2EHarmonic(t *testing.T) {
	input := bytes.NewBuffer([]byte(strings.Trim(`
.......
.a.....
..a....
.......
.......
.......
`, "\n")))
	in, err := day8.ReadInput(input)
	require.NoError(t, err)
	antinodes := day8.FindHarmonicAntinodes(in.Space, in.Antennas)

	require.Len(t, antinodes.List(), 4)
	fmt.Printf("Antinodes: %+v\n", antinodes.List())
	require.True(t, antinodes.Contains(field.Point{X: 0, Y: 0}))
	require.True(t, antinodes.Contains(field.Point{X: 3, Y: 3}))
	require.True(t, antinodes.Contains(field.Point{X: 4, Y: 4}))
	require.True(t, antinodes.Contains(field.Point{X: 5, Y: 5}))
	// require.True(t, antinodes.Contains(field.Point{X: 6, Y: 1}))
}

func TestE2EHarmonicSimple(t *testing.T) {
	input := bytes.NewBuffer([]byte(strings.Trim(`
....b..
.......
....b..
.......
.......
.......
`, "\n")))
	in, err := day8.ReadInput(input)
	require.NoError(t, err)
	antinodes := day8.FindHarmonicAntinodes(in.Space, in.Antennas)

	require.Len(t, antinodes.List(), 4)
	fmt.Printf("Antinodes: %+v\n", antinodes.List())
	require.True(t, antinodes.Contains(field.Point{X: 4, Y: 1}))
	require.True(t, antinodes.Contains(field.Point{X: 4, Y: 3}))
	require.True(t, antinodes.Contains(field.Point{X: 4, Y: 4}))
	require.True(t, antinodes.Contains(field.Point{X: 4, Y: 5}))
	// require.True(t, antinodes.Contains(field.Point{X: 6, Y: 1}))
}

func TestNPlusOne(t *testing.T) {
	pairs := [][2]int{}
	sl := []int{1, 2, 3, 4, 5}

	for i, v1 := range sl {
		for _, v2 := range sl[i+1:] {
			pairs = append(pairs, [2]int{v1, v2})
		}
	}

	require.Len(t, pairs, 10)
	require.Contains(t, pairs, [2]int{3, 4})
	require.Contains(t, pairs, [2]int{2, 5})
	require.Contains(t, pairs, [2]int{1, 3})
	require.NotContains(t, pairs, [2]int{3, 3})
	require.NotContains(t, pairs, [2]int{5, 2})
	require.NotContains(t, pairs, [2]int{2, 1})
}
