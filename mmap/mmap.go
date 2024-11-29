package mmap

import (
	"io"
	"os"
	"sync"

	"golang.org/x/exp/mmap"
)

type Chunk struct {
	Index int
	Data  []byte
}

// ReadFile reads in large files using memory-mapped I/O.
// It reads the file in chunks and returns the entire file as a byte slice.
//
// Parameters:
//   - file_path: string
//   - chunk_size: int (optional)
//
// Returns:
//   - []byte
//   - error
func ReadFile(file_path string, chunk_size ...int) ([]byte, error) {
	_chunk_size := 64 * 1024
	if len(chunk_size) > 0 {
		_chunk_size = chunk_size[0]
	}

	reader, err := mmap.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	file_info, err := os.Stat(file_path)
	if err != nil {
		return nil, err
	}

	file_size := file_info.Size()

	var wg sync.WaitGroup
	res := make(chan Chunk, 10)

	for offset, chunk_index := int64(0), 0; offset < file_size; chunk_index++ {
		remaining_size := file_size - offset
		current_chunk_size := int64(_chunk_size)
		if remaining_size < current_chunk_size {
			current_chunk_size = remaining_size
		}

		large_chunk := make([]byte, current_chunk_size+1024)
		n, err := reader.ReadAt(large_chunk, offset)
		if err != nil && err != io.EOF {
			return nil, err
		}

		actual_chunk_size := int64(n)
		for i := n - 1; i >= 0; i-- {
			if large_chunk[i] == '\n' {
				actual_chunk_size = int64(i + 1)
				break
			}
		}

		chunk := large_chunk[:actual_chunk_size]
		offset += actual_chunk_size

		wg.Add(1)
		go func(idx int, chunk_data []byte) {
			defer wg.Done()
			res <- Chunk{Index: idx, Data: chunk_data}
		}(chunk_index, chunk)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	chunk_map := make(map[int][]byte)
	for chunk := range res {
		chunk_map[chunk.Index] = chunk.Data
	}

	var result []byte
	for i := 0; i < len(chunk_map); i++ {
		result = append(result, chunk_map[i]...)
	}

	return result, nil
}
