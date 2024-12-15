package aoc

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/day14"
	"github.com/makarchuk/aoc2024/day15"
	_ "github.com/makarchuk/aoc2024/day15"
)

type entry struct {
	day  int
	part int
}

var days map[entry]func(io.Reader) (string, error) = make(map[entry]func(io.Reader) (string, error))

func Register(day int, part int, f func(io.Reader) (string, error)) {
	days[entry{day, part}] = f
}

func Call(day int, part int, r io.Reader) (string, error) {
	f, ok := days[entry{day, part}]
	if !ok {
		return "", fmt.Errorf("Day %d part %d not registered", day, part)
	}
	res, err := f(r)
	return res, err
}

func init() {
	Register(14, 1, day14.Part1)
	Register(14, 2, day14.Part2)
	Register(15, 1, day15.Part1)
	Register(15, 2, day15.Part2)
}
