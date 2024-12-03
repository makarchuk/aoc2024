package day2_test

import (
	"testing"

	"github.com/makarchuk/aoc2024/day2"
	"github.com/stretchr/testify/require"
)

func TestCheckSafe(t *testing.T) {
	require.True(t, day2.CheckReport([]int{1, 2, 3, 4}))
	require.True(t, day2.CheckReport([]int{4, 3, 2, 1}))
	require.True(t, day2.CheckReport([]int{17, 14, 11, 8}))
	require.False(t, day2.CheckReport([]int{17, 18, 2, 1}))
	require.False(t, day2.CheckReport([]int{17, 3, 2, 1}))
	require.False(t, day2.CheckReport([]int{1, 2, 4, 3}))
}

func TestCheckSafeDampened(t *testing.T) {
	require.True(t, day2.CheckReportsDampened([]int{7, 6, 4, 2, 1}))
	require.False(t, day2.CheckReportsDampened([]int{1, 2, 7, 8, 9}))
	require.False(t, day2.CheckReportsDampened([]int{9, 7, 6, 2, 1}))
	require.True(t, day2.CheckReportsDampened([]int{1, 3, 2, 4, 5}))
	require.True(t, day2.CheckReportsDampened([]int{8, 6, 4, 4, 1}))
	require.True(t, day2.CheckReportsDampened([]int{1, 3, 6, 7, 9}))
	require.True(t, day2.CheckReportsDampened([]int{17, 3, 6, 7, 9}))
	require.True(t, day2.CheckReportsDampened([]int{1, 3, 6, 7, 924}))
}

func TestArray(t *testing.T) {
	report := []int{1, 2, 3, 4, 5}
	index := 4
	newReport := append(report[:index], report[index+1:]...)
	require.EqualValues(t, []int{1, 2, 3, 4}, newReport)
}
