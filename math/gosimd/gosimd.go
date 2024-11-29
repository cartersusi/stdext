package gosimd

import "runtime"

const (
	small = 65536   //1 << 16
	med   = 262144  //1 << 18
	large = 1048576 //1 << 20

	outer_block_size = 32
	inner_block_size = 16
	cache_line_size  = 64
)

func getOptimalChunkSize(vectorSize int) int {
	numCPU := runtime.GOMAXPROCS(0)
	chunkSize := (vectorSize + numCPU - 1) / numCPU

	const minChunkSize = 1024
	if chunkSize < minChunkSize {
		chunkSize = minChunkSize
	}

	chunkSize = ((chunkSize + 15) / 16) * 16
	if chunkSize > vectorSize {
		chunkSize = vectorSize
	}

	return chunkSize
}
