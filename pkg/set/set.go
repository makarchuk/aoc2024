package set

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

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) List() []T {
	var l []T
	for k := range s.m {
		l = append(l, k)
	}
	return l
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}
