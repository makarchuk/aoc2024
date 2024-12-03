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

	for {
		char, err := reader.Next()
		if err != nil {
			break
		}
		if char != 'm' {
			continue
		}

		res, ok := reader.ReadMultiplicationCommand()
		if ok {
			total += res
		}
	}

	fmt.Println(total)
}
