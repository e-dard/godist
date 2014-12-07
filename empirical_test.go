package godist

import (
	"testing"
)

func Test_Empirical_Imp_Distribution(t *testing.T) {
	var _ Distribution = &Empirical{}
}

func Test_Empirical_Mean(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	e := EmptyDistributionError{
		s: "mean cannot be calculated on empty distribution.",
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.5, 3.0, 3.0}, out: 2.5},
		Example{in: []float64{}, err: e},
		Example{in: nil, err: e},
	}

	for _, ex := range examples {
		em := Empirical{}
		em.Add(ex.in...)
		actual, err := em.Mean()
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Empirical_Median(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	e := EmptyDistributionError{
		s: "median cannot be calculated on empty distribution.",
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.1, 3.1, 2.0}, out: 2.0},
		Example{in: []float64{1.1, 2.0, 3.0, 4.1}, out: 2.5},
		Example{in: []float64{}, err: e},
		Example{in: nil, err: e},
	}

	for _, ex := range examples {
		dist := Empirical{}
		dist.Add(ex.in...)

		actual, err := dist.Median()
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Empirical_SampleMode(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	e := EmptyDistributionError{
		s: "mode cannot be calculated on empty distribution.",
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1, 2.0}, out: 1.1},
		Example{in: []float64{1.1, 2.0, 1.1}, out: 1.1},
		Example{in: []float64{2.0, 1.1, 1.1}, out: 1.1},
		Example{in: []float64{2.0, 1.1, 1.1, 2.0}, out: 1.1},
		Example{in: []float64{1.1, 2.0, 1.1, 2.0, 2.0, 3.021}, out: 2.0},
		Example{in: []float64{}, err: e},
		Example{in: nil, err: e},
	}

	for _, ex := range examples {
		dist := Empirical{}
		dist.Add(ex.in...)
		actual, err := dist.Mode()

		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_Empirical_Variance(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	e := EmptyDistributionError{
		s: "variance cannot be calculated on empty distribution.",
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 0.0},
		Example{in: []float64{1.1, 1.1}, out: 0.0},
		Example{in: []float64{1.1, 1.1, 4.1}, out: 2.0},
		Example{in: []float64{}, err: e},
		Example{in: nil, err: e},
	}

	for _, ex := range examples {
		em := Empirical{}
		em.Add(ex.in...)
		actual, err := em.Variance()
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !floatsPicoEqual(actual, ex.out) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}
