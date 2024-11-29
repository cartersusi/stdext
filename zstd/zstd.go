package fops

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
)

// Compress compresses data and writes it to a file.
//
// Parameters:
//   - data: interface{}(Must be *bytes.Buffer | []byte)
//   - file: *os.File
//   - level: ...int
//
// Returns:
//   - error
func Compress(data interface{}, file *os.File, level ...int) error {
	compression_lvl := zstd.SpeedDefault
	if len(level) > 0 {
		switch level[0] {
		case 1:
			compression_lvl = zstd.SpeedFastest
		case 2:
			compression_lvl = zstd.SpeedDefault
		case 3:
			compression_lvl = zstd.SpeedBetterCompression
		case 4:
			compression_lvl = zstd.SpeedBestCompression
		default:
			compression_lvl = zstd.SpeedDefault
		}
	}
	opts := []zstd.EOption{zstd.WithEncoderLevel(compression_lvl)}

	var compression_data []byte
	switch v := data.(type) {
	case *bytes.Buffer:
		compression_data = v.Bytes()
	case []byte:
		compression_data = v
	default:
		return errors.New("Invalid data type, must be *bytes.Buffer or []byte")
	}

	enc, err := zstd.NewWriter(file, opts...)
	if err != nil {
		return err
	}
	defer enc.Close()

	_, err = enc.Write(compression_data)
	if err != nil {
		return err
	}

	return nil
}

// Decompress reads a compressed file and returns the decompressed data.
//
// Parameters:
//   - fpath: string
//
// Returns:
//   - []byte
//   - error
func Decompress(fpath string) ([]byte, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	zstdReader, err := zstd.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer zstdReader.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, zstdReader)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
