package flag

import (
	"errors"
	"fmt"

	stderr "github.com/cartersusi/stdext/errors"
)

// GetFlag checks the values of the long and short flags. If both flags are set to the default value or both
// are set to non-default values, it triggers an error. Otherwise, it returns the value of the flag that is set.
// Parameters:
//   - long_flag: Pointer to the long-form of the flag.
//   - short_flag: Pointer to the short-form of the flag.
//   - value: Default value for the flags.
//   - necessary: Whether the flag is necessary or not.
//   - flag_name: Optional name of the flag to be used in the error message.s
//
// Returns:
//   - The value of the specified flag (long_flag or short_flag).
func GetFlag[T string | int | bool](long_flag, short_flag *T, value T, necessary bool, flag_name ...string) T {
	flagstr := "flags"
	if len(flag_name) > 0 {
		flagstr = flag_name[0]
	}

	if *long_flag == value && *short_flag == value && necessary {
		if !necessary {
			return value
		}
		err := errors.New(fmt.Sprintf("Missing values for %s. Please provide one.", flagstr))
		stderr.HandleError(&err, stderr.PANIC)
	}

	if *long_flag != value && *short_flag != value {
		if !necessary {
			return value
		}
		err := errors.New(fmt.Sprintf("Both %s are set. Please provide only one.", flagstr))
		stderr.HandleError(&err, stderr.PANIC)
	}

	if *long_flag != value {
		return *long_flag
	}

	return *short_flag
}
