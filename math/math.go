package math

import (
	"github.com/cartersusi/stdext/math/gosimd"
)

type Numeric interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// DotProduct calculates the dot product of two vectors using NEON or AVX SIMD
// instructions if supported, otherwise it falls back to the standard
// implementation.
//
// Parameters:
//   - left: the first vector
//   - right: the second vector
//   - result: pointer to the result
func DotProduct(left, right []float32, result *float32) {
	gosimd.DotProduct(left, right, result)
}

// Max returns the maximum of two values of any numeric type.
//
// Parameters:
//   - a: the first value
//   - b: the second value
//
// Returns:
//   - the maximum of the two values
func Max[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two values of any numeric type.
//
// Parameters:
//   - a: the first value
//   - b: the second value
//
// Returns:
//   - the minimum of the two values
func Min[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Abs returns the absolute value of a number of any numeric type.
//
// Parameters:
//   - a: the value
//
// Returns:
//   - the absolute value of the number
func Abs[T Numeric](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
