package godist

type EmptyDistributionError struct{ s string }

func (e EmptyDistributionError) Error() string { return e.s }

type Distribution interface {
	// distribution mean
	Mean() (float64, error)

	// distribution median
	Median() (float64, error)

	// distribution mode
	Mode() (float64, error)

	// distribution variance
	Variance() (float64, error)

	// generate a random value according to the probability distribution
	Float64() (float64, error)
}
