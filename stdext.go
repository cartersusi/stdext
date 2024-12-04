package stdext

import (
	"os"

	stdscrypto "github.com/cartersusi/stdext/crypto"
	stderr "github.com/cartersusi/stdext/errors"
	stdexec "github.com/cartersusi/stdext/exec"
	stdext_ext "github.com/cartersusi/stdext/ext"
	stdfilepath "github.com/cartersusi/stdext/filepath"
	stdflag "github.com/cartersusi/stdext/flag"
	stdmath "github.com/cartersusi/stdext/math"
	stdmmap "github.com/cartersusi/stdext/mmap"
	stdzstd "github.com/cartersusi/stdext/zstd"
)

// import ("github.com/cartersusi/stdext/crypto") for full package documentation
func Encrypt(data []byte, key ...interface{}) ([]byte, error) {
	return stdscrypto.Encrypt(data, key...)
}

// import ("github.com/cartersusi/stdext/crypto") for full package documentation
func Decrypt(data []byte, key ...interface{}) ([]byte, error) {
	return stdscrypto.Decrypt(data, key...)
}

// import ("github.com/cartersusi/stdext/errors") for full package documentation
func HandleError(err *error, action ...interface{}) error {
	return stderr.HandleError(err, action...)
}

// import ("github.com/cartersusi/stdext/exec") for full package documentation
func RunCMD(fs string) error {
	return stdexec.RunCMD(fs)
}

// import ("github.com/cartersusi/stdext/exec") for full package documentation
func RunReturnCMD(fs string) (string, error) {
	return stdexec.RunReturnCMD(fs)
}

// import ("github.com/cartersusi/stdext/exec") for full package documentation
func RunLooseCMD(fs string) {
	stdexec.RunLooseCMD(fs)
}

// import ("github.com/cartersusi/stdext/exec") for full package documentation
func RunReallyLooseCMD(fs string) {
	stdexec.RunReallyLooseCMD(fs)
}

// import ("github.com/cartersusi/stdext/filepath") for full package documentation
func ListDir(fpath string, recurse ...bool) ([]string, error) {
	return stdfilepath.ListDir(fpath, recurse...)
}

// import ("github.com/cartersusi/stdext/flag") for full package documentation
func GetFlag[T string | int | bool](long_flag, short_flag *T, value T, necessary bool, flag_name ...string) T {
	return stdflag.GetFlag(long_flag, short_flag, value, necessary, flag_name...)
}

// import ("github.com/cartersusi/stdext/math") for full package documentation
func DotProduct(left, right []float32, result *float32) {
	stdmath.DotProduct(left, right, result)
}

// import ("github.com/cartersusi/stdext/math") for full package documentation
func Max[T stdmath.Numeric](a, b T) T {
	return stdmath.Max(a, b)
}

// import ("github.com/cartersusi/stdext/math") for full package documentation
func Min[T stdmath.Numeric](a, b T) T {
	return stdmath.Min(a, b)
}

// import ("github.com/cartersusi/stdext/mmap") for full package documentation
func ReadFile(file_path string, chunk_size ...int) ([]byte, error) {
	return stdmmap.ReadFile(file_path, chunk_size...)
}

// import ("github.com/cartersusi/stdext/zstd") for full package documentation
func Compress(data interface{}, file *os.File, level ...int) error {
	return stdzstd.Compress(data, file, level...)
}

// import ("github.com/cartersusi/stdext/zstd") for full package documentation
func Decompress(fpath string) ([]byte, error) {
	return stdzstd.Decompress(fpath)
}

// import ("github.com/cartersusi/stdext/ext") for full package documentation
func Ternary[T any](condition bool, first T, second T) T {
	return stdext_ext.Ternary(condition, first, second)
}

// import ("github.com/cartersusi/stdext/ext") for full package documentation
func Red(s string) string {
	return stdext_ext.Red(s)
}

// import ("github.com/cartersusi/stdext/ext") for full package documentation
func Green(s string) string {
	return stdext_ext.Green(s)
}

// import ("github.com/cartersusi/stdext/ext") for full package documentation
func Yellow(s string) string {
	return stdext_ext.Yellow(s)
}
