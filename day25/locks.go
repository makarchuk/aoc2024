package day25

import (
	"bufio"
	"fmt"
	"io"
)

type Key [5]int
type Lock [5]int

type Input struct {
	Keys  []Key
	Locks []Lock
}

const (
	keystart  = "....."
	lockstart = "#####"

	empty = '.'
	full  = '#'
)

func Part1(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	match := 0
	for _, key := range input.Keys {
		for _, lock := range input.Locks {
			if !Overlap(key, lock) {
				match++
			}
		}
	}

	return fmt.Sprint(match), nil
}

func ParseInput(in io.Reader) (Input, error) {
	input := Input{}

	scanner := bufio.NewScanner(in)

	for {
		if !scanner.Scan() {
			return input, fmt.Errorf("unexpected end of input")
		}
		switch scanner.Text() {
		case keystart:
			key := Key{-1, -1, -1, -1, -1}
			for row := 0; row < 6; row++ {
				if !scanner.Scan() {
					return input, fmt.Errorf("unexpected end of input")
				}
				for i, b := range scanner.Bytes() {
					if b == full && key[i] == -1 {
						key[i] = 5 - row
					}
				}
			}
			input.Keys = append(input.Keys, key)

		case lockstart:
			lock := Lock{-1, -1, -1, -1, -1}
			for row := 0; row < 6; row++ {
				if !scanner.Scan() {
					return input, fmt.Errorf("unexpected end of input")
				}
				for i, b := range scanner.Bytes() {
					if b == empty && lock[i] == -1 {
						lock[i] = row
					}
				}
			}
			input.Locks = append(input.Locks, lock)
		default:
			return input, fmt.Errorf("unexpected input: %s", scanner.Text())
		}

		if !scanner.Scan() && len(scanner.Bytes()) == 0 {
			break
		}
	}

	return input, nil
}

func Overlap(key Key, lock Lock) bool {
	for i := range key {
		if key[i]+lock[i] > 5 {
			return true
		}
	}
	return false
}
