//go:build arm64
// +build arm64

package neon

/*
#cgo CFLAGS: -O3

#include <stdbool.h>
float DotProduct(float *left, float *right, int len, float result);
bool Supported();
*/
import "C"
import "unsafe"

func DotProduct(left, right []float32, result float32) float32 {
	return float32(C.DotProduct(
		(*C.float)(unsafe.Pointer(&left[0])),
		(*C.float)(unsafe.Pointer(&right[0])),
		C.int(len(right)),
		C.float(result),
	))
}

func Supported() bool {
	return bool(C.Supported())
}
