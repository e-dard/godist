package dist

import (
	"fmt"
	"sort"
)

var EmptyInputError = fmt.Errorf("Input slice is empty")

// SampleMean returns the mean of the input values.
func SampleMean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, EmptyInputError
	}

	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values)), nil
}

// SampleMedian returns the median of the input values.
//
// In the case that the length of the input values is even, the mean of
// the two middle values is returned.
func SampleMedian(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, EmptyInputError
	}
	sort.Float64s(values)
	mid := len(values) / 2
	if len(values)%2 == 1 {
		return values[mid], nil
	}
	return (values[mid-1] + values[mid]) / 2.0, nil
}

// SampleMode returns the mode of the input values.
//
// If the input value distribution is multi-modal then the smallest mode
// is returned.
func SampleMode(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, EmptyInputError
	}
	sort.Float64s(values)
	modei, maxc := 0, 1
	for i := 0; i < len(values); i++ {
		count := 1
		for j := i + 1; j < len(values); j++ {
			if values[j] != values[i] {
				break
			}
			count++
		}

		if count > maxc {
			modei, maxc = i, count
		}
	}
	return values[modei], nil
}

// SampleVar returns the variance of the input values.
func SampleVar(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, EmptyInputError
	}
	mu, _ := SampleMean(values)
	var sum float64
	for _, v := range values {
		sum += (v - mu) * (v - mu)
	}
	return sum / float64(len(values)), nil
}
