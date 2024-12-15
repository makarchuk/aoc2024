package day15

import (
	"bytes"
	"fmt"
	"io"
)

func Part2(in io.Reader) (string, error) {
	raw, err := io.ReadAll(in)
	if err != nil {
		return "", err
	}

	for _, substitution := range [][2][]byte{
		{[]byte("#"), []byte("##")},
		{[]byte("."), []byte("..")},
		{[]byte("O"), []byte("[]")},
		{[]byte("@"), []byte("@.")},
	} {
		raw = bytes.ReplaceAll(
			raw, substitution[0], substitution[1],
		)
	}

	input, err := ParseInput(bytes.NewReader(raw))
	if err != nil {
		return "", err
	}

	state := State{
		sub:         input.submarine,
		boxes:       input.boxes,
		doubleBoxes: input.doubleBoxes,
	}

	// fmt.Println("\x1b[2J\x1b[H")
	// state.Print(input.warehouse)

	for _, cmd := range input.commands {
		state = state.ExecuteWithDoubleBoxes(cmd, input.warehouse)
		// os.Stdin.Read(make([]byte, 1))
		// fmt.Println("\x1b[2J\x1b[H")
		// fmt.Println("===========================")
		// fmt.Printf("cmd: %v\n", cmd)
		// state.Print(input.warehouse)
	}

	sum := 0

	for _, box := range state.doubleBoxes.List() {
		sum += 100*box.Y + box.X
	}

	return fmt.Sprintf("%d", sum), nil
}
