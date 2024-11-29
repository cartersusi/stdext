package errors

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

const (
	PRINT = iota
	LOG
	LOG_EXIT
	PANIC
	NOTHING
)

var (
	DEFAULT = PANIC
)

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func toNumeric(value interface{}) int {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(v.Uint())
	case reflect.Float32, reflect.Float64:
		return int(v.Float())
	default:
		log.Printf("Unhandled type in convert action to numeric value: %T. Defaulting...\n", value)
		return DEFAULT
	}
}

// ErrorHandler is a function that handles errors to reduce if-else statements in code.
// It accepts an error pointer and an optional action.
// The action can be a function that accepts an error, a function that accepts an error pointer, a function that accepts no arguments, or a numeric value.
//
// Parameters:
//   - err: *error
//   - action: interface{}
//
// Returns:
//   - error
func HandleError(err *error, action ...interface{}) error {
	if *err == nil {
		return nil
	}
	*err = errors.New(fmt.Sprintf("ERROR: %s", red((*err).Error())))

	var status int = DEFAULT

	if len(action) > 0 {
		switch v := action[0].(type) {
		case func(error):
			v(*err)
			return *err
		case func(*error):
			v(err)
			return *err
		case func():
			v()
			return *err
		default:
			status = toNumeric(v)
		}
	}

	switch status {
	case PRINT:
		fmt.Println(*err)
	case LOG:
		log.Default().Println(*err)
	case LOG_EXIT:
		log.Fatal(*err)
	case PANIC:
		panic(*err)
	}

	return *err
}
