package helpers

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
