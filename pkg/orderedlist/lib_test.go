package orderedlist_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/makarchuk/aoc2024/pkg/orderedlist"
)

func TestList_Add(t *testing.T) {
	var l = orderedlist.List[int]{}
	l.Add(3)
	l.Add(1)
	l.Add(2)
	l.Add(4)
	require.Equal(t, []int{1, 2, 3, 4}, l.Items())
}

func TestList_Contains(t *testing.T) {
	var l = orderedlist.List[int]{}
	l.Add(3)
	l.Add(4)
	l.Add(7)
	l.Add(28)

	require.True(t, l.Contains(3))
	require.True(t, l.Contains(4))
	require.True(t, l.Contains(7))
	require.True(t, l.Contains(28))
	require.False(t, l.Contains(2))
	require.False(t, l.Contains(193))
}
