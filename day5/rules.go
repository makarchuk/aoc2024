package day5

import "slices"

type Rules struct {
	mustBeBefore map[int][]int
}

func NewRules() *Rules {
	return &Rules{
		mustBeBefore: make(map[int][]int),
	}
}

func (r *Rules) AddRule(key, value int) {
	r.mustBeBefore[key] = append(r.mustBeBefore[key], value)
}

func (r *Rules) CanBeAfter(page int, after []int) bool {
	mustBeBefore := r.mustBeBefore[page]
	for _, page := range after {
		if slices.Contains(mustBeBefore, page) {
			return false
		}
	}
	return true
}

func (r *Rules) Part1Check(sequence []int) bool {
	for i, page := range sequence {
		if i == 0 {
			continue
		}
		if !r.CanBeAfter(page, sequence[:i]) {
			return false
		}
	}
	return true
}
