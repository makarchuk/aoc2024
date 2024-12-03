package day2

import (
	"github.com/makarchuk/aoc2024/pkg/helpers"
)

func CheckReport(report []int) bool {
	if len(report) < 2 {
		return true
	}

	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		if increasing && report[i] >= report[i+1] {
			return false
		}
		if !increasing && report[i] <= report[i+1] {
			return false
		}

		if helpers.Abs(report[i]-report[i+1]) > 3 {
			return false
		}
	}
	return true
}

func CheckReportsDampened(report []int) bool {

	increasesCount := map[bool]int{}

	increasesCount[report[1] > report[0]]++
	increasesCount[report[2] > report[1]]++
	increasesCount[report[3] > report[2]]++

	increasing := increasesCount[true] > increasesCount[false]

	badIndices := []int{}

	for i := 0; i < len(report)-1; i++ {
		if increasing && report[i] >= report[i+1] {
			badIndices = append(badIndices, i, i+1)
			break
		}
		if !increasing && report[i] <= report[i+1] {
			badIndices = append(badIndices, i, i+1)
			break
		}

		if helpers.Abs(report[i]-report[i+1]) > 3 {
			badIndices = append(badIndices, i, i+1)
			break
		}
	}
	if len(badIndices) == 0 {
		return true
	}

	for _, index := range badIndices {
		newReport := make([]int, len(report))
		copy(newReport, report)
		newReport = append(newReport[:index], newReport[index+1:]...)
		if CheckReport(newReport) {
			return true
		}
	}

	return false
}
