package day14_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day14"
	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	size := field.Point{X: 11, Y: 7}
	g := day14.Guard{
		Position: field.Point{X: 2, Y: 4},
		Velocity: field.Point{X: 2, Y: -3},
	}
	g = g.Move(size)
	require.Equal(t, field.Point{X: 4, Y: 1}, g.Position)
	g = g.Move(size)
	require.Equal(t, field.Point{X: 6, Y: 5}, g.Position)
}
