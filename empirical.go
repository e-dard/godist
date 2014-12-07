package godist

import "sort"

type Empirical struct {
	sample   []float64
	mean     float64
	median   float64
	mode     float64
	variance float64
	n        float64
	medStale bool
	modStale bool
}

func (e *Empirical) Add(values ...float64) {
	if len(values) == 0 {
		return
	}
	e.sample = append(e.sample, values...)

	// update moments
	for _, v := range values {
		if e.n == 0 {
			e.n = 1
			e.mean, e.median, e.mode = values[0], values[0], values[0]
			e.medStale, e.modStale = false, false
			continue
		}

		// update running mean and variance
		e.n++
		curmean := e.mean
		e.mean += (v - e.mean) / e.n
		e.variance += (v - curmean) * (v - e.mean)

		// check if we need to make the current median/mods values
		// stale.
		if v != e.median {
			e.medStale = true
		}

		if v != e.mode {
			e.modStale = true
		}
	}

}

// Mean returns the distribution mean.
func (e *Empirical) Mean() (float64, error) {
	if len(e.sample) == 0 {
		msg := "mean cannot be calculated on empty distribution."
		return 0.0, EmptyDistributionError{s: msg}
	}
	return e.mean, nil
}

// Median calculates the distribution median.
//
// Median returns a memoised median if either: (1) the distribution has
// not been updated since the last call to Median, or (2) all values
// added to the distribution since the last call are equal to the median
// of the distribution.
//
// In the case that the distribution sample size is even, the mean of
// the two middle values is returned.
func (e *Empirical) Median() (float64, error) {
	if len(e.sample) == 0 {
		msg := "median cannot be calculated on empty distribution."
		return 0.0, EmptyDistributionError{s: msg}
	}

	if !e.medStale {
		// no new values, or only values equal to current median added
		return e.median, nil
	}

	e.medStale = false
	// sort sample to find median value
	sort.Float64s(e.sample)
	mid := int64(e.n) / 2
	if int64(e.n)%2 == 1 {
		e.median = e.sample[mid]
		return e.median, nil
	}
	e.median = (e.sample[mid-1] + e.sample[mid]) / 2.0
	return e.median, nil
}

// Mode calculates the distribution mode.
//
// Mode returns a memoised if either: (1) the distribution has not been
// updated since the last call to Mode, or (2) all values added to the
// distribution since the last call are equal to the mode of the
// distribution.
//
// In the case that the distribution is multi-modal, the smallest mode
// is returned.
func (e *Empirical) Mode() (float64, error) {
	if len(e.sample) == 0 {
		msg := "mode cannot be calculated on empty distribution."
		return 0.0, EmptyDistributionError{s: msg}
	}

	if !e.modStale {
		// no new values, or only values equal to current median added
		return e.mode, nil
	}

	e.modStale = false
	sort.Float64s(e.sample)

	modei, maxc := 0, 1
	for i := 0; i < int(e.n); i++ {
		count := 1
		for j := i + 1; j < int(e.n); j++ {
			if e.sample[j] != e.sample[i] {
				break
			}
			count++
		}

		if count > maxc {
			modei, maxc = i, count
		}
	}
	e.mode = e.sample[modei]
	return e.mode, nil
}

// Variance returns the distribution variance.
func (e *Empirical) Variance() (float64, error) {
	if len(e.sample) == 0 {
		msg := "variance cannot be calculated on empty distribution."
		return 0.0, EmptyDistributionError{s: msg}
	}
	return e.variance / e.n, nil
}
