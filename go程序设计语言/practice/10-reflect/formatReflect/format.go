package format

import (
	"reflect"
	"strconv"
)

//Any return the format string of any value.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

//we can't expose func with reflect directly
func formatAtom(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.String:
		return strconv.Quote(value.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return value.Type().String() + " 0x" +
			strconv.FormatUint(uint64(value.Pointer()), 16)
	default:
		return value.Type().String() + " value"
	}
}
