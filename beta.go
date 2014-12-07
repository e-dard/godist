package godist

import (
	"fmt"
	"math"
	"math/rand"
)

// A Beta distribution is a continuous probability distribution on the
// range [0, 1], which can be formed using shape parameters α, β > 0.
//
// Beta distributions have many uses, with one of the more common ones
// being to model random variables.
type Beta struct {
	Alpha float64
	Beta  float64
}

// Mean returns the mean of the Beta distribution, i.e., α / (α + β)
func (beta Beta) Mean() (float64, error) {
	if ok, err := beta.valid(); !ok {
		return 0, err
	}
	return beta.Alpha / (beta.Alpha + beta.Beta), nil
}

// Median returns the median of the Beta distribution.
//
// Since there is no closed-form expression for the median of a Beta
// distribution, we use a number of known closed-form special cases, and
// an approximation for the general case where α > 1 and β > 1.
//
// Currently, Median cannot calculate the median where α and β are < 1
// unless α = β.
func (beta Beta) Median() (float64, error) {
	aa, bb := beta.Alpha, beta.Beta
	if ok, err := beta.valid(); !ok {
		return 0, err
	}

	if aa == 1 && bb > 0 {
		return 1.0 - math.Pow(2.0, -1/bb), nil
	} else if aa > 0 && bb == 1 {
		return math.Pow(2.0, -1/aa), nil
	} else if aa == 3 && bb == 2 {
		return 0.6142724318, nil
	} else if aa == 2 && bb == 3 {
		return 0.3857275681, nil
	}

	if aa == bb {
		return 0.5, nil
	} else if aa < 1 && bb < 1 {
		msg := fmt.Sprintf("Median not supported for Beta Distribution [α = %v, β = %v]", aa, bb)
		return 0, UnsupportedError{S: msg}
	}

	// when α > 1 and β > 1 use approximation
	// Kerman J (2011) - "A closed-form approximation for the median of
	// 					  the beta distribution"
	return (aa - 1.0/3.0) / (aa + bb - (2.0 / 3.0)), nil
}

// Mode returns the mode of the Beta distribution.
func (beta Beta) Mode() (float64, error) {
	aa, bb := beta.Alpha, beta.Beta
	if ok, err := beta.valid(); !ok {
		return 0, err
	}

	if aa <= 1 || bb <= 1 {
		msg := fmt.Sprintf("Mode not supported for Beta Distribution [α = %v, β = %v]", aa, bb)
		return 0, InvalidDistributionError{S: msg}
	}
	return (aa - 1) / (aa + bb - 2), nil
}

// Variance returns the variance of the Beta Distribution.
func (beta Beta) Variance() (float64, error) {
	aa, bb := beta.Alpha, beta.Beta
	if ok, err := beta.valid(); !ok {
		return 0, err
	}
	return (aa * bb) / (math.Pow(aa+bb, 2) * (aa + bb + 1)), nil
}

// Float64 returns a random variate from the Beta Distribution.
//
// Float64 makes use of four different algorithms for generating random
// variates, depending on the values of α and β. This implementation is
// based on the work done by Kevin Karplus in gen_beta.c
// (https://compbio.soe.ucsc.edu/gen_sequence/gen_beta.c)
func (beta Beta) Float64() (float64, error) {
	aa, bb := beta.Alpha, beta.Beta
	if ok, err := beta.valid(); !ok {
		return 0, err
	}
	a, b := math.Min(aa, bb), math.Max(aa, bb)

	if b < 0.5 {
		// Jöhnk (1964)
		return genBetaJohnk(aa, bb), nil
	} else if a <= 1.0 {
		// Cheng BC (1978)
		return genBetaChengBC(aa, bb, a, b), nil

	}
	// Cheng BB (1978)
	return genBetaChengBB(aa, bb, a, b), nil
}

// genBetaJohnk generates a random variate from a Beta distribution with
// shape parameters aa and bb, according to Jöhnk's algorithm, described
// by Dagpunar in "Principles of Random Variate Generation" (1988).
func genBetaJohnk(aa, bb float64) float64 {
	u, y := rand.Float64(), rand.Float64()
	return math.Pow(u, 1/aa) / (math.Pow(u, 1/aa) + math.Pow(y, 1/bb))
}

// genBetaChengBB generates a random variate from a Beta distribution
// with shape parameters aa and bb, according to Cheng's BB algorithm,
// described in "Generating beta variates with non-integral shape
// parameters" (1978).
func genBetaChengBB(aa, bb, a, b float64) float64 {
	alpha := a + b
	beta := math.Sqrt((alpha - 2.0) / (2.0*a*b - alpha))
	gamma := a + 1.0/beta

	var r, s, t, v, w, z float64
	complete := func() bool {
		u1, u2 := rand.Float64(), rand.Float64()

		v = beta * math.Log(u1/(1.0-u1))
		if v <= 709.78 {
			w = a * math.Exp(v)
			if math.IsInf(w, 0) {
				w = math.MaxFloat64
			}
		} else {
			w = math.MaxFloat64
		}

		z = u1 * u1 * u2
		r = gamma*v - 1.3862944
		s = a + r - w

		if s+2.609438 >= 5.0*z {
			return true
		}

		t = math.Log(z)
		return s > t
	}

	if !complete() {
		for r+alpha*math.Log(alpha/(b+w)) < t {
			if complete() {
				break
			}
		}
	}

	if aa != a {
		return b / (b + w)
	}
	return w / (b + w)
}

// genBetaChengBC generates a random variate from a Beta distribution
// with shape parameters aa and bb, according to Cheng's BC algorithm,
// described in "Generating beta variates with non-integral shape
// parameters" (1978).
func genBetaChengBC(aa, bb, a, b float64) float64 {
	var u1, u2, v, w, y, z float64
	alpha := a + b
	beta := 1.0 / a
	delta := 1.0 + b - a
	k1 := delta * (0.0138889 + 0.0416667*a) / (b*beta - 0.777778)
	k2 := 0.25 + (0.5+0.25/delta)*a

	setVW := func() {
		v = beta * math.Log(u1/(1.0-u1))
		if v <= 709.78 {
			w = b * math.Exp(v)
			if math.IsInf(w, 0) {
				w = math.MaxFloat64
			}
		} else {
			w = math.MaxFloat64
		}
	}

	for {
		u1, u2 = rand.Float64(), rand.Float64()
		if u1 < 0.5 {
			y = u1 * u2
			z = u1 * y
			if 0.25*u2+z-y >= k1 {
				continue
			}
		} else {
			z = u1 * u1 * u2
			if z <= 0.25 {
				setVW()
				break
			}
			if z >= k2 {
				continue
			}
		}

		setVW()

		if alpha*(math.Log(alpha/(a+w))+v)-1.3862944 >= math.Log(z) {
			break
		}
	}
	if aa == a {
		return a / (a + w)
	}
	return w / (a + w)
}

func (beta Beta) valid() (bool, error) {
	if beta.Alpha == 0 || beta.Beta == 0 {
		msg := fmt.Sprintf("Invalid Beta Distribution: [α = %v, β = %v]", beta.Alpha, beta.Beta)
		return false, InvalidDistributionError{S: msg}
	}
	return true, nil
}
