package day7_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day7"
	"github.com/stretchr/testify/require"
)

func TestBruteforce(t *testing.T) {
	require.True(
		t,
		day7.Expression{
			Result:   5,
			Operands: []int{2, 3},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}),
	)
	require.True(
		t,
		day7.Expression{
			Result:   190,
			Operands: []int{19, 10},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}),
	)

	require.True(
		t,
		day7.Expression{
			Result:   3267,
			Operands: []int{81, 40, 27},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}),
	)

	require.True(
		t,
		day7.Expression{
			Result:   292,
			Operands: []int{11, 6, 16, 20},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}),
	)

	require.False(
		t,
		day7.Expression{
			Result:   21037,
			Operands: []int{9, 7, 18, 13},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul}),
	)
	require.True(
		t,
		day7.Expression{
			Result:   7290,
			Operands: []int{6, 8, 6, 15},
		}.BruteforceOperators([]day7.Operator{day7.OperatorAdd, day7.OperatorMul, day7.OperatorConcat}),
	)
}

func TestConcat(t *testing.T) {
	require.Equal(t, 12, day7.Concat(1, 2))
	require.Equal(t, 123, day7.Concat(12, 3))
	require.Equal(t, 1234, day7.Concat(123, 4))
	require.Equal(t, 12345, day7.Concat(123, 45))
	require.Equal(t, 681928, day7.Concat(68, 1928))
}
