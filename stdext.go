package stdext

import (
	"os"

	stdscrypto "github.com/cartersusi/stdext/crypto"
	stderr "github.com/cartersusi/stdext/errors"
	stdexec "github.com/cartersusi/stdext/exec"
	stdfilepath "github.com/cartersusi/stdext/filepath"
	stdflag "github.com/cartersusi/stdext/flag"
	stdmath "github.com/cartersusi/stdext/math"
	stdmmap "github.com/cartersusi/stdext/mmap"
	stdzstd "github.com/cartersusi/stdext/zstd"
)

func Encrypt(data []byte, key ...interface{}) ([]byte, error) {
	return stdscrypto.Encrypt(data, key...)
}

func Decrypt(data []byte, key ...interface{}) ([]byte, error) {
	return stdscrypto.Decrypt(data, key...)
}

func HandleError(err *error, action ...interface{}) error {
	return stderr.HandleError(err, action...)
}

func RunCMD(fs string) error {
	return stdexec.RunCMD(fs)
}

func RunReturnCMD(fs string) (string, error) {
	return stdexec.RunReturnCMD(fs)
}

func RunLooseCMD(fs string) {
	stdexec.RunLooseCMD(fs)
}

func RunReallyLooseCMD(fs string) {
	stdexec.RunReallyLooseCMD(fs)
}

func ListDir(fpath string, recurse ...bool) ([]string, error) {
	return stdfilepath.ListDir(fpath, recurse...)
}

func GetFlag[T string | int | bool](long_flag, short_flag *T, value T, necessary bool, flag_name ...string) T {
	return stdflag.GetFlag(long_flag, short_flag, value, necessary, flag_name...)
}

func DotProduct(left, right []float32, result *float32) {
	stdmath.DotProduct(left, right, result)
}

func Max[T stdmath.Numeric](a, b T) T {
	return stdmath.Max(a, b)
}

func Min[T stdmath.Numeric](a, b T) T {
	return stdmath.Min(a, b)
}

func ReadFile(file_path string, chunk_size ...int) ([]byte, error) {
	return stdmmap.ReadFile(file_path, chunk_size...)
}

func Compress(data interface{}, file *os.File, level ...int) error {
	return stdzstd.Compress(data, file, level...)
}

func Decompress(fpath string) ([]byte, error) {
	return stdzstd.Decompress(fpath)
}
