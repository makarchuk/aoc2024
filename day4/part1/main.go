package main

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
	rows [][]byte
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

			if char != 'X' {
				continue
			}

			if checkDirection(input.rows, []byte("MAS"), x, y, 1, 0) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, 0, 1) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, 0, -1) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, -1, 0) {
				totalFound++
			}

			if checkDirection(input.rows, []byte("MAS"), x, y, 1, 1) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, 1, -1) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, -1, 1) {
				totalFound++
			}
			if checkDirection(input.rows, []byte("MAS"), x, y, -1, -1) {
				totalFound++
			}
		}
	}

	fmt.Println(totalFound)
}

func checkDirection(matrix [][]byte, word []byte, x, y, dx, dy int) bool {
	for _, expected := range word {
		x += dx
		y += dy
		char, ok := safeGet(matrix, x, y)
		if !ok {
			return false
		}
		if char != expected {
			return false
		}
	}
	return true
}

func safeGet(matrix [][]byte, x, y int) (byte, bool) {
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) {
		return 0, false
	}
	return matrix[y][x], true
}
