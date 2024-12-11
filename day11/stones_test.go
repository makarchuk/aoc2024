package day11_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day11"

	"github.com/stretchr/testify/require"
)

func TestCutInHalf(t *testing.T) {
	require.Equal(t, [2]int{2, 0}, day11.CutInHalf(20))
	require.Equal(t, [2]int{400, 1}, day11.CutInHalf(400001))
	require.Equal(t, [2]int{12, 34}, day11.CutInHalf(1234))
	require.Equal(t, [2]int{1111, 1111}, day11.CutInHalf(11111111))
	require.Equal(t, [2]int{10, 0}, day11.CutInHalf(1000))
}

func TestPart1(t *testing.T) {
	counter := day11.NewCounter(
		[]int{2, 54, 992917, 5270417, 2514, 28561, 0, 990},
	)

	for i := 0; i < 25; i++ {
		counter = counter.NextGeneration()
	}
	require.Equal(t, 222461, counter.Len())

}
