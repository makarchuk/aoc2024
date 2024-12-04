package main

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
	rows [][]byte
}

var allowedDiags = map[string]bool{
	"MS": true,
	"SM": true,
}

func main() {
	var input Input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := make([]byte, len(scanner.Bytes()))
		copy(row, scanner.Bytes())
		input.rows = append(input.rows, row)
	}

	totalFound := 0

	for y, row := range input.rows {
		for x := range row {
			char, ok := safeGet(input.rows, x, y)
			if !ok {
				panic("should not happen")
			}

			if char != 'A' {
				continue
			}

			mainDiag, ok := readDiagonal(input.rows, x, y, true)
			if !ok {
				continue
			}
			if !allowedDiags[string(mainDiag)] {
				continue
			}

			antiDiag, ok := readDiagonal(input.rows, x, y, false)
			if !ok {
				continue
			}
			if !allowedDiags[string(antiDiag)] {
				continue
			}

			totalFound++
		}
	}

	fmt.Println(totalFound)
}

func readDiagonal(matrix [][]byte, x, y int, main bool) ([]byte, bool) {
	var result []byte
	var points [][]int
	if main {
		points = [][]int{{x + 1, y + 1}, {x - 1, y - 1}}
	} else {
		points = [][]int{{x + 1, y - 1}, {x - 1, y + 1}}
	}

	for _, point := range points {
		char, ok := safeGet(matrix, point[0], point[1])
		if !ok {
			return nil, false
		}
		result = append(result, char)
	}
	return result, true
}

func safeGet(matrix [][]byte, x, y int) (byte, bool) {
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) {
		return 0, false
	}
	return matrix[y][x], true
}
