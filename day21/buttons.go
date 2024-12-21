package day21

import (
	"bufio"
	"fmt"
	"io"
	"strings"

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

	directPresser := NewManualPresser()

	historianRobot := NewRobotPresser(
		ArrowPadButtons,
		directPresser,
		"historian",
	)

	freezingRobot := NewRobotPresser(
		ArrowPadButtons,
		historianRobot,
		"frosty",
	)

	radiationRobot := NewRobotPresser(
		NumberPadButtons,
		freezingRobot,
		"radiation",
	)

	totalCost := 0

	for _, combination := range input.Combinations {
		res := EnterCombination([]byte(combination), radiationRobot)
		var numPart int
		_, err := fmt.Sscanf(string(combination), "%dA", &numPart)
		if err != nil {
			return "", err
		}
		fmt.Printf("Adding: %v*%v\n", len(res), numPart)
		totalCost += len(res) * numPart
	}

	return fmt.Sprint(totalCost), nil
}

func EnterCombination(combinatino []byte, presser PadPresser) []byte {
	total := []byte{}
	for _, c := range combinatino {
		res := presser.Press(byte(c))
		total = append(total, res...)
	}
	return total
}

type KeyPad interface {
	Coordinates(button byte) (field.Point, error)
	AllowedPoint(p field.Point) bool
}

type pad struct {
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

func NewPad(buttons map[byte]field.Point) KeyPad {
	reverseMap := make(map[field.Point]byte)
	for k, v := range buttons {
		if _, ok := reverseMap[v]; ok {
			panic(fmt.Sprintf("duplicate button %v", v))
		}
		reverseMap[v] = k
	}
	return &pad{
		buttons:    buttons,
		reverseMap: reverseMap,
	}
}

func (p *pad) Coordinates(button byte) (field.Point, error) {
	if v, ok := p.buttons[button]; ok {
		return v, nil
	}
	return field.Point{}, fmt.Errorf("button %v not found", button)
}

func (p *pad) AllowedPoint(pnt field.Point) bool {
	_, ok := p.reverseMap[pnt]
	return ok
}

// PadPresser describes a device that can press buttons
type PadPresser interface {
	Press(button byte) []byte
	Clone() PadPresser
}

type ManualPresser struct{}

func (m ManualPresser) Press(button byte) []byte {
	return []byte{button}
}

func (m ManualPresser) Clone() PadPresser {
	return ManualPresser{}
}

func NewManualPresser() PadPresser {
	return ManualPresser{}
}

type move [2]byte

type RobotPresser struct {
	pad             KeyPad
	presser         PadPresser
	cache           map[move][]byte
	currentPosition byte
	name            string
}

func (r *RobotPresser) Clone() PadPresser {
	return &RobotPresser{
		pad:             r.pad,
		presser:         r.presser.Clone(),
		cache:           map[move][]byte{},
		currentPosition: r.currentPosition,
		name:            r.name,
	}
}

func NewRobotPresser(buttons map[byte]field.Point, presser PadPresser, name string) *RobotPresser {
	return &RobotPresser{
		pad:     NewPad(buttons),
		presser: presser,
		cache:   map[move][]byte{},
		//default possition is always 'A'
		currentPosition: 'A',
		name:            name,
	}
}

// Press is always moved starting from A to our target button
func (r *RobotPresser) Press(button byte) []byte {
	// fmt.Printf("Want %v to press %c. Currently at %c\n", r.name, button, r.currentPosition)
	moveKey := move{r.currentPosition, button}
	fullPath, ok := r.cache[moveKey]
	if !ok {
		//move position on current keypad to `button`
		//press (presser presses A)
		var movesToButton []byte
		if r.currentPosition != button {
			movesToButton = r.pressSlow(r.currentPosition, button)
		}
		buttonPress := r.presser.Press('A')
		fullPath = append(movesToButton, buttonPress...)
		r.cache[moveKey] = fullPath
	}
	r.currentPosition = button
	// fmt.Printf("%v: moves from %c to %c: %v\n", r.name, r.currentPosition, button, PrintPresses(fullPath))
	return fullPath
}

type moverState struct {
	position    field.Point
	presser     PadPresser
	accumulated []byte
}

func (r *RobotPresser) pressSlow(from, to byte) []byte {
	currentCoordinates, err := r.pad.Coordinates(from)
	if err != nil {
		panic(fmt.Errorf("%v: could not find coordinates for button %c", r.name, from))
	}

	target, err := r.pad.Coordinates(to)
	if err != nil {
		panic(fmt.Errorf("%v; could not find coordinates for button %c", r.name, to))
	}

	var bestSolution []byte
	var bestPresser PadPresser

	frontier := []moverState{{
		position:    currentCoordinates,
		accumulated: []byte{},
		presser:     r.presser.Clone(),
	}}

	for {
		nextFrontier := []moverState{}
		for _, state := range frontier {

			if len(bestSolution) > 0 && len(state.accumulated) >= len(bestSolution) {
				continue
			}

			for _, moveButton := range moveButtons {
				next := state.position.Move(moveButton.direction)
				if !r.pad.AllowedPoint(next) {
					continue
				}

				newPresser := state.presser.Clone()

				neccessarryButtons := newPresser.Press(moveButton.button)

				newAccumulated := make([]byte, 0, len(state.accumulated)+len(neccessarryButtons))
				newAccumulated = append(newAccumulated, state.accumulated...)
				newAccumulated = append(newAccumulated, neccessarryButtons...)

				if next == target {
					if bestSolution == nil || len(newAccumulated) < len(bestSolution) {
						bestSolution = newAccumulated
						bestPresser = newPresser
					}
					continue
				}

				nextFrontier = append(nextFrontier, moverState{
					position:    next,
					accumulated: newAccumulated,
					presser:     newPresser,
				})
			}
		}
		frontier = nextFrontier
		if len(frontier) == 0 {
			break
		}
	}

	if bestSolution == nil {
		panic(fmt.Errorf("could not find solution from %v to %v", from, to))
	}
	r.presser = bestPresser
	return bestSolution
}

func PrintPresses(r []byte) string {
	s := strings.Builder{}
	s.WriteByte('[')
	for _, b := range r {
		s.WriteByte(b)
	}
	s.WriteByte(']')
	return s.String()
}
