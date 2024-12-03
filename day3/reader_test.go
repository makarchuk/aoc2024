package day3_test

import (
	"io"
	"testing"

	"github.com/makarchuk/aoc2024/day3"
	"github.com/stretchr/testify/require"
)

func TestFindFirst(t *testing.T) {
	reader := day3.NewReader([]byte("asdmul1mul2"))
	err := reader.FindFirst('m')
	require.NoError(t, err)
	require.True(t, reader.ConsumeNext("ul1"))
}

func TestFailingConsumeNotSkipping(t *testing.T) {
	reader := day3.NewReader([]byte("asdmmul1mul2"))
	require.NoError(t, reader.FindFirst('m'))
	require.False(t, reader.ConsumeNext("ul1"))
	require.NoError(t, reader.FindFirst('m'))
	require.True(t, reader.ConsumeNext("ul1"))
	require.NoError(t, reader.FindFirst('m'))
	require.True(t, reader.ConsumeNext("ul2"))
	require.ErrorIs(t, io.EOF, reader.FindFirst('m'))
}

func TestNumbersLoading(t *testing.T) {
	reader := day3.NewReader([]byte("asdm1m22m0m333m4444"))
	require.NoError(t, reader.FindFirst('m'))
	num, err := reader.ReadInt(3)
	require.NoError(t, err)
	require.Equal(t, 1, num)
	require.NoError(t, reader.FindFirst('m'))
	num, err = reader.ReadInt(3)
	require.NoError(t, err)
	require.Equal(t, 22, num)
	require.NoError(t, reader.FindFirst('m'))
	num, err = reader.ReadInt(3)
	require.NoError(t, err)
	require.Equal(t, 0, num)
	require.NoError(t, reader.FindFirst('m'))
	num, err = reader.ReadInt(3)
	require.NoError(t, err)
	require.Equal(t, 333, num)
	require.NoError(t, reader.FindFirst('m'))
	num, err = reader.ReadInt(3)
	require.NoError(t, err)
	require.Equal(t, 444, num)
	require.True(t, reader.ConsumeNext("4"))
}
