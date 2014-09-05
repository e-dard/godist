package util

import (
	"math"
)

// FloatsEqual determines if two values are within epsilon of each other.
func FloatsEqual(f1, f2, epsilon float64) bool {
	return math.Abs(f1-f2) < epsilon
}

// FloatsDeciEqual determines if two values are within 10^-1 of each other.
func FloatsDeciEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.1
}

// FloatsCentiEqual determines if two values are within 10^-2 of each other.
func FloatsCentiEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.01
}

// FloatsMilliEqual determines if two values are within 10^-3 of each other.
func FloatsMilliEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.001
}

// FloatsNanoEqual determines if two values are within 10^-9 of each other.
func FloatsNanoEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.000000001
}

// FloatsPicoEqual determines if two values are within 10^-12 of each other.
func FloatsPicoEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.000000000001
}
