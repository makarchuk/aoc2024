package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/aoc"
)

//go:embed input/*
var INPUTS embed.FS

func main() {
	day := flag.Int("day", 0, "Day of the challenge")
	part := flag.Int("part", 0, "Part of the challenge")
	customInput := flag.String("input", "", "Custom input string")

	flag.Parse()

	if *day == 0 {
		panic("Day is required")
	}
	if *part == 0 {
		*part = 1
	}

	var input io.Reader

	if *customInput != "" {
		if *customInput == "-" {
			input = os.Stdin
		} else {
			input = strings.NewReader(*customInput)
		}
	} else {
		f, err := INPUTS.Open(fmt.Sprintf("input/day%d", *day))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		input = f
	}

	result, err := aoc.Call(*day, *part, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
