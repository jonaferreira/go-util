package util

import (
	"math"
)

// Get the lowest of multiple int values.
func MinInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[1]
		} else {
			return a[0]
		}
	default:
		min := math.MaxInt32
		for _, i := range a {
			if i < min {
				min = i
			}
		}
		return min
	}
}

// Get the highest of multiple int values.
func MaxInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[0]
		} else {
			return a[1]
		}
	default:
		max := math.MinInt32
		for _, i := range a {
			if i > max {
				max = i
			}
		}
		return max
	}
}
