package witch

import (
    "fmt"
    "reflect"
    "strconv"
)

// Any formats x(any value) into a string
func Any(x interface{}) string {
    return formatX(reflect.ValueOf(x))
}

// formatX is the unexported implementation of the Any function
func formatX(x reflect.Value) string {
    switch x.Kind() {
    case reflect.Invalid:
        return "Invalid"
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return strconv.FormatInt(x.Int(), 10)
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return strconv.FormatUint(x.Uint(), 10)
    case reflect.Bool:
        return strconv.FormatBool(x.Bool())
    case reflect.String:
        return strconv.Quote(x.String())
    case reflect.Chan, reflect.Func, reflect.Slice, reflect.Ptr, reflect.Map:
        return x.Type().String() + " 0x" + strconv.FormatUint(uint64(x.Pointer()), 16)
    default:
        return x.Type().String() + " value"
    }
}

func Anyln(x interface{}) {
    fmt.Println(Any(x))
}
