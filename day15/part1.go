package day15

import (
	"fmt"
	"io"
)

func Part1(r io.Reader) (string, error) {
	in, err := ParseInput(r)
	if err != nil {
		return "", err
	}

	state := State{
		sub:   in.submarine,
		boxes: in.boxes,
	}

	for _, cmd := range in.commands {
		state = state.Execute(cmd, in.warehouse)
	}

	sum := 0

	for _, box := range state.boxes.List() {
		sum += 100*box.Y + box.X
	}

	return fmt.Sprintf("%d", sum), nil
}
