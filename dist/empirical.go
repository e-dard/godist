package dist

import (
	"fmt"
	"sort"
)

type Empirical struct {
	sample    []float64
	n         float64
	mean      float64
	median    float64
	medianIdx int
	variance  float64
}

var EmptyInputError = fmt.Errorf("No values in sample")

func (e *Empirical) Add(values ...float64) {
	if len(values) == 0 {
		return
	}

	e.sample = append(e.sample, values...)

	if len(e.sample) == 1 {

	}

	for _, v := range values {
		if e.n == 0 {
			e.mean, e.median = values[0], values[0]
			e.n++
			continue
		}

		oldM := e.mean
		e.n++
		e.mean += (v - e.mean) / e.n
		e.variance += (v - oldM) * (v - e.mean)
	}

}

// Mean returns the mean of the input values.
func (e Empirical) Mean() (float64, error) {
	if len(e.sample) == 0 {
		return 0.0, EmptyInputError
	}
	return e.mean, nil
}

// SampleMedian returns the median of the input values.
//
// In the case that the length of the input values is even, the mean of
// the two middle values is returned.
func (e Empirical) SampleMedian() (float64, error) {
	values := e.sample
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
func (e Empirical) SampleMode() (float64, error) {
	values := e.sample
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

// Variance returns the variance of the input values.
func (e Empirical) Variance() (float64, error) {
	if len(e.sample) == 0 {
		return 0.0, EmptyInputError
	}
	return e.variance / e.n, nil
}
