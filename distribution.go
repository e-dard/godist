// Package godist provides a collection of useful continuous and
// discrete probability distributions, as well as helpful associated
// methods around them.
//
package godist

type InvalidDistributionError struct{ S string }
type UnsupportedError struct{ S string }

func (e InvalidDistributionError) Error() string { return e.S }
func (e UnsupportedError) Error() string         { return e.S }

// Distribution is the interface that defines useful methods for
// understanding specific instances of distributions, and sampling
// random variates from them.
type Distribution interface {
	Mean() (float64, error)
	Median() (float64, error)
	Mode() (float64, error)
	Variance() (float64, error)

	// generate a random value according to the probability distribution
	Float64() (float64, error)
}
