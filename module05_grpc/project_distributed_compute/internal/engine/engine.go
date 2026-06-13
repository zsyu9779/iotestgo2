package engine

import (
	"fmt"
	"math"
	"sort"
)

type Operation string

const (
	OperationSum    Operation = "sum"
	OperationAvg    Operation = "avg"
	OperationMax    Operation = "max"
	OperationMin    Operation = "min"
	OperationStddev Operation = "stddev"
	OperationMedian Operation = "median"
)

func ParseOperation(value string) Operation {
	return Operation(value)
}

func Compute(operation Operation, numbers []int64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}

	switch operation {
	case OperationSum:
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		return float64(sum), nil
	case OperationAvg:
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		return float64(sum) / float64(len(numbers)), nil
	case OperationMax:
		max := numbers[0]
		for _, n := range numbers[1:] {
			if n > max {
				max = n
			}
		}
		return float64(max), nil
	case OperationMin:
		min := numbers[0]
		for _, n := range numbers[1:] {
			if n < min {
				min = n
			}
		}
		return float64(min), nil
	case OperationStddev:
		var mean float64
		for _, n := range numbers {
			mean += float64(n)
		}
		mean /= float64(len(numbers))
		var variance float64
		for _, n := range numbers {
			diff := float64(n) - mean
			variance += diff * diff
		}
		return math.Sqrt(variance / float64(len(numbers))), nil
	case OperationMedian:
		sorted := make([]int64, len(numbers))
		copy(sorted, numbers)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
		mid := len(sorted) / 2
		if len(sorted)%2 == 0 {
			return float64(sorted[mid-1]+sorted[mid]) / 2, nil
		}
		return float64(sorted[mid]), nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
