// package witch offer you a magic to change the value.
package witch

import (
    "fmt"
    "reflect"
)

type Errno int

func process(x, val interface{}, method string) (Errno, bool, error) {
    if !reflect.ValueOf(x).CanAddr() {
        x = reflect.ValueOf(x).Elem()
    }
    return Piu(x, val, method)
}

// Piu changes the value of given variable x, while setted is true,
// we change the value with method Set() otherwise we use pointer to
// change the that value
func Piu(x, val interface{}, method string) (Errno, bool, error) {
    ax := reflect.ValueOf(x).Elem()
    typ := reflect.TypeOf(x).Kind()

    // TODO: refactor belows code for common use
    switch typ {
    case reflect.Int:
        p := ax.Addr().Interface().(int)
        if method == "ptr" {
            if getType(x) == getType(val) {
                p = val.(int)
                return Errno(p), true, nil
            } else {
                return -1, false, UnmatchedType
            }
        }
    case reflect.Uint:
        p := ax.Addr().Interface().(uint)
        if method == "ptr" {
            if getType(x) == getType(val) {
                p = val.(uint)
                return Errno(p), true, nil
            } else {
                return -1, false, UnmatchedType
            }
        }
    }

    if method == "set" {
        res := reflect.ValueOf(val)
        ax.Set(res)
        return 0, true, nil
    } else {
        return -1, false, UnknownMethod
    }
    return -1, false, UnknownType
}

func set(a reflect.Value, v interface{}) (bool, error) {
    vv := reflect.ValueOf(v)
    switch vv.Kind() {
    case reflect.Int:
        a.SetInt(v.(int64))
        return true, nil
    case reflect.Uint:
        a.SetUint(v.(uint64))
        return true, nil
    case reflect.String:
        a.SetString(v.(string))
        return true, nil
    case reflect.Float64:
        a.SetFloat(v.(float64))
        return true, nil
    default:
        if IsSet(vv) {
            a.Set(v.(reflect.Value))
            return true, nil
        } else {
            return false, UnAddrAndSet
        }
    }
}

// TODO: Xiu() undos modification
func Xiu() (bool, error) {
    return true, nil
}

func IsAddr(x interface{}) bool {
    return reflect.ValueOf(x).CanAddr()
}

func IsSet(x interface{}) bool {
    return reflect.ValueOf(x).CanSet()
}

func getType(x interface{}) string {
    switch reflect.ValueOf(x).Kind() {
    case reflect.Int:
        return "int"
    case reflect.Uint:
        return "uint"
    case reflect.Float32:
        return "float32"
    case reflect.Float64:
        return "float64"
    case reflect.Bool:
        return "bool"
    case reflect.String:
        return "string"
    default:
        return ""
    }
}

// for now only support int type
// func ReadUnexported(unexported string) (int, error) {
//     stdout := reflect.ValueOf(os.Stdout).Elem()
//     result := stdout.FieldByName(unexported)
//     return result.Int()
// }

func main() {
    x := 5
    v := 2
    m := "ptr"

    r, t, _ := process(x, v, m)
    fmt.Printf("Variable changed %t: %d\n", t, r)
}
