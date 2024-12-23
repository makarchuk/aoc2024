package day23

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/set"
)

type Input struct {
	Connections map[string]set.Set[string]
}

func (i *Input) AddConnection(a, b string) {
	if i.Connections == nil {
		i.Connections = make(map[string]set.Set[string])
	}
	if _, ok := i.Connections[a]; !ok {
		i.Connections[a] = set.New[string]()
	}
	if _, ok := i.Connections[b]; !ok {
		i.Connections[b] = set.New[string]()
	}
	i.Connections[a].Add(b)
	i.Connections[b].Add(a)
}

func ParseInput(in *io.Reader) (*Input, error) {
	i := &Input{}
	scanner := bufio.NewScanner(*in)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input: %s", line)
		}
		a, b := parts[0], parts[1]
		i.AddConnection(a, b)
	}

	return i, nil
}

func Part1(in io.Reader) (string, error) {

	groupsOfThree := 0

	input, err := ParseInput(&in)
	if err != nil {
		return "", err
	}

	for a, connections := range input.Connections {
		connectionsList := connections.List()

		for i, b := range connectionsList {
			for _, c := range connectionsList[i:] {
				if input.Connections[b].Contains(c) {
					if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
						groupsOfThree++
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", groupsOfThree/3), nil
}
