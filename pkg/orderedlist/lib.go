package orderedlist

import "cmp"

type List[Item cmp.Ordered] struct {
	items []Item
}

func (l *List[Item]) Add(item Item) {
	for i, v := range l.items {
		if item < v {
			l.items = append(l.items[:i], append([]Item{item}, l.items[i:]...)...)
			return
		}
	}
	l.items = append(l.items, item)
}

func (l *List[Item]) Items() []Item {
	return l.items
}

func (l *List[Item]) Contains(item Item) bool {
	low := 0
	high := len(l.items) - 1

	for low <= high {
		median := (low + high) / 2
		if l.items[median] == item {
			return true
		} else if l.items[median] < item {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return false
}
