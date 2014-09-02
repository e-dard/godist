package dist

import (
	"github.com/e-dard/godist/util"
	"testing"
)

func Test_SampleMean(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.5, 3.0, 3.0}, out: 2.5},
		Example{in: []float64{}, err: EmptyInputError},
		Example{in: nil, err: EmptyInputError},
	}

	for _, ex := range examples {
		actual, err := SampleMean(ex.in)
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_SampleMedian(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.1, 3.1, 2.0}, out: 2.0},
		Example{in: []float64{1.1, 2.0, 3.0, 4.1}, out: 2.5},
		Example{in: []float64{}, err: EmptyInputError},
		Example{in: nil, err: EmptyInputError},
	}

	for _, ex := range examples {
		actual, err := SampleMedian(ex.in)
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_SampleMode(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1}, out: 1.1},
		Example{in: []float64{1.1, 1.1, 2.0}, out: 1.1},
		Example{in: []float64{1.1, 2.0, 1.1}, out: 1.1},
		Example{in: []float64{2.0, 1.1, 1.1}, out: 1.1},
		Example{in: []float64{2.0, 1.1, 1.1, 2.0}, out: 1.1},
		Example{in: []float64{1.1, 2.0, 1.1, 2.0, 2.0, 3.021}, out: 2.0},
		Example{in: []float64{}, err: EmptyInputError},
		Example{in: nil, err: EmptyInputError},
	}

	for _, ex := range examples {
		actual, err := SampleMode(ex.in)
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if actual != ex.out {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}

func Test_SampleVar(t *testing.T) {
	type Example struct {
		in  []float64
		err error
		out float64
	}

	examples := []Example{
		Example{in: []float64{1.1}, out: 0.0},
		Example{in: []float64{1.1, 1.1}, out: 0.0},
		Example{in: []float64{1.1, 1.1, 4.1}, out: 2.0},
		Example{in: []float64{}, err: EmptyInputError},
		Example{in: nil, err: EmptyInputError},
	}

	for _, ex := range examples {
		actual, err := SampleVar(ex.in)
		if err != ex.err {
			t.Fatalf("expected %v\n got %v\n", ex.err, err)
		}

		if !util.FloatsEqual(actual, ex.out, 0.02) {
			t.Fatalf("expected %v\n got %v\n", ex.out, actual)
		}
	}
}
