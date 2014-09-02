package util

import (
	"math"
)

// FloatsEqual determines if two values are within epsilon of each other.
func FloatsEqual(f1, f2, epsilon float64) bool {
	return math.Abs(f1-f2) < epsilon
}
