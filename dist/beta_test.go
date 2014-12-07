package dist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/e-dard/godist/util"
)

type betaExample struct {
	in  Beta
	err error
	out float64
}

func Test_Beta_Mean(t *testing.T) {
	examples := []betaExample{
		betaExample{in: Beta{Alpha: 1, Beta: 1}, out: 0.5},
		betaExample{in: Beta{Alpha: 2, Beta: 1}, out: 0.6666666666666666},
		betaExample{in: Beta{Alpha: 0.5, Beta: 10}, out: 0.047619047619047616},
		betaExample{
			in:  Beta{Alpha: 2.0, Beta: 0},
			err: fmt.Errorf("Invalid Beta Distribution: [α = 2, β = 0]"),
		},
		betaExample{
			in:  Beta{Alpha: 0, Beta: 2.0},
			err: fmt.Errorf("Invalid Beta Distribution: [α = 0, β = 2]"),
		},
	}

	for _, ex := range examples {
		actual, err := ex.in.Mean()
		if ex.err != nil && (err == nil || err.Error() != ex.err.Error()) {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !util.FloatsPicoEqual(actual, ex.out) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Beta_Median(t *testing.T) {
	examples := []betaExample{
		betaExample{in: Beta{Alpha: 1, Beta: 0.1}, out: 0.9990234375},
		betaExample{in: Beta{Alpha: 0.1, Beta: 1}, out: 0.0009765625},
		betaExample{in: Beta{Alpha: 1, Beta: 1}, out: 0.5},
		betaExample{in: Beta{Alpha: 2, Beta: 2}, out: 0.5},
		betaExample{in: Beta{Alpha: 3, Beta: 2}, out: 0.6142724318},
		betaExample{in: Beta{Alpha: 2, Beta: 3}, out: 0.3857275681},
		betaExample{in: Beta{Alpha: 20, Beta: 18}, out: 0.5267857142857143},
		betaExample{
			in:  Beta{Alpha: 0, Beta: 0},
			err: fmt.Errorf("Invalid Beta Distribution: [α = 0, β = 0]"),
		},
		betaExample{
			in:  Beta{Alpha: 0.1, Beta: 0.9},
			err: fmt.Errorf("Median not supported for Beta Distribution [α = 0.1, β = 0.9]"),
		},
	}

	for _, ex := range examples {
		actual, err := ex.in.Median()
		if ex.err != nil && (err == nil || err.Error() != ex.err.Error()) {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !util.FloatsPicoEqual(actual, ex.out) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Beta_Mode(t *testing.T) {
	examples := []betaExample{
		betaExample{in: Beta{Alpha: 2, Beta: 2}, out: 0.5},
		betaExample{in: Beta{Alpha: 20, Beta: 20}, out: 0.5},
		betaExample{in: Beta{Alpha: 200, Beta: 20}, out: 0.9128440366972477},
		betaExample{
			in:  Beta{Alpha: 0, Beta: 0},
			err: fmt.Errorf("Invalid Beta Distribution: [α = 0, β = 0]"),
		},
		betaExample{
			in:  Beta{Alpha: 1, Beta: 2},
			err: fmt.Errorf("Mode not supported for Beta Distribution [α = 1, β = 2]"),
		},
	}

	for _, ex := range examples {
		actual, err := ex.in.Mode()
		if ex.err != nil && (err == nil || err.Error() != ex.err.Error()) {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !util.FloatsPicoEqual(actual, ex.out) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Beta_Variance(t *testing.T) {
	examples := []betaExample{
		betaExample{in: Beta{Alpha: 1, Beta: 0.1}, out: 0.03935458480913026},
		betaExample{in: Beta{Alpha: 0.1, Beta: 1}, out: 0.03935458480913026},
		betaExample{in: Beta{Alpha: 1, Beta: 1}, out: 0.08333333333333333},
		betaExample{in: Beta{Alpha: 20, Beta: 4}, out: 0.005555555555555556},
		betaExample{
			in:  Beta{Alpha: 0, Beta: 0},
			err: fmt.Errorf("Invalid Beta Distribution: [α = 0, β = 0]"),
		},
	}

	for _, ex := range examples {
		actual, err := ex.in.Variance()
		if ex.err != nil && (err == nil || err.Error() != ex.err.Error()) {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !util.FloatsPicoEqual(actual, ex.out) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

// tests random variate generation for values using Jöhnk's algorithm
func Test_Float64(t *testing.T) {
	inputs := []Beta{
		// use Jöhnk
		Beta{Alpha: 0.45, Beta: 0.45},
		Beta{Alpha: 0.15, Beta: 0.15},
		// use Cheng BB
		Beta{Alpha: 10, Beta: 3},
		Beta{Alpha: 100, Beta: 30},
		Beta{Alpha: 1000, Beta: 300},
		Beta{Alpha: 10, Beta: 300.25},
		Beta{Alpha: 200, Beta: 3500},
		// use Cheng BC
		Beta{Alpha: 10, Beta: 1.0},
		Beta{Alpha: 1.0, Beta: 12.0},
		Beta{Alpha: 0.6, Beta: 0.6},
		Beta{Alpha: 0.75, Beta: 0.75},
	}

	for _, b := range inputs {
		actual := genBetaDist(b, 10001)
		expMean, _ := b.Mean()
		if !util.FloatsCentiEqual(actual.mean, expMean) {
			msg := "[Mean] expected %v got %v for %#v\n"
			t.Fatalf(msg, expMean, actual.mean, b)
		}

		expMed, _ := b.Median()
		if !util.FloatsDeciEqual(actual.median, expMed) {
			msg := "[Median] expected %v got %v for %#v\n"
			t.Fatalf(msg, expMed, actual.median, b)
		}

		expVar, _ := b.Variance()
		if !util.FloatsDeciEqual(actual.variance, expVar) {
			msg := "[Variance] expected %v got %v for %#v\n"
			t.Fatalf(msg, expVar, actual.variance, b)
		}
	}
}

type dist struct {
	mean     float64
	median   float64
	variance float64
}

func genBetaDist(b Beta, size int) dist {
	rand.Seed(int64(time.Now().Nanosecond()))
	var sample []float64
	for i := 0; i < size; i++ {
		v, _ := b.Float64()
		sample = append(sample, v)
	}

	d := dist{}
	ed := Empirical{}
	ed.Add(sample...)
	d.mean, _ = ed.Mean()
	d.median, _ = ed.Median()
	d.variance, _ = ed.Variance()
	return d
}
