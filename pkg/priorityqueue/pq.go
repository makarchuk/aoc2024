package priorityqueue

type Entry[T any] struct {
	Score int
	Value T
}

type PriorityQueue[T any] struct {
	entries []Entry[T]
}

func (pq *PriorityQueue[T]) Push(score int, value T) {
	pq.entries = append(pq.entries, Entry[T]{Score: score, Value: value})
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	var t T
	if len(pq.entries) == 0 {
		return t, false
	}

	maxIndex := 0
	for i, entry := range pq.entries {
		if entry.Score > pq.entries[maxIndex].Score {
			maxIndex = i
		}
	}

	value := pq.entries[maxIndex].Value

	pq.entries = append(pq.entries[:maxIndex], pq.entries[maxIndex+1:]...)
	return value, true
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.entries)
}
