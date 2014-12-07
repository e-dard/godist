godist
======

[![Build Status](https://drone.io/github.com/e-dard/godist/status.png)](https://drone.io/github.com/e-dard/godist/latest)

[![GoDoc](https://godoc.org/github.com/e-dard/godist?status.svg)](http://godoc.org/github.com/e-dard/godist)

`godist` provides some Go implementations of useful continuous and
discrete probability distributions, as well as some handy methods for
working with them.

The general idea is that I will add to these over time, but that each
distribution will implement the following interface:

```go
type Distribution interface{
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
```

In practice, distributions may also provide other useful methods, where
appropriate.

The intentions of `godist` is not to provide the fastest, most efficient
implementations, but instead to provide idiomatic Go implementations
that can be easily understood and extended. Having said that, where
there are useful and well-understood numerical tricks and tools to
improve performance, these have been utilised and documented.

Contributions welcome!

### Current Distributions

- Beta Distribution
- Empirical Distribution
