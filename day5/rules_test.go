package day5_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day5"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	rules := day5.NewRules()

	rules.AddRule(47, 53)
	rules.AddRule(97, 13)
	rules.AddRule(97, 61)
	rules.AddRule(97, 47)
	rules.AddRule(75, 29)
	rules.AddRule(61, 13)
	rules.AddRule(75, 53)
	rules.AddRule(29, 13)
	rules.AddRule(97, 29)
	rules.AddRule(53, 29)
	rules.AddRule(61, 53)
	rules.AddRule(97, 53)
	rules.AddRule(61, 29)
	rules.AddRule(47, 13)
	rules.AddRule(75, 47)
	rules.AddRule(97, 75)
	rules.AddRule(47, 61)
	rules.AddRule(75, 61)
	rules.AddRule(47, 29)
	rules.AddRule(75, 13)
	rules.AddRule(53, 13)

	require.True(t, rules.CanBeAfter(29, []int{75, 47, 61, 53}))

	require.True(t, rules.Part1Check([]int{75, 47, 61, 53, 29}))
	require.True(t, rules.Part1Check([]int{97, 61, 53, 29, 13}))
	require.True(t, rules.Part1Check([]int{75, 29, 13}))
	require.False(t, rules.Part1Check([]int{75, 97, 47, 61, 53}))
	require.False(t, rules.Part1Check([]int{61, 13, 29}))
	require.False(t, rules.Part1Check([]int{97, 13, 75, 29, 47}))

}
