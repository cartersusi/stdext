package gosimd

import (
	"runtime"
	"sync"
)

type simdInterface interface {
	dotProduct(left, right []float32, result *float32)
}

// DotProduct calculates the dot product of two vectors using NEON or AVX SIMD
// instructions if supported, otherwise it falls back to the standard
// implementation.
//
// Parameters:
//   - left: the first vector
//   - right: the second vector
//   - result: the initial value of the result
//
// Returns:
//   - the dot product of the two vectors
func DotProduct(left, right []float32, result *float32) {
	if len(left) != len(right) {
		panic("vectors must be the same length")
	}

	if len(left) < small {
		dotProduct_unroll(left, right, result)
		return
	}

	getSimdImplementation().dotProduct(left, right, result)
}

func dotProduct(left, right []float32, result *float32) {
	dotProduct_np(left, right, result)
}

func dotProduct_np(left, right []float32, result *float32) {
	n := len(left)

	n_cpu := runtime.GOMAXPROCS(0)
	chunk_size := getOptimalChunkSize(n)
	n_chunks := (n + chunk_size - 1) / chunk_size

	ps := make([]float32, n_chunks)
	var wg sync.WaitGroup
	workers := make(chan struct{}, n_cpu)

	for s, i := 0, 0; s < n; s += chunk_size {
		e := s + chunk_size
		if e > n {
			e = n
		}

		workers <- struct{}{}
		wg.Add(1)

		go func(s, e, i int) {
			defer func() {
				<-workers
				wg.Done()
			}()

			var dot float32 = 0
			kk := s

			for ; kk <= e-16; kk += 16 {
				dot += left[kk] * right[kk]
				dot += left[kk+1] * right[kk+1]
				dot += left[kk+2] * right[kk+2]
				dot += left[kk+3] * right[kk+3]
				dot += left[kk+4] * right[kk+4]
				dot += left[kk+5] * right[kk+5]
				dot += left[kk+6] * right[kk+6]
				dot += left[kk+7] * right[kk+7]
				dot += left[kk+8] * right[kk+8]
				dot += left[kk+9] * right[kk+9]
				dot += left[kk+10] * right[kk+10]
				dot += left[kk+11] * right[kk+11]
				dot += left[kk+12] * right[kk+12]
				dot += left[kk+13] * right[kk+13]
				dot += left[kk+14] * right[kk+14]
				dot += left[kk+15] * right[kk+15]
			}

			for ; kk < e; kk++ {
				dot += left[kk] * right[kk]
			}

			ps[i] = dot
		}(s, e, i)

		i++
	}

	wg.Wait()

	for _, sum := range ps {
		*result += sum
	}
}

func dotProduct_unroll(left, right []float32, result *float32) {
	for i := 0; i < len(left)-15; i += 16 {
		*result += left[i] * right[i]
		*result += left[i+1] * right[i+1]
		*result += left[i+2] * right[i+2]
		*result += left[i+3] * right[i+3]
		*result += left[i+4] * right[i+4]
		*result += left[i+5] * right[i+5]
		*result += left[i+6] * right[i+6]
		*result += left[i+7] * right[i+7]
		*result += left[i+8] * right[i+8]
		*result += left[i+9] * right[i+9]
		*result += left[i+10] * right[i+10]
		*result += left[i+11] * right[i+11]
		*result += left[i+12] * right[i+12]
		*result += left[i+13] * right[i+13]
		*result += left[i+14] * right[i+14]
		*result += left[i+15] * right[i+15]
	}

	for i := len(left) - (len(left) % 16); i < len(left); i++ {
		*result += left[i] * right[i]
	}
}
