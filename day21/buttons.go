package day21

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
)

type Input struct {
	Combinations []string
}

func ParseInput(in io.Reader) (Input, error) {
	var res Input
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		res.Combinations = append(res.Combinations, scanner.Text())
	}
	return res, scanner.Err()
}

func Part1(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	arrowPad := NewPad(ArrowPadButtons)
	numbersPad := NewPad(NumberPadButtons)

	var provider CostProvider = ManualPresser{}

	for _ = range 2 {
		provider = BuildPrecalculatedCostProvider(*arrowPad, provider)
		// fmt.Printf("Build provider for robot-%v\n", i)
	}

	finalProvider := BuildPrecalculatedCostProvider(*numbersPad, provider)

	totalCost := 0

	for _, combination := range input.Combinations {
		res := EnterCombinationPrecalculated(combination, finalProvider)
		var numPart int
		_, err := fmt.Sscanf(string(combination), "%dA", &numPart)
		if err != nil {
			return "", err
		}
		// fmt.Printf("Adding: %v*%v\n", res, numPart)
		totalCost += res * numPart
	}

	return fmt.Sprint(totalCost), nil
}

func Part2(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	arrowPad := NewPad(ArrowPadButtons)
	numbersPad := NewPad(NumberPadButtons)

	var provider CostProvider = ManualPresser{}

	for _ = range 25 {
		provider = BuildPrecalculatedCostProvider(*arrowPad, provider)
		// fmt.Printf("Build provider for robot-%v\n", i)
	}

	finalProvider := BuildPrecalculatedCostProvider(*numbersPad, provider)

	totalCost := 0

	for _, combination := range input.Combinations {
		res := EnterCombinationPrecalculated(combination, finalProvider)
		var numPart int
		_, err := fmt.Sscanf(string(combination), "%dA", &numPart)
		if err != nil {
			return "", err
		}
		// fmt.Printf("Adding: %v*%v\n", res, numPart)
		totalCost += res * numPart
	}

	return fmt.Sprint(totalCost), nil
}

type keypad struct {
	buttons    map[byte]field.Point
	reverseMap map[field.Point]byte
}

type button struct {
	direction field.Direction
	button    byte
}

var moveButtons []button = []button{
	{field.RIGHT, '>'},
	{field.DOWN, 'v'},
	{field.LEFT, '<'},
	{field.UP, '^'},
}

var NumberPadButtons map[byte]field.Point = map[byte]field.Point{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},

	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},

	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},

	'0': {1, 3},
	'A': {2, 3},
}

var ArrowPadButtons map[byte]field.Point = map[byte]field.Point{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

func NewPad(buttons map[byte]field.Point) *keypad {
	reverseMap := make(map[field.Point]byte)
	for k, v := range buttons {
		if _, ok := reverseMap[v]; ok {
			panic(fmt.Sprintf("duplicate button %v", v))
		}
		reverseMap[v] = k
	}
	return &keypad{
		buttons:    buttons,
		reverseMap: reverseMap,
	}
}

func (p *keypad) Coordinates(button byte) (field.Point, error) {
	if v, ok := p.buttons[button]; ok {
		return v, nil
	}
	return field.Point{}, fmt.Errorf("button %v not found", button)
}

func (p *keypad) AllowedPoint(pnt field.Point) bool {
	_, ok := p.reverseMap[pnt]
	return ok
}
