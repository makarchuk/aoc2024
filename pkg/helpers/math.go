package helpers

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Digits(num int) int {
	if num == 0 {
		return 1
	}
	digits := 0
	for num != 0 {
		digits++
		num /= 10
	}
	return digits
}
