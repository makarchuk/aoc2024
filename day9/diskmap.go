package day9

import (
	"bytes"
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/orderedlist"
)

const empty = -1

type DiskMap struct {
	compressed []byte
}

func New(input io.Reader) (*DiskMap, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	m := &DiskMap{}

	data = bytes.Trim(data, "\n")

	for _, b := range data {
		if b > '9' || b < '0' {
			return nil, fmt.Errorf("invalid byte %c", b)
		}
		m.compressed = append(m.compressed, b-'0')
	}

	return m, nil
}

func (d *DiskMap) Render() []int {
	return d.Defragmenter().Memory
}

func (d *DiskMap) Defragmenter() Defragmenter {
	defrag := Defragmenter{}

	data := true
	currentIndex := 0
	fileNum := 0
	totalSize := 0
	for _, b := range d.compressed {
		var num int
		totalSize += int(b)
		if data {
			num = fileNum
			defrag.DataChunks = append(defrag.DataChunks, DataChunk{
				FileNum:  fileNum,
				Length:   int(b),
				Position: currentIndex,
			})
			fileNum++
		} else {
			num = empty
			if defrag.EmptySpaces[b] == nil {
				defrag.EmptySpaces[b] = &orderedlist.List[int]{}
			}
			defrag.EmptySpaces[b].Add(currentIndex)
		}
		chunk := slices.Repeat([]int{num}, int(b))
		currentIndex += int(b)
		defrag.Memory = append(defrag.Memory, chunk...)
		data = !data
	}
	defrag.TotalSize = totalSize
	return defrag
}

type Defragmenter struct {
	DataChunks  []DataChunk
	EmptySpaces [10]*orderedlist.List[int]
	TotalSize   int
	Memory      []int
}

func (d *Defragmenter) Defragment() []int {
	result := make([]int, d.TotalSize)
	copy(result, d.Memory)

	for tail := len(d.DataChunks) - 1; tail >= 0; tail-- {
		tailChunk := d.DataChunks[tail]
		d.placeChunk(tailChunk, result)
		// fmt.Println(RenderMemory(result, false))
	}
	return result
}

func (d *Defragmenter) placeChunk(chunk DataChunk, result []int) {

	type possibleSlot struct {
		Index int
		Size  int
	}

	possibleSlots := []possibleSlot{}

	for slotSize := 9; slotSize > 0; slotSize-- {
		if chunk.Length > slotSize {
			break
		}

		if d.EmptySpaces[slotSize] == nil {
			continue
		}
		if len(d.EmptySpaces[slotSize].Items()) == 0 {
			continue
		}
		emptySlot := d.EmptySpaces[slotSize].Items()[0]

		if emptySlot > chunk.Position {
			continue
		}

		possibleSlots = append(possibleSlots, possibleSlot{Index: emptySlot, Size: slotSize})
	}

	if len(possibleSlots) == 0 {
		return
	}

	chosenSlot := slices.MinFunc(possibleSlots, func(a, b possibleSlot) int {
		return cmp.Compare(a.Index, b.Index)
	})
	d.EmptySpaces[chosenSlot.Size].Pop()

	remainingEmpty := chosenSlot.Size - chunk.Length
	if remainingEmpty > 0 {
		if d.EmptySpaces[remainingEmpty] == nil {
			d.EmptySpaces[remainingEmpty] = &orderedlist.List[int]{}
		}
		d.EmptySpaces[remainingEmpty].Add(chosenSlot.Index + chunk.Length)
	}

	for i := range chunk.Length {
		result[chosenSlot.Index+i] = chunk.FileNum
		result[chunk.Position+i] = empty
	}

	return
}

type DataChunk struct {
	FileNum  int
	Length   int
	Position int
}

func (d *DataChunk) Render() []int {
	return slices.Repeat([]int{d.FileNum}, d.Length)
}

func Defragment(disk []int) {
	head := 0
	tail := len(disk) - 1

	for head <= tail {
		emptyIndex := findEmpty(disk, head, tail)
		if emptyIndex == -1 {
			break
		}
		dataIndex := findData(disk, tail, head)
		if dataIndex == -1 {
			break
		}
		if emptyIndex >= dataIndex {
			break
		}
		head = emptyIndex
		tail = dataIndex
		disk[emptyIndex], disk[dataIndex] = disk[dataIndex], disk[emptyIndex]
	}
}

func CheckSum(disk []int) int64 {
	checksum := int64(0)
	for i, fileNum := range disk {
		if fileNum == empty {
			continue
		}
		checksum += int64(i) * int64(fileNum)
	}
	return checksum
}

func findEmpty(disk []int, head int, maxHead int) int {
	for i := head; i < maxHead; i++ {
		if disk[i] == empty {
			return i
		}
	}
	return -1
}

func findData(disk []int, tail int, minTail int) int {
	for i := tail; i > minTail; i-- {
		if disk[i] != empty {
			return i
		}
	}
	return -1
}

func RenderMemory(disk []int, escape bool) string {
	s := strings.Builder{}
	for _, i := range disk {
		if i == empty {
			_, _ = s.WriteRune('.')
		} else {
			if escape {
				_, _ = s.WriteString(fmt.Sprintf("[%d]", i))
			} else {
				_, _ = s.WriteString(fmt.Sprintf("%d", i))
			}
		}
	}
	return s.String()
}
