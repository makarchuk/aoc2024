package day15

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type Input struct {
	warehouse   field.Field
	submarine   field.Point
	boxes       set.Set[field.Point]
	doubleBoxes set.Set[field.Point]
	commands    []field.Direction
}

const (
	EMPTY          byte = '.'
	WALL           byte = '#'
	BOX            byte = 'O'
	SUB            byte = '@'
	DOUBLEBOXSTART byte = '['
	DOUBLEBOXEND   byte = ']'
)

type State struct {
	sub         field.Point
	boxes       set.Set[field.Point]
	doubleBoxes set.Set[field.Point]
}

func ParseInput(in io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(in)
	warehouse, err := field.Read(scanner)
	if err != nil {
		return nil, err
	}

	removed, newWh := warehouse.Extract([]byte{BOX, SUB, DOUBLEBOXSTART, DOUBLEBOXEND}, EMPTY)

	if len(removed['@']) != 1 {
		return nil, fmt.Errorf("submarine not found")
	}

	submarine := removed[SUB][0]
	boxes := set.From(removed[BOX])
	doubleBoxes := set.From(removed[DOUBLEBOXSTART])

	var commands []field.Direction
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			switch b {
			case '>':
				commands = append(commands, field.RIGHT)
			case '<':
				commands = append(commands, field.LEFT)
			case '^':
				commands = append(commands, field.UP)
			case 'v':
				commands = append(commands, field.DOWN)
			default:
				return nil, fmt.Errorf("invalid command: %v", b)
			}
		}
	}

	return &Input{
		warehouse:   newWh,
		commands:    commands,
		submarine:   submarine,
		boxes:       boxes,
		doubleBoxes: doubleBoxes,
	}, nil
}

func (s *State) Execute(cmd field.Direction, wh field.Field) State {
	newSub := s.sub.Move(cmd)

	val, ok := wh.Get(newSub)
	if !ok {
		panic(fmt.Sprintf("submarine out of bounds: %v", newSub))
	}
	if val == WALL {
		return *s
	}

	if !s.boxes.Contains(newSub) {
		return State{
			sub:   newSub,
			boxes: s.boxes,
		}
	}

	newBoxes := s.boxes.Clone()
	newBoxes.Remove(newSub)
	movingBox := newSub

	for {
		newBoxPosition := movingBox.Move(cmd)
		val, ok := wh.Get(newBoxPosition)
		if !ok {
			panic("box out of bounds")
		}
		if val == WALL {
			return *s
		}

		if !newBoxes.Contains(newBoxPosition) {
			newBoxes.Add(newBoxPosition)
			break
		}

		movingBox = newBoxPosition
	}

	return State{
		sub:   newSub,
		boxes: newBoxes,
	}
}

func (s *State) ExecuteWithDoubleBoxes(cmd field.Direction, wh field.Field) State {
	newSub := s.sub.Move(cmd)

	val, ok := wh.Get(newSub)
	if !ok {
		panic(fmt.Sprintf("submarine out of bounds: %v", newSub))
	}
	if val == WALL {
		return *s
	}

	if !s.doubleBoxes.Contains(newSub) && !s.doubleBoxes.Contains(newSub.Move(field.LEFT)) {
		return State{
			sub:         newSub,
			doubleBoxes: s.doubleBoxes,
		}
	}

	var movingBoxes set.Set[field.Point]
	newDoubleBoxes := s.doubleBoxes.Clone()
	if s.doubleBoxes.Contains(newSub) {
		newDoubleBoxes.Remove(newSub)
		movingBoxes = set.From([]field.Point{newSub.Move(cmd)})
	} else {
		newDoubleBoxes.Remove(newSub.Move(field.LEFT))
		movingBoxes = set.From([]field.Point{newSub.Move(field.LEFT).Move(cmd)})
	}

	for {
		newMovingBoxes := set.New[field.Point]()
		for _, movingBox := range movingBoxes.List() {
			for _, boxPosition := range []field.Point{movingBox, movingBox.Move(field.RIGHT)} {
				val, ok := wh.Get(boxPosition)
				if !ok {
					panic("box out of bounds")
				}
				if val == WALL {
					return *s
				}
			}
			collidesWith := s.doubleBoxCollision(movingBox, newDoubleBoxes)
			for _, collision := range collidesWith {
				newDoubleBoxes.Remove(collision)
				newMovingBoxes.Add(collision.Move(cmd))
			}
			newDoubleBoxes.Add(movingBox)
		}
		if newMovingBoxes.Len() == 0 {
			return State{
				sub:         newSub,
				doubleBoxes: newDoubleBoxes,
			}
		}
		movingBoxes = newMovingBoxes
	}
}

// get coordinates of the double box that double box possitioned at P would collide with
func (s *State) doubleBoxCollision(p field.Point, doubleBoxes set.Set[field.Point]) []field.Point {
	result := []field.Point{}
	if doubleBoxes.Contains(p) {
		result = append(result, p)
	}

	if doubleBoxes.Contains(p.Move(field.RIGHT)) {
		box := p.Move(field.RIGHT)
		result = append(result, box)
	}

	if doubleBoxes.Contains(p.Move(field.LEFT)) {
		box := p.Move(field.LEFT)
		result = append(result, box)
	}

	return result
}

func (s *State) Print(f field.Field) {
	for y := 0; y <= f.Size().Y; y++ {
		for x := 0; x <= f.Size().X; x++ {
			p := field.Point{X: x, Y: y}
			if s.sub == p {
				fmt.Print(string(SUB))
			} else if s.boxes.Contains(p) {
				fmt.Print(string(BOX))
			} else if s.doubleBoxes.Contains(p) {
				fmt.Print(string(DOUBLEBOXSTART))
			} else if s.doubleBoxes.Contains(p.Move(field.LEFT)) {
				fmt.Print(string(DOUBLEBOXEND))
			} else {
				val, _ := f.Get(p)
				fmt.Print(string(val))
			}
		}
		fmt.Println()
	}
}
