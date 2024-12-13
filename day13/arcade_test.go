package day13_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day13"
	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestArcadeMachine(t *testing.T) {
	type testcase struct {
		arcadeMachine day13.Arcade
		price         int
	}

	for _, tc := range []testcase{
		{
			arcadeMachine: day13.Arcade{
				OffsetA: field.Point{X: 94, Y: 34},
				OffsetB: field.Point{X: 22, Y: 67},
				Target:  field.Point{X: 8400, Y: 5400},
			},
			price: 280,
		},
	} {
		solution := tc.arcadeMachine.OptimalSolution()
		require.Equal(t, tc.price, solution)
	}
}
