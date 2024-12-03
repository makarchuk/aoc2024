package main

import (
	"fmt"
	"io"
	"os"

	"github.com/makarchuk/aoc2024/day3"
)

type Input struct {
	Memory []byte
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := Input{
		Memory: bytes,
	}

	reader := day3.NewReader(input.Memory)

	total := 0

	enabled := true
	for {
		char, err := reader.Next()
		if err != nil {
			break
		}

		switch char {
		case 'm':
			if !enabled {
				continue
			}
			res, ok := reader.ReadMultiplicationCommand()
			if ok {
				total += res
			}
		case 'd':
			if !reader.ConsumeNext("o") {
				continue
			}
			next := reader.Peak()
			switch next {
			case 'n':
				if reader.ConsumeNext("n't()") {
					enabled = false
				}
				continue
			case '(':
				if reader.ConsumeNext("()") {
					enabled = true
				}
				continue
			}
		default:
			continue
		}
	}

	fmt.Println(total)
}
