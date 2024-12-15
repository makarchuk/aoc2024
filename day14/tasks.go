package day14

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/field"
)

const wipeScreen = "\x1b[2J\x1b[H"

func Part1(in io.Reader) (string, error) {
	inp, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	inp.Size = field.Point{X: 101, Y: 103}

	for _ = range 100 {
		inp.Step()
	}

	return fmt.Sprintf("%v", inp.SafetyFactor()), nil
}

func Part2(in io.Reader) (string, error) {
	inp, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	inp.Size = field.Point{X: 101, Y: 103}
	fmt.Println(wipeScreen)
	inp.Print()
	step := 0
	for {
		step++
		inp.Step()
		fmt.Println(wipeScreen)
		fmt.Printf(`
===============================
			step=%d
===============================
`, step)
		inp.Print()
		cmd := make([]byte, 8)
		n, err := os.Stdin.Read(cmd)
		if err != nil {
			return "", err
		}

		if bytes.HasPrefix(cmd, []byte("y")) {
			return fmt.Sprintf("%d", step), nil
		}
		skip, err := strconv.Atoi(strings.Trim(string(cmd[:n]), "\n"))
		if err == nil {
			for i := 0; i < skip; i++ {
				step++
				inp.Step()
			}
		}
	}

}
