//go:build amd64

package gosimd

import "github.com/cartersusi/stdext/math/gosimd/avx"

func getSimdImplementation() simdInterface {
	return &avxImplementation{}
}

type avxImplementation struct{}

func (a *avxImplementation) dotProduct(left, right []float32, result *float32) {
	if avx.Supported() {
		res := avx.DotProduct(left, right, *result)
		*result = res
		return
	}

	dotProduct(left, right, result)
}
