package day9_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/makarchuk/aoc2024/day9"
	"github.com/stretchr/testify/require"
)

var defragmentedExample = []int{
	0, 0,
	9, 9,
	8,
	1, 1, 1,
	8, 8, 8,
	2,
	7, 7, 7,
	3, 3, 3,
	6,
	4, 4,
	6,
	5, 5, 5, 5,
	6, 6,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

func TestRender(t *testing.T) {
	expected := []int{
		0, 0,
		-1, -1, -1,
		1, 1, 1,
		-1, -1, -1,
		2,
		-1, -1, -1,
		3, 3, 3,
		-1,
		4, 4,
		-1,
		5, 5, 5, 5,
		-1,
		6, 6, 6, 6,
		-1,
		7, 7, 7,
		-1,
		8, 8, 8, 8,
		9, 9,
	}

	require.Equal(
		t,
		expected,
		diskFrom(t, "2333133121414131402"),
	)
}

func TestDefragmentation(t *testing.T) {
	disk := diskFrom(t, "2333133121414131402")

	day9.Defragment(disk)
	require.Equal(t, defragmentedExample, disk)
}

func TestChecksum(t *testing.T) {
	require.EqualValues(t, 1928, day9.CheckSum(defragmentedExample))
}

func TestE2EPart1(t *testing.T) {
	type testcase struct {
		input    string
		checksum int64
	}

	for _, tc := range []testcase{
		{
			input:    "12345",
			checksum: 60,
		},
		{
			input:    "1010101010101010101010",
			checksum: 385,
		},
		{
			input:    "252",
			checksum: 5,
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			input := bytes.NewBuffer([]byte(tc.input))
			dm, err := day9.New(input)
			require.NoError(t, err)
			rendered := dm.Render()
			day9.Defragment(rendered)
			fmt.Println(rendered)
			require.Equal(t, tc.checksum, day9.CheckSum(rendered))
		})
	}
}

func TestPart2Defragmentation(t *testing.T) {
	type testcase struct {
		input  string
		rawMem string
		defMem string
	}

	for _, tc := range []testcase{
		{
			input:  "2333133121414131402",
			rawMem: "00...111...2...333.44.5555.6666.777.888899",
			defMem: "00992111777.44.333....5555.6666.....8888..",
		},
		{
			input:  "1910663",
			rawMem: "0.........1222222......333",
			defMem: "03332222221...............",
		},
		{
			input:  "1910604",
			rawMem: "0.........12222223333",
			defMem: "033331.....222222....",
		},
		{
			input:  "1313132",
			rawMem: "0...1...2...33",
			defMem: "03321.........",
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			defrag := defragmenterFrom(t, tc.input)
			require.Equal(
				t,
				tc.rawMem,
				day9.RenderMemory(defrag.Memory, false),
			)

			require.Equal(
				t,
				tc.defMem,
				day9.RenderMemory(defrag.Defragment(), false),
			)
		})
	}
}

func TestPart2E2E(t *testing.T) {
	type testcase struct {
		input    string
		checksum int64
	}

	for _, tc := range []testcase{
		{
			input:    "2333133121414131402",
			checksum: 2858,
		},
		{
			input:    "1010101010101010101010",
			checksum: 385,
		},
		{
			input:    "12345",
			checksum: 132,
		},
		{
			input:    "354631466260",
			checksum: 1325,
		},
		{
			input:    "171010402",
			checksum: 88,
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			input := bytes.NewBuffer([]byte(tc.input))
			dm, err := day9.New(input)
			require.NoError(t, err)
			defrag := dm.Defragmenter()

			require.Equal(t, tc.checksum, day9.CheckSum(defrag.Defragment()))
		})
	}
}

func diskFrom(t *testing.T, data string) []int {
	t.Helper()
	input := bytes.NewBuffer([]byte(data))
	dm, err := day9.New(input)
	require.NoError(t, err)
	return dm.Render()
}

func defragmenterFrom(t *testing.T, data string) day9.Defragmenter {
	t.Helper()
	input := bytes.NewBuffer([]byte(data))
	dm, err := day9.New(input)
	require.NoError(t, err)
	return dm.Defragmenter()
}
