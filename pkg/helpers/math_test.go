package helpers_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/pkg/helpers"
	"github.com/stretchr/testify/require"
)

func TestDigits(t *testing.T) {
	require.Equal(t, 1, helpers.Digits(0))
	require.Equal(t, 1, helpers.Digits(7))
	require.Equal(t, 2, helpers.Digits(83))
	require.Equal(t, 3, helpers.Digits(300))
	require.Equal(t, 9, helpers.Digits(111111111))
	require.Equal(t, 4, helpers.Digits(1000))
}
