package helpers

import (
	"strconv"
	"strings"
)

func ParseIntsArray(input string, sep string) ([]int, error) {
	var result []int
	for _, s := range strings.Split(input, sep) {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
