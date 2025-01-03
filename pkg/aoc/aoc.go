package aoc

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/day14"
	"github.com/makarchuk/aoc2024/day15"
	"github.com/makarchuk/aoc2024/day16"
	"github.com/makarchuk/aoc2024/day17"
	"github.com/makarchuk/aoc2024/day18"
	"github.com/makarchuk/aoc2024/day19"
	"github.com/makarchuk/aoc2024/day20"
	"github.com/makarchuk/aoc2024/day21"
	"github.com/makarchuk/aoc2024/day22"
	"github.com/makarchuk/aoc2024/day23"
	"github.com/makarchuk/aoc2024/day24"
	"github.com/makarchuk/aoc2024/day25"
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
	Register(16, 1, day16.Part1)
	Register(16, 2, day16.Part2)
	Register(17, 1, day17.Part1)
	Register(17, 2, day17.Part2)
	Register(18, 1, day18.Part1)
	Register(18, 2, day18.Part2)
	Register(19, 1, day19.Part1)
	Register(19, 2, day19.Part2)
	Register(20, 1, day20.Part1)
	Register(20, 2, day20.Part2)
	Register(21, 1, day21.Part1)
	Register(21, 2, day21.Part2)
	Register(22, 1, day22.Part1)
	Register(22, 2, day22.Part2)
	Register(23, 1, day23.Part1)
	Register(23, 2, day23.Part2)
	Register(24, 1, day24.Part1)
	Register(24, 2, day24.Part2)
	Register(25, 1, day25.Part1)
}
