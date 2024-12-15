package field

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"slices"
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

func (p Point) Mul(n int) Point {
	return Point{p.X * n, p.Y * n}
}

func (p Point) Move(d Direction) Point {
	switch d {
	case UP:
		return Point{p.X, p.Y - 1}
	case RIGHT:
		return Point{p.X + 1, p.Y}
	case DOWN:
		return Point{p.X, p.Y + 1}
	case LEFT:
		return Point{p.X - 1, p.Y}
	default:
		panic(fmt.Sprintf("invalid direction: %v", d))
	}
}

type Direction byte

const (
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "^"
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	default:
		panic(fmt.Sprintf("invalid direction: %v", d))
	}
}

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

type Field struct {
	field [][]byte

	size Point
}

func (f Field) Size() Point {
	return f.size
}

func New(io io.Reader) (*Field, error) {
	scanner := bufio.NewScanner(io)
	field, err := Read(scanner)
	if err != nil {
		return nil, err
	}
	if scanner.Scan() {
		return nil, fmt.Errorf("extra input")
	}
	return field, nil
}

func Read(s *bufio.Scanner) (*Field, error) {
	field := &Field{}
	y := 0
	for s.Scan() {
		if len(s.Bytes()) == 0 {
			break
		}
		line := s.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		field.field = append(field.field, row)
		y++
	}
	field.size = Point{len(field.field[0]) - 1, len(field.field) - 1}
	return field, nil
}

func (f *Field) InField(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= f.size.X && p.Y <= f.size.Y
}

func (f *Field) Get(p Point) (byte, bool) {
	if !f.InField(p) {
		return 0, false
	}
	return f.field[p.Y][p.X], true
}

func (f *Field) Set(p Point, value byte) {
	f.field[p.Y][p.X] = value
}

func (f *Field) Replace(p Point, value byte) Field {
	newField := Field{
		field: make([][]byte, len(f.field)),
		size:  f.size,
	}

	for y, row := range f.field {
		newRow := make([]byte, len(row))
		copy(newRow, row)
		if y == p.Y {
			newRow[p.X] = value
		}
		newField.field[y] = newRow
	}
	return newField
}

func (f *Field) Iter() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for y, row := range f.field {
			for x := range row {
				if !yield(Point{x, y}) {
					return
				}
			}
		}
	}
}

// Extract removes all instances of charachters from the field, replacing them.
// Returns new field and map of removed points
func (f *Field) Extract(toRemove []byte, replacement byte) (map[byte][]Point, Field) {
	removed := make(map[byte][]Point)
	newField := Field{
		field: make([][]byte, len(f.field)),
		size:  f.size,
	}

	for y, row := range f.field {
		newRow := make([]byte, len(row))
		copy(newRow, row)
		for x, cell := range row {
			if slices.Contains(toRemove, cell) {
				removed[cell] = append(removed[cell], Point{x, y})
				newRow[x] = replacement
			}
		}
		newField.field[y] = newRow
	}
	return removed, newField
}
