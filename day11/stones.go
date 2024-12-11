package day11

import (
	"bytes"
	"io"
	"math"

	"github.com/makarchuk/aoc2024/pkg/helpers"
)

type Stones struct {
	stones []int
}

func NewStones(in io.Reader) (*Stones, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	data = bytes.Trim(data, "\n")
	stones, err := helpers.ParseIntsArray(string(data), " ")
	if err != nil {
		return nil, err
	}

	return &Stones{stones: stones}, nil
}

func (s *Stones) Stones() []int {
	return s.stones
}

func nextGeneration(s int) []int {
	switch {
	case s == 0:
		return []int{1}
	case helpers.Digits(s)%2 == 0:
		halfs := CutInHalf(s)
		return []int{halfs[0], halfs[1]}
	default:
		return []int{s * 2024}
	}
}

type Counter struct {
	stones map[int]int
}

func NewCounter(stones []int) Counter {
	counter := Counter{
		stones: make(map[int]int),
	}
	for _, s := range stones {
		counter.stones[s] += 1
	}
	return counter
}

func (c *Counter) NextGeneration() Counter {
	next := Counter{
		stones: make(map[int]int, len(c.stones)),
	}

	for s, count := range c.stones {
		for _, stone := range nextGeneration(s) {
			next.stones[stone] += count
		}
	}
	return next
}

func (c *Counter) Len() int {
	sum := 0
	for _, val := range c.stones {
		sum += val
	}
	return sum
}

func CutInHalf(num int) [2]int {
	res := [2]int{}
	digits := helpers.Digits(num)
	order := int(math.Pow10(digits / 2))
	res[0] = num / order
	res[1] = num % order
	return res
}
