//go:build amd64
// +build amd64

package avx

func Supported() bool
func DotProduct(left, right []float32, result float32) float32
