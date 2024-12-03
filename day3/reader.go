package day3

import (
	"errors"
	"io"
)

type MemoryReader struct {
	memory []byte
	index  int
}

func NewReader(memory []byte) *MemoryReader {
	return &MemoryReader{
		memory: memory,
	}
}

func (m *MemoryReader) FindFirst(char byte) error {
	for {
		next := m.next()
		if next == char {
			return nil
		}
		if next == 0 {
			return io.EOF
		}
	}
}

func (m *MemoryReader) isNext(char []byte) bool {
	if m.index+len(char) > len(m.memory) {
		return false
	}
	for i, c := range char {
		if m.memory[m.index+i] != c {
			return false
		}
	}
	return true
}

func (m *MemoryReader) Peak() byte {
	if m.index >= len(m.memory) {
		return 0
	}
	return m.memory[m.index]
}

func (m *MemoryReader) next() byte {
	if m.index >= len(m.memory) {
		return 0
	}
	m.index++
	return m.memory[m.index-1]
}

func (m *MemoryReader) ConsumeNext(s string) bool {
	input := []byte(s)
	if m.isNext(input) {
		m.index += len(input)
		return true
	}
	return false
}

func (m *MemoryReader) ReadInt(maxDigits int) (int, error) {
	num := 0
	for i := 0; i < maxDigits; i++ {
		next := m.Peak()
		if next < '0' || next > '9' {
			if i == 0 {
				return 0, errors.New("not a number")
			}
			return num, nil
		}
		num = num*10 + int(next-'0')
		m.index++
	}
	return num, nil
}

func (m *MemoryReader) Next() (byte, error) {
	if m.index >= len(m.memory) {
		return 0, io.EOF
	}
	m.index++
	return m.memory[m.index-1], nil
}

func (m *MemoryReader) ReadMultiplicationCommand() (int, bool) {
	if !m.ConsumeNext("ul") {
		return 0, false
	}

	if !m.ConsumeNext("(") {
		return 0, false
	}

	left, err := m.ReadInt(3)
	if err != nil {
		return 0, false
	}

	if !m.ConsumeNext(",") {
		return 0, false
	}

	right, err := m.ReadInt(3)
	if err != nil {
		return 0, false
	}

	if !m.ConsumeNext(")") {
		return 0, false
	}

	return left * right, true
}
