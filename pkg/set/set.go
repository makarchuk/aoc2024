package set

import (
	"encoding/json"
	"iter"
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]struct{})}
}

func From[T comparable](l []T) Set[T] {
	s := New[T]()
	for _, v := range l {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s.m)
}

func (s Set[T]) List() []T {
	var l []T
	for k := range s.m {
		l = append(l, k)
	}
	return l
}

func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s.m {
			if !yield(k) {
				return
			}
		}
	}
}

func (s Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s Set[T]) Clone() Set[T] {
	c := New[T]()
	for k := range s.m {
		c.Add(k)
	}
	return c
}

func (s Set[T]) Equal(o Set[T]) bool {
	if s.Len() != o.Len() {
		return false
	}

	for k := range s.m {
		if !o.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	i := New[T]()
	for k := range s.m {
		if other.Contains(k) {
			i.Add(k)
		}
	}
	return i
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.List())
}
