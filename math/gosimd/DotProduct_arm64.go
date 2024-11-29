//go:build arm64

package gosimd

import "github.com/cartersusi/stdext/math/gosimd/neon"

func getSimdImplementation() simdInterface {
	return &fallbackImplementation{}
}

type fallbackImplementation struct{}

func (f *fallbackImplementation) dotProduct(left, right []float32, result *float32) {
	if neon.Supported() {
		res := neon.DotProduct(left, right, *result)
		*result = res
		return
	}
	dotProduct(left, right, result)
}
